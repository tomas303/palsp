package discover

import (
	"fmt"
	"os"
	"palsp/internal/log"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

// Preprocessor handles include files and compiler defines
type Preprocessor struct {
	defines     map[string]bool
	searchPaths []string
	mutex       sync.RWMutex
}

// SourcePosition tracks the original position of a token
type SourcePosition struct {
	File   string
	Line   int
	Column int
	Length int
}

// PreprocessedContent contains the processed content with position mapping
type PreprocessedContent struct {
	Content     string
	PositionMap []SourcePosition // Maps virtual position to real source position
}

// DirectiveType represents different compiler directive types
type DirectiveType int

const (
	IncludeDirective DirectiveType = iota
	DefineDirective
	UndefDirective
	IfdefDirective
	IfndefDirective
	ElseDirective
	EndifDirective
)

// CompilerDirective represents a parsed compiler directive
type CompilerDirective struct {
	Type      DirectiveType
	Parameter string
	Position  int
	Length    int
}

var globalPreprocessor *Preprocessor
var preprocessorOnce sync.Once

// GetPreprocessor returns the global preprocessor instance
func GetPreprocessor() *Preprocessor {
	preprocessorOnce.Do(func() {
		globalPreprocessor = &Preprocessor{
			defines:     make(map[string]bool),
			searchPaths: make([]string, 0),
		}
	})
	return globalPreprocessor
}

// SetDefines sets the active compiler defines
func (p *Preprocessor) SetDefines(defines []string) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.defines = make(map[string]bool)
	for _, define := range defines {
		p.defines[strings.ToUpper(define)] = true
	}
	log.Main.Info().Msgf("Set %d compiler defines", len(defines))
}

// AddSearchPath adds a search path for include files
func (p *Preprocessor) AddSearchPath(path string) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// Avoid duplicates
	for _, existing := range p.searchPaths {
		if existing == path {
			return
		}
	}
	p.searchPaths = append(p.searchPaths, path)
}

// PreprocessContent processes the content handling includes and defines
func (p *Preprocessor) PreprocessContent(content string, baseFile string) (*PreprocessedContent, error) {
	result := &PreprocessedContent{
		PositionMap: make([]SourcePosition, 0),
	}

	var builder strings.Builder
	defineStack := make([]bool, 0) // Stack for nested ifdef/endif
	currentlyActive := true

	// Find all compiler directives first
	directives := p.parseDirectives(content)

	lastPos := 0
	currentLine := 1
	currentColumn := 1

	// Calculate initial line and column for more accurate position tracking
	for i := 0; i < lastPos && i < len(content); i++ {
		if content[i] == '\n' {
			currentLine++
			currentColumn = 1
		} else {
			currentColumn++
		}
	}

	for _, directive := range directives {
		// Add content before directive if currently active
		if currentlyActive {
			beforeDirective := content[lastPos:directive.Position]
			if len(beforeDirective) > 0 {
				builder.WriteString(beforeDirective)
				p.addPositionMappingSimple(result, beforeDirective, baseFile, currentLine, currentColumn)
			}
		}

		// Update position tracking for the skipped content
		for i := lastPos; i < directive.Position && i < len(content); i++ {
			if content[i] == '\n' {
				currentLine++
				currentColumn = 1
			} else {
				currentColumn++
			}
		}

		// Process the directive
		switch directive.Type {
		case IncludeDirective:
			if currentlyActive {
				if err := p.processInclude(&builder, result, directive.Parameter, baseFile); err != nil {
					log.Main.Warn().Err(err).Msgf("Failed to process include: %s", directive.Parameter)
				}
			}
		case DefineDirective:
			if currentlyActive {
				p.mutex.Lock()
				p.defines[strings.ToUpper(directive.Parameter)] = true
				p.mutex.Unlock()
			}
		case UndefDirective:
			if currentlyActive {
				p.mutex.Lock()
				delete(p.defines, strings.ToUpper(directive.Parameter))
				p.mutex.Unlock()
			}
		case IfdefDirective:
			defineStack = append(defineStack, currentlyActive)
			if currentlyActive {
				p.mutex.RLock()
				currentlyActive = p.defines[strings.ToUpper(directive.Parameter)]
				p.mutex.RUnlock()
			} else {
				currentlyActive = false
			}
		case IfndefDirective:
			defineStack = append(defineStack, currentlyActive)
			if currentlyActive {
				p.mutex.RLock()
				currentlyActive = !p.defines[strings.ToUpper(directive.Parameter)]
				p.mutex.RUnlock()
			} else {
				currentlyActive = false
			}
		case ElseDirective:
			if len(defineStack) > 0 {
				parentActive := defineStack[len(defineStack)-1]
				currentlyActive = parentActive && !currentlyActive
			}
		case EndifDirective:
			if len(defineStack) > 0 {
				currentlyActive = defineStack[len(defineStack)-1]
				defineStack = defineStack[:len(defineStack)-1]
			}
		}

		// Update position tracking past the directive
		for i := directive.Position; i < directive.Position+directive.Length && i < len(content); i++ {
			if content[i] == '\n' {
				currentLine++
				currentColumn = 1
			} else {
				currentColumn++
			}
		}

		lastPos = directive.Position + directive.Length
	}

	// Add remaining content
	if currentlyActive && lastPos < len(content) {
		remaining := content[lastPos:]
		if len(remaining) > 0 {
			builder.WriteString(remaining)
			p.addPositionMappingSimple(result, remaining, baseFile, currentLine, currentColumn)
		}
	}

	result.Content = builder.String()
	return result, nil
}

// parseDirectives finds all compiler directives in the content
func (p *Preprocessor) parseDirectives(content string) []CompilerDirective {
	var directives []CompilerDirective

	// Regex patterns for different directive types
	patterns := map[DirectiveType]*regexp.Regexp{
		IncludeDirective: regexp.MustCompile(`\{\$(?:I|INCLUDE)\s+([^}]+)\}`),
		DefineDirective:  regexp.MustCompile(`\{\$DEFINE\s+([^}]+)\}`),
		UndefDirective:   regexp.MustCompile(`\{\$UNDEF\s+([^}]+)\}`),
		IfdefDirective:   regexp.MustCompile(`\{\$IFDEF\s+([^}]+)\}`),
		IfndefDirective:  regexp.MustCompile(`\{\$IFNDEF\s+([^}]+)\}`),
		ElseDirective:    regexp.MustCompile(`\{\$ELSE\s*\}`),
		EndifDirective:   regexp.MustCompile(`\{\$ENDIF\s*\}`),
	}

	for directiveType, pattern := range patterns {
		matches := pattern.FindAllStringSubmatchIndex(content, -1)
		for _, match := range matches {
			directive := CompilerDirective{
				Type:     directiveType,
				Position: match[0],
				Length:   match[1] - match[0],
			}

			if len(match) > 2 && match[2] != -1 {
				directive.Parameter = strings.TrimSpace(content[match[2]:match[3]])
			}

			directives = append(directives, directive)
		}
	}

	// Sort by position
	for i := 0; i < len(directives)-1; i++ {
		for j := i + 1; j < len(directives); j++ {
			if directives[i].Position > directives[j].Position {
				directives[i], directives[j] = directives[j], directives[i]
			}
		}
	}

	return directives
}

// processInclude handles include file processing
func (p *Preprocessor) processInclude(builder *strings.Builder, result *PreprocessedContent, filename string, baseFile string) error {
	// Clean up filename
	filename = strings.Trim(filename, " \t\r\n'\"")

	// Resolve include file path
	includePath, err := p.resolveIncludePath(filename, baseFile)
	if err != nil {
		return err
	}

	// Read include file content
	includeContent, err := os.ReadFile(includePath)
	if err != nil {
		return fmt.Errorf("failed to read include file %s: %w", includePath, err)
	}

	// Recursively preprocess include content
	preprocessed, err := p.PreprocessContent(string(includeContent), includePath)
	if err != nil {
		return fmt.Errorf("failed to preprocess include file %s: %w", includePath, err)
	}

	// Add preprocessed include content
	builder.WriteString(preprocessed.Content)

	// Add position mappings for included content - preserve original include file positions
	for _, pos := range preprocessed.PositionMap {
		result.PositionMap = append(result.PositionMap, SourcePosition{
			File:   pos.File,   // This will be the include file path
			Line:   pos.Line,   // This will be the line in the include file
			Column: pos.Column, // This will be the column in the include file
			Length: pos.Length,
		})
	}

	return nil
}

// resolveIncludePath resolves the full path of an include file
func (p *Preprocessor) resolveIncludePath(filename string, baseFile string) (string, error) {
	// Try relative to current file first
	if baseFile != "" {
		dir := filepath.Dir(baseFile)
		candidatePath := filepath.Join(dir, filename)
		if _, err := os.Stat(candidatePath); err == nil {
			return candidatePath, nil
		}
	}

	// Try search paths including subdirectories
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	for _, searchPath := range p.searchPaths {
		// First try direct path in search directory
		candidatePath := filepath.Join(searchPath, filename)
		if _, err := os.Stat(candidatePath); err == nil {
			return candidatePath, nil
		}

		// Then search recursively in subdirectories
		if foundPath, err := p.searchInSubdirectories(searchPath, filename); err == nil {
			return foundPath, nil
		}
	}

	return "", fmt.Errorf("include file not found: %s", filename)
}

// searchInSubdirectories recursively searches for a file in subdirectories
func (p *Preprocessor) searchInSubdirectories(rootPath string, filename string) (string, error) {
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

// addPositionMapping adds position mapping information
func (p *Preprocessor) addPositionMapping(result *PreprocessedContent, content string, sourceFile string, originalOffset int, _ int) {
	lines := strings.Split(content, "\n")
	currentLine := 1
	currentColumn := 1

	// Calculate starting line and column from original offset with bounds checking
	if originalOffset > 0 && originalOffset < len(content) {
		// This is a simplified calculation - in a real implementation you might want to maintain
		// more precise position tracking
		beforeOffset := strings.Count(content[:originalOffset], "\n")
		currentLine += beforeOffset
		if beforeOffset > 0 {
			lastNewline := strings.LastIndex(content[:originalOffset], "\n")
			currentColumn = originalOffset - lastNewline
		} else {
			currentColumn = originalOffset + 1
		}
	}

	for i, line := range lines {
		if i > 0 {
			currentLine++
			currentColumn = 1
		}

		if len(line) > 0 {
			result.PositionMap = append(result.PositionMap, SourcePosition{
				File:   sourceFile,
				Line:   currentLine,
				Column: currentColumn,
				Length: len(line),
			})
		}
	}
}

// addPositionMappingSimple adds position mapping for simple cases without complex logic
func (p *Preprocessor) addPositionMappingSimple(result *PreprocessedContent, content string, sourceFile string, startLine int, startColumn int) {
	lines := strings.Split(content, "\n")
	currentLine := startLine
	currentColumn := startColumn

	for i, line := range lines {
		if i > 0 {
			currentLine++
			currentColumn = 1
		}

		if len(line) > 0 {
			result.PositionMap = append(result.PositionMap, SourcePosition{
				File:   sourceFile,
				Line:   currentLine,
				Column: currentColumn,
				Length: len(line),
			})
		}
	}
}

// MapPosition maps a virtual position back to original source position
func (result *PreprocessedContent) MapPosition(virtualPos int) *SourcePosition {
	if virtualPos < 0 || virtualPos >= len(result.PositionMap) {
		return nil
	}

	pos := result.PositionMap[virtualPos]
	return &pos
}
