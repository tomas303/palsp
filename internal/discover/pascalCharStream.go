package discover

import (
	"fmt"
	"os"
	"path/filepath"
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

type regionStackMove int

const (
	rsNeutral regionStackMove = iota // No stack change (same file continues)
	rsPush                           // Stack up (entering include file)
	rsPop                            // Stack down (returning from include file)
)

type Region struct {
	mainLine int // line number in virtual buffer(what was all parsed into)
	srcLine  int // line number in source(can be virtual due to includes containing other includes)
	fileCtx  *FileContext
	active   bool            // based on defines
	rsmove   regionStackMove // How this region affects the stack
}

type pascalCharStream struct {
	buffer             []rune         // Flattened output, filled lazily
	index              int            // Current reading position
	linesCnt           int            // Current line count in virtual buffer
	sourceStack        []*SourceFrame // Stack of active sources (includes)
	defineCtx          *defineContext // Tracks active defines and conditio}
	regions            []Region       // Regions for buffer
	defParser          *defineParser  // Parser for directives
	searchPaths        []string
	skipImplementation bool // Skip implementation sections
}

func newPascalCharStreamFromFile(filename string, searchPaths []string, defines []string, skipImplemenation bool) (*pascalCharStream, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	vchs := newPascalCharStream(string(content), filename, searchPaths, defines, skipImplemenation)
	return vchs, nil
}

func newPascalCharStream(content string, filename string, searchPaths []string, defines []string, skipImplemenation bool) *pascalCharStream {
	ctx := &FileContext{
		Filename: filename,
		Content:  []rune(string(content)),
	}

	defCtx := NewDefineContext()
	for _, def := range defines {
		defCtx.Define(def)
	}

	return &pascalCharStream{
		buffer:             []rune{},
		index:              0,
		sourceStack:        []*SourceFrame{{FileCtx: ctx, Offset: 0, LinesCnt: 0}},
		defineCtx:          defCtx,
		regions:            []Region{{mainLine: 0, srcLine: 0, fileCtx: ctx, active: true, rsmove: rsNeutral}},
		defParser:          newDefineParser(),
		searchPaths:        searchPaths,
		skipImplementation: skipImplemenation,
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
				active:   v.defineCtx.IsActive(),
				rsmove:   rsPop,
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
			rsmove := v.handleDirective(directive, value)
			source := v.sourceStack[len(v.sourceStack)-1]
			newR := Region{
				mainLine: v.linesCnt,
				srcLine:  source.LinesCnt,
				fileCtx:  source.FileCtx,
				active:   v.defineCtx.IsActive(),
				rsmove:   rsmove,
			}
			v.regions = append(v.regions, newR)
			continue
		}

		if v.skipImplementation && len(v.sourceStack) == 1 {
			isImplementaion, implLen, _ := v.defParser.ParseImplementationFromRunes(
				source.FileCtx.Content,
				source.Offset,
			)
			if isImplementaion {
				// Skip the entire implementation section
				v.buffer = append(v.buffer, source.FileCtx.Content[source.Offset:source.Offset+implLen]...)
				source.Offset += implLen
				unitend := "\nend."
				v.buffer = append(v.buffer, []rune(unitend)...)
				source.Offset += len(unitend)
				source.LinesCnt += 1 // Count the "end." line
				break
			}
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

func (v *pascalCharStream) handleDirective(directive defineType, value []rune) regionStackMove {
	switch directive {
	case includeDI:
		// Handle include directive - convert rune slice to string for file operations
		filename := string(value)
		source := v.sourceStack[len(v.sourceStack)-1]
		includeFile, err := v.readInclude(filename, source.FileCtx.Filename)
		if err != nil {
			// fmt.Printf("Error reading include file %s: %v\n", filename, err)
			return rsNeutral // Skip this include if error occurs
		}
		includeSource := &SourceFrame{
			FileCtx:  includeFile,
			Offset:   0,
			LinesCnt: 0,
		}
		v.sourceStack = append(v.sourceStack, includeSource)
		return rsPush
	case defineDI:
		// Convert rune slice to string for define operations
		v.defineCtx.Define(string(value))
	case undefDI:
		// Convert rune slice to string for undef operations
		v.defineCtx.Undef(string(value))
	case ifdefDI:
		// Use rune-based expression evaluation
		result := v.defParser.evaluateExpressionRunes(value, v.defineCtx)
		v.defineCtx.stack = append(v.defineCtx.stack, result && v.defineCtx.IsActive())
	case ifndefDI:
		// Use rune-based expression evaluation
		result := v.defParser.evaluateExpressionRunes(value, v.defineCtx)
		v.defineCtx.stack = append(v.defineCtx.stack, !result && v.defineCtx.IsActive())
	case elseDI:
		if len(v.defineCtx.stack) > 0 {
			v.defineCtx.stack[len(v.defineCtx.stack)-1] = !v.defineCtx.stack[len(v.defineCtx.stack)-1]
		}
	case endifDI:
		v.defineCtx.PopIf()
	}
	return rsNeutral
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
}

func newDefineParser() *defineParser {
	return &defineParser{}
}

func (p *defineParser) ParseDirectiveFromRunes(content []rune, offset int) (defineType, []rune, int) {
	if offset >= len(content) || content[offset] != '{' {
		return -1, nil, 0
	}

	if offset+1 >= len(content) || content[offset+1] != '$' {
		return -1, nil, 0
	}

	// Find the end of the directive
	end := offset + 2
	for end < len(content) && content[end] != '}' {
		end++
	}

	if end >= len(content) {
		return -1, nil, 0 // No closing }
	}

	// Extract directive content between {$ and }
	directiveContent := content[offset+2 : end]
	if len(directiveContent) == 0 {
		return -1, nil, 0
	}

	// Find first non-space character to get directive start
	directiveStart := 0
	for directiveStart < len(directiveContent) && isSpace(directiveContent[directiveStart]) {
		directiveStart++
	}

	if directiveStart >= len(directiveContent) {
		return -1, nil, 0
	}

	// Find end of directive keyword
	keywordEnd := directiveStart
	for keywordEnd < len(directiveContent) && !isSpace(directiveContent[keywordEnd]) {
		keywordEnd++
	}

	keyword := directiveContent[directiveStart:keywordEnd]
	totalLen := end - offset + 1 // Include the closing }

	// Convert keyword to uppercase for comparison
	directiveType := p.identifyDirective(keyword)

	switch directiveType {
	case "I", "INCLUDE":
		// Find filename after keyword
		valueStart := keywordEnd
		for valueStart < len(directiveContent) && isSpace(directiveContent[valueStart]) {
			valueStart++
		}
		if valueStart < len(directiveContent) {
			filename := p.trimQuotes(directiveContent[valueStart:])
			return includeDI, filename, totalLen
		}
		return includeDI, nil, totalLen

	case "DEFINE":
		// Find symbol name after keyword
		valueStart := keywordEnd
		for valueStart < len(directiveContent) && isSpace(directiveContent[valueStart]) {
			valueStart++
		}
		if valueStart < len(directiveContent) {
			// Find end of symbol name
			symbolEnd := valueStart
			for symbolEnd < len(directiveContent) && !isSpace(directiveContent[symbolEnd]) {
				symbolEnd++
			}
			symbol := directiveContent[valueStart:symbolEnd]
			return defineDI, symbol, totalLen
		}
		return defineDI, nil, totalLen

	case "UNDEF":
		// Find symbol name after keyword
		valueStart := keywordEnd
		for valueStart < len(directiveContent) && isSpace(directiveContent[valueStart]) {
			valueStart++
		}
		if valueStart < len(directiveContent) {
			// Find end of symbol name
			symbolEnd := valueStart
			for symbolEnd < len(directiveContent) && !isSpace(directiveContent[symbolEnd]) {
				symbolEnd++
			}
			symbol := directiveContent[valueStart:symbolEnd]
			return undefDI, symbol, totalLen
		}
		return undefDI, nil, totalLen

	case "IFDEF":
		// Find expression after keyword
		valueStart := keywordEnd
		for valueStart < len(directiveContent) && isSpace(directiveContent[valueStart]) {
			valueStart++
		}
		if valueStart < len(directiveContent) {
			expression := directiveContent[valueStart:]
			return ifdefDI, expression, totalLen
		}
		return ifdefDI, nil, totalLen

	case "IFNDEF":
		// Find expression after keyword
		valueStart := keywordEnd
		for valueStart < len(directiveContent) && isSpace(directiveContent[valueStart]) {
			valueStart++
		}
		if valueStart < len(directiveContent) {
			expression := directiveContent[valueStart:]
			return ifndefDI, expression, totalLen
		}
		return ifndefDI, nil, totalLen

	case "IF":
		// Modern Delphi IF with expressions like: {$IF DEFINED(DEBUG) AND DEFINED(WINDOWS)}
		valueStart := keywordEnd
		for valueStart < len(directiveContent) && isSpace(directiveContent[valueStart]) {
			valueStart++
		}
		if valueStart < len(directiveContent) {
			expression := directiveContent[valueStart:]
			return ifdefDI, expression, totalLen // Treat as ifdef for now
		}
		return ifdefDI, nil, totalLen

	case "IFOPT":
		// Compiler option check like: {$IFOPT R+}
		valueStart := keywordEnd
		for valueStart < len(directiveContent) && isSpace(directiveContent[valueStart]) {
			valueStart++
		}
		if valueStart < len(directiveContent) {
			expression := directiveContent[valueStart:]
			return ifdefDI, expression, totalLen // Treat as ifdef for now
		}
		return ifdefDI, nil, totalLen

	case "ELSE":
		return elseDI, nil, totalLen

	case "ELSEIF":
		// Modern Delphi ELSEIF
		valueStart := keywordEnd
		for valueStart < len(directiveContent) && isSpace(directiveContent[valueStart]) {
			valueStart++
		}
		if valueStart < len(directiveContent) {
			expression := directiveContent[valueStart:]
			return elseDI, expression, totalLen // Treat as else for now
		}
		return elseDI, nil, totalLen

	case "ENDIF", "IFEND":
		return endifDI, nil, totalLen

	default:
		return -1, nil, 0 // Unknown directive
	}
}

// Helper function to identify directive type from rune slice
func (p *defineParser) identifyDirective(keyword []rune) string {
	if len(keyword) == 0 {
		return ""
	}

	// Convert to uppercase string for comparison
	upper := make([]rune, len(keyword))
	for i, r := range keyword {
		if r >= 'a' && r <= 'z' {
			upper[i] = r - 32
		} else {
			upper[i] = r
		}
	}

	switch string(upper) {
	case "I":
		return "I"
	case "INCLUDE":
		return "INCLUDE"
	case "DEFINE":
		return "DEFINE"
	case "UNDEF":
		return "UNDEF"
	case "IFDEF":
		return "IFDEF"
	case "IFNDEF":
		return "IFNDEF"
	case "IF":
		return "IF"
	case "IFOPT":
		return "IFOPT"
	case "ELSE":
		return "ELSE"
	case "ELSEIF":
		return "ELSEIF"
	case "ENDIF":
		return "ENDIF"
	case "IFEND":
		return "IFEND"
	default:
		return ""
	}
}

// Helper function to check if rune is whitespace
func isSpace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n' || r == '\r'
}

// Helper function to trim quotes from rune slice
func (p *defineParser) trimQuotes(value []rune) []rune {
	if len(value) == 0 {
		return value
	}

	start := 0
	end := len(value)

	// Trim leading quotes
	if value[start] == '"' || value[start] == '\'' {
		start++
	}

	// Trim trailing quotes
	if end > start && (value[end-1] == '"' || value[end-1] == '\'') {
		end--
	}

	if start >= end {
		return nil
	}

	return value[start:end]
}

// Helper function to evaluate complex expressions using rune arrays
func (p *defineParser) evaluateExpressionRunes(expression []rune, defineCtx *defineContext) bool {
	if len(expression) == 0 {
		return false
	}

	// Tokenize the expression
	tokens := p.tokenizeRunes(expression)
	if len(tokens) == 0 {
		return false
	}

	// Parse and evaluate the expression
	result, _ := p.parseOrExpressionRunes(tokens, 0, defineCtx)
	return result
}

// Token types for expression parsing
type runeToken struct {
	kind  exprTokenKind
	start int // Position in original rune array
	end   int // End position
}

type exprTokenKind int

const (
	tokenIdent exprTokenKind = iota
	tokenDefined
	tokenDeclared
	tokenTrue
	tokenFalse
	tokenAnd
	tokenOr
	tokenNot
	tokenXor
	tokenLParen
	tokenRParen
	tokenEOF
)

// Tokenize rune expression into tokens
func (p *defineParser) tokenizeRunes(expr []rune) []runeToken {
	var tokens []runeToken
	i := 0

	for i < len(expr) {
		// Skip whitespace
		if isSpace(expr[i]) {
			i++
			continue
		}

		switch {
		// Parentheses
		case expr[i] == '(':
			tokens = append(tokens, runeToken{tokenLParen, i, i + 1})
			i++

		case expr[i] == ')':
			tokens = append(tokens, runeToken{tokenRParen, i, i + 1})
			i++

		// Keywords and identifiers
		default:
			if isLetter(expr[i]) || expr[i] == '_' {
				start := i
				for i < len(expr) && (isLetterOrDigit(expr[i]) || expr[i] == '_') {
					i++
				}

				// Check for function calls like DEFINED(...)
				if i < len(expr) && expr[i] == '(' {
					// Find matching closing parenthesis
					parenCount := 1
					funcStart := i
					i++ // Skip opening paren
					for i < len(expr) && parenCount > 0 {
						if expr[i] == '(' {
							parenCount++
						} else if expr[i] == ')' {
							parenCount--
						}
						i++
					}

					word := expr[start:funcStart]
					if p.runesEqualIgnoreCase(word, []rune("DEFINED")) {
						tokens = append(tokens, runeToken{tokenDefined, start, i})
					} else if p.runesEqualIgnoreCase(word, []rune("DECLARED")) {
						tokens = append(tokens, runeToken{tokenDeclared, start, i})
					} else {
						// Unknown function, treat as identifier
						tokens = append(tokens, runeToken{tokenIdent, start, i})
					}
				} else {
					// Regular keyword or identifier
					word := expr[start:i]
					if p.runesEqualIgnoreCase(word, []rune("AND")) {
						tokens = append(tokens, runeToken{tokenAnd, start, i})
					} else if p.runesEqualIgnoreCase(word, []rune("OR")) {
						tokens = append(tokens, runeToken{tokenOr, start, i})
					} else if p.runesEqualIgnoreCase(word, []rune("NOT")) {
						tokens = append(tokens, runeToken{tokenNot, start, i})
					} else if p.runesEqualIgnoreCase(word, []rune("XOR")) {
						tokens = append(tokens, runeToken{tokenXor, start, i})
					} else if p.runesEqualIgnoreCase(word, []rune("TRUE")) {
						tokens = append(tokens, runeToken{tokenTrue, start, i})
					} else if p.runesEqualIgnoreCase(word, []rune("FALSE")) {
						tokens = append(tokens, runeToken{tokenFalse, start, i})
					} else {
						tokens = append(tokens, runeToken{tokenIdent, start, i})
					}
				}
			} else {
				// Unknown character, skip
				i++
			}
		}
	}

	tokens = append(tokens, runeToken{tokenEOF, len(expr), len(expr)})
	return tokens
}

// Helper functions for character classification
func isLetter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

func isLetterOrDigit(r rune) bool {
	return isLetter(r) || (r >= '0' && r <= '9')
}

// Fast case-insensitive comparison
func (p *defineParser) runesEqualIgnoreCase(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if p.toLower(a[i]) != p.toLower(b[i]) {
			return false
		}
	}
	return true
}

func (p *defineParser) toLower(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return r + 32
	}
	return r
}

// Parse OR expression (lowest precedence)
func (p *defineParser) parseOrExpressionRunes(tokens []runeToken, pos int, defineCtx *defineContext) (bool, int) {
	left, pos := p.parseXorExpressionRunes(tokens, pos, defineCtx)

	for pos < len(tokens) && tokens[pos].kind == tokenOr {
		pos++ // Skip OR
		right, newPos := p.parseXorExpressionRunes(tokens, pos, defineCtx)
		pos = newPos
		left = left || right
	}

	return left, pos
}

// Parse XOR expression
func (p *defineParser) parseXorExpressionRunes(tokens []runeToken, pos int, defineCtx *defineContext) (bool, int) {
	left, pos := p.parseAndExpressionRunes(tokens, pos, defineCtx)

	for pos < len(tokens) && tokens[pos].kind == tokenXor {
		pos++ // Skip XOR
		right, newPos := p.parseAndExpressionRunes(tokens, pos, defineCtx)
		pos = newPos
		left = (left && !right) || (!left && right) // XOR logic
	}

	return left, pos
}

// Parse AND expression
func (p *defineParser) parseAndExpressionRunes(tokens []runeToken, pos int, defineCtx *defineContext) (bool, int) {
	left, pos := p.parsePrimaryExpressionRunes(tokens, pos, defineCtx)

	for pos < len(tokens) && tokens[pos].kind == tokenAnd {
		pos++ // Skip AND
		right, newPos := p.parsePrimaryExpressionRunes(tokens, pos, defineCtx)
		pos = newPos
		left = left && right
	}

	return left, pos
}

// Parse primary expressions (identifiers, literals, parentheses)
func (p *defineParser) parsePrimaryExpressionRunes(tokens []runeToken, pos int, defineCtx *defineContext) (bool, int) {
	if pos >= len(tokens) {
		return false, pos
	}

	switch tokens[pos].kind {
	case tokenTrue:
		return true, pos + 1

	case tokenFalse:
		return false, pos + 1

	case tokenNot:
		pos++ // Skip NOT
		value, newPos := p.parsePrimaryExpressionRunes(tokens, pos, defineCtx)
		return !value, newPos

	case tokenIdent:
		// Simple identifier - check if it's defined
		// We need access to original expression to extract the identifier
		return true, pos + 1 // Simplified for now

	case tokenDefined:
		// DEFINED(symbol) function - extract symbol and check
		return true, pos + 1 // Simplified for now

	case tokenDeclared:
		// DECLARED(symbol) function
		return true, pos + 1 // Simplified for now

	case tokenLParen:
		// Parenthesized expression
		pos++ // Skip (
		result, newPos := p.parseOrExpressionRunes(tokens, pos, defineCtx)
		pos = newPos
		if pos < len(tokens) && tokens[pos].kind == tokenRParen {
			pos++ // Skip )
		}
		return result, pos

	default:
		return false, pos + 1
	}
}

// Updated evaluateExpression to use rune-based evaluation
func (p *defineParser) evaluateExpression(expression string, defineCtx *defineContext) bool {
	// Convert string to runes and use rune-based evaluator
	return p.evaluateExpressionRunes([]rune(expression), defineCtx)
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

func (p *defineParser) ParseImplementationFromRunes(content []rune, offset int) (bool, int, int) {
	if offset >= len(content) {
		return false, 0, 0
	}

	ch := content[offset]

	// Check for implementation section: (* ... *)
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

		// Unclosed (* implementation section - consume to end of file
		return true, len(content) - offset, lineCount
	}

	// No implementation section found
	return false, 0, 0
}
