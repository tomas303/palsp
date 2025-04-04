package discover

import (
	"database/sql"
	"log"
	"os" // added to read files
	"path/filepath"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	_ "github.com/glebarez/go-sqlite" // Registers the sqlite driver under "sqlite"
)

var db *symDB

type symDB struct {
	conn        *sql.DB
	searchPaths []string
}

// Define an interface for symDB operations
type SymbolDatabase interface {
	AddSearchPath(path string)
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
		log.Fatalf("Failed to initialize database: %v", err)
	}

	err = createTables(db)
	if err != nil {
		log.Fatalf("Failed to create tables: %v", err)
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
	row := db.conn.QueryRow(insertUnitSQL, unitname, unitpath, modTime, modTime)
	err = row.Scan(&unitID)
	if err == sql.ErrNoRows {
		selectUnitIDSQL := `
		SELECT id FROM units
		WHERE unitname = ? AND unitpath = ?;`
		err = db.conn.QueryRow(selectUnitIDSQL, unitname, unitpath).Scan(&unitID)
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
	_, err = db.conn.Exec(insertSymbolSQL, unitID, strings.ToLower(symbol), scope, kind, definition)
	return err
}

// GetUnitContent returns the unit id and locates the file path for the given unit name,
// reads the file (assumed UTF-8 encoded), and returns its id and content as a string.
func (db *symDB) GetUnitContent(unit string) (int, string, error) {
	var unitID int
	var unitpath string
	// Updated query to select both unit id and unitpath with case insensitive comparison
	query := "SELECT id, unitpath FROM units WHERE unitname = ? COLLATE NOCASE"
	if err := db.conn.QueryRow(query, unit).Scan(&unitID, &unitpath); err != nil {
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
	err = db.conn.QueryRow(query, unit).Scan(&unitID, &unitpath, &lastModified, &scanned)
	if err != nil {
		return nil, err
	}

	// Check the current file modification time
	currentModTime, err := getFileModTime(unitpath)
	if err != nil {
		return []Symbol{}, err
	}

	// Refresh symbols if file was modified or not yet scanned
	if currentModTime > lastModified || scanned == 0 {
		err = db.dropSymbols(unitID)
		if err != nil {
			return []Symbol{}, err
		}
		err = db.fillSymbols(unitID, unitpath, unit)
		if err != nil {
			return []Symbol{}, err
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

	// rows, err := db.conn.Query(searchQuery, unitID, "%"+searchTerm+"%")
	// rows, err := db.conn.Query(searchQuery, unitID, "%%")
	rows, err := db.conn.Query(searchQuery, unitID, searchTerm)
	if err != nil {
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
	_, err := db.conn.Exec("DELETE FROM symbols WHERE unit_id = ?", unitID)
	return err
}

func (db *symDB) fillSymbols(unitID int, unitpath string, unit string) error {
	// Get the file's current modification time
	modTime, err := getFileModTime(unitpath)
	if err != nil {
		return err
	}

	content, err := os.ReadFile(unitpath)
	if err != nil {
		return err
	}

	db.collectSymbols(unitID, string(content))

	// Mark this unit as scanned and update the last_modified timestamp
	_, err = db.conn.Exec("UPDATE units SET scanned = 1, last_modified = ? WHERE id = ?", modTime, unitID)
	return err
}

func (db *symDB) collectSymbols(unitID int, content string) {
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
	cst := ParseCST(content)
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

func (db *symDB) searchUnits(folder string) {
	fc := fileCrawler{}
	fc.processPasFiles(folder,
		func(path string) {
			filename := filepath.Base(path)
			ext := filepath.Ext(path)
			unitName := strings.TrimSuffix(filename, ext)
			println("Unit found:", unitName)
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
	db := &symDB{}
	var err error
	db.conn, err = sql.Open("sqlite", "file::memory:?cache=shared")
	return db, err
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
	_, err := db.conn.Exec(createUnitsTableSQL)
	if err != nil {
		return err
	}

	createUnitNameIndexSQL := `
	CREATE INDEX IF NOT EXISTS idx_unitname ON units (
		unitname COLLATE NOCASE
	);`
	_, err = db.conn.Exec(createUnitNameIndexSQL)
	if err != nil {
		return err
	}

	createUnitPathIndexSQL := `
	CREATE INDEX IF NOT EXISTS idx_unitpath ON units (
		unitpath COLLATE NOCASE
	);`
	_, err = db.conn.Exec(createUnitPathIndexSQL)
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
	_, err = db.conn.Exec(createSymbolsTableSQL)
	if err != nil {
		return err
	}

	createIndexSQL := `
	CREATE INDEX IF NOT EXISTS idx_unitname_scope ON symbols (
		unit_id,
		scope COLLATE NOCASE
	);`
	_, err = db.conn.Exec(createIndexSQL)
	if err != nil {
		return err
	}

	return nil
}
