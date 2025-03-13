package discover

import (
	"database/sql"
	"log"
	"os" // added to read files
	"strings"

	_ "github.com/glebarez/go-sqlite" // Registers the sqlite driver under "sqlite"
)

var db *symDB

type symDB struct {
	conn          *sql.DB
	searchFolders []string
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

	createPathsTableSQL := `
	CREATE TABLE IF NOT EXISTS paths (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		path TEXT UNIQUE
	);`
	_, err = db.conn.Exec(createPathsTableSQL)
	if err != nil {
		return err
	}

	return nil
}

func SymDB() *symDB {
	return db
}

// getFileModTime returns the Unix timestamp of the file's last modification time
func getFileModTime(filepath string) (int64, error) {
	fileInfo, err := os.Stat(filepath)
	if err != nil {
		return 0, err
	}
	return fileInfo.ModTime().Unix(), nil
}

func (db *symDB) insertUnit(unitname, unitpath string) (int, error) {
	// Get the file's modification time
	modTime, err := getFileModTime(unitpath)
	if err != nil {
		return 0, err
	}

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
		// Use COLLATE NOCASE for explicit case insensitive matching on unitname
		selectUnitIDSQL := `
		SELECT id FROM units
		WHERE unitname = ? COLLATE NOCASE AND unitpath = ?;`
		err = db.conn.QueryRow(selectUnitIDSQL, unitname, unitpath).Scan(&unitID)
		if err != nil {
			return 0, err
		}
	} else if err != nil {
		return 0, err
	}
	return unitID, nil
}

func (db *symDB) insertSymbol(unitID int, symbol, scope string, kind int, definition string) error {
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

// IsUnitLoaded checks if a unit with the given name exists in the database.
func (db *symDB) IsUnitLoaded(unit string) bool {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM units WHERE unitname = ? COLLATE NOCASE)"
	err := db.conn.QueryRow(query, unit).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}

func (db *symDB) DropSymbols(unitID int) error {
	_, err := db.conn.Exec("DELETE FROM symbols WHERE unit_id = ?", unitID)
	return err
}

func (db *symDB) FillSymbols(unitID int, unitpath string, unit string) error {
	// Get the file's current modification time
	modTime, err := getFileModTime(unitpath)
	if err != nil {
		return err
	}

	content, err := os.ReadFile(unitpath)
	if err != nil {
		return err
	}

	l := &publicSymbolsListener{unit_id: unitID, unitName: unit}
	parseFromContent(string(content), l, defaultOptions())

	// Mark this unit as scanned and update the last_modified timestamp
	_, err = db.conn.Exec("UPDATE units SET scanned = 1, last_modified = ? WHERE id = ?", modTime, unitID)
	return err
}

// SearchSymbolsWithinUnit searches for symbols within a specific unit that match the search term.
// It returns a slice of matching symbol information.
func (db *symDB) SearchSymbolsWithinUnit(unit, searchTerm string) ([]Symbol, error) {
	var unitID int
	var unitpath string
	var lastModified int64
	var scanned int

	query := "SELECT id, unitpath, last_modified, scanned FROM units WHERE unitname = ? COLLATE NOCASE"
	err := db.conn.QueryRow(query, unit).Scan(&unitID, &unitpath, &lastModified, &scanned)
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
		err = db.DropSymbols(unitID)
		if err != nil {
			return []Symbol{}, err
		}
		err = db.FillSymbols(unitID, unitpath, unit)
		if err != nil {
			return []Symbol{}, err
		}
	}

	// First attempt to fetch symbols
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

	rows, err := db.conn.Query(searchQuery, unitID, "%"+searchTerm+"%")
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

// PathExists checks if a path already exists in the paths table
func (db *symDB) PathExists(path string) bool {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM paths WHERE path = ?)"
	err := db.conn.QueryRow(query, path).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}

// AddPath inserts a path into the paths table if it doesn't already exist
func (db *symDB) AddPath(path string) error {
	insertPathSQL := `
	INSERT INTO paths (path)
	VALUES (?)
	ON CONFLICT(path) DO NOTHING;`
	_, err := db.conn.Exec(insertPathSQL, path)
	return err
}

// GetAllPaths returns all paths stored in the database
func (db *symDB) GetAllPaths() ([]string, error) {
	query := "SELECT path FROM paths"
	rows, err := db.conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var paths []string
	for rows.Next() {
		var path string
		if err := rows.Scan(&path); err != nil {
			return nil, err
		}
		paths = append(paths, path)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return paths, nil
}

// ClearPaths removes all paths from the paths table
func (db *symDB) ClearPaths() error {
	_, err := db.conn.Exec("DELETE FROM paths")
	return err
}

func (db *symDB) SetSearchFolders(folders []string) {
	db.searchFolders = folders
	db.RescanUnits()
}

func (db *symDB) RescanUnits() {
	d := &Discover{}
	for _, folder := range db.searchFolders {
		d.Units(folder)
	}
}

func (db *symDB) RescanPublicSymbols(unit string) {
	d := &Discover{}
	d.PublicSymbols(unit)
}
