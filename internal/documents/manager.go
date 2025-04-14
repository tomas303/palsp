package edit

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"palsp/internal/discover"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

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
}

type files struct {
	// Path to the file.
	fileDict map[string]file
}

type Manager struct {
	fls *files
}

func (mgr *Manager) Init(searchFolders []string) OpResult {
	for _, folder := range searchFolders {
		discover.SymDB().AddSearchPath(folder)
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
	var err error
	var f *file

	if f, err = mgr.locateFile(uri, text, version); err != nil {
		return OpFailure(fmt.Sprintf("unable to locate file %s", uri), nil)
	}

	var hoverText string
	if hoverText, err = f.findText(line, character); err != nil {
		return OpFailure(fmt.Sprintf("cannot find text on position - URI: %s, line: %d, chr: %d", uri, line, character), nil)
	}

	pos := discover.NewPosition(line, character)
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

// One-time definition at package level
type SymbolWriterFunc func(sym *discover.Symbol) error

func (f SymbolWriterFunc) WriteSymbol(sym *discover.Symbol) error {
	return f(sym)
}

func (mgr *Manager) Completion(uri string, text string, version int, line int, character int) OpResult {
	var err error
	var f *file

	if f, err = mgr.locateFile(uri, text, version); err != nil {
		return OpFailure(fmt.Sprintf("unable to locate file %s", uri), nil)
	}

	var hoverText string
	if hoverText, err = f.findText(line, character); err != nil {
		return OpFailure(fmt.Sprintf("cannot find text on position - URI: %s, line: %d, chr: %d", uri, line, character), nil)
	}
	pos := discover.NewPosition(line, character)

	items := make([]CompletionItem, 0, 100)
	writeSymbol := func(sym *discover.Symbol) error {
		item := CompletionItem{
			Label:         sym.Name,
			Kind:          symbolKindToCompletionKind(discover.SymbolKind(sym.Kind)),
			Detail:        sym.Definition,
			Documentation: fmt.Sprintf("scope: %s", sym.Scope),
		}
		items = append(items, item)
		return nil
	}
	writer := SymbolWriterFunc(writeSymbol)
	err = f.scope.LocateSimilarSymbols(hoverText, pos, writer)
	if err != nil {
		return OpFailure(fmt.Sprintf("cannot find symbol - URI: %s, line: %d, chr: %d", uri, line, character), nil)
	}

	if f.scope.IsInImplementation(pos) {
		searchSymbolInUnits2("%"+hoverText+"%", f.scope.ImplementationUses(), writer)
	}
	searchSymbolInUnits2("%"+hoverText+"%", f.scope.InteraceUsese(), writer)

	cl := CompletionList{
		IsIncomplete: false,
		Items:        items,
	}
	return OpSuccessWith(cl)
}

// searchSymbolInUnits looks for a symbol by name in the given list of units
// and returns formatted information if found
func searchSymbolInUnits(symbolName string, units []string) string {
	if len(units) == 0 {
		return ""
	}

	// Limit concurrency to number of CPU cores
	maxWorkers := runtime.NumCPU()

	type searchResult struct {
		unit    string
		symbols []discover.Symbol
		err     error
	}

	// Context for cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create channel for results with buffer
	resultCh := make(chan searchResult, maxWorkers) // Buffer only needs to be as large as max concurrent workers

	// Process units channel for worker scheduling
	unitsCh := make(chan string, len(units))
	for _, unit := range units {
		unitsCh <- unit
	}
	close(unitsCh)

	// Launch only maxWorkers goroutines
	var wg sync.WaitGroup
	for i := 0; i < maxWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			var unitName string
			var ok bool
			for {
				// Get next unit to process
				select {
				case unitName, ok = <-unitsCh:
					if !ok {
						// No more units to process
						return
					}
				case <-ctx.Done():
					// Work cancelled
					return
				}
				// Process the unit
				symbols, err := discover.SymDB().SearchSymbol(unitName, symbolName)
				resultCh <- searchResult{unit: unitName, symbols: symbols, err: err}
				if err == nil && len(symbols) > 0 {
					// Found symbols, cancel other searches
					cancel()
				}
			}
		}()
	}

	// Close result channel when all workers are done
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	// Process results in original unit order (unchanged)
	results := make(map[string][]discover.Symbol, len(units))
	errors := make(map[string]error, len(units))

	for result := range resultCh {
		if result.err == nil {
			results[result.unit] = result.symbols
		} else {
			errors[result.unit] = result.err
		}
	}

	// Check units in original order (unchanged)
	for _, unit := range units {
		if symbols, ok := results[unit]; ok && len(symbols) > 0 {
			return symbols[0].HoverInfo()
		}
	}

	return ""
}

func searchSymbolInUnits2(symbolName string, units []string, writer discover.SymbolWriter) {

	if len(units) == 0 {
		return
	}

	// Limit concurrency to number of CPU cores
	maxWorkers := runtime.NumCPU()

	type searchResult struct {
		unit    string
		symbols []discover.Symbol
		err     error
	}

	// Context for cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create channel for results with buffer
	resultCh := make(chan searchResult, maxWorkers) // Buffer only needs to be as large as max concurrent workers

	// Process units channel for worker scheduling
	unitsCh := make(chan string, len(units))
	for _, unit := range units {
		unitsCh <- unit
	}
	close(unitsCh)

	// Launch only maxWorkers goroutines
	var wg sync.WaitGroup
	for i := 0; i < maxWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			var unitName string
			var ok bool
			for {
				// Get next unit to process
				select {
				case unitName, ok = <-unitsCh:
					if !ok {
						// No more units to process
						return
					}
				case <-ctx.Done():
					// Work cancelled
					return
				}
				// Process the unit
				symbols, err := discover.SymDB().SearchSymbol(unitName, symbolName)
				resultCh <- searchResult{unit: unitName, symbols: symbols, err: err}
			}
		}()
	}

	// Close result channel when all workers are done
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	// Process results in original unit order (unchanged)
	results := make(map[string][]discover.Symbol, len(units))
	errors := make(map[string]error, len(units))

	for result := range resultCh {
		if result.err == nil {
			results[result.unit] = result.symbols
		} else {
			errors[result.unit] = result.err
		}
	}

	// Check units in original order (unchanged)
	for _, unit := range units {
		if symbols, ok := results[unit]; ok {
			for _, symbol := range symbols {
				writer.WriteSymbol(&symbol)
			}
		}
	}

}

func (mgr *Manager) locateFile(uri string, text string, version int) (*file, error) {
	var err error
	if text == "" {
		// Read file content from URI when text is not provided
		var parsed *url.URL
		if parsed, err = url.Parse(uri); err != nil {
			return nil, err
		}
		var content []byte
		if content, err = os.ReadFile(parsed.Path); err != nil {
			return nil, err
		}
		text = string(content) // Assuming UTF-8 encoding
	}
	f, ok := mgr.fls.fileDict[uri]
	if !ok {
		cst := discover.ParseCST(text, uri)
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
		f.cst = discover.ParseCST(text, uri)
		f.scope = newScope(f.cst)
	}
	return &f, nil
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
		discover.SymDB().AddSearchPath(dir)
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

func (f *file) findText(line int, character int) (string, error) {
	node, err := f.findNode(line, character)
	if err != nil {
		return "", err
	}
	return strings.ToLower(node.GetText()), nil
}

func newScope(cst antlr.Tree) discover.TopScope {
	collector := discover.NewMemorySymbolCollector()
	sl := discover.NewUnifiedListener(collector)
	antlr.ParseTreeWalkerDefault.Walk(sl, cst)
	return collector.GetScope()
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
