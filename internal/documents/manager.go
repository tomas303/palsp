package edit

import (
	"fmt"
	"palsp/internal/discover"

	"github.com/antlr4-go/antlr/v4"
)

var Mgr *Manager

func init() {
	Mgr = &Manager{}
}

type Manager struct {
}

func (mgr *Manager) Init(searchFolders []string, unitScopeNames []string) OpResult {
	for _, folder := range searchFolders {
		discover.SymDB().AddSearchPath(folder)
	}
	discover.SymDB().SetUnitScopeNames(unitScopeNames)
	resp := InitializeResult{
		Capabilities: map[string]interface{}{
			"textDocumentSync":   1, // Full document sync
			"hoverProvider":      true,
			"definitionProvider": true, // Add support for the definition feature
		},
	}
	return OpSuccessWith(resp)
}

func (mgr *Manager) DidOpen(uri string, text string, version int) OpResult {
	if _, err := discover.EditFileCache().Open(uri, text, version); err != nil {
		return OpFailure(fmt.Sprintf("unable to open file %s", uri), err)
	} else {
		mgr.addPath(uri)
		return OpSuccess()
	}
}

func (mgr *Manager) DidChange(uri string, text string, version int) OpResult {
	if _, err := discover.EditFileCache().Open(uri, text, version); err != nil {
		return OpFailure(fmt.Sprintf("unable to open file %s", uri), err)
	} else {
		mgr.addPath(uri)
		return OpSuccess()
	}
}

func (mgr *Manager) DidClose(uri string) OpResult {
	if _, err := discover.EditFileCache().Close(uri); err != nil {
		return OpFailure(fmt.Sprintf("unable to close file %s", uri), err)
	} else {
		return OpSuccess()
	}
}

func (mgr *Manager) Hover(uri string, text string, version int, line int, character int) OpResult {
	var found bool
	var err error
	var fci *discover.FileCacheItem

	if fci, err = discover.EditFileCache().Open(uri, text, version); err != nil {
		return OpFailure(fmt.Sprintf("unable to locate file %s", uri), err)
	}

	var hoverText string
	if hoverText, found = fci.FindText(line, character); !found {
		return OpFailure(fmt.Sprintf("cannot find text on position - URI: %s, line: %d, chr: %d", uri, line, character), err)
	}

	pos := discover.NewPosition(line, character)

	var info string
	writeSymbol := func(sym *discover.Symbol) error {
		info = sym.HoverInfo()
		return discover.ErrFirstSymbolWriten
	}
	writer := discover.SymbolWriterFunc(writeSymbol)
	err = fci.LocateSymbolsByName(hoverText, pos, writer)

	if err == discover.ErrFirstSymbolWriten {
		hoverResp := Hover{
			Contents: MarkupContent{
				Kind:  "plaintext",
				Value: info,
			},
		}
		return OpSuccessWith(hoverResp)
	}

	if err == nil {
		hoverResp := Hover{
			Contents: MarkupContent{
				Kind:  "plaintext",
				Value: "No information found for " + hoverText,
			},
		}
		return OpSuccessWith(hoverResp)
	}

	return OpFailure(fmt.Sprintf("hover error - URI: %s, line: %d, chr: %d", uri, line, character), err)
}

func (mgr *Manager) Completion(uri string, text string, version int, line int, character int) OpResult {
	var found bool
	var err error
	var fci *discover.FileCacheItem

	if fci, err = discover.EditFileCache().Open(uri, text, version); err != nil {
		return OpFailure(fmt.Sprintf("unable to locate file %s", uri), err)
	}

	var hoverText string
	if hoverText, found = fci.FindText(line, character); !found {
		return OpFailure(fmt.Sprintf("cannot find text on position - URI: %s, line: %d, chr: %d", uri, line, character), err)
	}
	pos := discover.NewPosition(line, character)

	items := make([]CompletionItem, 0, 100)
	writeSymbol := func(sym *discover.Symbol) error {
		item := CompletionItem{
			Label:         sym.Name,
			Kind:          symbolKindToCompletionKind(discover.SymbolKind(sym.Kind)),
			Detail:        sym.Definition,
			Documentation: fmt.Sprintf("scope: %s", sym.Path),
		}
		items = append(items, item)
		return nil
	}
	writer := discover.SymbolWriterFunc(writeSymbol)
	err = fci.LocateSymbolsByName(".*"+hoverText+".*", pos, writer)

	if err == nil {
		cl := CompletionList{
			IsIncomplete: false,
			Items:        items,
		}
		return OpSuccessWith(cl)
	}

	return OpFailure(fmt.Sprintf("completion error - URI: %s, line: %d, chr: %d", uri, line, character), err)
}

func (mgr *Manager) Definition(uri string, text string, version int, line int, character int) OpResult {
	var err error
	var fci *discover.FileCacheItem

	if fci, err = discover.EditFileCache().Open(uri, text, version); err != nil {
		return OpFailure(fmt.Sprintf("unable to locate file %s", uri), err)
	}

	var hoverText string
	var found bool
	if hoverText, found = fci.FindText(line, character); !found {
		return OpFailure(fmt.Sprintf("cannot find text on position - URI: %s, line: %d, chr: %d", uri, line, character), err)
	}
	pos := discover.NewPosition(line, character)

	var locations []interface{}
	writeSymbol := func(sym *discover.Symbol) error {
		filePath, err := discover.SymDB().GetUnitPath(sym.Unitname)
		if err != nil {
			return err
		}
		location := map[string]interface{}{
			"uri": "file://" + filePath,
			"range": map[string]interface{}{
				"start": map[string]interface{}{
					"line":      sym.Position.Line - 1,
					"character": sym.Position.Character - 1,
				},
				"end": map[string]interface{}{
					"line":      sym.Position.Line - 1,
					"character": sym.Position.Character - 1 + len(sym.Name),
				},
			},
		}
		locations = append(locations, location)
		return discover.ErrFirstSymbolWriten
	}
	writer := discover.SymbolWriterFunc(writeSymbol)
	err = fci.LocateSymbolsByName(hoverText, pos, writer)

	if err == discover.ErrFirstSymbolWriten {
		return OpSuccessWith(locations)
	}

	if err == nil {
		return OpSuccessWith([]interface{}{})
	}

	return OpFailure(fmt.Sprintf("definition error - URI: %s, line: %d, chr: %d", uri, line, character), err)
}

func (mgr *Manager) getDir(uri string) (string, error) {
	pathElements := discover.DecodePath(uri)
	return pathElements.Dir(), nil
}

func (mgr *Manager) addPath(uri string) {
	dir, err := mgr.getDir(uri)
	if err == nil {
		discover.SymDB().AddSearchPath(dir)
	}
}

func newScope(cst antlr.Tree, unitName string) discover.TopScope {
	collector := discover.NewMemorySymbolCollector(unitName)
	sl := discover.NewScopesListener(collector)
	antlr.ParseTreeWalkerDefault.Walk(sl, cst)
	scope := collector.GetScope()
	scope.WriteToLog()
	return scope
}

// Helper function to convert symbol kinds to LSP completion item kinds
func symbolKindToCompletionKind(kind discover.SymbolKind) int {
	switch kind {
	case discover.FunctionSymbol:
		return 3 // Function
	case discover.ProcedureSymbol:
		return 3 // Function
	case discover.VariableSymbol:
		return 6 // Variable
	case discover.ConstantSymbol:
		return 21 // Constant
	// case discover.TypeSymbol:
	// 	return 22 // Struct (or 7 for Class)
	case discover.ClassSymbol:
		return 7 // Struct (or 7 for Class)
	case discover.UnitReference:
		return 9 // Module
	// case discover.PropertySymbol:
	// 	return 10 // Property
	case discover.ClassVariable:
		return 5 // Field
	case discover.ParameterSymbol:
		return 6 // Variable (for parameters)
	// case discover.EnumValueSymbol:
	// 	return 20 // EnumMember
	default:
		return 1 // Text as fallback
	}
}
