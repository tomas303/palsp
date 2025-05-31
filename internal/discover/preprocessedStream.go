package discover

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

type CharMeta struct {
	Source  *FileContext
	Line    int
	Column  int
	Defines []string
}

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

type DefineContext struct {
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
	// source  *SourceFrame
}

func NewDefineContext() *DefineContext {
	return &DefineContext{
		defined: make(map[string]bool),
		stack:   []bool{true},
	}
}

func (d *DefineContext) IsActive() bool {
	return d.stack[len(d.stack)-1]
}

func (d *DefineContext) Define(name string) {
	d.defined[name] = true
}

func (d *DefineContext) Undef(name string) {
	delete(d.defined, name)
}

func (d *DefineContext) PushIf(name string) {
	d.stack = append(d.stack, d.defined[name])
}

func (d *DefineContext) PopIf() {
	if len(d.stack) > 1 {
		d.stack = d.stack[:len(d.stack)-1]
	}
}

func (d *DefineContext) CurrentDefines() []string {
	var keys []string
	for k := range d.defined {
		keys = append(keys, k)
	}
	return keys
}

type VirtualCharStream struct {
	buffer       []rune // Flattened output, filled lazily
	linesCnt     int
	columnsCnt   int
	deshuntSize  int            // Size of the deshunt buffer, used for offset adjustments
	deshuntLines int            // Number of lines in the deshunt buffer
	metaMap      []CharMeta     // Parallel to buffer (for tooling)
	index        int            // Current reading position
	sourceStack  []*SourceFrame // Stack of active sources (includes)
	defineCtx    *DefineContext // Tracks active defines and conditio}
	regions      []Region       // Regions for buffer
	defParser    *defineParser  // Parser for directives
	searchPaths  []string
}

func NewVirtualCharStreamFromFile(filename string) (*VirtualCharStream, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	vchs := NewVirtualCharStream(string(content), filename)
	return vchs, nil
}

func NewVirtualCharStream(content string, filename string) *VirtualCharStream {
	ctx := &FileContext{
		Filename: filename,
		Content:  []rune(string(content)),
	}

	return &VirtualCharStream{
		buffer:      []rune{},
		metaMap:     []CharMeta{},
		index:       0,
		sourceStack: []*SourceFrame{{FileCtx: ctx, Offset: 0, Line: 1, Column: 1}},
		defineCtx:   NewDefineContext(),
		regions:     []Region{{start: TrackPos{}, deshunt: TrackPos{}, active: true}},
		defParser:   newDefineParser(),
	}
}

func (v *VirtualCharStream) LA(n int) int {
	// LA (and LT) are 1-based in ANTLR, so 1 means the current character
	target := v.index + n - 1
	v.fillTo(target)

	if target >= len(v.buffer) {
		return -1 // EOF
	}
	return int(v.buffer[target])
}

func (v *VirtualCharStream) Consume() {
	v.index++
}

func (v *VirtualCharStream) Mark() int { return -1 }

func (v *VirtualCharStream) Release(marker int) {}

func (v *VirtualCharStream) GetSourceName() string {
	return "virtual"
}

// Required by ANTLR
func (v *VirtualCharStream) Index() int {
	return v.index
}

func (v *VirtualCharStream) Size() int {
	return len(v.buffer)
}

func (v *VirtualCharStream) Seek(index int) {
	v.index = index
}

func (v *VirtualCharStream) GetText(start, stop int) string {
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

func (v *VirtualCharStream) GetTextFromTokens(start, end antlr.Token) string {
	return v.GetText(start.GetStart(), end.GetStop())
}

func (v *VirtualCharStream) GetTextFromInterval(antlr.Interval) string {
	return ""
}

func (v *VirtualCharStream) fillTo(target int) {
	for len(v.buffer) <= target {
		source := v.sourceStack[len(v.sourceStack)-1]
		source.Offset += 1

		if source.Offset >= len(source.FileCtx.Content) {
			if len(v.sourceStack) == 1 {
				break
			}
			v.sourceStack = v.sourceStack[:len(v.sourceStack)-1]
			continue
		}

		ch := source.FileCtx.Content[source.Offset]
		if ch == '{' {
			directive, value, matchLen := v.defParser.ParseDirectiveFromRunes(
				source.FileCtx.Content,
				source.Offset,
			)
			if directive != -1 {
				switch directive {
				case includeDI:
					source.Offset += matchLen - 1 // -1 because we'll increment at loop start
					// Handle include directive
					includeFile, err := v.readInclude(value, source.FileCtx.Filename)
					if err != nil {
						// fmt.Printf("Error reading include file %s: %v\n", value, err)
						continue
					}
					v.sourceStack = append(v.sourceStack, &SourceFrame{
						FileCtx: includeFile,
						Offset:  -1,
						Line:    1,
						Column:  1,
					})
					v.deshuntSize += len(includeFile.Content)
					// v.deshuntLines += 1 // Assuming each include starts on a new line
					// v.deshuntColumns = 0 // Reset columns for new include
					newR := Region{
						start:   TrackPos{pos: len(v.buffer), line: v.linesCnt, column: v.columnsCnt},
						deshunt: TrackPos{pos: v.deshuntSize, line: v.deshuntLines, column: 0},
						active:  v.defineCtx.IsActive(),
					}
					v.regions = append(v.regions, newR)
					continue
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
				source.Offset += matchLen - 1 // -1 because we'll increment at loop start

				newR := Region{
					start:   TrackPos{pos: len(v.buffer), line: v.linesCnt, column: v.columnsCnt},
					deshunt: TrackPos{pos: v.deshuntSize, line: v.deshuntLines, column: 0},
					active:  v.defineCtx.IsActive(),
				}
				v.regions = append(v.regions, newR)

				continue
			}
		}

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

	}
}

func (v *VirtualCharStream) readInclude(filename string, baseFile string) (*FileContext, error) {
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

func (v *VirtualCharStream) resolveIncludePath(filename string, baseFile string) (string, error) {
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

func (v *VirtualCharStream) searchInSubdirectories(rootPath string, filename string) (string, error) {
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

	directive := strings.ToUpper(parts[0])
	totalLen := end - offset + 1 // Include the closing }

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
			return defineDI, parts[1], totalLen
		}
		return defineDI, "", totalLen

	case "UNDEF":
		if len(parts) > 1 {
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
func (p *defineParser) evaluateExpression(expression string, defineCtx *DefineContext) bool {
	// For now, simple implementation - just check if the symbol is defined
	// Later you can expand this to handle AND, OR, NOT, DEFINED(), etc.

	// Remove common function calls and operators for simple cases
	expr := strings.TrimSpace(expression)
	expr = strings.ToUpper(expr)

	// Handle DEFINED(symbol) function
	if strings.HasPrefix(expr, "DEFINED(") && strings.HasSuffix(expr, ")") {
		symbol := strings.TrimSpace(expr[8 : len(expr)-1])
		return defineCtx.defined[symbol]
	}

	// Handle simple symbol name
	if strings.Contains(expr, " ") {
		// Complex expression with AND/OR - for now just return true
		// TODO: Implement proper expression parser
		return true
	}

	// Simple symbol check
	return defineCtx.defined[expr]
}
