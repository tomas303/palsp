package discover

import (
	"fmt"
	"testing"
)

func TestDiscoverUnits(t *testing.T) {
	// Create a Discover instance.
	d := &Discover{}

	// Run Units, and recover from any panics caused by our listener.
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered panic:", r)
			}
		}()
		// Provide a valid directory (can be current directory) for testing.
		// d.Units("/home/tomas/testsrcs/short")
		d.Units("/home/tomas/testsrcs/source")
	}()

	// Query the units table.
	dbConn := SymDB().conn
	rows, err := dbConn.Query("SELECT id, unitname, unitpath FROM units")
	if err != nil {
		t.Fatalf("Query failed: %v", err)
	}
	defer rows.Close()

	fmt.Println("Units Table:")
	for rows.Next() {
		var id int
		var unitname, unitpath string
		if err := rows.Scan(&id, &unitname, &unitpath); err != nil {
			t.Fatalf("Scan failed: %v", err)
		}
		fmt.Printf("ID: %d, Unit: %s, Path: %s\n", id, unitname, unitpath)
	}
	if err := rows.Err(); err != nil {
		t.Fatalf("Row error: %v", err)
	}
}

func TestDiscoverPublicSymbols(t *testing.T) {
	// Create a Discover instance.

	// Run PublicSymbols, and recover from any panics caused by our listener.
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered panic:", r)
			}
		}()
		d := &Discover{}
		d.Units("/home/tomas/testsrcs/source")
		// d.PublicSymbols("iNALCrypt")
		{
			dbConn := SymDB().conn
			unitRows, err := dbConn.Query("SELECT unitname FROM units")
			if err != nil {
				t.Fatalf("Query for units failed: %v", err)
			}
			defer unitRows.Close()
			for unitRows.Next() {
				var unitname string
				if err := unitRows.Scan(&unitname); err != nil {
					t.Fatalf("Scan of unitname failed: %v", err)
				}
				d.PublicSymbols(unitname)
			}
		}

	}()

	// Query the symbols table with a left join on units to get unitname.
	dbConn := SymDB().conn
	rows, err := dbConn.Query("SELECT s.symbol, s.scope, s.kind, s.definition, u.unitname FROM symbols s LEFT JOIN units u ON s.unit_id = u.id")
	if err != nil {
		t.Fatalf("Query failed: %v", err)
	}
	defer rows.Close()

	fmt.Println("Symbols Table:")
	fmt.Printf("%-20s %-20s %-5s %-20s %-50s\n", "symbol", "unitname", "kind", "scope", "definition")
	totalRows := 0
	for rows.Next() {
		totalRows++
		var kind int
		var symbol, unitname, definition, scope string
		if err := rows.Scan(&symbol, &scope, &kind, &definition, &unitname); err != nil {
			t.Fatalf("Scan failed: %v", err)
		}
		fmt.Printf("%-20s %-20s %-5d %-20s %-50s\n", symbol, unitname, kind, scope, definition)
	}
	if err := rows.Err(); err != nil {
		t.Fatalf("Row error: %v", err)
	}
	fmt.Printf("Total rows: %d\n", totalRows)

}

func TestDiscoverScopeSymbols(t *testing.T) {
	// Create a Discover instance.

	// Run PublicSymbols, and recover from any panics caused by our listener.
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered panic:", r)
			}
		}()
		d := &Discover{}
		d.Units("/home/tomas/testsrcs/pubtest")
		// sc := d.ScopeSymbols("iNALCrypt")
		// {
		// 	dbConn := SymDB().conn
		// 	unitRows, err := dbConn.Query("SELECT unitname FROM units")
		// 	if err != nil {
		// 		t.Fatalf("Query for units failed: %v", err)
		// 	}
		// 	defer unitRows.Close()
		// 	for unitRows.Next() {
		// 		var unitname string
		// 		if err := unitRows.Scan(&unitname); err != nil {
		// 			t.Fatalf("Scan of unitname failed: %v", err)
		// 		}
		// 		d.ScopeSymbols(unitname)
		// 	}
		// }
		println("\nRESULTS:\n")
		// showUnitScope(sc)
	}()

}
