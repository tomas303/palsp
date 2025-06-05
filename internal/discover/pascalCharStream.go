package discover

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"

	"github.com/antlr4-go/antlr/v4"
)

type FileContext struct {
	Filename string
	Content  []rune
}

type SourceFrame struct {
	FileCtx  *FileContext
	Offset   int
	LinesCnt int
}

type defineContext struct {
	defined map[string]bool
	stack   []bool
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

type Region struct {
	mainLine int // line number in virtual buffer(what was all parsed into)
	srcLine  int // line number in source(can be virtual due to includes containing other includes)
	delta    int // mapping of virtual line to source line
	fileCtx  *FileContext
	active   bool // based on defines
}

type pascalCharStream struct {
	buffer      []rune         // Flattened output, filled lazily
	index       int            // Current reading position
	linesCnt    int            // Current line count in virtual buffer
	sourceStack []*SourceFrame // Stack of active sources (includes)
	defineCtx   *defineContext // Tracks active defines and conditio}
	regions     []Region       // Regions for buffer
	defParser   *defineParser  // Parser for directives
	searchPaths []string
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
		sourceStack: []*SourceFrame{{FileCtx: ctx, Offset: 0, LinesCnt: 0}},
		defineCtx:   defCtx,
		regions:     []Region{{mainLine: 0, srcLine: 0, fileCtx: ctx, delta: 0, active: true}},
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
	source := v.sourceStack[len(v.sourceStack)-1]
	for len(v.buffer) <= target {

		if source.Offset >= len(source.FileCtx.Content) {
			if len(v.sourceStack) == 1 {
				break
			}

			v.sourceStack = v.sourceStack[:len(v.sourceStack)-1]
			source := v.sourceStack[len(v.sourceStack)-1]
			newR := Region{
				mainLine: v.linesCnt,
				srcLine:  source.LinesCnt,
				fileCtx:  source.FileCtx,
				delta:    source.LinesCnt,
				active:   v.defineCtx.IsActive(),
			}
			v.regions = append(v.regions, newR)

			continue
		}

		ch := source.FileCtx.Content[source.Offset]

		// Check for comments
		isComment, commentLen, commentLines := v.defParser.ParseCommentFromRunes(
			source.FileCtx.Content,
			source.Offset,
		)
		if isComment {
			// Skip the entire comment
			v.buffer = append(v.buffer, source.FileCtx.Content[source.Offset:source.Offset+commentLen]...)
			source.Offset += commentLen

			// Update line counts based on newlines in the comment
			v.linesCnt += commentLines
			source.LinesCnt += commentLines

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
			source := v.sourceStack[len(v.sourceStack)-1]
			newR := Region{
				mainLine: v.linesCnt,
				srcLine:  source.LinesCnt,
				fileCtx:  source.FileCtx,
				delta:    0,
				active:   v.defineCtx.IsActive(),
			}
			v.regions = append(v.regions, newR)
			continue
		}

		// normal character processing
		if v.defineCtx.IsActive() {
			v.buffer = append(v.buffer, ch)
		} else {
			if unicode.IsSpace(ch) {
				v.buffer = append(v.buffer, ch)
			} else {
				v.buffer = append(v.buffer, ' ')
			}
		}
		if ch == '\n' {
			v.linesCnt++
			source.LinesCnt++
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
			FileCtx:  includeFile,
			Offset:   0,
			LinesCnt: 0,
		}
		v.sourceStack = append(v.sourceStack, includeSource)
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
	// Check if filename contains path separators (relative or absolute path)
	hasPath := strings.Contains(filename, "/") || strings.Contains(filename, "\\") || filepath.IsAbs(filename)

	if hasPath {
		// Filename contains a path (relative or absolute)
		var candidatePath string

		if filepath.IsAbs(filename) {
			// Absolute path - use as is
			candidatePath = filename
		} else {
			// Relative path - resolve relative to current file
			if baseFile != "" {
				dir := filepath.Dir(baseFile)
				candidatePath = filepath.Join(dir, filename)
			} else {
				candidatePath = filename
			}
		}

		// Check if the resolved path exists
		if _, err := os.Stat(candidatePath); err == nil {
			return candidatePath, nil
		}

		// Path specified but file not found - don't search in searchPaths
		return "", fmt.Errorf("include file not found: %s", filename)
	}

	// Filename is just a filename without path - first try relative to current file
	if baseFile != "" {
		dir := filepath.Dir(baseFile)
		candidatePath := filepath.Join(dir, filename)
		if _, err := os.Stat(candidatePath); err == nil {
			return candidatePath, nil
		}
	}

	// File not found relative to current file, now search in searchPaths
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

func (p *defineParser) ParseCommentFromRunes(content []rune, offset int) (bool, int, int) {
	if offset >= len(content) {
		return false, 0, 0
	}

	ch := content[offset]

	// Check for line comment: //
	if ch == '/' && offset+1 < len(content) && content[offset+1] == '/' {
		// Find end of line comment (newline or EOF)
		end := offset + 2
		lineCount := 0

		for end < len(content) && content[end] != '\n' && content[end] != '\r' {
			end++
		}

		// Include the newline character if present and count it
		if end < len(content) && (content[end] == '\n' || content[end] == '\r') {
			lineCount++
			end++
			// Handle \r\n sequences
			if end < len(content) && content[end-1] == '\r' && content[end] == '\n' {
				end++
			}
		}

		return true, end - offset, lineCount
	}

	// Check for block comment: { ... }
	if ch == '{' {
		// Check if it's a directive (starts with {$)
		if offset+1 < len(content) && content[offset+1] == '$' {
			return false, 0, 0 // This is a directive, not a comment
		}

		// Find end of block comment and count newlines
		end := offset + 1
		lineCount := 0

		for end < len(content) && content[end] != '}' {
			if content[end] == '\n' {
				lineCount++
			} else if content[end] == '\r' {
				lineCount++
				// Handle \r\n sequences - don't double count
				if end+1 < len(content) && content[end+1] == '\n' {
					end++
				}
			}
			end++
		}

		if end >= len(content) {
			// Unclosed block comment - consume to end of file
			return true, len(content) - offset, lineCount
		}

		// Include the closing }
		end++
		return true, end - offset, lineCount
	}

	// Check for alternative block comment: (* ... *)
	if ch == '(' && offset+1 < len(content) && content[offset+1] == '*' {
		end := offset + 2
		lineCount := 0

		// Look for closing *)
		for end+1 < len(content) {
			if content[end] == '*' && content[end+1] == ')' {
				end += 2 // Include the closing *)
				return true, end - offset, lineCount
			}

			if content[end] == '\n' {
				lineCount++
			} else if content[end] == '\r' {
				lineCount++
				// Handle \r\n sequences - don't double count
				if end+1 < len(content) && content[end+1] == '\n' {
					end++
				}
			}
			end++
		}

		// Unclosed (* comment - consume to end of file
		return true, len(content) - offset, lineCount
	}

	// No comment found
	return false, 0, 0
}
