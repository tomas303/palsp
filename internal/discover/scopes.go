package discover

type Scope interface {
	getName() string
	print()
	findSymbol(position Position) *Symbol
}

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

type Symbol struct {
	Name       string
	Definition string
	Kind       int
	Position   Position
}

type commonScope struct {
	name        string
	symbolStack stack[Symbol]
	scopeStack  stack[Scope]
	parentSWM   int
	position    Position
}

type commonScopeBuilder struct {
	cmsc commonScope
}

type UnitScope struct {
	Scope
	usesStack stack[string]
}

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

func (s *commonScope) getName() string {
	return s.name
}

func (s *UnitScope) Print() {
	println("Name: ", s.getName())
	println("uses: ")
	for _, unit := range s.usesStack.all() {
		print(unit)
	}
	println("symbols:")
	s.Scope.print()
}

func (s *UnitScope) FindSymbol(position Position) *Symbol {
	return s.Scope.findSymbol(position)
}

func (b *commonScopeBuilder) addSymbol(name string, definition string, kind int, position Position) *commonScopeBuilder {
	smb := Symbol{Name: name, Definition: definition, Kind: kind, Position: position}
	b.cmsc.symbolStack.push(smb)
	return b
}

func (b *commonScopeBuilder) addScope(sc Scope) *commonScopeBuilder {
	b.cmsc.scopeStack.push(sc)
	return b
}

func (b *commonScopeBuilder) parentSWM(swm int) *commonScopeBuilder {
	b.cmsc.parentSWM = swm
	return b
}

func (b *commonScopeBuilder) finish() Scope {
	return &b.cmsc
}

func (b *commonScopeBuilder) setName(name string) *commonScopeBuilder {
	b.cmsc.name = name
	return b
}

func (b *commonScopeBuilder) getName() string {
	return b.cmsc.name
}

func (b *commonScopeBuilder) symbolStackLast() int {
	return b.cmsc.symbolStack.length() - 1
}

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
