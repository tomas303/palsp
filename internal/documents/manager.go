package edit

import (
	"fmt"
	"os"
	"palsp/internal/discover"
	"strings"

	"palsp/internal/log"

	"github.com/antlr4-go/antlr/v4"
)

var Mgr *Manager

func init() {
	// Create a new lsp instance.
	Mgr = &Manager{
		fls: &files{
			fileDict: make(map[string]file),
		},
	}
}

type file struct {
	uri     string
	version int
	text    string
	scope   discover.TopScope
	cst     antlr.Tree
	stream  antlr.TokenStream
}

type files struct {
	// Path to the file.
	fileDict map[string]file
}

type Manager struct {
	fls *files
}

func (mgr *Manager) Init(searchFolders []string, unitScopeNames []string) OpResult {
	for _, folder := range searchFolders {
		discover.SymDB().AddSearchPath(folder)
	}
	discover.SymDB().SetUnitScopeNames(unitScopeNames)
	resp := InitializeResult{
		Capabilities: map[string]interface{}{
			"textDocumentSync": 1, // Full document sync
			"hoverProvider":    true,
		},
	}
	return OpSuccessWith(resp)
}

func (mgr *Manager) DidOpen(uri string, text string, version int) OpResult {
	var err error
	if _, err = mgr.locateFile(uri, text, version); err != nil {
		return OpFailure(fmt.Sprintf("unable to locate file %s", uri), err)
	}
	mgr.addPath(uri)
	discover.SymDB().DropSymbolsFromPath(uri)
	return OpSuccess()
}

func (mgr *Manager) DidChange(uri string, text string, version int) OpResult {
	mgr.locateFile(uri, text, version)
	return OpSuccess()
}

func (mgr *Manager) DidClose(uri string) OpResult {
	mgr.dropFile(uri)
	return OpSuccess()
}

func (mgr *Manager) Hover(uri string, text string, version int, line int, character int) OpResult {
	var found bool
	var err error
	var f *file

	log.Logger.Debug().Str("file", uri).Msg("Hover requested")
	if f, err = mgr.locateFile(uri, text, version); err != nil {
		return OpFailure(fmt.Sprintf("unable to locate file %s", uri), err)
	}

	var hoverText string
	if hoverText, found = f.findText(line, character); !found {
		return OpFailure(fmt.Sprintf("cannot find text on position - URI: %s, line: %d, chr: %d", uri, line, character), err)
	}

	pos := discover.NewPosition(line, character)

	var info string
	writeSymbol := func(sym *discover.Symbol) error {
		info = sym.HoverInfo()
		return discover.ErrFirstSymbolWriten
	}
	writer := discover.SymbolWriterFunc(writeSymbol)
	err = f.scope.LocateSymbolsByName(hoverText, pos, writer)
	if err != discover.ErrFirstSymbolWriten {
		info = "No information found for " + hoverText
	}

	hoverResp := Hover{
		Contents: MarkupContent{
			Kind:  "plaintext",
			Value: info,
		},
	}
	return OpSuccessWith(hoverResp)
}

func (mgr *Manager) Completion(uri string, text string, version int, line int, character int) OpResult {
	var found bool
	var err error
	var f *file

	if f, err = mgr.locateFile(uri, text, version); err != nil {
		return OpFailure(fmt.Sprintf("unable to locate file %s", uri), err)
	}

	var hoverText string
	if hoverText, found = f.findText(line, character); !found {
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
	err = f.scope.LocateSymbolsByName(".*"+hoverText+".*", pos, writer)

	cl := CompletionList{
		IsIncomplete: false,
		Items:        items,
	}
	return OpSuccessWith(cl)
}

func (mgr *Manager) locateFile(uri string, text string, version int) (*file, error) {
	var err error
	pathElements := discover.DecodePath(uri)
	if text == "" {
		// Read file content from URI when text is not provided
		var content []byte
		if content, err = os.ReadFile(pathElements.Path()); err != nil {
			return nil, err
		}
		text = string(content) // Assuming UTF-8 encoding
	}
	f, ok := mgr.fls.fileDict[uri]
	if !ok {
		cst, stream := discover.ParseCST(text, uri)
		scope := newScope(cst, strings.ToLower(pathElements.Name()))
		f = file{
			uri:     uri,
			version: version,
			text:    text,
			scope:   scope,
			cst:     cst,
			stream:  stream,
		}
		mgr.fls.fileDict[uri] = f
	} else if f.version < version {
		cst, stream := discover.ParseCST(text, uri)
		f.text = text
		f.version = version
		f.scope = newScope(cst, strings.ToLower(pathElements.Name()))
		f.cst = cst
		f.stream = stream
	}
	return &f, nil
}

func (mgr *Manager) dropFile(uri string) {
	delete(mgr.fls.fileDict, uri)
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

func (f *file) walk(l antlr.ParseTreeListener) {
	antlr.ParseTreeWalkerDefault.Walk(l, f.cst)
}

func (f *file) findText(line int, character int) (string, bool) {
	for i := 0; i < f.stream.Size(); i++ {
		token := f.stream.Get(i)
		if token.GetLine() == line &&
			token.GetColumn() <= character &&
			(token.GetColumn()+len(token.GetText()) >= character) {
			return strings.ToLower(token.GetText()), true
		}
	}
	return "", false
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
