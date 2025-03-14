package discover

import (
	"log"

	"github.com/antlr4-go/antlr/v4"
)

type Discover struct{}

type DiscoverError struct {
	Message string
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
