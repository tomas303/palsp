package discover

import (
	"context"
	"database/sql"
	"os" // added to read files
	"palsp/internal/log"
	"path/filepath"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	_ "github.com/glebarez/go-sqlite" // Registers the sqlite driver under "sqlite"
)

var db *symDB

type symDB struct {
	db          *sql.DB
	con         *sql.Conn
	searchPaths []string
}

// Define an interface for symDB operations
type SymbolDatabase interface {
	AddSearchPath(path string)
	DropSymbolsFromPath(path string)
	GetUnitContent(unit string) (int, string, error)
	InsertSymbol(unitID int, symbol, scope string, kind int, definition string) error
	SearchSymbol(unit, searchTerm string) ([]Symbol, error)
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
		log.Logger.Fatal().Err(err).Msg("Failed to initialize database")
	}

	err = createTables(db)
	if err != nil {
		log.Logger.Fatal().Err(err).Msg("Failed to create tables")
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

func (db *symDB) InsertSymbol(unitID int, symbol, scope string, kind int, definition string) error {
	insertSymbolSQL := `
	INSERT INTO symbols (unit_id, symbol, scope, kind, definition)
	VALUES (?, ?, ?, ?, ?);`
	var err error
	_, err = db.Exec(insertSymbolSQL, unitID, strings.ToLower(symbol), scope, kind, definition)
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

// SearchSymbol searches for symbols within a specific unit that match the search term.
// It returns a slice of matching symbol information.
func (db *symDB) SearchSymbol(unit, searchTerm string) ([]Symbol, error) {
	var unitID int
	var unitpath string
	var lastModified int64
	var scanned int

	unit = strings.ToLower(unit)

	query := "SELECT id, unitpath, last_modified, scanned FROM units WHERE unitname = ?"
	var err error
	err = db.QueryRow(query, unit).Scan(&unitID, &unitpath, &lastModified, &scanned)
	if err != nil {
		log.Logger.Warn().Err(err).Msgf("SearchSymbol error unit %s not found", unit)
		return []Symbol{}, nil
	}

	// Check the current file modification time
	currentModTime, err := getFileModTime(unitpath)
	if err != nil {
		log.Logger.Warn().Err(err).Msgf("SearchSymbol error path %s errored obtaining file time", unitpath)
		return []Symbol{}, nil
	}

	// Refresh symbols if file was modified or not yet scanned
	if currentModTime > lastModified || scanned == 0 {
		err = db.dropSymbols(unitID)
		if err != nil {
			log.Logger.Warn().Err(err).Msg("dropping symbols error")
			return []Symbol{}, nil
		}
		err = db.fillSymbols(unitID, unitpath)
		if err != nil {
			log.Logger.Warn().Err(err).Msg("filling symbols error")
			return []Symbol{}, nil
		}
	}

	// fetch symbols
	results, err := db.fetchSymbolsFromUnit(unitID, searchTerm)
	if err != nil {
		return nil, err
	}

	return results, nil
}

// fetchSymbolsFromUnit queries the database for symbols matching the search term in a specific unit
func (db *symDB) fetchSymbolsFromUnit(unitID int, searchTerm string) ([]Symbol, error) {
	searchQuery := `
	SELECT symbol, scope, kind, definition 
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

	var results []Symbol
	for rows.Next() {
		var sym Symbol
		if err := rows.Scan(&sym.Name, &sym.Scope, &sym.Kind, &sym.Definition); err != nil {
			return nil, err
		}

		results = append(results, sym)
	}

	if err = rows.Err(); err != nil {
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
	sl := NewUnifiedListener(collector)
	cst := ParseCST(content, fileName)
	antlr.ParseTreeWalkerDefault.Walk(sl, cst)
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
		log.Logger.Warn().Err(err).Msgf("DropSymbolsFromPath error: couldn't query units for %s", unitName)
		return
	}
	defer rows.Close()

	// Delete symbols for each matching unit
	for rows.Next() {
		var unitID int
		if err := rows.Scan(&unitID); err != nil {
			log.Logger.Warn().Err(err).Msg("Error scanning unit ID")
			continue
		}

		if err := db.dropSymbols(unitID); err != nil {
			log.Logger.Warn().Err(err).Msgf("Error dropping symbols for unit ID %d", unitID)
		}
		// Reset the scanned flag
		_, err = db.Exec("UPDATE units SET scanned = 0 WHERE id = ?", unitID)
		if err != nil {
			log.Logger.Warn().Err(err).Msgf("Error resetting scanned flag for unit ID %d", unitID)
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

func getFileModTime(filepath string) (int64, error) {
	fileInfo, err := os.Stat(filepath)
	if err != nil {
		return 0, err
	}
	return fileInfo.ModTime().Unix(), nil
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
	symdb := &symDB{db: db, con: con}
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
