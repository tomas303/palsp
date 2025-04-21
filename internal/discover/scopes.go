// Package discover provides functionality to discover and navigate code structures
package discover

import (
	"fmt"
	"palsp/internal/log"
	"regexp"
	"strings"
)

// Scope represents a code scope that can be searched for symbols
type Scope interface {
	getName() string
	getPosition() Position
	getParentSWM() int
	locateSimilarSymbols(name string, position Position, writer SymbolWriter) error
	print()
	findSymbol(position Position) *Symbol
}

// TopScope represents a top-level scope with public methods for interaction
type TopScope interface {
	Scope
	Print()
	FindSymbol(position Position) *Symbol
	LocateSimilarSymbols(name string, position Position, writer SymbolWriter) error
	IsInImplementation(position Position) bool
	InteraceUsese() []string
	ImplementationUses() []string
}

type SymbolWriter interface {
	WriteSymbol(sym *Symbol) error
}

// Position represents a position in source code
type Position struct {
	Line      int
	Character int
}

// NewPosition creates a new Position instance
func NewPosition(line, character int) Position {
	return Position{Line: line, Character: character}
}

// Compare compares this position with another position
// Returns:
//
//	-1 if this position is before other
//	 0 if positions are equal
//	+1 if this position is after other
func (p Position) Compare(other Position) int {
	if p.Line < other.Line {
		return -1
	}
	if p.Line > other.Line {
		return 1
	}
	// Lines are equal, compare characters
	if p.Character < other.Character {
		return -1
	}
	if p.Character > other.Character {
		return 1
	}
	return 0 // Positions are equal
}

func (p Position) Equal(other Position) bool {
	return p.Compare(other) == 0
}

func (p Position) Before(other Position) bool {
	return p.Compare(other) < 0
}

func (p Position) After(other Position) bool {
	return p.Compare(other) > 0
}

// Symbol represents a code symbol with its metadata
type Symbol struct {
	Name       string
	Definition string
	Position   Position
	Kind       int
	Scope      string
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
	interfaceUses      stack[string]
	implementationUses stack[string]
	implementationPos  Position
}

// UnitScopeBuilder provides methods to construct a UnitScope
type UnitScopeBuilder struct {
	*commonScopeBuilder
	interfaceUses      stack[string]
	implementationUses stack[string]
	implementationPos  Position
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

// getPosition returns the scoping position of the scope.
// It provides access to the location information within the source code.
func (s *commonScope) getPosition() Position {
	return s.position
}

// getParentSWM returns the parent SWM identifier for the scope
func (s *commonScope) getParentSWM() int {
	return s.parentSWM
}

// print outputs the scope hierarchy to standard output
func (s *commonScope) print() {
	log.Logger.Debug().Msgf("Scope: %s", s.getName())
	log.Logger.Debug().Msg("symbols:")
	for _, symbol := range s.symbolStack.all() {
		log.Logger.Debug().Msg(symbol.Name)
	}
	log.Logger.Debug().Msg("scopes:")
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

func (s *commonScope) locateSimilarSymbols(name string, position Position, writer SymbolWriter) error {
	// Find the most specific scope for the position
	var hitScope Scope
	for i := s.scopeStack.length() - 1; i >= 0; i-- {
		scope := s.scopeStack.get(i)
		if scope.getPosition().Line <= position.Line {
			hitScope = scope
			break
		}
	}

	pattern := "(?i)^" + name + "$"
	re, err := regexp.Compile(pattern)
	if err != nil {
		return fmt.Errorf("invalid search pattern: %w", err)
	}

	var watermark int
	// If we found a more specific scope, search there first
	if hitScope != nil {
		if err := hitScope.locateSimilarSymbols(name, position, writer); err != nil {
			return err
		}
		watermark = hitScope.getParentSWM()
	} else {
		watermark = s.symbolStack.length() - 1
	}

	// Search in the current scope from watermark up
	for i := watermark; i >= 0; i-- {
		sym := s.symbolStack.get(i)
		if re.MatchString(sym.Name) {
			if err := writer.WriteSymbol(&sym); err != nil {
				return err
			}
		}
	}

	return nil
}

// FindSymbol implements TopScope interface to locate a symbol at the given position within the unit scope
func (s *UnitScope) FindSymbol(position Position) *Symbol {
	return s.Scope.findSymbol(position)
}

// LocateSimilarSymbols search for symbol based on regex pattern given by name parameter (search is case
// insensitive and over all subject).
func (s *UnitScope) LocateSimilarSymbols(name string, position Position, writer SymbolWriter) error {
	return s.Scope.locateSimilarSymbols(name, position, writer)
}

// Print outputs the unit scope to standard output, implementing TopScope interface
func (s *UnitScope) Print() {
	log.Logger.Debug().Msgf("Name: %s", s.getName())
	log.Logger.Debug().Msg("uses: ")
	for _, unit := range s.interfaceUses.all() {
		log.Logger.Debug().Msg(unit)
	}
	for _, unit := range s.implementationUses.all() {
		log.Logger.Debug().Msg(unit)
	}
	log.Logger.Debug().Msg("symbols:")
	s.Scope.print()
}

// IsInImplementation checks if the position is within the implementation part of the unit
func (s *UnitScope) IsInImplementation(position Position) bool {
	return position.Line >= s.implementationPos.Line
}

// InteraceUsese returns the list of interface unit dependencies
func (s *UnitScope) InteraceUsese() []string {
	return s.interfaceUses.all()
}

// ImplementationUses returns the list of implementation unit dependencies
func (s *UnitScope) ImplementationUses() []string {
	return s.implementationUses.all()
}

// addSymbol adds a symbol to the scope being built
func (b *commonScopeBuilder) addSymbol(name string, definition string, kind int, position Position, scope string) *commonScopeBuilder {
	smb := Symbol{Name: strings.ToLower(name), Definition: definition, Kind: kind, Position: position, Scope: scope}
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

// addUses adds a interface unit dependency to the UnitScope being built
func (b *UnitScopeBuilder) addInterfaceUses(unit string) *UnitScopeBuilder {
	b.interfaceUses.push(unit)
	return b
}

// addUses adds a implementation unit dependency to the UnitScope being built
func (b *UnitScopeBuilder) addImplementationUses(unit string) *UnitScopeBuilder {
	b.implementationUses.push(unit)
	return b
}

func (b *UnitScopeBuilder) setImplementationPos(pos Position) *UnitScopeBuilder {
	b.implementationPos = pos
	return b
}

func (b *UnitScopeBuilder) finish() TopScope {
	return &UnitScope{
		Scope:              b.commonScopeBuilder.finish(),
		interfaceUses:      b.interfaceUses,
		implementationUses: b.implementationUses,
		implementationPos:  b.implementationPos,
	}
}

func (smb *Symbol) String() string {
	return smb.Name
}

func (smb *Symbol) HoverInfo() string {
	var result strings.Builder

	result.WriteString("kind: ")
	result.WriteString(SymbolKindToString(SymbolKind(smb.Kind)))
	result.WriteString("\n")
	result.WriteString("scope: ")
	result.WriteString(smb.Scope)
	result.WriteString("\n")
	result.WriteString(smb.Name)
	if smb.Definition != "" {
		result.WriteString(": ")
		result.WriteString(smb.Definition)
	}

	return result.String()
}
