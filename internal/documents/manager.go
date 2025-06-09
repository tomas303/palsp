package edit

import (
	"fmt"
	"palsp/internal/discover"
	"palsp/internal/log"
	"strings"
	"sync"
)

type Manager struct {
}

var mgr *Manager
var mgrOnce sync.Once

func GetManager() *Manager {
	mgrOnce.Do(func() {
		mgr = &Manager{}
	})
	return mgr
}

func (mgr *Manager) Init(searchFolders []string, unitScopeNames []string, prefetchUnits bool, defines []string) OpResult {
	for _, folder := range searchFolders {
		discover.SymDB().AddSearchPath(folder)
	}
	discover.SymDB().SetUnitScopeNames(unitScopeNames)

	// Set compiler defines if provided
	if len(defines) > 0 {
		discover.SymDB().SetDefines(defines)
		log.Main.Info().Msgf("Initialized with %d compiler defines: %v", len(defines), defines)
	}

	discover.GetFetcher().Start()
	if prefetchUnits {
		for _, unit := range discover.SymDB().UnscannedUnits() {
			discover.GetFetcher().AddNormal(unit)
		}
	}
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

func (mgr *Manager) Hover(uri string, line int, character int) OpResult {
	var found bool
	var err error
	var fci *discover.FileCacheItem

	if fci, err = discover.EditFileCache().OpenActual(uri); err != nil {
		return OpFailure(fmt.Sprintf("unable to locate file %s", uri), err)
	}

	var hoverText string
	if hoverText, found, line = fci.FindText(line, character); !found {
		return OpFailure(fmt.Sprintf("cannot find text on position - URI: %s, line: %d, chr: %d", uri, line, character), err)
	}
	if hoverText == "" {
		return OpSuccessWith(Hover{
			Contents: MarkupContent{
				Kind:  "plaintext",
				Value: "",
			},
		})
	}

	pos := discover.NewPosition(line, character)

	var info string
	writeSymbol := func(sym *discover.Symbol) error {
		info = mgr.HoverInfo(sym, fci.PData)
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

func (mgr *Manager) Completion(uri string, line int, character int) OpResult {
	var found bool
	var err error
	var fci *discover.FileCacheItem

	if fci, err = discover.EditFileCache().OpenActual(uri); err != nil {
		return OpFailure(fmt.Sprintf("unable to locate file %s", uri), err)
	}

	var hoverText string
	if hoverText, found, line = fci.FindText(line, character); !found {
		return OpFailure(fmt.Sprintf("cannot find text on position - URI: %s, line: %d, chr: %d", uri, line, character), err)
	}
	if hoverText == "" {
		return OpSuccessWith(CompletionList{
			IsIncomplete: false,
			Items:        []CompletionItem{},
		},
		)
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

func (mgr *Manager) Definition(uri string, line int, character int) OpResult {
	var err error
	var fci *discover.FileCacheItem

	if fci, err = discover.EditFileCache().OpenActual(uri); err != nil {
		return OpFailure(fmt.Sprintf("unable to locate file %s", uri), err)
	}

	var hoverText string
	var found bool
	if hoverText, found, line = fci.FindText(line, character); !found {
		return OpFailure(fmt.Sprintf("cannot find text on position - URI: %s, line: %d, chr: %d", uri, line, character), err)
	}
	if hoverText == "" {
		return OpSuccessWith(map[string]interface{}{})
	}
	pos := discover.NewPosition(line, character)

	var locations []interface{}
	writeSymbol := func(sym *discover.Symbol) error {
		var filePath string
		var err error
		if discover.SymbolKind(sym.Kind) == discover.UnitReference {
			filePath, err = discover.SymDB().GetUnitPath(sym.Name)
		} else {
			filePath, err = discover.SymDB().GetUnitPath(sym.Unitname)
		}
		if err != nil {
			if err == discover.ErrUnitNotFound {
				filePath = ""
			} else {
				return err
			}
		}

		line, found, fileCtx := fci.PData.FindOriginalLine(sym.Position.Line)
		if !found {
			line = sym.Position.Line
		} else {
			// todo: later add better indication of inclcude files
			if fileCtx != nil && strings.HasSuffix(fileCtx.Filename, ".inc") {
				filePath = fileCtx.Filename
			}
		}

		location := map[string]interface{}{
			"uri": discover.FormatFileURI(filePath),
			"range": map[string]interface{}{
				"start": map[string]interface{}{
					"line":      line,
					"character": sym.Position.Character,
				},
				"end": map[string]interface{}{
					"line":      line,
					"character": sym.Position.Character + len(sym.Name),
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

func (mgr *Manager) DumpScopes(uri string) OpResult {
	var err error
	var fci *discover.FileCacheItem

	if fci, err = discover.EditFileCache().OpenActual(uri); err != nil {
		return OpFailure(fmt.Sprintf("unable to locate file %s", uri), err)
	}
	var sb strings.Builder
	fci.DumpScopes(&sb)

	return OpSuccessWith(DumpScopesResult{Dump: sb.String()})
}

// DumpDBScopes dumps the database structure for a given unit
func (m *Manager) DumpDBScopes(uri string) OpResult {
	// Extract unit name from URI
	pathElements := discover.DecodePath(uri)
	unitName := pathElements.Name()

	result, err := discover.SymDB().DumpDBScopes(unitName)
	if err != nil {
		return OpFailure("Failed to dump database scopes", err)
	}

	return OpSuccessWith(DumpScopesResult{Dump: result})
}

// ExecuteSQLQuery executes an arbitrary SQL query against the database
func (m *Manager) ExecuteSQLQuery(sqlQuery string) OpResult {
	result, err := discover.SymDB().ExecuteSQLQuery(sqlQuery)
	if err != nil {
		return OpFailure("Failed to execute SQL query", err)
	}

	return OpSuccessWith(DumpScopesResult{Dump: result})
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

func (mgr *Manager) HoverInfo(smb *discover.Symbol, pdata *discover.ParsedData) string {
	var result strings.Builder

	line, found, fileCtx := pdata.FindOriginalLine(smb.Position.Line)
	if !found {
		line = smb.Position.Line
	}

	result.WriteString("position: ")
	result.WriteString(fmt.Sprintf("%d:%d", line+1, smb.Position.Character+1))
	result.WriteString("\n")
	result.WriteString("kind: ")
	result.WriteString(discover.SymbolKindToString(discover.SymbolKind(smb.Kind)))
	result.WriteString("\n")
	result.WriteString("scope: ")
	result.WriteString(smb.Path)
	result.WriteString("\n")
	result.WriteString(smb.Name)
	if smb.Definition != "" {
		result.WriteString(": ")
		result.WriteString(smb.Definition)
	}
	if fileCtx != nil {
		result.WriteString("\n")
		result.WriteString(fileCtx.Filename)
	}
	return result.String()
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
