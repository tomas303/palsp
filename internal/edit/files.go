package edit

import (
	"fmt"
	"net/url"
	"palsp/internal/discover"
	dsc "palsp/internal/discover"
	"path/filepath"

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
	scope *discover.UnitScope
	path  string
	ast   antlr.Tree
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
	l.fls.fileDict[uri] = file{
		scope: d.ScopeSymbols(unitName),
		path:  uri,
		ast:   d.AST(unitName),
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

	var info string
	switch n := node.(type) {
	case antlr.TerminalNode:
		info = n.GetText()
	case antlr.RuleNode:
		info = n.GetText()
	default:
		info = "Unknown node type"
	}

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
	antlr.ParseTreeWalkerDefault.Walk(l, f.ast)
}

func (f *file) findOnPos(line int, character int) (antlr.ParseTree, error) {
	l := newLsnFindOnPos(line, character)
	defer func() {
		if r := recover(); r != nil {
			if r != "found" {
				panic(r)
			}
		}
	}()
	f.walk(l)
	return l.GetFound(), nil
}
