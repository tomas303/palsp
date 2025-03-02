package discover

import "fmt"

// Position represents a position in source code
type Position struct {
	Line      int
	Character int
}

type scope interface {
	addSymbol(name string, definition string, kind int, position Position)
	addScope(name string, position Position) scope
	symbolStackLast() int
	getName() string
	setName(name string)
	// findSymbol(name string) symbol
}

type topscope interface {
	scope
	addUses(unit string)
}

type symbol struct {
	name       string
	definition string
	kind       int
	position   Position
}

type commonScope struct {
	name        string
	symbolStack stack[symbol]
	scopeStack  stack[commonScope]
	parentSWM   int
	position    Position
}

type UnitScope struct {
	commonScope
	usesStack stack[string]
}

func showScope(level int, scope *commonScope) {
	fmt.Printf("%-*sscope name -> %s\n", level*4, "-", scope.getName())
	for _, symbol := range scope.symbolStack.all() {
		fmt.Printf("%-*s %v %d\n", level*4, "", symbol, level)
		// fmt.Printf("%-*s%+v\n", level*2, "", symbol)

	}
	for _, sc := range scope.scopeStack.all() {
		showScope(level+1, &sc)
	}

}

func showUnitScope(scope *UnitScope) {
	for _, unit := range scope.usesStack.all() {
		fmt.Printf("Uses %s\n", unit)
	}
	showScope(0, &scope.commonScope)
}

func newCommonScope(name string, parentSWM int, position Position) *commonScope {
	return &commonScope{
		name:        name,
		symbolStack: stack[symbol]{},
		scopeStack:  stack[commonScope]{},
		parentSWM:   parentSWM,
		position:    position,
	}
}

func newUnitScope(unit string) topscope {
	return &UnitScope{
		commonScope: *newCommonScope(unit, 0, Position{Line: 0, Character: 0}),
		usesStack:   stack[string]{},
	}
}

func (s *commonScope) addSymbol(name string, definition string, kind int, position Position) {
	s.symbolStack.push(symbol{name: name, definition: definition, kind: kind, position: position})
}

func (s *commonScope) addScope(name string, position Position) scope {
	new := newCommonScope(name, s.symbolStack.length()-1, position)
	s.scopeStack.push(*new)
	return new
}

func (s *commonScope) symbolStackLast() int {
	return s.symbolStack.length() - 1
}

func (s *UnitScope) addUses(unit string) {
	s.usesStack.push(unit)
}

func (s *commonScope) getName() string {
	return s.name
}

func (s *commonScope) setName(name string) {
	s.name = name
}
