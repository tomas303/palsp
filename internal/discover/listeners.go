package discover

import (
	"errors"
	"palsp/internal/parser"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

// Interface for any object that has a GetStart() method returning something with GetLine() and GetColumn()
type positionable interface {
	GetStart() antlr.Token
}

// Helper function to create a Position from any context that has GetStart()
func positionFromCtx(ctx positionable) Position {
	if ctx == nil || ctx.GetStart() == nil {
		return NewPosition(0, 0)
	}
	return NewPosition(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
}

// Scope visibility levels
type AccessSpec int

const (
	AccUnknown AccessSpec = iota
	AccPrivate
	AccStrictPrivate
	AccProtected
	AccStrictProtected
	AccPublic
	AccPublished
)

var ErrListenerBreak = errors.New("listener breaks gracefully")

// helper information for scope (can differ based on scope type)
type scopeInfo struct {
	position Position
	ancestor *string
}

// SymbolCollector defines an interface for collecting symbols
type SymbolCollector interface {
	AddSymbol(name string, kind SymbolKind, definition string, position Position)
	BeginScope(name string, info scopeInfo)
	EndScope(name string)
	EnterImplementation(position Position)
	AddUseUnit(unit string)
	AccessSpecifier(as AccessSpec)
}

// dbCollector implements SymbolCollector and store symbols in database
type dbCollector struct {
	unitID           int
	db               *symDB
	currentScope     stack[string]
	curentAccessSpec AccessSpec
}

func NewDBSymbolCollector(unitID int, db *symDB) *dbCollector {
	return &dbCollector{
		unitID:       unitID,
		db:           db,
		currentScope: *newStack[string](),
	}
}

func (dc *dbCollector) BeginScope(name string, scopeInfo scopeInfo) {
	if dc.currentScope.length() == 0 {
		dc.currentScope.push(strings.ToLower(name))
	} else {
		dc.currentScope.push(dc.currentScope.peek() + "." + strings.ToLower(name))
	}
}

func (dc *dbCollector) EndScope(name string) {
	dc.currentScope.pop()
}

func (dc *dbCollector) EnterImplementation(position Position) {
	// for now only public symbols are stored in the database - symbols are searched using uses unit
	// references so only public symbols are needed
	panic(ErrListenerBreak)
}

func (dc *dbCollector) AddUseUnit(unit string) {
	dc.db.InsertSymbol(dc.unitID, unit, dc.currentScope.peek(), int(UnitReference), unit)
}

func (dc *dbCollector) AddSymbol(name string, kind SymbolKind, definition string, position Position) {
	dc.db.InsertSymbol(dc.unitID, name, dc.currentScope.peek(), int(kind), definition)
}

func (dc *dbCollector) AccessSpecifier(as AccessSpec) {
	dc.curentAccessSpec = as
}

// memoryCollector implements SymbolCollector and creates structure of scopes
// and returns it as TopScope
type memoryCollector struct {
	unitScope        *unitScope
	scopeStack       *scopeStack
	inImplementation bool
	// todo - maybe replace with some support of listener (listener can track current scope internally ... its just names of scopes split by ".")
	currentScope stack[string]
	unitName     string
}

func NewMemorySymbolCollector(unitName string) *memoryCollector {
	rootScope := commonScope{name: "root"}
	scopeStack := newScopeStack()
	scopeStack.push(&rootScope)
	unitScope := &unitScope{
		interfaceUses:      *newStack[string](),
		implementationUses: *newStack[string](),
	}
	currentScope := stack[string]{}
	return &memoryCollector{
		unitScope:    unitScope,
		scopeStack:   scopeStack,
		currentScope: currentScope,
		unitName:     unitName,
	}
}

func (mc *memoryCollector) BeginScope(name string, scopeInfo scopeInfo) {
	// todo - create some specific logger for parsing scopes(too much infos so to be able to allow only something per demand and use it in listener)
	// log.Logger.Debug().Str("begin scope", name).Int("line", position.Line).Int("chr", position.Character).Send()
	var parentScope *commonScope
	if mc.scopeStack.length() == 0 {
		parentScope = nil
	} else {
		parentScope = mc.scopeStack.peek()
	}
	newScope := commonScope{name: strings.ToLower(name), info: scopeInfo, parentScope: parentScope}
	mc.scopeStack.push(&newScope)
	if mc.currentScope.length() == 0 {
		mc.currentScope.push(strings.ToLower(name))
	} else {
		mc.currentScope.push(mc.currentScope.peek() + "." + strings.ToLower(name))
	}
}

func (mc *memoryCollector) EndScope(name string) {
	// log.Logger.Debug().Str("end scope", name).Send()
	scope := mc.scopeStack.pop()
	parentscope := mc.scopeStack.peek()
	scope.parentSWM = parentscope.symbolStack.length() - 1
	parentscope.scopeStack.push(scope)
	mc.currentScope.pop()
}

func (mc *memoryCollector) EnterImplementation(position Position) {
	mc.inImplementation = true
	mc.unitScope.implementationPos = position
}

func (mc *memoryCollector) AddUseUnit(unit string) {
	if mc.inImplementation {
		mc.unitScope.implementationUses.push(unit)
	} else {
		mc.unitScope.interfaceUses.push(unit)
	}
}

func (mc *memoryCollector) AddSymbol(name string, kind SymbolKind, definition string, position Position) {
	name = strings.ToLower(name)
	smb := Symbol{Name: name, Definition: definition, Kind: int(kind), Position: position, Scope: mc.currentScope.peek(), Unitname: mc.unitName}
	scope := mc.scopeStack.current()
	scope.symbolStack.push(smb)
	if scope.scopeStack.length() > 0 && scope.scopeStack.peek().getName() == name {
		scope.scopeStack.peek().setParentSWM(scope.symbolStack.length() - 1)
	}
}

func (mc *memoryCollector) GetScope() TopScope {
	// todo later try to have it like root from the beginning
	topCommon := mc.scopeStack.peek().scopeStack.peek()
	topCommon.parentScope = mc.unitScope
	mc.unitScope.scope = topCommon
	return mc.unitScope
}

func (dc *memoryCollector) AccessSpecifier(as AccessSpec) {
	// todo - implement access specifier(must be in db too and locate symbols will have some filtering function to reflect various conditions)
}

// scopeStack manages a stack of scope builders
type scopeStack struct {
	items []*commonScope
}

// newScopeStack creates a new scope builder stack
func newScopeStack() *scopeStack {
	return &scopeStack{
		items: make([]*commonScope, 0),
	}
}

// push adds a scope builder to the top of the stack
func (s *scopeStack) push(item *commonScope) {
	s.items = append(s.items, item)
}

// pop removes and returns the top scope builder from the stack
func (s *scopeStack) pop() *commonScope {
	if len(s.items) == 0 {
		return nil
	}

	n := len(s.items) - 1
	item := s.items[n]
	s.items = s.items[:n]
	return item
}

// peek returns the top scope builder without removing it
func (s *scopeStack) peek() *commonScope {
	if len(s.items) == 0 {
		return nil
	}
	return s.items[len(s.items)-1]
}

// current returns the current scope builder (alias for peek)
func (s *scopeStack) current() *commonScope {
	return s.peek()
}

// length returns the number of items in the stack
func (s *scopeStack) length() int {
	return len(s.items)
}

// scopesListener separates beginings and endings of scopes and collects symbols
// in between
type scopesListener struct {
	parser.BasepascalListener
	collector     SymbolCollector
	inDeclaration bool
}

func NewScopesListener(collector SymbolCollector) *scopesListener {
	return &scopesListener{
		collector: collector,
	}
}

func (s *scopesListener) EnterImplementationSection(ctx *parser.ImplementationSectionContext) {
	s.collector.EnterImplementation(positionFromCtx(ctx))
}

func (s *scopesListener) ExitUsesUnits(ctx *parser.UsesUnitsContext) {
	for _, identifier := range ctx.IdentifierList().AllIdentifier() {
		s.collector.AddUseUnit(buildIdentifier(identifier))
		s.collector.AddSymbol(buildIdentifier(identifier), UnitReference, identifier.GetText(), positionFromCtx(identifier))
	}
}

func (s *scopesListener) ExitVariableDeclaration(ctx *parser.VariableDeclarationContext) {
	typedef := buildUnderscoreTypeDef(ctx.TypedIdentifierList().Type_())
	for _, identifier := range ctx.TypedIdentifierList().IdentifierList().AllIdentifier() {
		s.collector.AddSymbol(buildIdentifier(identifier), VariableSymbol, typedef, positionFromCtx(identifier))
	}
}

func (s *scopesListener) ExitConstantDefinition(ctx *parser.ConstantDefinitionContext) {
	fieldtype := ""
	if ctx.TypeIdentifier() != nil {
		fieldtype = buildTypeIdentifier(ctx.TypeIdentifier())
	}
	if fieldtype == "" {
		if ctx.Constant().String_() != nil {
			fieldtype = "string"
		} else if ctx.Constant().UnsignedNumber() != nil {
			fieldtype = "integer"
		} else if ctx.Constant().Sign() != nil {
			fieldtype = "integer"
		}
	}
	s.collector.AddSymbol(buildIdentifier(ctx.Identifier()), ConstantSymbol, fieldtype, positionFromCtx(ctx.Identifier()))
}

func (s *scopesListener) ExitFormalParameterList(ctx *parser.FormalParameterListContext) {
	for _, parSecCtx := range ctx.AllFormalParameterSection() {
		parType := ""
		if parSecCtx.ParameterGroup() == nil {
			continue
		}
		if parSecCtx.ParameterGroup().TypeIdentifier() != nil {
			parType = parSecCtx.ParameterGroup().TypeIdentifier().GetText()
		}
		// todo - builders have only buildParameterGroup ... this must be changed to be able to create only one parameter (each symbol is independent)
		// if ctx.DefaultValue() != nil {
		// 	result += " = " + ctx.DefaultValue().GetText()
		// }
		for _, id := range parSecCtx.ParameterGroup().IdentifierList().AllIdentifier() {
			s.collector.AddSymbol(buildIdentifier(id), ParameterSymbol, parType, positionFromCtx(id))
		}
		//
	}
}

func (s *scopesListener) ExitClassDeclarationPart(ctx *parser.ClassDeclarationPartContext) {
	if ctx.TypedIdentifierList() != nil {
		typedef := buildUnderscoreTypeDef(ctx.TypedIdentifierList().Type_())
		for _, id := range ctx.TypedIdentifierList().IdentifierList().AllIdentifier() {
			s.collector.AddSymbol(buildIdentifier(id), ClassVariable, typedef, positionFromCtx(id))
		}
	}
}

func (s *scopesListener) EnterUnit(ctx *parser.UnitContext) {
	s.collector.BeginScope(buildIdentifier(ctx.Identifier()), scopeInfo{position: positionFromCtx(ctx.Identifier())})
}

func (s *scopesListener) ExitUnit(ctx *parser.UnitContext) {
	s.collector.EndScope(buildIdentifier(ctx.Identifier()))
}

func (s *scopesListener) EnterProcedureHeader(ctx *parser.ProcedureHeaderContext) {
	if s.inDeclaration {
		return
	}
	s.collector.BeginScope(buildIdentifier(ctx.Identifier()), scopeInfo{position: positionFromCtx(ctx.Identifier())})
}

func (s *scopesListener) ExitProcedureHeader(ctx *parser.ProcedureHeaderContext) {
	if s.inDeclaration {
		s.collector.AddSymbol(getLastIdent(ctx.Identifier()), ProcedureSymbol, buildProcedureHeader(ctx), positionFromCtx(ctx.Identifier()))
		return
	}
	s.collector.EndScope(buildIdentifier(ctx.Identifier()))
	s.collector.AddSymbol(getLastIdent(ctx.Identifier()), ProcedureSymbol, buildProcedureHeader(ctx), positionFromCtx(ctx.Identifier()))
}

func (s *scopesListener) EnterProcedureDeclaration(ctx *parser.ProcedureDeclarationContext) {
	s.inDeclaration = true
	s.collector.BeginScope(buildIdentifier(ctx.ProcedureHeader().Identifier()), scopeInfo{position: positionFromCtx(ctx.ProcedureHeader().Identifier())})
}

func (s *scopesListener) ExitProcedureDeclaration(ctx *parser.ProcedureDeclarationContext) {
	s.collector.EndScope(buildIdentifier(ctx.ProcedureHeader().Identifier()))
	s.inDeclaration = false
}

func (s *scopesListener) EnterFunctionHeader(ctx *parser.FunctionHeaderContext) {
	if s.inDeclaration {
		return
	}
	s.collector.BeginScope(buildIdentifier(ctx.Identifier()), scopeInfo{position: positionFromCtx(ctx.Identifier())})
}

func (s *scopesListener) ExitFunctionHeader(ctx *parser.FunctionHeaderContext) {
	if s.inDeclaration {
		if ctx.ResultType() != nil {
			s.collector.AddSymbol("result", FunctionResult, buildTypeIdentifier(ctx.ResultType().TypeIdentifier()), positionFromCtx(ctx.Identifier()))
		}
		s.collector.AddSymbol(getLastIdent(ctx.Identifier()), FunctionSymbol, buildFunctionHeader(ctx), positionFromCtx(ctx.Identifier()))
		return
	}
	if ctx.ResultType() != nil {
		s.collector.AddSymbol("result", FunctionResult, buildTypeIdentifier(ctx.ResultType().TypeIdentifier()), positionFromCtx(ctx.Identifier()))
	}
	s.collector.EndScope(buildIdentifier(ctx.Identifier()))
	s.collector.AddSymbol(getLastIdent(ctx.Identifier()), FunctionSymbol, buildFunctionHeader(ctx), positionFromCtx(ctx.Identifier()))
}

func (s *scopesListener) EnterFunctionDeclaration(ctx *parser.FunctionDeclarationContext) {
	s.inDeclaration = true
	s.collector.BeginScope(buildIdentifier(ctx.FunctionHeader().Identifier()), scopeInfo{position: positionFromCtx(ctx.FunctionHeader().Identifier())})
}

func (s *scopesListener) ExitFunctionDeclaration(ctx *parser.FunctionDeclarationContext) {
	s.collector.EndScope(buildIdentifier(ctx.FunctionHeader().Identifier()))
	s.inDeclaration = false
}

func (s *scopesListener) EnterTypeDefinition(ctx *parser.TypeDefinitionContext) {
	var ancestor *string
	if ctx.Type_() != nil && ctx.Type_().StructuredType() != nil && ctx.Type_().StructuredType().ClassType() != nil {
		tmp := buildIdentifier(ctx.Type_().StructuredType().ClassType().Identifier())
		ancestor = &tmp
	} else {
		ancestor = nil
	}
	s.collector.BeginScope(buildIdentifier(ctx.Identifier()), scopeInfo{position: positionFromCtx(ctx.Identifier()), ancestor: ancestor})
}

func (s *scopesListener) ExitTypeDefinition(ctx *parser.TypeDefinitionContext) {
	s.collector.EndScope(buildIdentifier(ctx.Identifier()))
	s.collector.AddSymbol(buildIdentifier(ctx.Identifier()), TypeSymbol, buildTypeDef(ctx), positionFromCtx(ctx.Identifier()))
}
