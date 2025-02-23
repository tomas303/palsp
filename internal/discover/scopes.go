package discover

import "fmt"

type scope interface {
	addSymbol(name string, definition string, kind int)
	addScope(name string) scope
	parent() scope
	symbolStackLast() int
	getName() string
	setName(name string)
}

type topscope interface {
	scope
	addUses(unit string)
}

type symbol struct {
	name       string
	definition string
	kind       int
}

type commonScope struct {
	name        string
	symbolStack stack[symbol]
	scopeStack  stack[scope]
	parentSWM   int
	parentScope scope
}

type UnitScope struct {
	scope
	usesStack stack[string]
}

func showScope(level int, scope *commonScope) {
	fmt.Printf("%-*sscope name -> %s\n", level*4, "-", scope.getName())
	for _, symbol := range scope.symbolStack.all() {
		fmt.Printf("%-*s %v %d\n", level*4, "", symbol, level)
		// fmt.Printf("%-*s%+v\n", level*2, "", symbol)

	}
	for _, sc := range scope.scopeStack.all() {
		showScope(level+1, sc.(*commonScope))
	}

}

func showUnitScope(scope *UnitScope) {
	for _, unit := range scope.usesStack.all() {
		fmt.Printf("Uses %s\n", unit)
	}
	showScope(0, scope.scope.(*commonScope))
}

func newCommonScope(name string, parent scope, parentSWM int) scope {
	return &commonScope{
		name:        name,
		symbolStack: stack[symbol]{},
		scopeStack:  stack[scope]{},
		parentSWM:   parentSWM,
		parentScope: parent,
	}
}

func newUnitScope(unit string) topscope {
	return &UnitScope{
		scope:     newCommonScope(unit, nil, 0),
		usesStack: stack[string]{},
	}
}

func (s *commonScope) addSymbol(name string, definition string, kind int) {
	s.symbolStack.push(symbol{name: name, definition: definition, kind: kind})
}

func (s *commonScope) addScope(name string) scope {
	new := newCommonScope(name, s, s.symbolStack.length()-1)
	s.scopeStack.push(new)
	return new
}

func (s *commonScope) parent() scope {
	return s.parentScope
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
