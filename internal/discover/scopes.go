// Package discover provides functionality to discover and navigate code structures
package discover

// Scope represents a code scope that can be searched for symbols
type Scope interface {
	getName() string
	print()
	findSymbol(position Position) *Symbol
}

// TopScope represents a top-level scope with public methods for interaction
type TopScope interface {
	Scope
	Print()
	FindSymbol(position Position) *Symbol
}

// Position represents a position in source code
type Position struct {
	Line      int
	Character int
}

// Symbol represents a code symbol with its metadata
type Symbol struct {
	Name       string
	Definition string
	Position   Position
	Kind       int
}

// commonScope implements basic scope functionality
type commonScope struct {
	name        string
	symbolStack stack[Symbol]
	scopeStack  stack[Scope]
	parentSWM   int
	position    Position
}

// commonScopeBuilder provides methods to construct a commonScope
type commonScopeBuilder struct {
	cmsc commonScope
}

// UnitScope represents a program unit scope (like a file or module)
type UnitScope struct {
	Scope
	usesStack stack[string]
}

// UnitScopeBuilder provides methods to construct a UnitScope
type UnitScopeBuilder struct {
	*commonScopeBuilder
	usesStack stack[string]
}

// func showScope(level int, scope *commonScope) {
// 	fmt.Printf("%-*sscope name -> %s\n", level*4, "-", scope.getName())
// 	for _, symbol := range scope.symbolStack.all() {
// 		fmt.Printf("%-*s %v %d\n", level*4, "", symbol, level)
// 		// fmt.Printf("%-*s%+v\n", level*2, "", symbol)

// 	}
// 	for _, sc := range scope.scopeStack.all() {
// 		showScope(level+1, &sc)
// 	}

// }

// func showUnitScope(scope *UnitScope) {
// 	for _, unit := range scope.usesStack.all() {
// 		fmt.Printf("Uses %s\n", unit)
// 	}
// 	showScope(0, &scope.commonScope)
// }

// getName returns the name of the scope
func (s *commonScope) getName() string {
	return s.name
}

// print outputs the scope hierarchy to standard output
func (s *commonScope) print() {
	println("Name: ", s.getName())
	println("symbols:")
	for _, symbol := range s.symbolStack.all() {
		println(symbol.Name)
	}
	println("scopes:")
	for _, scope := range s.scopeStack.all() {
		scope.print()
	}
}

// findSymbol locates a symbol at the given position within the scope
func (s *commonScope) findSymbol(position Position) *Symbol {
	for i := s.symbolStack.length() - 1; i >= 0; i-- {
		sym := s.symbolStack.get(i)
		if sym.Position.Line == position.Line &&
			position.Character >= sym.Position.Character &&
			position.Character < sym.Position.Character+len(sym.Name) {
			return &sym
		}
	}
	return nil
}

// FindSymbol implements TopScope interface to locate a symbol at the given position within the unit scope
func (s *UnitScope) FindSymbol(position Position) *Symbol {
	return s.Scope.findSymbol(position)
}

// Print outputs the unit scope to standard output, implementing TopScope interface
func (s *UnitScope) Print() {
	println("Name: ", s.getName())
	println("uses: ")
	for _, unit := range s.usesStack.all() {
		print(unit)
	}
	println("symbols:")
	s.Scope.print()
}

// addSymbol adds a symbol to the scope being built
func (b *commonScopeBuilder) addSymbol(name string, definition string, kind int, position Position) *commonScopeBuilder {
	smb := Symbol{Name: name, Definition: definition, Kind: kind, Position: position}
	b.cmsc.symbolStack.push(smb)
	return b
}

// addScope adds a child scope to the scope being built
func (b *commonScopeBuilder) addScope(sc Scope) *commonScopeBuilder {
	b.cmsc.scopeStack.push(sc)
	return b
}

// parentSWM sets the parent SWM identifier for the scope
func (b *commonScopeBuilder) parentSWM(swm int) *commonScopeBuilder {
	b.cmsc.parentSWM = swm
	return b
}

// setName sets the name of the scope being built
func (b *commonScopeBuilder) setName(name string) *commonScopeBuilder {
	b.cmsc.name = name
	return b
}

// getName gets the name of the scope being built
func (b *commonScopeBuilder) getName() string {
	return b.cmsc.name
}

// symbolStackLast returns the index of the last symbol in the stack
func (b *commonScopeBuilder) symbolStackLast() int {
	return b.cmsc.symbolStack.length() - 1
}

// finish completes the scope building process and returns the built Scope
func (b *commonScopeBuilder) finish() Scope {
	return &b.cmsc
}

// addUses adds a unit dependency to the UnitScope being built
func (b *UnitScopeBuilder) addUses(unit string) *UnitScopeBuilder {
	b.usesStack.push(unit)
	return b
}

func (b *UnitScopeBuilder) finish() TopScope {
	return &UnitScope{
		Scope:     b.commonScopeBuilder.finish(),
		usesStack: b.usesStack,
	}
}
