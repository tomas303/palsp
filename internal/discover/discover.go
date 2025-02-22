package discover

import (
	"log"
	"path/filepath"
	"strings"
)

type Discover struct{}

type DiscoverError struct {
	Message string
}

func (d *Discover) Units(rootDir string) {
	fc := fileCrawler{}
	// fc.processPasListeners(rootDir,
	// 	func() antlr.ParseTreeListener {
	// 		return &unitNameListener{}
	// 	},
	// 	func(listener antlr.ParseTreeListener, path string) {
	// 		unitNameListener := listener.(*unitNameListener)
	// 		if unitNameListener.IsUnit() {
	// 			fmt.Println("Unit found:", unitNameListener.UnitName())
	// 			SymDB().insertUnit(unitNameListener.UnitName(), path)
	// 		}
	// 	})
	fc.processPasFiles(rootDir,
		func(path string) {
			filename := filepath.Base(path)
			ext := filepath.Ext(path)
			unitName := strings.TrimSuffix(filename, ext)
			println("Unit found:", unitName)
			SymDB().insertUnit(unitName, path)
		})

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

func (d *Discover) ScopeSymbols(unit string) *unitScope {

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
	l := newScopeListener(unit)
	parseFromContent(content, l, fullDebugOptions())
	return l.unitScope.(*unitScope)

}

func (e *DiscoverError) Error() string {
	return e.Message
}
