package edit

import (
	"fmt"
	"net/url"
	dsc "palsp/internal/discover"
	"path/filepath"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

var Lspi *lsp

func init() {
	// Create a new lsp instance.
	Lspi = &lsp{
		fls: &files{
			fileDict: make(map[string]file),
		},
	}
}

type file struct {
	scope dsc.TopScope
	path  string
	cst   antlr.Tree
}

type files struct {
	// Path to the file.
	fileDict map[string]file
}

type lsp struct {
	fls *files
}

// splitURI splits a URI into its path, name (without extension), and extension.
func getUnitName(uri string) (name string, err error) {
	parsed, err := url.Parse(uri)
	if err != nil {
		return "", err
	}
	dir := parsed.Path
	base := filepath.Base(dir)
	ext := filepath.Ext(base)
	name = base[:len(base)-len(ext)]
	return name, nil
}

func (l *lsp) Init(searchFolders []string) OpResult {
	dsc.SymDB()
	d := &dsc.Discover{}
	for _, folder := range searchFolders {
		d.Units(folder)
	}
	resp := InitializeResult{
		Capabilities: map[string]interface{}{
			"textDocumentSync": 1, // Full document sync
			"hoverProvider":    true,
		},
	}
	return OpSuccessWith(resp)
}

func (l *lsp) DidOpen(uri string, text string) OpResult {
	d := &dsc.Discover{}
	unitName, err := getUnitName(uri)
	if err != nil {
		return OpFailure(fmt.Sprintf("Invalid URI: %s", uri), err)
	}
	delete(l.fls.fileDict, uri)

	parsed, err := url.Parse(uri)
	if err != nil {
		return OpFailure(fmt.Sprintf("Invalid URI: %s", uri), err)
	}
	dir := filepath.Dir(parsed.Path)
	d.Units(dir)

	cst := d.CST(unitName)

	sl := dsc.NewScopeListener("")
	antlr.ParseTreeWalkerDefault.Walk(sl, cst)
	scope := sl.GetScope()

	l.fls.fileDict[uri] = file{
		scope: scope,
		path:  uri,
		cst:   cst,
	}
	return OpSuccess()
}

func (l *lsp) DidChange(uri string, text string) OpResult {
	return l.DidOpen(uri, text)
}

func (l *lsp) DidClose(uri string) OpResult {
	delete(l.fls.fileDict, uri)
	return OpSuccess()
}

func (l *lsp) Hover(uri string, line int, character int) OpResult {

	f, ok := l.fls.fileDict[uri]
	if !ok {
		l.DidOpen(uri, "")
		f = l.fls.fileDict[uri]
	}

	node, err := f.findOnPos(line, character)
	if err != nil {
		return OpFailure(fmt.Sprintf("problem locate position URI: %s, line: %d, chr: %d", uri, line, character), err)
	}
	text := strings.ToLower(node.GetText())

	var info string
	sym := f.scope.LocateSymbol(text, dsc.Position{Line: line, Character: character})
	if sym != nil {
		info = sym.Name + " " + sym.Definition
	} else {
		info = ""

		if f.scope.IsInImplementation(dsc.Position{Line: line, Character: character}) {
			info = searchSymbolInUnits(text, f.scope.ImplementationUses())
		}

		if info == "" {
			info = searchSymbolInUnits(text, f.scope.InteraceUsese())
		}
	}
	if info == "" {
		info = "No information found for " + text
	}

	// fmt.Printf("f.scope: %v\n", f.scope.print())
	// f.scope.Print()

	// sl := discover.NewScopeListener("")
	// f.walk(sl)
	// ts := sl.GetScope()
	// ts.Print()

	hoverResp := Hover{
		Contents: MarkupContent{
			Kind:  "plaintext",
			Value: fmt.Sprintf("Hover info for %s", info),
		},
		// Range: &Range{
		// 	Start: Position{Line: line, Character: character},
		// 	End:   Position{Line: line, Character: character},
		// },
	}
	return OpSuccessWith(hoverResp)
}

// searchSymbolInUnits looks for a symbol by name in the given list of units
// and returns formatted information if found
func searchSymbolInUnits(symbolName string, units []string) string {
	for _, unit := range units {
		if !dsc.SymDB().IsUnitLoaded(unit) {
			dsc.SymDB().RescanUnits()
		}
		d := &dsc.Discover{}
		d.PublicSymbols(unit)
		symbols, err := dsc.SymDB().SearchSymbolsWithinUnit(unit, symbolName)
		if err != nil {
			continue
		}
		for _, sym := range symbols {
			if sym.Name == symbolName {
				return sym.Name + " " + sym.Definition
			}
		}
	}
	return ""
}

func (l *lsp) Completion(uri string, line int, character int) OpResult {
	// Dummy completion data implementation
	items := []CompletionItem{
		{
			Label:         "dummyCompletion",
			Kind:          1,
			Detail:        "Dummy detail",
			Documentation: "Documentation for dummy completion",
		},
	}
	cl := CompletionList{
		IsIncomplete: false,
		Items:        items,
	}
	return OpSuccessWith(cl)
}

func (f *file) walk(l antlr.ParseTreeListener) {
	antlr.ParseTreeWalkerDefault.Walk(l, f.cst)
}

func (f *file) findOnPos(line int, character int) (antlr.TerminalNode, error) {
	var result antlr.TerminalNode
	var err error

	// Function to be executed in a deferred context
	func() {
		// Defer the recovery
		defer func() {
			if r := recover(); r != nil {
				switch v := r.(type) {
				case error:
					// If the panic value is an error, store it
					err = v
				case string:
					// If the panic value is a string, create an error from it
					err = fmt.Errorf("%s", v)
				default:
					// For any other type, create a generic error
					err = fmt.Errorf("unexpected panic: %v", r)
				}
			}
		}()

		// Create the listener and walk the tree
		l := newLsnFindOnPos(line, character)
		f.walk(l)
		result = l.GetFound()
	}()

	// Check if we found a result or if there was an error
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, fmt.Errorf("no node found at position line: %d, character: %d", line, character)
	}

	return result, nil
}
