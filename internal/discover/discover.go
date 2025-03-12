package discover

import (
	"log"
	"path/filepath"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

type Discover struct{}

type DiscoverError struct {
	Message string
}

func (d *Discover) Units(rootDir string) {
	if SymDB().PathExists(rootDir) {
		log.Printf("Path %s already scanned", rootDir)
		return
	}
	fc := fileCrawler{}
	fc.processPasFiles(rootDir,
		func(path string) {
			filename := filepath.Base(path)
			ext := filepath.Ext(path)
			unitName := strings.TrimSuffix(filename, ext)
			println("Unit found:", unitName)
			SymDB().insertUnit(unitName, path)
		})
	SymDB().AddPath(rootDir)
}

func (d *Discover) PublicSymbols(unit string) {

	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(*finishError); ok {
			} else {
				log.Printf("Error collection public symbols %s: %v", unit, r)
			}
		}
	}()

	unit_id, content, err := SymDB().GetUnitContent(unit)
	if err != nil {
		panic(DiscoverError{Message: err.Error()})
	}

	l := &publicSymbolsListener{unit_id: unit_id, unitName: unit}
	parseFromContent(content, l, defaultOptions())

}

func (d *Discover) ScopeSymbols(unit string) TopScope {

	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(*finishError); ok {
			} else {
				log.Printf("Error collection public symbols %s: %v", unit, r)
			}
		}
	}()

	_, content, err := SymDB().GetUnitContent(unit)
	if err != nil {
		panic(DiscoverError{Message: err.Error()})
	}

	// l := &scopeSymbolsListener{unit_id: unit_id, unitName: unit}
	// scopeparseFromContent(content, l, fullDebugOptions())
	l := NewScopeListener(unit)
	parseFromContent(content, l, fullDebugOptions())
	// return l.unitScope.(*UnitScope)
	return l.usb.finish()

}

func (d *Discover) CST(unit string) antlr.Tree {

	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(*finishError); ok {
			} else {
				log.Printf("Error collection ast %s: %v", unit, r)
			}
		}
	}()

	_, content, err := SymDB().GetUnitContent(unit)
	if err != nil {
		panic(DiscoverError{Message: err.Error()})
	}

	tree := ParseCST(content)
	return tree

}

func (e *DiscoverError) Error() string {
	return e.Message
}
