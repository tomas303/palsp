package discover

import (
	"context"
	"database/sql" // added for formatted strings
	"os"           // added to read files
	"palsp/internal/log"
	"path/filepath"
	"strings"
	"sync"

	"github.com/antlr4-go/antlr/v4"
	_ "github.com/glebarez/go-sqlite" // Registers the sqlite driver under "sqlite"
)

var db *symDB

type symDB struct {
	db             *sql.DB
	con            *sql.Conn
	searchPaths    []string
	unitScopeNames []string
}

type UnitString string
type UnitID int

// Define an interface for symDB operations
type SymbolDatabase interface {
	AddSearchPath(path string)
	SetUnitScopeNames(unitScopeNames []string)
	DropSymbolsFromPath(path string)
	GetUnitContent(unit string) (int, string, error)
	InsertSymbol(unitID int, symbol, scope string, kind int, definition string, position Position) error
	SearchSymbol(unit, searchTerm string) ([]Symbol, error)
	SearchSymbolByKind(unit string, kind int) ([]Symbol, error)
	RetriveUnit(unit string) (int, string, error)
	GetUnitPath(unit string) (string, error)
	LocateSymbolsInScope(name string, unit string, scope string, writer SymbolWriter) error
}

// SymbolKind represents the kind of public symbol as an integer.
type SymbolKind int

const (
	ProcedureSymbol SymbolKind = iota // 0
	FunctionSymbol                    // 1
	ConstantSymbol                    // 2
	VariableSymbol                    // 3
	ClassSymbol                       // 4
	TypeSymbol                        // 5
	ParameterSymbol                   // 6
	FunctionResult                    // 7
	ClassVariable                     // 8
	UnitReference                     // 9
	TypeIdentifier                    // 10
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
	case TypeSymbol:
		return "type"
	case ParameterSymbol:
		return "parameter"
	case FunctionResult:
		return "result"
	case ClassVariable:
		return "field"
	case UnitReference:
		return "unit"
	case TypeIdentifier:
		return "type ident"
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

// RetriveUnit retrieves the unit ID for a given unit name. In case of changes it refreshes the unit content.
func (db *symDB) RetriveUnit(unit string) (int, string, error) {
	unitID, unitpath, lastModified, scanned, unitname, err := db.findUnitInfo(unit)
	if err != nil {
		return 0, "", err
	}

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
		if err == sql.ErrNoRows {
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
		if err == sql.ErrNoRows {
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

	// If we get here, we couldn't find the unit
	log.Main.Warn().Err(err).Msgf("Unit %s not found, even with scope prefixes", unit)
	return 0, "", 0, 0, "", err
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
	sl := NewScopesListener(collector)
	cst, _ := ParseCST(content, fileName)
	antlr.ParseTreeWalkerDefault.Walk(sl, cst)
	db.writeToLog(strings.ToLower(DecodePath(fileName).name))
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

func (db *symDB) SetUnitScopeNames(unitScopeNames []string) {
	for _, name := range unitScopeNames {
		db.unitScopeNames = append(db.unitScopeNames, strings.ToLower(name))
	}
}

func (db *symDB) DropSymbolsFromPath(path string) {
	// Extract filename from path (could be in URI format)
	path = filepath.ToSlash(path) // Convert to forward slashes for consistency
	fileName := filepath.Base(path)
	ext := filepath.Ext(fileName)
	unitName := strings.TrimSuffix(fileName, ext)
	// Get all unit IDs matching this name
	query := "SELECT id FROM units WHERE unitname = ? COLLATE NOCASE"
	rows, err := db.Query(query, unitName)
	if err != nil {
		log.Main.Warn().Err(err).Msgf("DropSymbolsFromPath error: couldn't query units for %s", unitName)
		return
	}
	defer rows.Close()

	// Delete symbols for each matching unit
	for rows.Next() {
		var unitID int
		if err := rows.Scan(&unitID); err != nil {
			log.Main.Warn().Err(err).Msg("Error scanning unit ID")
			continue
		}

		if err := db.dropSymbols(unitID); err != nil {
			log.Main.Warn().Err(err).Msgf("Error dropping symbols for unit ID %d", unitID)
		}
		// Reset the scanned flag
		_, err = db.Exec("UPDATE units SET scanned = 0 WHERE id = ?", unitID)
		if err != nil {
			log.Main.Warn().Err(err).Msgf("Error resetting scanned flag for unit ID %d", unitID)
		}
	}
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

// Mutex to synchronize access to the writeToLog function
var writeToLogMutex sync.Mutex

func (db *symDB) writeToLog(unitName string) {
	// Acquire lock to ensure only one goroutine can execute this method at a time
	writeToLogMutex.Lock()
	defer writeToLogMutex.Unlock()

	if !log.Structure.Debug().Enabled() {
		return
	}

	// Get the unit file path from the database
	var unitPath string
	pathQuery := `
	SELECT unitpath 
	FROM units 
	WHERE unitname = ? COLLATE NOCASE 
	LIMIT 1`

	err := db.QueryRow(pathQuery, unitName).Scan(&unitPath)
	if err != nil {
		log.Main.Warn().Err(err).Msgf("Error retrieving unit path for logging: %s", unitName)
		// Continue execution even if we couldn't get the path
	}

	// Query to get all symbols for the unit, grouped by path
	query := `
	SELECT s.symbol, s.scope, s.kind
	FROM symbols s
	JOIN units u ON s.unit_id = u.id
	WHERE u.unitname = ? COLLATE NOCASE
	ORDER BY s.scope, s.symbol COLLATE NOCASE`

	rows, err := db.Query(query, unitName)
	if err != nil {
		log.Main.Warn().Err(err).Msgf("Error retrieving symbols for logging: %s", unitName)
		return
	}
	defer rows.Close()

	// Track the current path to create appropriate indentation
	currentPath := ""

	// Log the header with unit name and file path
	if unitPath != "" {
		log.Structure.Debug().Msgf("Symbol structure for unit: %s (%s)", unitName, unitPath)
	} else {
		log.Structure.Debug().Msgf("Symbol structure for unit: %s", unitName)
	}

	for rows.Next() {
		var name, path string
		var kind int

		if err := rows.Scan(&name, &path, &kind); err != nil {
			log.Main.Warn().Err(err).Msg("Error scanning symbol row for logging")
			continue
		}

		// If the path changed, print the new path
		if path != currentPath {
			currentPath = path
			// Calculate indentation based on path nesting
			indent := strings.Repeat("  ", strings.Count(path, "."))
			if path == "" {
				log.Structure.Debug().Msgf("└── Unit level symbols:")
			} else {
				log.Structure.Debug().Msgf("%s├── %s:", indent, path)
			}
		}

		// Print the symbol with indentation
		kindStr := SymbolKindToString(SymbolKind(kind))
		indent := strings.Repeat("  ", strings.Count(path, ".")+1)
		log.Structure.Debug().Msgf("   %s├── %s (%s)", indent, name, kindStr)
	}

	if err = rows.Err(); err != nil {
		log.Main.Warn().Err(err).Msg("Error iterating symbol rows for logging")
	}
}

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
	symdb := &symDB{db: db, con: con, searchPaths: []string{}, unitScopeNames: []string{}}
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
