package edit

import (
	"fmt"
	"net/url"
	"palsp/internal/discover"
	dsc "palsp/internal/discover"
	"path/filepath"
	"strings"

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
	scope   dsc.TopScope
	cst     antlr.Tree
}

type files struct {
	// Path to the file.
	fileDict map[string]file
}

type Manager struct {
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

func (mgr *Manager) Init(searchFolders []string) OpResult {
	for _, folder := range searchFolders {
		dsc.SymDB().AddSearchPath(folder)
	}
	resp := InitializeResult{
		Capabilities: map[string]interface{}{
			"textDocumentSync": 1, // Full document sync
			"hoverProvider":    true,
		},
	}
	return OpSuccessWith(resp)
}

func (mgr *Manager) DidOpen(uri string, text string, version int) OpResult {
	mgr.locateFile(uri, text, version)
	mgr.addPath(uri)
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

	f := mgr.locateFile(uri, text, version)
	if f == nil {
		return OpFailure(fmt.Sprintf("unable to locate file %s", uri), nil)
	}

	hoverText := f.findText(line, character)
	if hoverText == "" {
		return OpFailure(fmt.Sprintf("cannot find text on position - URI: %s, line: %d, chr: %d", uri, line, character), nil)
	}

	pos := dsc.NewPosition(line, character)
	var info string

	symbol := f.scope.LocateSymbol(hoverText, pos)
	if symbol != nil {
		info = symbol.HoverInfo()
	} else {
		if f.scope.IsInImplementation(pos) {
			info = searchSymbolInUnits(hoverText, f.scope.ImplementationUses())
		}
		if info == "" {
			info = searchSymbolInUnits(hoverText, f.scope.InteraceUsese())
		}
	}
	if info == "" {
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

func (mgr *Manager) Completion(uri string, line int, character int) OpResult {
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

// searchSymbolInUnits looks for a symbol by name in the given list of units
// and returns formatted information if found
func searchSymbolInUnits(symbolName string, units []string) string {
	for _, unit := range units {
		symbols, err := dsc.SymDB().SearchSymbolsWithinUnit(unit, symbolName)
		if err != nil {
			continue
		}
		for _, sym := range symbols {
			if sym.Name == symbolName {
				return sym.HoverInfo()
			}
		}
	}
	return ""
}

func (mgr *Manager) locateFile(uri string, text string, version int) *file {
	f, ok := mgr.fls.fileDict[uri]
	if !ok {
		cst := discover.ParseCST(text)
		scope := newScope(cst)
		f = file{
			uri:     uri,
			version: version,
			text:    text,
			scope:   scope,
			cst:     cst,
		}
		mgr.fls.fileDict[uri] = f
	} else if f.version < version {
		f.text = text
		f.version = version
		f.cst = discover.ParseCST(text)
		f.scope = newScope(f.cst)
	}
	return &f
}

func (mgr *Manager) dropFile(uri string) {
	delete(mgr.fls.fileDict, uri)
}

func (mgr *Manager) getDir(uri string) (string, error) {
	parsed, err := url.Parse(uri)
	if err != nil {
		return "", err
	}
	return filepath.Dir(parsed.Path), nil
}

func (mgr *Manager) addPath(uri string) {
	dir, err := mgr.getDir(uri)
	if err == nil {
		dsc.SymDB().AddSearchPath(dir)
	}
}

func (f *file) walk(l antlr.ParseTreeListener) {
	antlr.ParseTreeWalkerDefault.Walk(l, f.cst)
}

func (f *file) findNode(line int, character int) (antlr.TerminalNode, error) {
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

func (f *file) findText(line int, character int) string {
	node, err := f.findNode(line, character)
	if err != nil {
		return ""
	}
	return strings.ToLower(node.GetText())
}

func newScope(cst antlr.Tree) dsc.TopScope {
	sl := dsc.NewScopeListener("")
	antlr.ParseTreeWalkerDefault.Walk(sl, cst)
	return sl.GetScope()
}
