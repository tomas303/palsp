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
	GetStop() antlr.Token
}

type identfiable interface {
	Identifier() parser.IIdentifierContext
}

// Helper function to create a Position from any context that has GetStart()
func ctxStartPos(ctx positionable) Position {
	if ctx == nil || ctx.GetStart() == nil {
		return NewPosition(0, 0)
	}
	// Get the line and column from the token (1-based from ANTLR)
	line := ctx.GetStart().GetLine()
	column := ctx.GetStart().GetColumn()

	// Convert to 0-based for our Position system
	return NewPosition(line-1, column)
}

func ctxStopPos(ctx positionable) Position {
	if ctx == nil || ctx.GetStop() == nil {
		return NewPosition(0, 0)
	}
	// Get the line and column from the token (1-based from ANTLR)
	line := ctx.GetStop().GetLine()
	column := ctx.GetStop().GetColumn()

	// Convert to 0-based for our Position system
	return NewPosition(line-1, column)
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
	startPos Position
	stopPos  Position
	ancestor *string
	path     string
	kind     SymbolKind
}

// SymbolCollector defines an interface for collecting symbols
type SymbolCollector interface {
	AddSymbol(name string, kind SymbolKind, definition string, position Position)
	BeginScope(name string, info *scopeInfo)
	EndScope(name string, info *scopeInfo)
	EnterImplementation(position Position)
	AddUseUnit(unit string)
	AccessSpecifier(as AccessSpec)
}

// dbCollector implements SymbolCollector and store symbols in database
type dbCollector struct {
	unitID           int
	db               *symDB
	curentAccessSpec AccessSpec
	currentInfo      scopeInfo
}

func NewDBSymbolCollector(unitID int, db *symDB) *dbCollector {
	return &dbCollector{
		unitID: unitID,
		db:     db,
	}
}

func (dc *dbCollector) BeginScope(name string, si *scopeInfo) {
	dc.currentInfo = *si
}

func (dc *dbCollector) EndScope(name string, si *scopeInfo) {
	if si == nil {
		// in case of error and difference in scope stack
		dc.currentInfo = scopeInfo{}
		return
	}
	dc.currentInfo = *si
}

func (dc *dbCollector) EnterImplementation(position Position) {
	// for now only public symbols are stored in the database - symbols are searched using uses unit
	// references so only public symbols are needed
	panic(ErrListenerBreak)
}

func (dc *dbCollector) AddUseUnit(unit string) {
	// is added as symbol too
}

func (dc *dbCollector) AddSymbol(name string, kind SymbolKind, definition string, position Position) {
	err := dc.db.InsertSymbol(dc.unitID, name, dc.currentInfo.path, int(kind), definition, position)
	if err != nil {
		panic(err)
	}
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
	unitName         string
}

func NewMemorySymbolCollector(unitName string) *memoryCollector {
	scopeStack := newScopeStack()
	unitScope := &unitScope{
		interfaceUses:      *newStack[string](),
		implementationUses: *newStack[string](),
	}
	return &memoryCollector{
		unitScope:  unitScope,
		scopeStack: scopeStack,
		unitName:   unitName,
	}
}

func (mc *memoryCollector) BeginScope(name string, scopeInfo *scopeInfo) {
	var parentScope scope
	if mc.scopeStack.length() == 0 {
		parentScope = mc.unitScope
	} else {
		parentScope = mc.scopeStack.peek()
	}
	newScope := commonScope{name: strings.ToLower(name), info: *scopeInfo, parentScope: parentScope}
	mc.scopeStack.push(&newScope)
}

func (mc *memoryCollector) EndScope(name string, scopeInfo *scopeInfo) {
	scope := mc.scopeStack.pop()
	if mc.scopeStack.length() == 0 {
		mc.unitScope.scope = scope
	} else {
		parentscope := mc.scopeStack.peek()
		scope.parentSWM = parentscope.symbolStack.length() - 1
		parentscope.scopeStack.push(scope)
	}
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
	scope := mc.scopeStack.current()
	smb := Symbol{Name: name, Definition: definition, Kind: int(kind), Position: position, Path: scope.info.path, Unitname: mc.unitName}
	scope.symbolStack.push(smb)
	if scope.scopeStack.length() > 0 && scope.scopeStack.peek().getName() == name {
		scope.scopeStack.peek().setParentSWM(scope.symbolStack.length() - 1)
	}
}

func (mc *memoryCollector) GetScope() TopScope {
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
	scopePath     *stack[string]
	infoStack     *stack[*scopeInfo]
	debugInfo     string // Add debugInfo to track which file we're processing
	pdata         *ParsedData
}

func NewScopesListener(collector SymbolCollector, pdata *ParsedData) *scopesListener {
	return &scopesListener{
		collector: collector,
		scopePath: newStack[string](),
		infoStack: newStack[*scopeInfo](),
		pdata:     pdata,
	}
}

// SetDebugInfo sets the debug info for position mapping
func (s *scopesListener) SetDebugInfo(debugInfo string) {
	s.debugInfo = debugInfo
}

func (s *scopesListener) beginScope(ctxScope positionable, ctxIdentifier identfiable, kind SymbolKind) {
	id := strings.ToLower(buildIdentifier(ctxIdentifier.Identifier()))
	s.scopePath.push(id)
	s.infoStack.push(&scopeInfo{startPos: ctxStartPos(ctxScope), stopPos: ctxStopPos(ctxScope), path: s.scopePath.joinByDot(), kind: kind})
	s.collector.BeginScope(id, s.infoStack.peek())
}

func (s *scopesListener) beginTypeScope(ctxScope positionable, ctxIdentifier identfiable, ancestor *string, kind SymbolKind) {
	id := strings.ToLower(buildIdentifier(ctxIdentifier.Identifier()))
	s.scopePath.push(id)
	s.infoStack.push(&scopeInfo{startPos: ctxStartPos(ctxScope), stopPos: ctxStopPos(ctxScope), ancestor: ancestor, path: s.scopePath.joinByDot(), kind: kind})
	s.collector.BeginScope(buildIdentifier(ctxIdentifier.Identifier()), s.infoStack.peek())
}

func (s *scopesListener) endScope() *scopeInfo {
	id := s.scopePath.pop()
	si := s.infoStack.pop()
	s.collector.EndScope(id, s.infoStack.peek())
	return si
}

func (s *scopesListener) EnterImplementationSection(ctx *parser.ImplementationSectionContext) {
	s.collector.EnterImplementation(ctxStartPos(ctx))
}

func (s *scopesListener) ExitUsesUnits(ctx *parser.UsesUnitsContext) {
	for _, identifier := range ctx.IdentifierList().AllIdentifier() {
		s.collector.AddUseUnit(buildIdentifier(identifier))
		s.collector.AddSymbol(buildIdentifier(identifier), UnitReference, identifier.GetText(), ctxStartPos(identifier))
	}
}

func (s *scopesListener) ExitVariableDeclaration(ctx *parser.VariableDeclarationContext) {
	typedef := buildUnderscoreTypeDef(ctx.TypedIdentifierList().Type_())
	for _, identifier := range ctx.TypedIdentifierList().IdentifierList().AllIdentifier() {
		s.collector.AddSymbol(buildIdentifier(identifier), VariableSymbol, typedef, ctxStartPos(identifier))
	}
}

func (s *scopesListener) ExitVariableDeclarationStatement(ctx *parser.VariableDeclarationStatementContext) {
	typedef := ""
	if ctx.Type_() != nil {
		typedef = buildUnderscoreTypeDef(ctx.Type_())
	}
	if ctx.Expression() != nil {
		// todo - if there is no type definition then it should be taken from expression
		typedef += " := " + ctx.Expression().GetText()
	}
	for _, identifier := range ctx.IdentifierList().AllIdentifier() {
		s.collector.AddSymbol(buildIdentifier(identifier), VariableSymbol, typedef, ctxStartPos(identifier))
	}
}

func (s *scopesListener) ExitConstantDefinition(ctx *parser.ConstantDefinitionContext) {
	fieldtype := ""
	if ctx.TypeIdentifier() != nil {
		fieldtype = buildTypeIdentifier(ctx.TypeIdentifier())
	}
	if fieldtype == "" {
		if ctx.Constant() != nil {
			if ctx.Constant().UnsignedConstant() != nil {
				if ctx.Constant().UnsignedConstant().String_() != nil {
					fieldtype = "string"
				} else if ctx.Constant().UnsignedConstant().UnsignedNumber() != nil {
					fieldtype = "integer"
				}
			}
		}
	}
	s.collector.AddSymbol(buildIdentifier(ctx.Identifier()), ConstantSymbol, fieldtype, ctxStartPos(ctx.Identifier()))
}

func (s *scopesListener) ExitFormalParameterList(ctx *parser.FormalParameterListContext) {
	for _, parSecCtx := range ctx.AllFormalParameterSection() {
		if parSecCtx.ParameterGroup() != nil {
			for _, id := range parSecCtx.ParameterGroup().IdentifierList().AllIdentifier() {
				s.collector.AddSymbol(buildIdentifier(id), ParameterSymbol, buildOneParameter(parSecCtx, id), ctxStartPos(id))
			}
		}
	}
}

func (s *scopesListener) ExitClassDeclarationPart(ctx *parser.ClassDeclarationPartContext) {
	if ctx.TypedIdentifierList() != nil {
		typedef := buildUnderscoreTypeDef(ctx.TypedIdentifierList().Type_())
		for _, id := range ctx.TypedIdentifierList().IdentifierList().AllIdentifier() {
			s.collector.AddSymbol(buildIdentifier(id), ClassVariable, typedef, ctxStartPos(id))
		}
	}
}

func (s *scopesListener) ExitPropertyDeclaration(ctx *parser.PropertyDeclarationContext) {
	s.collector.AddSymbol(buildIdentifier(ctx.Identifier()), PropertySymbol, buildProperty(ctx), ctxStartPos(ctx.Identifier()))
}

func (s *scopesListener) ExitForStatement(ctx *parser.ForStatementContext) {
	if ctx.VAR() != nil && ctx.Identifier() != nil {
		s.collector.AddSymbol(buildIdentifier(ctx.Identifier()), VariableSymbol, "(for loop inferred)", ctxStartPos(ctx.Identifier()))
	}
}

func (s *scopesListener) EnterUnit(ctx *parser.UnitContext) {
	s.beginScope(ctx, ctx, Unit)
}

func (s *scopesListener) ExitUnit(ctx *parser.UnitContext) {
	s.endScope()
}

func (s *scopesListener) EnterProcedureHeader(ctx *parser.ProcedureHeaderContext) {
	if s.inDeclaration {
		return
	}
	s.beginScope(ctx, ctx, ProcedureSymbol)
}

func (s *scopesListener) ExitProcedureHeader(ctx *parser.ProcedureHeaderContext) {
	if s.inDeclaration {
		s.collector.AddSymbol(getLastIdent(ctx.Identifier()), ProcedureSymbol, buildProcedureHeader(ctx), ctxStartPos(ctx.Identifier()))
		return
	}
	s.endScope()
	s.collector.AddSymbol(getLastIdent(ctx.Identifier()), ProcedureSymbol, buildProcedureHeader(ctx), ctxStartPos(ctx.Identifier()))
}

func (s *scopesListener) EnterProcedureDeclaration(ctx *parser.ProcedureDeclarationContext) {
	s.inDeclaration = true
	s.beginScope(ctx, ctx.ProcedureHeader(), ProcedureSymbol)
}

func (s *scopesListener) ExitProcedureDeclaration(ctx *parser.ProcedureDeclarationContext) {
	s.endScope()
	s.inDeclaration = false
}

func (s *scopesListener) EnterFunctionHeader(ctx *parser.FunctionHeaderContext) {
	if s.inDeclaration {
		return
	}
	s.beginScope(ctx, ctx, FunctionSymbol)
}

func (s *scopesListener) ExitFunctionHeader(ctx *parser.FunctionHeaderContext) {
	if s.inDeclaration {
		if ctx.ResultType() != nil {
			s.collector.AddSymbol("result", FunctionResult, buildTypeIdentifier(ctx.ResultType().TypeIdentifier()), ctxStartPos(ctx.Identifier()))
		}
		s.collector.AddSymbol(getLastIdent(ctx.Identifier()), FunctionSymbol, buildFunctionHeader(ctx), ctxStartPos(ctx.Identifier()))
		return
	}
	if ctx.ResultType() != nil {
		s.collector.AddSymbol("result", FunctionResult, buildTypeIdentifier(ctx.ResultType().TypeIdentifier()), ctxStartPos(ctx.Identifier()))
	}
	s.endScope()
	s.collector.AddSymbol(getLastIdent(ctx.Identifier()), FunctionSymbol, buildFunctionHeader(ctx), ctxStartPos(ctx.Identifier()))
}

func (s *scopesListener) EnterFunctionDeclaration(ctx *parser.FunctionDeclarationContext) {
	s.inDeclaration = true
	s.beginScope(ctx, ctx.FunctionHeader(), FunctionSymbol)
}

func (s *scopesListener) ExitFunctionDeclaration(ctx *parser.FunctionDeclarationContext) {
	s.endScope()
	s.inDeclaration = false
}

func (s *scopesListener) EnterTypeDefinition(ctx *parser.TypeDefinitionContext) {
	var ancestor *string
	var kind SymbolKind
	if ctx.Type_() != nil && ctx.Type_().StructuredType() != nil {
		if ctx.Type_().StructuredType().HelperType() != nil {
			tmp := buildTypeIdentifier(ctx.Type_().StructuredType().HelperType().TypeIdentifier())
			ancestor = &tmp
			kind = HelperSymbol
		} else if ctx.Type_().StructuredType().ClassType() != nil {
			tmp := buildIdentifier(ctx.Type_().StructuredType().ClassType().Identifier())
			ancestor = &tmp
			kind = ClassSymbol
		} else if ctx.Type_().StructuredType().InterfaceType() != nil {
			tmp := buildIdentifier(ctx.Type_().StructuredType().InterfaceType().Identifier())
			ancestor = &tmp
			kind = InterfaceSymbol
		} else {
			ancestor = nil
			kind = TypeSymbol
		}
	} else {
		ancestor = nil
		kind = TypeSymbol
	}
	s.beginTypeScope(ctx, ctx, ancestor, kind)
}

func (s *scopesListener) ExitTypeDefinition(ctx *parser.TypeDefinitionContext) {
	si := s.endScope()
	s.collector.AddSymbol(buildIdentifier(ctx.Identifier()), si.kind, buildTypeDef(ctx), ctxStartPos(ctx.Identifier()))
}

// Helper function to get min of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
