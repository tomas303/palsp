package discover

import (
	"fmt"
	"os"
	"palsp/internal/log"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

var editFileCache *FileCache

func init() {
	editFileCache = &FileCache{
		fileDict: make(map[string]*FileCacheItem),
	}
}

func EditFileCache() *FileCache {
	return editFileCache
}

type FileCacheItem struct {
	uri      string
	unitName string
	version  int
	text     string
	scope    TopScope
	PData    *ParsedData
	active   bool
	modTime  int64
}

func (f *FileCacheItem) parseGenericTemplate(fromIndex int) string {
	// Add safety check for nil stream
	if f.PData.Stream == nil {
		log.Main.Debug().Msg("stream is nil in parseGenericTemplate")
		return ""
	}

	var beginSymbolID int
	var endSymbolID int
	var semicolonSymbolID int
	var found bool
	if beginSymbolID, found = findParserSymbolicNameID("BEGIN"); !found {
		return ""
	}
	if endSymbolID, found = findParserSymbolicNameID("BEGIN"); !found {
		return ""
	}
	if semicolonSymbolID, found = findParserSymbolicNameID("BEGIN"); !found {
		return ""
	}

	pattern := ""
	depth := 0
	for i := fromIndex; i < f.PData.Stream.Size(); i++ {
		token := f.PData.Stream.Get(i)

		if token.GetTokenType() == beginSymbolID || token.GetTokenType() == endSymbolID ||
			token.GetTokenType() == semicolonSymbolID {
			//unexpected end of template, invalid template
			return pattern
		}

		text := token.GetText()
		switch text {
		case " ", "\t", "\n", "\r":
			// Skip whitespace characters
		case "<":
			depth++
			pattern += "<.*" // for now just placeholde to match all
		case ">":
			if depth > 0 {
				depth--
				pattern += ">"
			} else {
				// Mismatched closing bracket, invalid template
				return pattern
			}
		case ",":
			if depth > 0 {
				pattern += ",.*"
			}
		}
		if depth == 0 {
			// parsing is complete
			break
		}
	}

	return pattern
}

func (f *FileCacheItem) FindText(line int, character int) (string, bool, int) {
	// Add safety check for nil stream
	if f.PData.Stream == nil {
		log.Main.Debug().Msg("stream is nil in FindText")
		return "", false, 0
	}

	var identSymbolID int
	var found bool
	if identSymbolID, found = findParserSymbolicNameID("IDENT"); !found {
		return "", false, 0
	}
	orline, found := f.PData.FindParsedLine(line)
	if found {
		line = orline
	}

	for i := 0; i < f.PData.Stream.Size(); i++ {
		token := f.PData.Stream.Get(i)
		tokenLine := token.GetLine() - 1 // ANTLR lines are 1-based and columns 0-based
		if tokenLine == line &&
			token.GetColumn() <= character &&
			(token.GetColumn()+len(token.GetText()) >= character) {
			if token.GetTokenType() != identSymbolID {
				return "", true, line
			}
			text := token.GetText()
			text += f.parseGenericTemplate(i + 1)
			return text, true, line
		}
	}
	log.Main.Debug().Int("line", line).Int("character", character).Msg("no text found on pos line char")
	return "", true, line
}

func (f *FileCacheItem) LocateSymbolsByName(name string, position Position, writer SymbolWriter) error {
	return f.scope.LocateSymbolsOnPos(name, position, writer)
}

func (f *FileCacheItem) DumpScopes(sb *strings.Builder) {
	f.scope.Dump(sb)
}

func (fci *FileCacheItem) isStale() bool {
	if fileExists(fci.uri) {
		pathElements := DecodePath(fci.uri)
		modTime, err := getFileModTime(pathElements.Path())
		if err != nil {
			return false
		}
		return fci.modTime < modTime
	} else {
		return false
	}
}

type FileCache struct {
	// Path to the file.
	fileDict map[string]*FileCacheItem
}

func (fc *FileCache) Open(uri string, text string, version int) (*FileCacheItem, error) {
	fcitem, err := fc.locateFile(uri, text, version)
	if err != nil {
		return nil, err
	}
	fcitem.active = true
	for _, unit := range fcitem.scope.GetUnits() {
		GetFetcher().AddPrioritized(unit)
	}
	return fcitem, nil
}

func (fc *FileCache) OpenActual(uri string) (*FileCacheItem, error) {
	fcitem, ok := fc.fileDict[uri]
	if !ok {
		return nil, fmt.Errorf("FileCache: file %s not found", uri)
	}
	return fcitem, nil
}

func (fc *FileCache) Close(uri string) (*FileCacheItem, error) {
	fcitem, ok := fc.fileDict[uri]
	if ok {
		fcitem.active = false
		return fcitem, nil
	}
	return nil, fmt.Errorf("FileCache: file %s not found", uri)
}

func (fc *FileCache) FindByUnit(unit string) *FileCacheItem {
	unit = strings.ToLower(unit)
	for _, fcitem := range fc.fileDict {
		if fcitem.unitName == unit {
			return fcitem
		}
	}
	return nil
}

func (fc *FileCache) locateFile(uri string, text string, version int) (*FileCacheItem, error) {
	var err error
	fcitem, ok := fc.fileDict[uri]
	if !ok || fcitem.version < version || text == "" && fcitem.isStale() {
		fcitem, err = newFileCacheItem(uri, text, version)
		if err != nil {
			return nil, err
		}
		fc.fileDict[uri] = fcitem
		return fcitem, nil
	} else {
		return fcitem, nil
	}
}

func (fc *FileCache) dropFile(uri string) {
	delete(fc.fileDict, uri)
}

func getFileContent(uri string, text string) (string, error) {
	if text == "" {
		pathElements := DecodePath(uri)
		if content, err := os.ReadFile(pathElements.Path()); err != nil {
			return "", err
		} else {
			return string(content), nil // Assuming UTF-8 encoding
		}
	} else {
		return text, nil
	}
}

func newFileCacheItem(uri string, text string, version int) (*FileCacheItem, error) {
	var err error
	var content string
	if content, err = getFileContent(uri, text); err != nil {
		return &FileCacheItem{}, err
	} else {
		pathElements := DecodePath(uri)
		var modTime int64

		modTime, err = getFileModTime(pathElements.Path())
		if err != nil {
			return &FileCacheItem{}, err
		}
		pdata := ParseCST(content, uri, false)
		unitName := strings.ToLower(pathElements.Name())
		scope := newScope(pdata, unitName, pathElements.DebugInfo())
		fci := FileCacheItem{
			uri:      uri,
			unitName: unitName,
			version:  version,
			text:     text,
			scope:    scope,
			PData:    pdata,
			active:   true,
			modTime:  modTime,
		}
		return &fci, nil
	}
}

func newScope(pdata *ParsedData, unitName string, debugInfo string) TopScope {
	collector := NewMemorySymbolCollector(unitName)
	sl := NewScopesListener(collector, pdata)
	sl.SetDebugInfo(debugInfo) // Set debug info for position mapping
	antlr.ParseTreeWalkerDefault.Walk(sl, pdata.Tree)
	scope := collector.GetScope()
	return scope
}
