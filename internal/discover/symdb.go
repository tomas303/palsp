package discover

import (
	"database/sql"
	"log"
	"os" // added to read files

	_ "github.com/glebarez/go-sqlite" // Registers the sqlite driver under "sqlite"
)

var db *symDB

type symDB struct {
	conn *sql.DB
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
		unitname TEXT,
		unitpath TEXT,
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

func SymDB() *symDB {
	return db
}

func (db *symDB) insertUnit(unitname, unitpath string) (int, error) {
	insertUnitSQL := `
	INSERT INTO units (unitname, unitpath)
	VALUES (?, ?)
	ON CONFLICT(unitname, unitpath) DO NOTHING
	RETURNING id;`
	var unitID int
	row := db.conn.QueryRow(insertUnitSQL, unitname, unitpath)
	err := row.Scan(&unitID)
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
	_, err = db.conn.Exec(insertSymbolSQL, unitID, symbol, scope, kind, definition)
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
