// Package discover provides functionality to discover and navigate code structures
package discover

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"runtime"
	"strings"
	"sync"
)

// SymbolWriter is used to collect symbols from different searches.
type SymbolWriter interface {
	WriteSymbol(sym *Symbol) error
}

// SymbolWriterFunc type is a helper type used for making anonymous
// function usable as a SymbolWriter interface.
type SymbolWriterFunc func(sym *Symbol) error

func (f SymbolWriterFunc) WriteSymbol(sym *Symbol) error {
	return f(sym)
}

// Error commonly used to bail out searching on first symbol located.
var ErrFirstSymbolWriten = errors.New("first symbol writen")

// Position represents a position in source code(line and character
// starts from zero).
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

// TopScope represents a top-level scope and an entry to subscopes to be
// searched for.
// Methods starting with Find* are used to just find symbols in scopes.
// Methods starting with Locate* follow common pattern to find symbol
// in scopes and write it to the writer. Locate methods bail out on
// whichever error - this can be used by writer implementation to stop
// locating symbols early.
type TopScope interface {
	scope
	Dump(sb *strings.Builder)
	// FindSymbolOnPosition(position Position) *Symbol
	LocateSymbolsOnPos(name string, position Position, writer SymbolWriter) error
	LocateSymbols(name string, scope string, writer SymbolWriter) error
	GetUnits() []string
}

// scope represents a scope in an hierarchy of scopes that can be searched
// for symbols.
//
// getParentSWM() - returns so called scope watermark - this is the position
// of the last symbol in parent scope when this scope was created.
// This means that in this scope symbols from parent scopes can be used but
// only those that are declared before this scope - so from the watermark up
// (that is so in pascal).
//
// findScopeGlobally(name string) - when type supports inheritance, this method
// will try to find scope of its ancestor. This is used during symbol location
// when stepping outside of methods and symbols from such scope are procceed too.
type scope interface {
	getName() string
	getParentSWM() int
	locateSymbolsOnPos(name string, position Position, writer SymbolWriter) error
	dump(sb *strings.Builder, prefix string)
	findSymbolOnPosition(position Position) *Symbol
	findAncestorScope(ancestorName string) (inheritenceScope, error)
	locateSymbols(name string, scope string, writer SymbolWriter) error
}

// helper scope to be used to crawl up the scope by inheritance hierarchy
// and locate symbols there.
type inheritenceScope interface {
	getAncestorScope() (inheritenceScope, error)
	locateSymbolsViaInheritance(name string, writer SymbolWriter) error
}

// Symbol represents a metadata of a code symbol
type Symbol struct {
	Name       string
	Definition string
	Position   Position
	Kind       int
	Path       string
	Unitname   string
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
	result.WriteString(smb.Path)
	result.WriteString("\n")
	result.WriteString(smb.Name)
	if smb.Definition != "" {
		result.WriteString(": ")
		result.WriteString(smb.Definition)
	}

	return result.String()
}

func (smb *Symbol) String() string {
	return smb.Name
}

// unitScope represents a unit scope - a top scope of this unit
// scope is a root scope of an unit and an entry point to the subscopes.
type unitScope struct {
	scope
	interfaceUses      stack[string]
	implementationUses stack[string]
	implementationPos  Position
}

func (s *unitScope) Dump(sb *strings.Builder) {
	sb.WriteString(fmt.Sprintf("Top scope: %s\n", s.getName()))
	sb.WriteString("uses:\n")
	for _, unit := range s.interfaceUses.all() {
		sb.WriteString(fmt.Sprintf("  %s\n", unit))
	}
	for _, unit := range s.implementationUses.all() {
		sb.WriteString(fmt.Sprintf("  %s\n", unit))
	}
	s.scope.dump(sb, "  ")
}

func (s *unitScope) FindSymbolOnPosition(position Position) *Symbol {
	return s.scope.findSymbolOnPosition(position)
}

// todo - actually this looks globally without scope, but later manager must this way find only begining of dot expression and later via LocateSymbolsInScope
// todo - this is special case without scope - is searched based on scope shadowing and uses units. With scope situation is a bit different
// todo - so rename it appropriately ... this basically find info about symbol without scope on a position, some starting point
func (s *unitScope) LocateSymbolsOnPos(name string, position Position, writer SymbolWriter) error {
	if err := s.scope.locateSymbolsOnPos(name, position, writer); err != nil {
		return err
	}
	if s.isInImplementation(position) {
		// if err := s.locateSymbolsByNameInDB(name, s.getImplementationUses(), writer); err != nil {
		// 	return err
		// }
		if err := s.locateSymbolsInUnits(name, s.getImplementationUses(), writer); err != nil {
			return err
		}
	}
	// if err := s.locateSymbolsByNameInDB(name, s.getInterfaceUses(), writer); err != nil {
	// 	return err
	// }
	if err := s.locateSymbolsInUnits(name, s.getInterfaceUses(), writer); err != nil {
		return err
	}
	return nil
}

func (s *unitScope) LocateSymbols(name string, scope string, writer SymbolWriter) error {
	topScopeName := s.getName()
	if scope == topScopeName {
		scope = ""
	} else if strings.HasPrefix(scope, topScopeName+".") {
		scope = scope[len(topScopeName)+1:]
	}
	if err := s.scope.locateSymbols(name, scope, writer); err != nil {
		return err
	}
	return nil
}

func (s *unitScope) locateSymbolsInUnits(name string, units []string, writer SymbolWriter) error {
	var err error
	for _, unit := range units {
		// todo - add here suport for prefixed units. Actaully expand unit based on what is in database(unit or firs combination with scope units)
		if fci := EditFileCache().FindByUnit(unit); fci != nil {
			if fci.active {
				err = fci.scope.LocateSymbols(name, strings.ToLower(unit), writer)
				if err != nil {
					return err
				}
			}
		}
		err = SymDB().LocateSymbolsInScope(name, unit, strings.ToLower(unit), writer)
		if err != nil {
			return err
		}
	}
	return nil
}

// todo - generalize idea - this is parallel mainly because of units dont't need to be parsed yet (but parallel processing can be usefull on more places)
func (s *unitScope) locateSymbolsByNameInDB(name string, units []string, writer SymbolWriter) error {

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
				var symbols []Symbol
				var err error

				// Use defer/recover to catch any panic
				func() {
					defer func() {
						if r := recover(); r != nil {
							// Convert panic to error
							switch x := r.(type) {
							case string:
								err = errors.New("panic: " + x)
							case error:
								err = fmt.Errorf("panic: %w", x)
							default:
								err = fmt.Errorf("panic: %v", r)
							}
						}
					}()
					symbols, err = SymDB().SearchSymbol(unitName, name)
				}()

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

func (s *unitScope) isInImplementation(position Position) bool {
	return position.Line >= s.implementationPos.Line
}

func (s *unitScope) getInterfaceUses() []string {
	return s.interfaceUses.all()
}

func (s *unitScope) getImplementationUses() []string {
	return s.implementationUses.all()
}

func (s *unitScope) findAncestorScope(name string) (inheritenceScope, error) {
	var nameSym *Symbol
	writeSymbol := func(sym *Symbol) error {
		nameSym = sym
		return ErrFirstSymbolWriten
	}
	writer := SymbolWriterFunc(writeSymbol)
	err := s.locateSymbolsByNameInDB(name, s.getInterfaceUses(), writer)
	if err != ErrFirstSymbolWriten {
		return nil, err
	}

	result := NewdbInheritenceScope(nameSym.Unitname, nameSym.Name, nameSym.Definition)
	return result, nil
}

func (s *unitScope) GetUnits() []string {
	var units []string
	units = append(units, s.implementationUses.reverse()...)
	units = append(units, s.interfaceUses.reverse()...)
	return units
}

// commonScope represents one scope inside other scopes(functions, structured types, etc.)
//
// name - name of the scope(e.g. function name)
// symbolStack - stack of symbols in this scope
// scopeStack - stack of scopes in this scope
// parentscope - parent scope of this scope(inside which this scope is declared)
// parentSWM - watermark of the parent scope - this is the position of the last symbol
// in parent scope when this scope was created(usefull when locating symbols in this scope).
// scopeInfo - various help information about the scope.
type commonScope struct {
	name        string
	symbolStack stack[Symbol]
	scopeStack  stack[*commonScope]
	parentScope scope
	parentSWM   int
	info        scopeInfo
}

func (s *commonScope) getName() string {
	return s.name
}

func (s *commonScope) getPosition() Position {
	return s.info.startPos
}

func (s *commonScope) getParentSWM() int {
	return s.parentSWM
}

func (s *commonScope) setParentSWM(swm int) {
	s.parentSWM = swm
}

func (s *commonScope) dump(sb *strings.Builder, prefix string) {
	sb.WriteString(fmt.Sprintf("%sscope name: %s\n", prefix, s.getName()))
	sb.WriteString(prefix + "--symbols:\n")
	for _, symbol := range s.symbolStack.all() {
		sb.WriteString(prefix + "----" + symbol.Name + "\n")
	}
	sb.WriteString(prefix + "--scopes:\n")
	for _, scope := range s.scopeStack.all() {
		scope.dump(sb, prefix+"----")
	}
}

func (s *commonScope) findSymbolOnPosition(position Position) *Symbol {
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

func (s *commonScope) getAncestorScope() (inheritenceScope, error) {
	if s.info.ancestor != nil {
		return s.findAncestorScope(*s.info.ancestor)
	}
	return nil, nil
}

func (s *commonScope) locateSymbolsViaInheritance(name string, writer SymbolWriter) error {
	if err := s.locateSymbolByName(name, writer); err != nil {
		return err
	}
	if ancestor, err := s.getAncestorScope(); err != nil {
		return err
	} else if ancestor != nil {
		return ancestor.locateSymbolsViaInheritance(name, writer)
	}
	return nil
}

func (s *commonScope) findLocalScope(name string) *commonScope {
	for i := 0; i < s.scopeStack.length(); i++ {
		scope := s.scopeStack.get(i)
		if name == scope.getName() {
			return scope
		}
	}
	return nil
}

func (s *commonScope) findAncestorScope(name string) (inheritenceScope, error) {

	// check if the name is in the current scope
	if scope := s.findLocalScope(name); scope != nil {
		return scope, nil
	}

	// continue searching in the parent scope
	if s.parentScope != nil {
		return s.parentScope.findAncestorScope(name)
	} else {
		return nil, fmt.Errorf("scope %s not found", name)
	}
}

func (s *commonScope) locateInClass(name string, prefixes []string, functionName string, writer SymbolWriter) error {
	if len(prefixes) == 0 {
		// parameters can be declared in header only
		functionScope := s.findLocalScope(functionName)
		if functionScope != nil {
			if err := functionScope.locateSymbolByName(name, writer); err != nil {
				return err
			}
		}
	} else {
		partScope := s.findLocalScope(prefixes[0])
		if partScope != nil {
			if err := partScope.locateInClass(name, prefixes[1:], functionName, writer); err != nil {
				return err
			}
			if len(prefixes) == 1 {
				// todo - this is class where method is declared, for going up scope some helper to name should be passed like to what context belongs
				if err := partScope.locateSymbolsViaInheritance(name, writer); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (s *commonScope) locateSymbolByName(name string, writer SymbolWriter) error {
	return s.locateSymbolByNameFrom(name, s.symbolStack.length()-1, writer)
}

func (s *commonScope) locateSymbolByNameFrom(name string, watermark int, writer SymbolWriter) error {

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

func (s *commonScope) locateSymbolsOnPos(name string, position Position, writer SymbolWriter) error {
	// Find the most specific scope for the position
	var hitScope scope
	for i := s.scopeStack.length() - 1; i >= 0; i-- {
		scope := s.scopeStack.get(i)
		if scope.info.startPos.Compare(position) <= 0 && scope.info.stopPos.Compare(position) >= 0 {
			hitScope = scope
			break
		}

	}

	var watermark int
	// If we found a more specific scope, search there first
	if hitScope != nil {
		if err := hitScope.locateSymbolsOnPos(name, position, writer); err != nil {
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
	if err := s.locateSymbolByNameFrom(name, watermark, writer); err != nil {
		return err
	}
	return nil
}

func (s *commonScope) locateSymbols(name string, scope string, writer SymbolWriter) error {
	if scope == "" {
		return s.locateSymbolByName(name, writer)
	}
	var lscope *commonScope
	dotIndex := strings.Index(scope, ".")
	if dotIndex == -1 {
		lscope = s.findLocalScope(scope)
	} else {
		lscope = s.findLocalScope(scope[:dotIndex])
	}
	if lscope == nil {
		return nil
	}
	return lscope.locateSymbols(name, scope[dotIndex+1:], writer)
}

// dbInheritenceScope encapsulate scope stored in database which can locate symbols
// further up the inheritance hierarchy.

type dbInheritenceScope struct {
	unitName   string
	name       string
	definition string
}

func NewdbInheritenceScope(unitName, name, definition string) *dbInheritenceScope {
	return &dbInheritenceScope{
		unitName:   unitName,
		name:       name,
		definition: definition,
	}
}

// will construct a scope from the database which is ancestor of the current scope
func (dbsh *dbInheritenceScope) getAncestorScope() (inheritenceScope, error) {
	// todo - try find out better solution then parsing the definition(some helper info in the database too)
	re, err := regexp.Compile("class\\s*=\\s*\\(\\s*(\\w*)")
	if err != nil {
		return nil, err
	}
	if re.MatchString(dbsh.definition) {
		matches := re.FindStringSubmatch(dbsh.definition)

		if len(matches) > 1 {
			ancestorName := matches[1]

			// try to find ancestor in same unit first
			symbols, err := SymDB().SearchSymbol(dbsh.unitName, ancestorName)
			if err != nil {
				return nil, err
			} else if len(symbols) > 0 {
				ancSymbol := symbols[0]
				result := NewdbInheritenceScope(ancSymbol.Unitname, ancSymbol.Name, ancSymbol.Definition)
				return result, nil
			}

			// otherwise try to find ancestor in units this unit references
			// todo sort based order - later add order column and fill there some order of geting from tree
			// todo - actually all interface and implementation uses units are searched, split UnitReference to two different kinds, here only from interface section should be used
			// todoe - later use here generalized multi threading algorithm, but which is able to maintaing order of searched results
			symbols, err = SymDB().SearchSymbolByKind(dbsh.unitName, int(UnitReference))
			if err != nil {
				return nil, err
			}
			for _, useUnitSmb := range symbols {
				ancestors, err := SymDB().SearchSymbol(useUnitSmb.Name, ancestorName)
				if err != nil {
					return nil, err
				} else if len(ancestors) > 0 {
					ancSymbol := symbols[0]
					result := NewdbInheritenceScope(ancSymbol.Unitname, ancSymbol.Name, ancSymbol.Definition)
					return result, nil
				}
			}
		}
	}
	return nil, nil
}

func (dbsh *dbInheritenceScope) locateSymbolsViaInheritance(name string, writer SymbolWriter) error {
	SymDB().LocateSymbolsInScope(name, dbsh.unitName, dbsh.name, writer)
	if ancestor, err := dbsh.getAncestorScope(); err != nil {
		return err
	} else if ancestor != nil {
		return ancestor.locateSymbolsViaInheritance(name, writer)
	}
	return nil
}
