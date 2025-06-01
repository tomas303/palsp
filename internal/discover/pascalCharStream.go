package discover

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

type FileContext struct {
	Filename string
	Content  []rune
}

type SourceFrame struct {
	FileCtx *FileContext
	Offset  int
	Line    int
	Column  int
}

type defineContext struct {
	defined map[string]bool
	stack   []bool
}

type TrackPos struct {
	pos    int
	line   int
	column int
}

type Region struct {
	start   TrackPos // Start position in the preprocessed buffer
	deshunt TrackPos // Offset adjustment to get to original position
	active  bool
	source  *SourceFrame
}

func NewDefineContext() *defineContext {
	return &defineContext{
		defined: make(map[string]bool),
		stack:   []bool{true},
	}
}

func (d *defineContext) IsActive() bool {
	return d.stack[len(d.stack)-1]
}

func (d *defineContext) Define(name string) {
	d.defined[strings.ToUpper(name)] = true
}

func (d *defineContext) Undef(name string) {
	delete(d.defined, strings.ToUpper(name))
}

func (d *defineContext) IsDefined(name string) bool {
	return d.defined[strings.ToUpper(name)]
}

func (d *defineContext) PushIf(name string) {
	d.stack = append(d.stack, d.defined[name])
}

func (d *defineContext) PopIf() {
	if len(d.stack) > 1 {
		d.stack = d.stack[:len(d.stack)-1]
	}
}

func (d *defineContext) CurrentDefines() []string {
	var keys []string
	for k := range d.defined {
		keys = append(keys, k)
	}
	return keys
}

type pascalCharStream struct {
	buffer       []rune // Flattened output, filled lazily
	linesCnt     int
	columnsCnt   int
	deshuntSize  int            // Size of the deshunt buffer, used for offset adjustments
	deshuntLines int            // Number of lines in the deshunt buffer
	index        int            // Current reading position
	sourceStack  []*SourceFrame // Stack of active sources (includes)
	defineCtx    *defineContext // Tracks active defines and conditio}
	regions      []Region       // Regions for buffer
	defParser    *defineParser  // Parser for directives
	searchPaths  []string
}

func newPascalCharStreamFromFile(filename string, searchPaths []string, defines []string) (*pascalCharStream, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	vchs := newPascalCharStream(string(content), filename, searchPaths, defines)
	return vchs, nil
}

func newPascalCharStream(content string, filename string, searchPaths []string, defines []string) *pascalCharStream {
	ctx := &FileContext{
		Filename: filename,
		Content:  []rune(string(content)),
	}

	defCtx := NewDefineContext()
	for _, def := range defines {
		defCtx.Define(def)
	}

	return &pascalCharStream{
		buffer:      []rune{},
		index:       0,
		sourceStack: []*SourceFrame{{FileCtx: ctx, Offset: 0, Line: 1, Column: 1}},
		defineCtx:   defCtx,
		regions:     []Region{{start: TrackPos{}, deshunt: TrackPos{}, active: true}},
		defParser:   newDefineParser(),
		searchPaths: searchPaths,
	}
}

func (v *pascalCharStream) LA(n int) int {
	// LA (and LT) are 1-based in ANTLR, so 1 means the current character
	target := v.index + n - 1
	v.fillTo(target)

	if target >= len(v.buffer) {
		return -1 // EOF
	}
	return int(v.buffer[target])
}

func (v *pascalCharStream) Consume() {
	v.index++
}

func (v *pascalCharStream) Mark() int { return -1 }

func (v *pascalCharStream) Release(marker int) {}

func (v *pascalCharStream) GetSourceName() string {
	// Get the source file for the current buffer position
	if v.index < len(v.buffer) {
		// Find which region contains the current position
		for i := len(v.regions) - 1; i >= 0; i-- {
			region := v.regions[i]
			if v.index >= region.start.pos {
				// Map back to original source using deshunt info
				return region.source.FileCtx.Filename
			}
		}
	}

	// Fallback to current source
	if len(v.sourceStack) > 0 {
		return v.sourceStack[len(v.sourceStack)-1].FileCtx.Filename
	}
	return "unknown"
}

// Required by ANTLR
func (v *pascalCharStream) Index() int {
	return v.index
}

func (v *pascalCharStream) Size() int {
	return len(v.buffer)
}

func (v *pascalCharStream) Seek(index int) {
	v.index = index
}

func (v *pascalCharStream) GetText(start, stop int) string {
	if start < 0 {
		start = 0
	}
	if stop >= len(v.buffer) {
		stop = len(v.buffer) - 1
	}
	if start > stop {
		return ""
	}
	return string(v.buffer[start : stop+1])
}

func (v *pascalCharStream) GetTextFromTokens(start, end antlr.Token) string {
	if start == nil || end == nil {
		return ""
	}

	startPos := start.GetStart()
	endPos := end.GetStop()

	if startPos > endPos {
		return "" // Invalid range
	}

	return v.GetText(startPos, endPos)
}

func (v *pascalCharStream) GetTextFromInterval(interval antlr.Interval) string {
	start := interval.Start
	stop := interval.Stop

	// Additional validation
	if start < 0 || stop < 0 || start > stop {
		return ""
	}

	// Ensure we don't exceed buffer bounds
	if start >= len(v.buffer) {
		return ""
	}

	return v.GetText(start, stop)
}

func (v *pascalCharStream) fillTo(target int) {
	for len(v.buffer) <= target {
		source := v.sourceStack[len(v.sourceStack)-1]

		if source.Offset >= len(source.FileCtx.Content) {
			if len(v.sourceStack) == 1 {
				break
			}
			v.sourceStack = v.sourceStack[:len(v.sourceStack)-1]
			continue
		}

		ch := source.FileCtx.Content[source.Offset]

		// Check for comments
		isComment, commentLen := v.defParser.ParseCommentFromRunes(
			source.FileCtx.Content,
			source.Offset,
		)
		if isComment {
			// Skip the entire comment
			v.buffer = append(v.buffer, source.FileCtx.Content[source.Offset:source.Offset+commentLen]...)
			source.Offset += commentLen
			continue
		}

		// Check for directives
		directive, value, matchLen := v.defParser.ParseDirectiveFromRunes(
			source.FileCtx.Content,
			source.Offset,
		)
		if directive != -1 {
			v.buffer = append(v.buffer, source.FileCtx.Content[source.Offset:source.Offset+matchLen]...)
			source.Offset += matchLen
			v.handleDirective(directive, value)
			newR := Region{
				start:   TrackPos{pos: len(v.buffer), line: v.linesCnt, column: v.columnsCnt},
				deshunt: TrackPos{pos: v.deshuntSize, line: v.deshuntLines, column: 0},
				active:  v.defineCtx.IsActive(),
				source:  source,
			}
			v.regions = append(v.regions, newR)
			continue
		}

		// normal character processing
		if ch == '\n' {
			v.buffer = append(v.buffer, ch)
			v.linesCnt++
			v.columnsCnt = 0

			if len(v.sourceStack) > 1 {
				v.deshuntLines++
			}

		} else {
			if v.defineCtx.IsActive() {
				v.buffer = append(v.buffer, ch)
			} else {
				v.buffer = append(v.buffer, ' ')
			}
		}
		source.Offset += 1

	}
}

func (v *pascalCharStream) handleDirective(directive defineType, value string) {
	switch directive {
	case includeDI:
		// Handle include directive
		source := v.sourceStack[len(v.sourceStack)-1]
		includeFile, err := v.readInclude(value, source.FileCtx.Filename)
		if err != nil {
			// fmt.Printf("Error reading include file %s: %v\n", value, err)
			return
		}
		includeSource := &SourceFrame{
			FileCtx: includeFile,
			Offset:  0,
			Line:    1,
			Column:  1,
		}
		v.sourceStack = append(v.sourceStack, includeSource)
		v.deshuntSize += len(includeFile.Content)
		// v.deshuntLines += 1 // Assuming each include starts on a new line
		// v.deshuntColumns = 0 // Reset columns for new include
		newR := Region{
			start:   TrackPos{pos: len(v.buffer), line: v.linesCnt, column: v.columnsCnt},
			deshunt: TrackPos{pos: v.deshuntSize, line: v.deshuntLines, column: 0},
			active:  v.defineCtx.IsActive(),
			source:  includeSource,
		}
		v.regions = append(v.regions, newR)
	case defineDI:
		v.defineCtx.Define(value)
	case undefDI:
		v.defineCtx.Undef(value)
	case ifdefDI:
		// Use the helper to evaluate complex expressions
		result := v.defParser.evaluateExpression(value, v.defineCtx)
		v.defineCtx.stack = append(v.defineCtx.stack, result && v.defineCtx.IsActive())
	case ifndefDI:
		result := v.defParser.evaluateExpression(value, v.defineCtx)
		v.defineCtx.stack = append(v.defineCtx.stack, !result && v.defineCtx.IsActive())
	case elseDI:
		if len(v.defineCtx.stack) > 0 {
			v.defineCtx.stack[len(v.defineCtx.stack)-1] = !v.defineCtx.stack[len(v.defineCtx.stack)-1]
		}
	case endifDI:
		v.defineCtx.PopIf()
	}

}

func (v *pascalCharStream) readInclude(filename string, baseFile string) (*FileContext, error) {
	// Clean up filename
	filename = strings.Trim(filename, " \t\r\n'\"")

	// Resolve include file path
	includePath, err := v.resolveIncludePath(filename, baseFile)
	if err != nil {
		return nil, err
	}

	// Read include file content
	includeContent, err := os.ReadFile(includePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read include file %s: %w", includePath, err)
	}

	// Create FileContext for the include file
	includeCtx := &FileContext{
		Filename: includePath,
		Content:  []rune(string(includeContent)),
	}
	return includeCtx, nil
}

func (v *pascalCharStream) resolveIncludePath(filename string, baseFile string) (string, error) {
	// Try relative to current file first
	if baseFile != "" {
		dir := filepath.Dir(baseFile)
		candidatePath := filepath.Join(dir, filename)
		if _, err := os.Stat(candidatePath); err == nil {
			return candidatePath, nil
		}
	}

	// Try search paths including subdirectories

	for _, searchPath := range v.searchPaths {
		// First try direct path in search directory
		candidatePath := filepath.Join(searchPath, filename)
		if _, err := os.Stat(candidatePath); err == nil {
			return candidatePath, nil
		}

		// Then search recursively in subdirectories
		if foundPath, err := v.searchInSubdirectories(searchPath, filename); err == nil {
			return foundPath, nil
		}
	}

	return "", fmt.Errorf("include file not found: %s", filename)
}

func (v *pascalCharStream) searchInSubdirectories(rootPath string, filename string) (string, error) {
	var foundPath string

	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			// Skip directories we can't access
			return nil
		}

		// Skip if it's a directory
		if info.IsDir() {
			return nil
		}

		// Check if this is the file we're looking for
		if info.Name() == filename {
			foundPath = path
			return filepath.SkipAll // Stop walking once found
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	if foundPath != "" {
		return foundPath, nil
	}

	return "", fmt.Errorf("file not found in subdirectories")
}

// DirectiveType represents different compiler directive types
type defineType int

const (
	includeDI defineType = iota
	defineDI
	undefDI
	ifdefDI
	ifndefDI
	elseDI
	endifDI
)

type defineParser struct {
	// Regex patterns for different directive types
	patterns map[defineType]*regexp.Regexp
}

func newDefineParser() *defineParser {
	return &defineParser{
		patterns: map[defineType]*regexp.Regexp{
			includeDI: regexp.MustCompile(`^\{\$(?:I|INCLUDE)\s+([^}]+)\}`),
			defineDI:  regexp.MustCompile(`^\{\$DEFINE\s+([^}]+)\}`),
			undefDI:   regexp.MustCompile(`^\{\$UNDEF\s+([^}]+)\}`),
			ifdefDI:   regexp.MustCompile(`^\{\$IFDEF\s+([^}]+)\}`),
			ifndefDI:  regexp.MustCompile(`^\{\$IFNDEF\s+([^}]+)\}`),
			elseDI:    regexp.MustCompile(`^\{\$ELSE\s*\}`),
			endifDI:   regexp.MustCompile(`^\{\$ENDIF\s*\}`),
		},
	}
}
func (p *defineParser) ParseDirective(text string) (defineType, string) {
	for dt, re := range p.patterns {
		if matches := re.FindStringSubmatch(text); matches != nil {
			if len(matches) > 1 {
				return dt, matches[1]
			}
			return dt, ""
		}
	}
	return -1, "" // No match found
}

func (p *defineParser) ParseDirectiveFromRunes(content []rune, offset int) (defineType, string, int) {
	if offset >= len(content) || content[offset] != '{' {
		return -1, "", 0
	}

	if offset+1 >= len(content) || content[offset+1] != '$' {
		return -1, "", 0
	}

	// Find the end of the directive
	end := offset + 2
	for end < len(content) && content[end] != '}' {
		end++
	}

	if end >= len(content) {
		return -1, "", 0 // No closing }
	}

	// Extract directive content between {$ and }
	directiveContent := strings.TrimSpace(string(content[offset+2 : end]))
	if directiveContent == "" {
		return -1, "", 0
	}

	// Split into parts
	parts := strings.Fields(directiveContent)
	if len(parts) == 0 {
		return -1, "", 0
	}

	directive := strings.ToUpper(parts[0]) // Make directive case-insensitive
	totalLen := end - offset + 1           // Include the closing }

	switch directive {
	case "I", "INCLUDE":
		if len(parts) > 1 {
			// Join all parts after INCLUDE as filename
			filename := strings.Join(parts[1:], " ")
			return includeDI, strings.Trim(filename, `"'`), totalLen
		}
		return includeDI, "", totalLen

	case "DEFINE":
		if len(parts) > 1 {
			// Keep the original case of the define name for proper Pascal behavior
			return defineDI, parts[1], totalLen
		}
		return defineDI, "", totalLen

	case "UNDEF":
		if len(parts) > 1 {
			// Keep the original case of the define name
			return undefDI, parts[1], totalLen
		}
		return undefDI, "", totalLen

	case "IFDEF":
		if len(parts) > 1 {
			// Join the rest as expression for complex conditions
			expression := strings.Join(parts[1:], " ")
			return ifdefDI, expression, totalLen
		}
		return ifdefDI, "", totalLen

	case "IFNDEF":
		if len(parts) > 1 {
			// Join the rest as expression for complex conditions
			expression := strings.Join(parts[1:], " ")
			return ifndefDI, expression, totalLen
		}
		return ifndefDI, "", totalLen

	case "IF":
		// Modern Delphi IF with expressions like: {$IF DEFINED(DEBUG) AND DEFINED(WINDOWS)}
		if len(parts) > 1 {
			expression := strings.Join(parts[1:], " ")
			return ifdefDI, expression, totalLen // Treat as ifdef for now
		}
		return ifdefDI, "", totalLen

	case "IFOPT":
		// Compiler option check like: {$IFOPT R+}
		if len(parts) > 1 {
			expression := strings.Join(parts[1:], " ")
			return ifdefDI, expression, totalLen // Treat as ifdef for now
		}
		return ifdefDI, "", totalLen

	case "ELSE":
		return elseDI, "", totalLen

	case "ELSEIF":
		// Modern Delphi ELSEIF
		if len(parts) > 1 {
			expression := strings.Join(parts[1:], " ")
			return elseDI, expression, totalLen // Treat as else for now
		}
		return elseDI, "", totalLen

	case "ENDIF", "IFEND":
		return endifDI, "", totalLen

	default:
		return -1, "", 0 // Unknown directive
	}
}

// Helper function to evaluate complex expressions (for future use)
func (p *defineParser) evaluateExpression(expression string, defineCtx *defineContext) bool {
	expr := strings.TrimSpace(expression)
	expr = strings.ToUpper(expr) // Make expression case-insensitive

	// Handle DEFINED(symbol) function
	if strings.HasPrefix(expr, "DEFINED(") && strings.HasSuffix(expr, ")") {
		symbol := strings.TrimSpace(expr[8 : len(expr)-1])
		return defineCtx.defined[symbol] // symbol is already uppercase
	}

	// Simple symbol check
	return defineCtx.defined[expr] // expr is already uppercase
}

func (p *defineParser) ParseCommentFromRunes(content []rune, offset int) (bool, int) {
	if offset >= len(content) {
		return false, 0
	}

	ch := content[offset]

	// Check for line comment: //
	if ch == '/' && offset+1 < len(content) && content[offset+1] == '/' {
		// Find end of line comment (newline or EOF)
		end := offset + 2
		for end < len(content) && content[end] != '\n' && content[end] != '\r' {
			end++
		}

		// Include the newline character if present
		if end < len(content) && (content[end] == '\n' || content[end] == '\r') {
			end++
			// Handle \r\n sequences
			if end < len(content) && content[end-1] == '\r' && content[end] == '\n' {
				end++
			}
		}

		return true, end - offset
	}

	// Check for block comment: { ... }
	if ch == '{' {
		// Check if it's a directive (starts with {$)
		if offset+1 < len(content) && content[offset+1] == '$' {
			return false, 0 // This is a directive, not a comment
		}

		// Find end of block comment
		end := offset + 1
		for end < len(content) && content[end] != '}' {
			end++
		}

		if end >= len(content) {
			// Unclosed block comment - consume to end of file
			return true, len(content) - offset
		}

		// Include the closing }
		end++
		return true, end - offset
	}

	// Check for alternative block comment: (* ... *)
	if ch == '(' && offset+1 < len(content) && content[offset+1] == '*' {
		end := offset + 2

		// Look for closing *)
		for end+1 < len(content) {
			if content[end] == '*' && content[end+1] == ')' {
				end += 2 // Include the closing *)
				return true, end - offset
			}
			end++
		}

		// Unclosed (* comment - consume to end of file
		return true, len(content) - offset
	}

	// No comment found
	return false, 0
}
