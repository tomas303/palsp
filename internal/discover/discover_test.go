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
