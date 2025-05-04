// Package discover provides functionality to discover and navigate code structures
package discover

import (
	"context"
	"fmt"
	"palsp/internal/log"
	"regexp"
	"runtime"
	"strings"
	"sync"
)

// Scope represents a code scope that can be searched for symbols
type Scope interface {
	getName() string
	getPosition() Position
	getParentSWM() int
	setParentSWM(swm int)
	locateSimilarSymbols(name string, position Position, writer SymbolWriter) error
	writeToLog(prefix string)
	findSymbol(position Position) *Symbol
}

// TopScope represents a top-level scope with public methods for interaction
type TopScope interface {
	Scope
	WriteToLog()
	FindSymbol(position Position) *Symbol
	LocateSimilarSymbols(name string, position Position, writer SymbolWriter) error
	IsInImplementation(position Position) bool
	InterfaceUses() []string
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
	scopeStack  stack[*commonScope]
	parentSWM   int
	scopeInfo   ScopeInfo
}

// UnitScope represents a program unit scope (like a file or module)
type UnitScope struct {
	Scope
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
	return s.scopeInfo.Position
}

// getParentSWM returns the parent SWM identifier for the scope
func (s *commonScope) getParentSWM() int {
	return s.parentSWM
}

func (s *commonScope) setParentSWM(swm int) {
	s.parentSWM = swm
}

// print outputs the scope hierarchy to standard output
func (s *commonScope) writeToLog(prefix string) {
	log.Logger.Debug().Msgf(prefix+"scope name: %s", s.getName())
	log.Logger.Debug().Msg(prefix + "--symbols:")
	for _, symbol := range s.symbolStack.all() {
		log.Logger.Debug().Msg(prefix + "----" + symbol.Name)
	}
	log.Logger.Debug().Msg(prefix + "--scopes:")
	for _, scope := range s.scopeStack.all() {
		scope.writeToLog(prefix + "----")
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

func (s *commonScope) findScopeByName(name string) *commonScope {
	for i := 0; i < s.scopeStack.length(); i++ {
		scope := s.scopeStack.get(i)
		if name == scope.getName() {
			return scope
		}
	}
	return nil
}

func (s *commonScope) locateInClass(name string, prefixes []string, methodname string, writer SymbolWriter) error {
	if len(prefixes) == 0 {
		methodScope := s.findScopeByName(methodname)
		if methodScope != nil {
			if err := methodScope.locateSymbolByName(name, methodScope.symbolStack.length()-1, writer); err != nil {
				return err
			}
		}
	} else {
		partScope := s.findScopeByName(prefixes[0])
		if partScope != nil {
			if err := partScope.locateInClass(name, prefixes[1:], methodname, writer); err != nil {
				return err
			}
			if len(prefixes) == 1 {
				// todo - this is class where method is declared, for going up scope some helper to name should be passed like to what context belongs
				if err := partScope.locateSymbolByName(name, partScope.symbolStack.length()-1, writer); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (s *commonScope) locateSymbolByName(name string, watermark int, writer SymbolWriter) error {

	pattern := "(?i)^" + name + "$"
	re, err := regexp.Compile(pattern)
	if err != nil {
		return fmt.Errorf("invalid search pattern: %w", err)
	}

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

func (s *commonScope) locateSimilarSymbols(name string, position Position, writer SymbolWriter) error {
	// Find the most specific scope for the position
	var hitScope Scope
	for i := s.scopeStack.length() - 1; i >= 0; i-- {
		scope := s.scopeStack.get(i)
		scopeLine := scope.getPosition().Line
		if scopeLine > 0 && scopeLine <= position.Line {
			hitScope = scope
			break
		}
	}

	var watermark int
	// If we found a more specific scope, search there first
	if hitScope != nil {
		if err := hitScope.locateSimilarSymbols(name, position, writer); err != nil {
			return err
		}
		// if inside a method scope, look inside class or record
		prefixes, methodName := SplitQualifiedName(hitScope.getName())
		if len(prefixes) > 0 {
			if err := s.locateInClass(name, prefixes, methodName, writer); err != nil {
				return err
			}
		}
		watermark = hitScope.getParentSWM()

	} else {
		watermark = s.symbolStack.length() - 1
	}

	// continue up the current symbols(in pascal symbol must be declared before usage)
	if err := s.locateSymbolByName(name, watermark, writer); err != nil {
		return err
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
	if err := s.Scope.locateSimilarSymbols(name, position, writer); err != nil {
		return err
	}
	if s.IsInImplementation(position) {
		if err := s.searchSymbolInUnits(name, s.ImplementationUses(), writer); err != nil {
			return err
		}
	}
	if err := s.searchSymbolInUnits(name, s.InterfaceUses(), writer); err != nil {
		return err
	}
	return nil
}

func (s *UnitScope) searchSymbolInUnits(name string, units []string, writer SymbolWriter) error {

	if len(units) == 0 {
		return nil
	}

	// Limit concurrency to number of CPU cores
	maxWorkers := runtime.NumCPU()

	type searchResult struct {
		unit    string
		symbols []Symbol
		err     error
	}

	// Context for cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create channel for results with buffer
	resultCh := make(chan searchResult, len(units))

	// Process units channel for worker scheduling
	unitsCh := make(chan string, len(units))
	for _, unit := range units {
		unitsCh <- unit
	}
	close(unitsCh)

	// Launch only maxWorkers goroutines
	var wg sync.WaitGroup
	for i := 0; i < maxWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			var unitName string
			var ok bool
			for {
				// Get next unit to process
				select {
				case unitName, ok = <-unitsCh:
					if !ok {
						// No more units to process
						return
					}
				case <-ctx.Done():
					// Work cancelled
					return
				}
				// Process the unit
				symbols, err := SymDB().SearchSymbol(unitName, name)
				resultCh <- searchResult{unit: unitName, symbols: symbols, err: err}
			}
		}()
	}

	// Close result channel when all workers are done
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	results := make(map[string]searchResult, len(units))
	idx := 0
	for result := range resultCh {
		results[result.unit] = result
		for {
			rs, ok := results[units[idx]]
			if !ok {
				break
			}
			if rs.err != nil {
				return rs.err
			}
			for _, symbol := range rs.symbols {
				err := writer.WriteSymbol(&symbol)
				if err != nil {
					return err
				}
			}
			idx++
			if idx >= len(units) {
				break
			}
		}
	}

	return nil

}

func (s *UnitScope) WriteToLog() {
	log.Logger.Debug().Msgf("Top scope: %s", s.getName())
	log.Logger.Debug().Msg("uses: ")
	for _, unit := range s.interfaceUses.all() {
		log.Logger.Debug().Msg(unit)
	}
	for _, unit := range s.implementationUses.all() {
		log.Logger.Debug().Msg(unit)
	}
	s.Scope.writeToLog("  ")
}

// IsInImplementation checks if the position is within the implementation part of the unit
func (s *UnitScope) IsInImplementation(position Position) bool {
	return position.Line >= s.implementationPos.Line
}

// InterfaceUses returns the list of interface unit dependencies
func (s *UnitScope) InterfaceUses() []string {
	return s.interfaceUses.all()
}

// ImplementationUses returns the list of implementation unit dependencies
func (s *UnitScope) ImplementationUses() []string {
	return s.implementationUses.all()
}

func (smb *Symbol) String() string {
	return smb.Name
}

func (smb *Symbol) HoverInfo() string {
	var result strings.Builder

	result.WriteString("position: ")
	result.WriteString(fmt.Sprintf("%d:%d", smb.Position.Line, smb.Position.Character))
	result.WriteString("\n")
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
