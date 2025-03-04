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

	var name string
	switch n := node.(type) {
	case antlr.TerminalNode:
		name = strings.ToLower(n.GetText())
	case antlr.RuleNode:
		name = strings.ToLower(n.GetText())
	default:
		return OpFailure(fmt.Sprintf("found node of unexpected type: %T", node), fmt.Errorf("unexpected node type"))
	}

	var info string
	sym := f.scope.LocateSymbol(name, dsc.Position{Line: line, Character: character})
	if sym != nil {
		info = sym.Name + " " + sym.Definition
	} else {
		info = ""
		// based on position find out if is in implementation ,,, probably enough to find out implementation position
		// then go through topscope uses and try to find out. if not exists then load and then find out public symbols
		// classes will be more dificult. Problem is that when loading sql query will be searching again and again.
		// that is not true ... I can put there where condition. So next time.

		//discover.SymDB().IsUnitLoaded(name)
		if f.scope.IsInImplementation(dsc.Position{Line: line, Character: character}) {
			for _, unit := range f.scope.ImplementationUses() {
				if !dsc.SymDB().IsUnitLoaded(unit) {
					d := &dsc.Discover{}
					d.PublicSymbols(unit)
				}
				symbols, err := dsc.SymDB().SearchSymbolsWithinUnit(unit, name)
				if err != nil {
					return OpFailure(fmt.Sprintf("failed to search symbols: %v", err), err)
				}
				for _, sym := range symbols {
					if sym.Name == name {
						info = sym.Name + " " + sym.Definition
						break
					}
				}
			}
		}
		if info == "" {
			for _, unit := range f.scope.InteraceUsese() {
				if !dsc.SymDB().IsUnitLoaded(unit) {
					d := &dsc.Discover{}
					d.PublicSymbols(unit)
				}
				symbols, err := dsc.SymDB().SearchSymbolsWithinUnit(unit, name)
				if err != nil {
					return OpFailure(fmt.Sprintf("failed to search symbols: %v", err), err)
				}
				for _, sym := range symbols {
					if sym.Name == name {
						info = sym.Name + " " + sym.Definition
						break
					}
				}
			}
		}
	}
	if info == "" {
		info = "No information found"
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

func (f *file) findOnPos(line int, character int) (antlr.ParseTree, error) {
	var result antlr.ParseTree
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
