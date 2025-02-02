package discover

import (
	"database/sql"
	"log"

	_ "github.com/glebarez/go-sqlite" // Registers the sqlite driver under "sqlite"
)

var db *symDB

type symDB struct {
	conn *sql.DB
}

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
	db.conn, err = sql.Open("sqlite", ":memory:")
	return db, err
}

func createTables(db *symDB) error {
	createUnitsTableSQL := `
	CREATE TABLE IF NOT EXISTS units (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		unitname TEXT,
		unitpath TEXT
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

func (db *symDB) insertSymbol(unitname, unitpath, symbol, scope string, kind int) error {
	unitID, err := db.insertUnit(unitname, unitpath)
	if err != nil {
		return err
	}

	insertSymbolSQL := `
	INSERT INTO symbols (unit_id, symbol, scope, kind)
	VALUES (?, ?, ?, ?, ?);`
	_, err = db.conn.Exec(insertSymbolSQL, unitID, symbol, scope, kind)
	return err
}
