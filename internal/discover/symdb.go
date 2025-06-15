package discover

import (
	"context"
	"database/sql" // added for formatted strings
	"errors"
	"fmt"
	"os" // added to read files
	"palsp/internal/log"
	"path/filepath"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	_ "github.com/glebarez/go-sqlite" // Registers the sqlite driver under "sqlite"
)

var db *symDB
var ErrUnitNotFound = errors.New("unit not found")

// todo : later move some commone ctx fields to separate global context
type symDB struct {
	db               *sql.DB
	con              *sql.Conn
	searchPaths      []string
	unitScopeNames   []string
	defines          []string // Add defines support
	retrieveUnitLock *KeyLock[int]
}

type UnitString string
type UnitID int

// Define an interface for symDB operations
type SymbolDatabase interface {
	AddSearchPath(path string)
	GetSearchPaths() []string
	SetUnitScopeNames(unitScopeNames []string)
	SetDefines(defines []string) // Add SetDefines method
	GetDefines() []string
	GetUnitContent(unit string) (int, string, error)
	InsertSymbol(unitID int, symbol, scope string, kind int, definition string, position Position) error
	SearchSymbol(unit, searchTerm string) ([]Symbol, error)
	SearchSymbolByKind(unit string, kind int) ([]Symbol, error)
	RetriveUnit(unit string) (int, string, error)
	GetUnitPath(unit string) (string, error)
	LocateSymbolsInScope(name string, unit string, scope string, writer SymbolWriter) error
	UnscannedUnits() []string
	DumpDBScopes(unitName string) (string, error)
	ExecuteSQLQuery(sqlQuery string) (string, error)
}

// SymbolKind represents the kind of public symbol as an integer.
type SymbolKind int

const (
	UnknownSymbol        SymbolKind = iota // 0
	ProcedureSymbol                        // 1
	FunctionSymbol                         // 2
	ConstantSymbol                         // 3
	VariableSymbol                         // 4
	ClassSymbol                            // 5
	InterfaceSymbol                        // 6
	TypeSymbol                             // 7
	ParameterSymbol                        // 8
	FunctionResult                         // 9
	ClassVariable                          // 10
	UnitReference                          // 11
	TypeIdentifier                         // 12
	PropertySymbol                         // 13
	HelperSymbol                           // 14
	ResourceStringSymbol                   // 15
	Unit                                   // 16
)

func init() {
	var err error
	db, err = newSymDB()
	if err != nil {
		log.Main.Fatal().Err(err).Msg("Failed to initialize database")
	}

	err = createTables(db)
	if err != nil {
		log.Main.Fatal().Err(err).Msg("Failed to create tables")
	}
}

// SymbolKindToString converts a SymbolKind constant to its string representation
func SymbolKindToString(kind SymbolKind) string {
	switch kind {
	case ProcedureSymbol:
		return "procedure"
	case FunctionSymbol:
		return "function"
	case ConstantSymbol:
		return "constant"
	case VariableSymbol:
		return "variable"
	case ClassSymbol:
		return "class"
	case InterfaceSymbol:
		return "interface"
	case TypeSymbol:
		return "type"
	case ParameterSymbol:
		return "parameter"
	case FunctionResult:
		return "result"
	case ClassVariable:
		return "field"
	case UnitReference:
		return "use unit"
	case TypeIdentifier:
		return "type ident"
	case PropertySymbol:
		return "property"
	case HelperSymbol:
		return "class helper"
	case ResourceStringSymbol:
		return "resourcestring"
	case Unit:
		return "unit"
	default:
		return "unknown"
	}
}

func SymDB() SymbolDatabase {
	return db
}

func (db *symDB) Exec(query string, args ...any) (sql.Result, error) {
	// log.Logger.Debug().Msgf("START exec: %s", query)
	result, err := db.con.ExecContext(context.Background(), query, args...)
	// log.Logger.Debug().Msgf("STOP exec: %s", query)
	return result, err
}

func (db *symDB) QueryRow(query string, args ...any) *sql.Row {
	// log.Logger.Debug().Msgf("START queryrow: %s", query)
	row := db.con.QueryRowContext(context.Background(), query, args...)
	// log.Logger.Debug().Msgf("STOP queryrow: %s", query)
	return row
}

func (db *symDB) Query(query string, args ...any) (*sql.Rows, error) {
	// log.Logger.Debug().Msgf("START query: %s", query)
	result, err := db.con.QueryContext(context.Background(), query, args...)
	// log.Logger.Debug().Msgf("STOP query %s", query)
	return result, err
}

func (db *symDB) insertUnit(unitname, unitpath string) (int, error) {
	// Get the file's modification time
	modTime, err := getFileModTime(unitpath)
	if err != nil {
		return 0, err
	}

	unitname = strings.ToLower(unitname)

	insertUnitSQL := `
	INSERT INTO units (unitname, unitpath, last_modified, scanned)
	VALUES (?, ?, ?, 0)
	ON CONFLICT(unitname, unitpath) DO UPDATE SET
		last_modified = ?
	RETURNING id;`

	var unitID int
	row := db.QueryRow(insertUnitSQL, unitname, unitpath, modTime, modTime)
	err = row.Scan(&unitID)
	if err == sql.ErrNoRows {
		selectUnitIDSQL := `
		SELECT id FROM units
		WHERE unitname = ? AND unitpath = ?;`
		err = db.QueryRow(selectUnitIDSQL, unitname, unitpath).Scan(&unitID)
		if err != nil {
			return 0, err
		}
	} else if err != nil {
		return 0, err
	}
	return unitID, nil
}

func (db *symDB) InsertSymbol(unitID int, symbol, scope string, kind int, definition string, position Position) error {
	insertSymbolSQL := `
	INSERT INTO symbols (unit_id, symbol, scope, kind, definition, line, column)
	VALUES (?, ?, ?, ?, ?, ?, ?);`
	var err error
	_, err = db.Exec(insertSymbolSQL, unitID, strings.ToLower(symbol), scope, kind, definition, position.Line, position.Character)
	return err
}

// GetUnitContent returns the unit id and locates the file path for the given unit name,
// reads the file (assumed UTF-8 encoded), and returns its id and content as a string.
func (db *symDB) GetUnitContent(unit string) (int, string, error) {
	var unitID int
	var unitpath string
	// Updated query to select both unit id and unitpath with case insensitive comparison
	query := "SELECT id, unitpath FROM units WHERE unitname = ? COLLATE NOCASE"
	if err := db.QueryRow(query, unit).Scan(&unitID, &unitpath); err != nil {
		return 0, "", err
	}
	data, err := os.ReadFile(unitpath)
	if err != nil {
		return 0, "", err
	}
	return unitID, string(data), nil
}

func (db *symDB) UnscannedUnits() []string {
	var units []string
	query := "SELECT unitname FROM units WHERE scanned = 0"
	rows, err := db.Query(query)
	if err != nil {
		log.Main.Warn().Err(err).Msg("UnscannedUnits error")
		return units
	}
	defer rows.Close()
	for rows.Next() {
		var unitname string
		if err := rows.Scan(&unitname); err != nil {
			log.Main.Warn().Err(err).Msg("UnscannedUnits error scanning")
			continue
		}
		units = append(units, unitname)
	}
	if err := rows.Err(); err != nil {
		log.Main.Warn().Err(err).Msg("UnscannedUnits error iterating")
	}
	if len(units) == 0 {
		log.Main.Debug().Msg("No unscanned units found")
	} else {
		log.Main.Debug().Msgf("Unscanned units: %v", units)
	}
	return units
}

// RetriveUnit retrieves the unit ID for a given unit name. In case of changes it refreshes the unit content.
func (db *symDB) RetriveUnit(unit string) (int, string, error) {

	unitID, unitpath, lastModified, scanned, unitname, err := db.findUnitInfo(unit)
	if err != nil {
		return 0, "", err
	}

	db.retrieveUnitLock.Lock(unitID)
	defer db.retrieveUnitLock.Unlock(unitID)

	// Check the current file modification time
	currentModTime, err := getFileModTime(unitpath)
	if err != nil {
		log.Main.Warn().Err(err).Msgf("SearchSymbol error path %s errored obtaining file time", unitpath)
		return 0, unitname, err
	}

	// Refresh symbols if file was modified or not yet scanned
	if currentModTime > lastModified || scanned == 0 {
		err = db.dropSymbols(unitID)
		if err != nil {
			log.Main.Warn().Err(err).Msg("dropping symbols error")
			return 0, unitname, err
		}
		err = db.fillSymbols(unitID, unitpath)
		if err != nil {
			log.Main.Warn().Err(err).Msg("filling symbols error")
			return 0, unitname, err
		}
	}
	return unitID, unitname, nil
}

func (db *symDB) GetUnitPath(unit string) (string, error) {
	_, unitpath, _, _, _, err := db.findUnitInfo(unit)
	if err != nil {
		return "", err
	}
	return unitpath, nil
}

// SearchSymbol searches for symbols within a specific unit that match the search term.
// It returns a slice of matching symbol information.
func (db *symDB) SearchSymbol(unit, searchTerm string) ([]Symbol, error) {

	unitID, unitname, err := db.RetriveUnit(unit)
	if err != nil {
		if err == ErrUnitNotFound {
			log.Main.Warn().Err(err).Msgf("SearchSymbol error path %s not found", unit)
			return []Symbol{}, nil
		}
		return nil, err
	}

	searchQuery := `
	SELECT symbol, scope, kind, definition, line, column 
	FROM symbols 
	WHERE unit_id = ? AND symbol LIKE ? COLLATE NOCASE
	ORDER BY symbol COLLATE NOCASE`

	rows, err := db.Query(searchQuery, unitID, searchTerm)
	if err != nil {
		if err == sql.ErrNoRows {
			return []Symbol{}, nil
		}
		return nil, err
	}
	defer rows.Close()

	// fetch symbols
	results, err := db.fetchSymbolsFromRows(rows, unitname)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (db *symDB) SearchSymbolByKind(unit string, kind int) ([]Symbol, error) {
	unitID, unitname, err := db.RetriveUnit(unit)
	if err != nil {
		return nil, err
	}

	searchQuery := `
	SELECT symbol, scope, kind, definition, line, column
	FROM symbols
	WHERE unit_id = ? AND kind = ?
	ORDER BY symbol COLLATE NOCASE`
	rows, err := db.Query(searchQuery, unitID, kind)
	if err != nil {
		if err == sql.ErrNoRows {
			return []Symbol{}, nil
		}
		return nil, err
	}
	defer rows.Close()
	// fetch symbols
	results, err := db.fetchSymbolsFromRows(rows, unitname)
	if err != nil {
		return nil, err
	}
	return results, nil

}

func (db *symDB) LocateSymbolsInScope(name string, unit string, scope string, writer SymbolWriter) error {

	unitID, _, err := db.RetriveUnit(unit)
	if err != nil {
		if err == ErrUnitNotFound {
			return nil
		}
		return err
	}

	// if unit == "uPackageBusinessObjects" {
	// 	q := `
	// 	SELECT s.symbol, s.scope, s.kind, s.definition, s.line, s.column
	// 	FROM symbols s
	// 	WHERE s.unit_id = ? AND s.symbol LIKE ? COLLATE NOCASE
	// 	ORDER BY s.symbol COLLATE NOCASE`
	// 	rows, err := db.Query(q, unitID, name)
	// 	if err != nil {
	// 		if err == sql.ErrNoRows {
	// 			return nil
	// 		}
	// 		return err
	// 	}
	// 	defer rows.Close()
	// 	for rows.Next() {
	// 		sym := Symbol{Unitname: unit}
	// 		if err := rows.Scan(&sym.Name, &sym.Path, &sym.Kind, &sym.Definition, &sym.Position.Line, &sym.Position.Character); err != nil {
	// 			return err
	// 		}
	// 		log.Main.Debug().Msgf("Found symbol: %s", sym.Name)
	// 	}

	// }

	// Prepare the query to search for symbols in a specific scope
	// query := `
	// SELECT s.symbol, s.scope, s.kind, s.definition, s.line, s.column
	// FROM symbols s
	// // JOIN units u ON s.unit_id = u.id
	// // WHERE s.scope = ? AND u.unitname = ? AND s.symbol LIKE ? COLLATE NOCASE
	// WHERE unit_id = ? and s.scope = ? AND s.symbol LIKE ? COLLATE NOCASE
	// ORDER BY s.symbol COLLATE NOCASE`
	query := `
	SELECT s.symbol, s.scope, s.kind, s.definition, s.line, s.column
	FROM symbols s
	WHERE s.unit_id = ? and s.scope = ? AND s.symbol LIKE ? COLLATE NOCASE
	ORDER BY s.symbol COLLATE NOCASE`

	sym := Symbol{Unitname: unit}
	// rows, err := db.Query(query, scope, unit, name)
	rows, err := db.Query(query, unitID, scope, name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&sym.Name, &sym.Path, &sym.Kind, &sym.Definition, &sym.Position.Line, &sym.Position.Character); err != nil {
			return err
		}
		if err = writer.WriteSymbol(&sym); err != nil {
			return err
		}
	}

	if err = rows.Err(); err != nil {
		return err
	}

	return nil
}

// findUnitInfo looks up unit information by name, trying with scope prefixes if the direct lookup fails
func (db *symDB) findUnitInfo(unit string) (unitID int, unitpath string, lastModified int64, scanned int, unitname string, err error) {
	unit = strings.ToLower(unit)

	// First try direct lookup
	query := "SELECT id, unitpath, last_modified, scanned FROM units WHERE unitname = ? COLLATE NOCASE"
	err = db.QueryRow(query, unit).Scan(&unitID, &unitpath, &lastModified, &scanned)
	if err == nil {
		return unitID, unitpath, lastModified, scanned, unit, nil
	}

	// If not found, try with scope prefixes
	for _, scope := range db.unitScopeNames {
		scopedUnit := scope + "." + unit
		log.Main.Debug().Str("original", unit).Str("scoped", scopedUnit).Msg("Trying with scope prefix")

		err = db.QueryRow(query, scopedUnit).Scan(&unitID, &unitpath, &lastModified, &scanned)
		if err == nil {
			return unitID, unitpath, lastModified, scanned, scopedUnit, nil
		}
	}

	if err == sql.ErrNoRows {
		log.Main.Warn().Err(err).Msgf("Unit %s not found in database, even with scope prefixes", unit)
		return 0, "", 0, 0, "", ErrUnitNotFound
	} else {
		log.Main.Warn().Err(err).Msgf("Unit %s not found, even with scope prefixes", unit)
		return 0, "", 0, 0, "", err
	}
}

func (db *symDB) fetchSymbolsFromRows(rows *sql.Rows, unitname string) ([]Symbol, error) {

	var results []Symbol
	for rows.Next() {
		sym := Symbol{Unitname: unitname}
		if err := rows.Scan(&sym.Name, &sym.Path, &sym.Kind, &sym.Definition, &sym.Position.Line, &sym.Position.Character); err != nil {
			return nil, err
		}

		results = append(results, sym)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (db *symDB) dropSymbols(unitID int) error {
	_, err := db.Exec("DELETE FROM symbols WHERE unit_id = ?", unitID)
	return err
}

func (db *symDB) fillSymbols(unitID int, unitpath string) error {
	// Get the file's current modification time
	modTime, err := getFileModTime(unitpath)
	if err != nil {
		return err
	}

	content, err := os.ReadFile(unitpath)
	if err != nil {
		return err
	}

	db.collectSymbols(unitID, string(content), unitpath)

	// Mark this unit as scanned and update the last_modified timestamp
	_, err = db.Exec("UPDATE units SET scanned = 1, last_modified = ? WHERE id = ?", modTime, unitID)
	return err
}

func (db *symDB) collectSymbols(unitID int, content string, fileName string) {
	defer func() {
		if r := recover(); r != nil {
			if r == ErrListenerBreak {
				return
			}
			panic(r) // Re-panic for all other errors
		}
	}()
	collector := NewDBSymbolCollector(unitID, db)
	pdata := ParseCST(content, fileName, true)
	sl := NewScopesListener(collector, pdata)
	antlr.ParseTreeWalkerDefault.Walk(sl, pdata.Tree)
}

func (db *symDB) AddSearchPath(path string) {
	for _, existingPath := range db.searchPaths {
		if existingPath == path {
			return
		}
	}
	db.searchPaths = append(db.searchPaths, path)
	db.searchUnits(path)
}
func (db *symDB) GetSearchPaths() []string {
	return db.searchPaths
}

func (db *symDB) SetUnitScopeNames(unitScopeNames []string) {
	for _, name := range unitScopeNames {
		db.unitScopeNames = append(db.unitScopeNames, strings.ToLower(name))
	}
}

func (db *symDB) SetDefines(defines []string) {
	db.defines = make([]string, len(defines))
	for i, define := range defines {
		db.defines[i] = strings.ToUpper(define) // Normalize to uppercase
	}
	log.Main.Info().Msgf("Set %d compiler defines in symDB", len(defines))
}

func (db *symDB) GetDefines() []string {
	return db.defines
}

func (db *symDB) searchUnits(folder string) {
	fc := fileCrawler{}
	fc.processPasFiles(folder,
		func(path string) {
			filename := filepath.Base(path)
			ext := filepath.Ext(path)
			unitName := strings.TrimSuffix(filename, ext)
			db.insertUnit(unitName, path)
		})
}

// todo: rewrite to custom command
// Mutex to synchronize access to the writeToLog function
// var writeToLogMutex sync.Mutex

// func (db *symDB) writeToLog(unitName string) {
// 	// Acquire lock to ensure only one goroutine can execute this method at a time
// 	writeToLogMutex.Lock()
// 	defer writeToLogMutex.Unlock()

// 	if !log.Structure.Debug().Enabled() {
// 		return
// 	}

// 	// Get the unit file path from the database
// 	var unitPath string
// 	pathQuery := `
// 	SELECT unitpath
// 	FROM units
// 	WHERE unitname = ? COLLATE NOCASE
// 	LIMIT 1`

// 	err := db.QueryRow(pathQuery, unitName).Scan(&unitPath)
// 	if err != nil {
// 		log.Main.Warn().Err(err).Msgf("Error retrieving unit path for logging: %s", unitName)
// 		// Continue execution even if we couldn't get the path
// 	}

// 	// Query to get all symbols for the unit, grouped by path
// 	query := `
// 	SELECT s.symbol, s.scope, s.kind
// 	FROM symbols s
// 	JOIN units u ON s.unit_id = u.id
// 	WHERE u.unitname = ? COLLATE NOCASE
// 	ORDER BY s.scope, s.symbol COLLATE NOCASE`

// 	rows, err := db.Query(query, unitName)
// 	if err != nil {
// 		log.Main.Warn().Err(err).Msgf("Error retrieving symbols for logging: %s", unitName)
// 		return
// 	}
// 	defer rows.Close()

// 	// Track the current path to create appropriate indentation
// 	currentPath := ""

// 	// Log the header with unit name and file path
// 	if unitPath != "" {
// 		log.Structure.Debug().Msgf("Symbol structure for unit: %s (%s)", unitName, unitPath)
// 	} else {
// 		log.Structure.Debug().Msgf("Symbol structure for unit: %s", unitName)
// 	}

// 	for rows.Next() {
// 		var name, path string
// 		var kind int

// 		if err := rows.Scan(&name, &path, &kind); err != nil {
// 			log.Main.Warn().Err(err).Msg("Error scanning symbol row for logging")
// 			continue
// 		}

// 		// If the path changed, print the new path
// 		if path != currentPath {
// 			currentPath = path
// 			// Calculate indentation based on path nesting
// 			indent := strings.Repeat("  ", strings.Count(path, "."))
// 			if path == "" {
// 				log.Structure.Debug().Msgf("└── Unit level symbols:")
// 			} else {
// 				log.Structure.Debug().Msgf("%s├── %s:", indent, path)
// 			}
// 		}

// 		// Print the symbol with indentation
// 		kindStr := SymbolKindToString(SymbolKind(kind))
// 		indent := strings.Repeat("  ", strings.Count(path, ".")+1)
// 		log.Structure.Debug().Msgf("   %s├── %s (%s)", indent, name, kindStr)
// 	}

// 	if err = rows.Err(); err != nil {
// 		log.Main.Warn().Err(err).Msg("Error iterating symbol rows for logging")
// 	}
// }

func newSymDB() (*symDB, error) {
	var err error
	var db *sql.DB
	var con *sql.Conn
	db, err = sql.Open("sqlite", "file::memory:?cache=shared")
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(1)
	con, err = db.Conn(context.Background())
	if err != nil {
		return nil, err
	}
	symdb := &symDB{db: db, con: con, searchPaths: []string{}, unitScopeNames: []string{}, retrieveUnitLock: NewKeyLock[int]()}
	return symdb, err
}

func createTables(db *symDB) error {
	createUnitsTableSQL := `
	CREATE TABLE IF NOT EXISTS units (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		unitname TEXT NOT NULL,
		unitpath TEXT NOT NULL,
		last_modified INTEGER NOT NULL DEFAULT (strftime('%s', 'now')),
		scanned INTEGER NOT NULL DEFAULT 0,
		UNIQUE(unitname, unitpath)
	);`
	_, err := db.Exec(createUnitsTableSQL)
	if err != nil {
		return err
	}

	createUnitNameIndexSQL := `
	CREATE INDEX IF NOT EXISTS idx_unitname ON units (
		unitname COLLATE NOCASE
	);`
	_, err = db.Exec(createUnitNameIndexSQL)
	if err != nil {
		return err
	}

	createUnitPathIndexSQL := `
	CREATE INDEX IF NOT EXISTS idx_unitpath ON units (
		unitpath COLLATE NOCASE
	);`
	_, err = db.Exec(createUnitPathIndexSQL)
	if err != nil {
		return err
	}

	createSymbolsTableSQL := `
	CREATE TABLE IF NOT EXISTS symbols (
		unit_id INTEGER,
		symbol TEXT,
		scope TEXT,
		kind INTEGER,
		definition TEXT,
		line INTEGER,
		column INTEGER,
		FOREIGN KEY(unit_id) REFERENCES units(id)
	);`
	_, err = db.Exec(createSymbolsTableSQL)
	if err != nil {
		return err
	}

	createIndexSQL := `
	CREATE INDEX IF NOT EXISTS idx_unitname_scope ON symbols (
		unit_id,
		scope COLLATE NOCASE
	);`
	_, err = db.Exec(createIndexSQL)
	if err != nil {
		return err
	}

	return nil
}

// DumpDBScopes dumps the database structure for a given unit in a tree format
func (db *symDB) DumpDBScopes(unitName string) (string, error) {
	var sb strings.Builder

	// Use RetriveUnit to ensure the unit is parsed and up-to-date
	unitID, actualUnitName, err := db.RetriveUnit(unitName)
	if err != nil {
		if err == ErrUnitNotFound {
			sb.WriteString(fmt.Sprintf("Unit '%s' not found in database\n", unitName))
			return sb.String(), nil
		}
		return "", err
	}

	// Get the unit path for display
	unitpath, err := db.GetUnitPath(actualUnitName)
	if err != nil {
		return "", err
	}

	// Write header with unit info
	sb.WriteString(fmt.Sprintf("Database Structure for Unit: %s\n", actualUnitName))
	sb.WriteString(fmt.Sprintf("File: %s\n", unitpath))
	sb.WriteString(fmt.Sprintf("Unit ID: %d\n", unitID))
	sb.WriteString("\nScope Structure:\n")

	// Query to get all symbols for the unit, ordered by scope path and symbol name
	query := `
	SELECT s.symbol, s.scope, s.kind, s.definition, s.line, s.column
	FROM symbols s
	WHERE s.unit_id = ?
	ORDER BY 
		CASE WHEN s.scope = '' THEN 0 ELSE 1 END,
		LENGTH(s.scope) - LENGTH(REPLACE(s.scope, '.', '')) ASC,
		s.scope COLLATE NOCASE,
		s.symbol COLLATE NOCASE`

	rows, err := db.Query(query, unitID)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	// Track scopes we've already printed headers for
	printedScopes := make(map[string]bool)
	currentScope := ""

	for rows.Next() {
		var symbol, scope, definition string
		var kind, line, column int

		if err := rows.Scan(&symbol, &scope, &kind, &definition, &line, &column); err != nil {
			return "", err
		}

		// If we're in a new scope, print the scope header
		if scope != currentScope {
			currentScope = scope

			if scope == "" {
				// Unit-level symbols
				if !printedScopes[""] {
					sb.WriteString("├─ Unit Level\n")
					printedScopes[""] = true
				}
			} else {
				// Nested scope - print the full scope path hierarchy
				db.printScopeHierarchy(&sb, scope, printedScopes)
			}
		}

		// Print the symbol with appropriate indentation (without definition)
		indent := db.getScopeIndent(scope)
		kindStr := SymbolKindToString(SymbolKind(kind))
		sb.WriteString(fmt.Sprintf("%s│    • %s (%s) (pos: %d:%d)\n",
			indent, symbol, kindStr, line, column))
	}

	if err = rows.Err(); err != nil {
		return "", err
	}

	return sb.String(), nil
}

// printScopeHierarchy prints the scope hierarchy for a given scope path
func (db *symDB) printScopeHierarchy(sb *strings.Builder, scopePath string, printedScopes map[string]bool) {
	if scopePath == "" {
		return
	}

	parts := strings.Split(scopePath, ".")
	currentPath := ""

	for i, part := range parts {
		if i > 0 {
			currentPath += "."
		}
		currentPath += part

		if !printedScopes[currentPath] {
			indent := strings.Repeat("│    ", i)
			sb.WriteString(fmt.Sprintf("%s├─ %s\n", indent, part))
			printedScopes[currentPath] = true
		}
	}
}

// getScopeIndent returns the appropriate indentation for a given scope
func (db *symDB) getScopeIndent(scope string) string {
	if scope == "" {
		return ""
	}
	depth := strings.Count(scope, ".") + 1
	return strings.Repeat("│    ", depth)
}

// ExecuteSQLQuery executes an arbitrary SQL query and returns the result as formatted text
func (db *symDB) ExecuteSQLQuery(sqlQuery string) (string, error) {
	var sb strings.Builder

	// Trim whitespace and check for empty query
	sqlQuery = strings.TrimSpace(sqlQuery)
	if sqlQuery == "" {
		return "Error: Empty SQL query", nil
	}

	// Execute the query
	rows, err := db.Query(sqlQuery)
	if err != nil {
		return fmt.Sprintf("Error executing query: %v", err), nil
	}
	defer rows.Close()

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		return fmt.Sprintf("Error getting columns: %v", err), nil
	}

	// Write header
	sb.WriteString("Query Result:\n")
	sb.WriteString(strings.Repeat("=", 50) + "\n")

	// Write column headers
	for i, col := range columns {
		if i > 0 {
			sb.WriteString(" | ")
		}
		sb.WriteString(fmt.Sprintf("%-15s", col))
	}
	sb.WriteString("\n")
	sb.WriteString(strings.Repeat("-", 15*len(columns)+3*(len(columns)-1)) + "\n")

	// Read and format rows
	rowCount := 0
	for rows.Next() {
		// Create slice to hold column values
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))
		for i := range values {
			valuePtrs[i] = &values[i]
		}

		// Scan the row
		if err := rows.Scan(valuePtrs...); err != nil {
			return fmt.Sprintf("Error scanning row: %v", err), nil
		}

		// Format and write the row
		for i, val := range values {
			if i > 0 {
				sb.WriteString(" | ")
			}

			// Convert value to string, handling NULL values
			var strVal string
			if val == nil {
				strVal = "NULL"
			} else {
				strVal = fmt.Sprintf("%v", val)
			}

			// Truncate if too long for display
			if len(strVal) > 15 {
				strVal = strVal[:12] + "..."
			}

			sb.WriteString(fmt.Sprintf("%-15s", strVal))
		}
		sb.WriteString("\n")
		rowCount++
	}

	if err = rows.Err(); err != nil {
		return fmt.Sprintf("Error iterating rows: %v", err), nil
	}

	// Write footer with row count
	sb.WriteString(strings.Repeat("=", 50) + "\n")
	sb.WriteString(fmt.Sprintf("Rows returned: %d\n", rowCount))

	return sb.String(), nil
}
