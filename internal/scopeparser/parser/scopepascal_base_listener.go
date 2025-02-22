// Code generated from /home/tomas/development/tomas303/projects/palsp/internal/scopeparser/scopepascal.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // scopepascal

import "github.com/antlr4-go/antlr/v4"

// BasescopepascalListener is a complete listener for a parse tree produced by scopepascalParser.
type BasescopepascalListener struct{}

var _ scopepascalListener = &BasescopepascalListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasescopepascalListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasescopepascalListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasescopepascalListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasescopepascalListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSource is called when production source is entered.
func (s *BasescopepascalListener) EnterSource(ctx *SourceContext) {}

// ExitSource is called when production source is exited.
func (s *BasescopepascalListener) ExitSource(ctx *SourceContext) {}

// EnterProgram is called when production program is entered.
func (s *BasescopepascalListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BasescopepascalListener) ExitProgram(ctx *ProgramContext) {}

// EnterUnit is called when production unit is entered.
func (s *BasescopepascalListener) EnterUnit(ctx *UnitContext) {}

// ExitUnit is called when production unit is exited.
func (s *BasescopepascalListener) ExitUnit(ctx *UnitContext) {}

// EnterInterfaceSection is called when production interfaceSection is entered.
func (s *BasescopepascalListener) EnterInterfaceSection(ctx *InterfaceSectionContext) {}

// ExitInterfaceSection is called when production interfaceSection is exited.
func (s *BasescopepascalListener) ExitInterfaceSection(ctx *InterfaceSectionContext) {}

// EnterImplementationSection is called when production implementationSection is entered.
func (s *BasescopepascalListener) EnterImplementationSection(ctx *ImplementationSectionContext) {}

// ExitImplementationSection is called when production implementationSection is exited.
func (s *BasescopepascalListener) ExitImplementationSection(ctx *ImplementationSectionContext) {}

// EnterInitializationSection is called when production initializationSection is entered.
func (s *BasescopepascalListener) EnterInitializationSection(ctx *InitializationSectionContext) {}

// ExitInitializationSection is called when production initializationSection is exited.
func (s *BasescopepascalListener) ExitInitializationSection(ctx *InitializationSectionContext) {}

// EnterFinalizationSection is called when production finalizationSection is entered.
func (s *BasescopepascalListener) EnterFinalizationSection(ctx *FinalizationSectionContext) {}

// ExitFinalizationSection is called when production finalizationSection is exited.
func (s *BasescopepascalListener) ExitFinalizationSection(ctx *FinalizationSectionContext) {}

// EnterInterfaceBlock is called when production interfaceBlock is entered.
func (s *BasescopepascalListener) EnterInterfaceBlock(ctx *InterfaceBlockContext) {}

// ExitInterfaceBlock is called when production interfaceBlock is exited.
func (s *BasescopepascalListener) ExitInterfaceBlock(ctx *InterfaceBlockContext) {}

// EnterImplementationBlock is called when production implementationBlock is entered.
func (s *BasescopepascalListener) EnterImplementationBlock(ctx *ImplementationBlockContext) {}

// ExitImplementationBlock is called when production implementationBlock is exited.
func (s *BasescopepascalListener) ExitImplementationBlock(ctx *ImplementationBlockContext) {}

// EnterUnitList is called when production unitList is entered.
func (s *BasescopepascalListener) EnterUnitList(ctx *UnitListContext) {}

// ExitUnitList is called when production unitList is exited.
func (s *BasescopepascalListener) ExitUnitList(ctx *UnitListContext) {}

// EnterLabelDeclaration is called when production labelDeclaration is entered.
func (s *BasescopepascalListener) EnterLabelDeclaration(ctx *LabelDeclarationContext) {}

// ExitLabelDeclaration is called when production labelDeclaration is exited.
func (s *BasescopepascalListener) ExitLabelDeclaration(ctx *LabelDeclarationContext) {}

// EnterConstSection is called when production constSection is entered.
func (s *BasescopepascalListener) EnterConstSection(ctx *ConstSectionContext) {}

// ExitConstSection is called when production constSection is exited.
func (s *BasescopepascalListener) ExitConstSection(ctx *ConstSectionContext) {}

// EnterResourceSection is called when production resourceSection is entered.
func (s *BasescopepascalListener) EnterResourceSection(ctx *ResourceSectionContext) {}

// ExitResourceSection is called when production resourceSection is exited.
func (s *BasescopepascalListener) ExitResourceSection(ctx *ResourceSectionContext) {}

// EnterTypeSection is called when production typeSection is entered.
func (s *BasescopepascalListener) EnterTypeSection(ctx *TypeSectionContext) {}

// ExitTypeSection is called when production typeSection is exited.
func (s *BasescopepascalListener) ExitTypeSection(ctx *TypeSectionContext) {}

// EnterTypeBlock is called when production typeBlock is entered.
func (s *BasescopepascalListener) EnterTypeBlock(ctx *TypeBlockContext) {}

// ExitTypeBlock is called when production typeBlock is exited.
func (s *BasescopepascalListener) ExitTypeBlock(ctx *TypeBlockContext) {}

// EnterType is called when production type is entered.
func (s *BasescopepascalListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BasescopepascalListener) ExitType(ctx *TypeContext) {}

// EnterAccessSpecifier is called when production accessSpecifier is entered.
func (s *BasescopepascalListener) EnterAccessSpecifier(ctx *AccessSpecifierContext) {}

// ExitAccessSpecifier is called when production accessSpecifier is exited.
func (s *BasescopepascalListener) ExitAccessSpecifier(ctx *AccessSpecifierContext) {}

// EnterProcedureOrFunctionModifiers is called when production procedureOrFunctionModifiers is entered.
func (s *BasescopepascalListener) EnterProcedureOrFunctionModifiers(ctx *ProcedureOrFunctionModifiersContext) {
}

// ExitProcedureOrFunctionModifiers is called when production procedureOrFunctionModifiers is exited.
func (s *BasescopepascalListener) ExitProcedureOrFunctionModifiers(ctx *ProcedureOrFunctionModifiersContext) {
}

// EnterClassForwardDeclaration is called when production classForwardDeclaration is entered.
func (s *BasescopepascalListener) EnterClassForwardDeclaration(ctx *ClassForwardDeclarationContext) {}

// ExitClassForwardDeclaration is called when production classForwardDeclaration is exited.
func (s *BasescopepascalListener) ExitClassForwardDeclaration(ctx *ClassForwardDeclarationContext) {}

// EnterClassType is called when production classType is entered.
func (s *BasescopepascalListener) EnterClassType(ctx *ClassTypeContext) {}

// ExitClassType is called when production classType is exited.
func (s *BasescopepascalListener) ExitClassType(ctx *ClassTypeContext) {}

// EnterRecordType is called when production recordType is entered.
func (s *BasescopepascalListener) EnterRecordType(ctx *RecordTypeContext) {}

// ExitRecordType is called when production recordType is exited.
func (s *BasescopepascalListener) ExitRecordType(ctx *RecordTypeContext) {}

// EnterPropertyDeclaration is called when production propertyDeclaration is entered.
func (s *BasescopepascalListener) EnterPropertyDeclaration(ctx *PropertyDeclarationContext) {}

// ExitPropertyDeclaration is called when production propertyDeclaration is exited.
func (s *BasescopepascalListener) ExitPropertyDeclaration(ctx *PropertyDeclarationContext) {}

// EnterPropertyReadDeclaration is called when production propertyReadDeclaration is entered.
func (s *BasescopepascalListener) EnterPropertyReadDeclaration(ctx *PropertyReadDeclarationContext) {}

// ExitPropertyReadDeclaration is called when production propertyReadDeclaration is exited.
func (s *BasescopepascalListener) ExitPropertyReadDeclaration(ctx *PropertyReadDeclarationContext) {}

// EnterPropertyWriteDeclaration is called when production propertyWriteDeclaration is entered.
func (s *BasescopepascalListener) EnterPropertyWriteDeclaration(ctx *PropertyWriteDeclarationContext) {
}

// ExitPropertyWriteDeclaration is called when production propertyWriteDeclaration is exited.
func (s *BasescopepascalListener) ExitPropertyWriteDeclaration(ctx *PropertyWriteDeclarationContext) {
}

// EnterPropertyDefaultValueDeclaration is called when production propertyDefaultValueDeclaration is entered.
func (s *BasescopepascalListener) EnterPropertyDefaultValueDeclaration(ctx *PropertyDefaultValueDeclarationContext) {
}

// ExitPropertyDefaultValueDeclaration is called when production propertyDefaultValueDeclaration is exited.
func (s *BasescopepascalListener) ExitPropertyDefaultValueDeclaration(ctx *PropertyDefaultValueDeclarationContext) {
}

// EnterPropertyIndexDeclaration is called when production propertyIndexDeclaration is entered.
func (s *BasescopepascalListener) EnterPropertyIndexDeclaration(ctx *PropertyIndexDeclarationContext) {
}

// ExitPropertyIndexDeclaration is called when production propertyIndexDeclaration is exited.
func (s *BasescopepascalListener) ExitPropertyIndexDeclaration(ctx *PropertyIndexDeclarationContext) {
}

// EnterPropertyIndexParameters is called when production propertyIndexParameters is entered.
func (s *BasescopepascalListener) EnterPropertyIndexParameters(ctx *PropertyIndexParametersContext) {}

// ExitPropertyIndexParameters is called when production propertyIndexParameters is exited.
func (s *BasescopepascalListener) ExitPropertyIndexParameters(ctx *PropertyIndexParametersContext) {}

// EnterPropertyIndexParametersList is called when production propertyIndexParametersList is entered.
func (s *BasescopepascalListener) EnterPropertyIndexParametersList(ctx *PropertyIndexParametersListContext) {
}

// ExitPropertyIndexParametersList is called when production propertyIndexParametersList is exited.
func (s *BasescopepascalListener) ExitPropertyIndexParametersList(ctx *PropertyIndexParametersListContext) {
}

// EnterArrayType is called when production arrayType is entered.
func (s *BasescopepascalListener) EnterArrayType(ctx *ArrayTypeContext) {}

// ExitArrayType is called when production arrayType is exited.
func (s *BasescopepascalListener) ExitArrayType(ctx *ArrayTypeContext) {}

// EnterPointerType is called when production pointerType is entered.
func (s *BasescopepascalListener) EnterPointerType(ctx *PointerTypeContext) {}

// ExitPointerType is called when production pointerType is exited.
func (s *BasescopepascalListener) ExitPointerType(ctx *PointerTypeContext) {}

// EnterSetType is called when production setType is entered.
func (s *BasescopepascalListener) EnterSetType(ctx *SetTypeContext) {}

// ExitSetType is called when production setType is exited.
func (s *BasescopepascalListener) ExitSetType(ctx *SetTypeContext) {}

// EnterFileType is called when production fileType is entered.
func (s *BasescopepascalListener) EnterFileType(ctx *FileTypeContext) {}

// ExitFileType is called when production fileType is exited.
func (s *BasescopepascalListener) ExitFileType(ctx *FileTypeContext) {}

// EnterScalarType is called when production scalarType is entered.
func (s *BasescopepascalListener) EnterScalarType(ctx *ScalarTypeContext) {}

// ExitScalarType is called when production scalarType is exited.
func (s *BasescopepascalListener) ExitScalarType(ctx *ScalarTypeContext) {}

// EnterSubrangeType is called when production subrangeType is entered.
func (s *BasescopepascalListener) EnterSubrangeType(ctx *SubrangeTypeContext) {}

// ExitSubrangeType is called when production subrangeType is exited.
func (s *BasescopepascalListener) ExitSubrangeType(ctx *SubrangeTypeContext) {}

// EnterBlockDeclaration is called when production blockDeclaration is entered.
func (s *BasescopepascalListener) EnterBlockDeclaration(ctx *BlockDeclarationContext) {}

// ExitBlockDeclaration is called when production blockDeclaration is exited.
func (s *BasescopepascalListener) ExitBlockDeclaration(ctx *BlockDeclarationContext) {}

// EnterFunctionDeclaration is called when production functionDeclaration is entered.
func (s *BasescopepascalListener) EnterFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// ExitFunctionDeclaration is called when production functionDeclaration is exited.
func (s *BasescopepascalListener) ExitFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// EnterProcedureDeclaration is called when production procedureDeclaration is entered.
func (s *BasescopepascalListener) EnterProcedureDeclaration(ctx *ProcedureDeclarationContext) {}

// ExitProcedureDeclaration is called when production procedureDeclaration is exited.
func (s *BasescopepascalListener) ExitProcedureDeclaration(ctx *ProcedureDeclarationContext) {}

// EnterFunctionOrProcedureDeclaration is called when production functionOrProcedureDeclaration is entered.
func (s *BasescopepascalListener) EnterFunctionOrProcedureDeclaration(ctx *FunctionOrProcedureDeclarationContext) {
}

// ExitFunctionOrProcedureDeclaration is called when production functionOrProcedureDeclaration is exited.
func (s *BasescopepascalListener) ExitFunctionOrProcedureDeclaration(ctx *FunctionOrProcedureDeclarationContext) {
}

// EnterFunctionOrProcedure is called when production functionOrProcedure is entered.
func (s *BasescopepascalListener) EnterFunctionOrProcedure(ctx *FunctionOrProcedureContext) {}

// ExitFunctionOrProcedure is called when production functionOrProcedure is exited.
func (s *BasescopepascalListener) ExitFunctionOrProcedure(ctx *FunctionOrProcedureContext) {}

// EnterBlockStatement is called when production blockStatement is entered.
func (s *BasescopepascalListener) EnterBlockStatement(ctx *BlockStatementContext) {}

// ExitBlockStatement is called when production blockStatement is exited.
func (s *BasescopepascalListener) ExitBlockStatement(ctx *BlockStatementContext) {}

// EnterRecordVariantDeclaration is called when production recordVariantDeclaration is entered.
func (s *BasescopepascalListener) EnterRecordVariantDeclaration(ctx *RecordVariantDeclarationContext) {
}

// ExitRecordVariantDeclaration is called when production recordVariantDeclaration is exited.
func (s *BasescopepascalListener) ExitRecordVariantDeclaration(ctx *RecordVariantDeclarationContext) {
}

// EnterRecordVariant is called when production recordVariant is entered.
func (s *BasescopepascalListener) EnterRecordVariant(ctx *RecordVariantContext) {}

// ExitRecordVariant is called when production recordVariant is exited.
func (s *BasescopepascalListener) ExitRecordVariant(ctx *RecordVariantContext) {}

// EnterStatementError is called when production statementError is entered.
func (s *BasescopepascalListener) EnterStatementError(ctx *StatementErrorContext) {}

// ExitStatementError is called when production statementError is exited.
func (s *BasescopepascalListener) ExitStatementError(ctx *StatementErrorContext) {}

// EnterStatement is called when production statement is entered.
func (s *BasescopepascalListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BasescopepascalListener) ExitStatement(ctx *StatementContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BasescopepascalListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BasescopepascalListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterIdentifierList is called when production identifierList is entered.
func (s *BasescopepascalListener) EnterIdentifierList(ctx *IdentifierListContext) {}

// ExitIdentifierList is called when production identifierList is exited.
func (s *BasescopepascalListener) ExitIdentifierList(ctx *IdentifierListContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *BasescopepascalListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BasescopepascalListener) ExitExpressionList(ctx *ExpressionListContext) {}

// EnterParamsDeclaration is called when production paramsDeclaration is entered.
func (s *BasescopepascalListener) EnterParamsDeclaration(ctx *ParamsDeclarationContext) {}

// ExitParamsDeclaration is called when production paramsDeclaration is exited.
func (s *BasescopepascalListener) ExitParamsDeclaration(ctx *ParamsDeclarationContext) {}

// EnterParamsDeclarationSection is called when production paramsDeclarationSection is entered.
func (s *BasescopepascalListener) EnterParamsDeclarationSection(ctx *ParamsDeclarationSectionContext) {
}

// ExitParamsDeclarationSection is called when production paramsDeclarationSection is exited.
func (s *BasescopepascalListener) ExitParamsDeclarationSection(ctx *ParamsDeclarationSectionContext) {
}

// EnterParamSpecifier is called when production paramSpecifier is entered.
func (s *BasescopepascalListener) EnterParamSpecifier(ctx *ParamSpecifierContext) {}

// ExitParamSpecifier is called when production paramSpecifier is exited.
func (s *BasescopepascalListener) ExitParamSpecifier(ctx *ParamSpecifierContext) {}

// EnterVarSection is called when production varSection is entered.
func (s *BasescopepascalListener) EnterVarSection(ctx *VarSectionContext) {}

// ExitVarSection is called when production varSection is exited.
func (s *BasescopepascalListener) ExitVarSection(ctx *VarSectionContext) {}

// EnterVarDeclaration is called when production varDeclaration is entered.
func (s *BasescopepascalListener) EnterVarDeclaration(ctx *VarDeclarationContext) {}

// ExitVarDeclaration is called when production varDeclaration is exited.
func (s *BasescopepascalListener) ExitVarDeclaration(ctx *VarDeclarationContext) {}

// EnterInlinedVarDeclaration is called when production inlinedVarDeclaration is entered.
func (s *BasescopepascalListener) EnterInlinedVarDeclaration(ctx *InlinedVarDeclarationContext) {}

// ExitInlinedVarDeclaration is called when production inlinedVarDeclaration is exited.
func (s *BasescopepascalListener) ExitInlinedVarDeclaration(ctx *InlinedVarDeclarationContext) {}

// EnterRelationaloperator is called when production relationaloperator is entered.
func (s *BasescopepascalListener) EnterRelationaloperator(ctx *RelationaloperatorContext) {}

// ExitRelationaloperator is called when production relationaloperator is exited.
func (s *BasescopepascalListener) ExitRelationaloperator(ctx *RelationaloperatorContext) {}

// EnterAdditiveoperator is called when production additiveoperator is entered.
func (s *BasescopepascalListener) EnterAdditiveoperator(ctx *AdditiveoperatorContext) {}

// ExitAdditiveoperator is called when production additiveoperator is exited.
func (s *BasescopepascalListener) ExitAdditiveoperator(ctx *AdditiveoperatorContext) {}

// EnterMultiplicativeoperator is called when production multiplicativeoperator is entered.
func (s *BasescopepascalListener) EnterMultiplicativeoperator(ctx *MultiplicativeoperatorContext) {}

// ExitMultiplicativeoperator is called when production multiplicativeoperator is exited.
func (s *BasescopepascalListener) ExitMultiplicativeoperator(ctx *MultiplicativeoperatorContext) {}

// EnterOperator is called when production operator is entered.
func (s *BasescopepascalListener) EnterOperator(ctx *OperatorContext) {}

// ExitOperator is called when production operator is exited.
func (s *BasescopepascalListener) ExitOperator(ctx *OperatorContext) {}

// EnterString is called when production string is entered.
func (s *BasescopepascalListener) EnterString(ctx *StringContext) {}

// ExitString is called when production string is exited.
func (s *BasescopepascalListener) ExitString(ctx *StringContext) {}

// EnterNumber is called when production number is entered.
func (s *BasescopepascalListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BasescopepascalListener) ExitNumber(ctx *NumberContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasescopepascalListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasescopepascalListener) ExitExpression(ctx *ExpressionContext) {}

// EnterTerm is called when production term is entered.
func (s *BasescopepascalListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BasescopepascalListener) ExitTerm(ctx *TermContext) {}

// EnterFunctionExpression is called when production functionExpression is entered.
func (s *BasescopepascalListener) EnterFunctionExpression(ctx *FunctionExpressionContext) {}

// ExitFunctionExpression is called when production functionExpression is exited.
func (s *BasescopepascalListener) ExitFunctionExpression(ctx *FunctionExpressionContext) {}

// EnterProcedureExpression is called when production procedureExpression is entered.
func (s *BasescopepascalListener) EnterProcedureExpression(ctx *ProcedureExpressionContext) {}

// ExitProcedureExpression is called when production procedureExpression is exited.
func (s *BasescopepascalListener) ExitProcedureExpression(ctx *ProcedureExpressionContext) {}

// EnterFunctionDesignator is called when production functionDesignator is entered.
func (s *BasescopepascalListener) EnterFunctionDesignator(ctx *FunctionDesignatorContext) {}

// ExitFunctionDesignator is called when production functionDesignator is exited.
func (s *BasescopepascalListener) ExitFunctionDesignator(ctx *FunctionDesignatorContext) {}

// EnterErrorExpression is called when production errorExpression is entered.
func (s *BasescopepascalListener) EnterErrorExpression(ctx *ErrorExpressionContext) {}

// ExitErrorExpression is called when production errorExpression is exited.
func (s *BasescopepascalListener) ExitErrorExpression(ctx *ErrorExpressionContext) {}
