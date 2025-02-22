// Code generated from /home/tomas/development/tomas303/projects/palsp/internal/scopeparser/scopepascal.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // scopepascal

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by scopepascalParser.
type scopepascalVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by scopepascalParser#source.
	VisitSource(ctx *SourceContext) interface{}

	// Visit a parse tree produced by scopepascalParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by scopepascalParser#unit.
	VisitUnit(ctx *UnitContext) interface{}

	// Visit a parse tree produced by scopepascalParser#interfaceSection.
	VisitInterfaceSection(ctx *InterfaceSectionContext) interface{}

	// Visit a parse tree produced by scopepascalParser#implementationSection.
	VisitImplementationSection(ctx *ImplementationSectionContext) interface{}

	// Visit a parse tree produced by scopepascalParser#initializationSection.
	VisitInitializationSection(ctx *InitializationSectionContext) interface{}

	// Visit a parse tree produced by scopepascalParser#finalizationSection.
	VisitFinalizationSection(ctx *FinalizationSectionContext) interface{}

	// Visit a parse tree produced by scopepascalParser#interfaceBlock.
	VisitInterfaceBlock(ctx *InterfaceBlockContext) interface{}

	// Visit a parse tree produced by scopepascalParser#implementationBlock.
	VisitImplementationBlock(ctx *ImplementationBlockContext) interface{}

	// Visit a parse tree produced by scopepascalParser#unitList.
	VisitUnitList(ctx *UnitListContext) interface{}

	// Visit a parse tree produced by scopepascalParser#labelDeclaration.
	VisitLabelDeclaration(ctx *LabelDeclarationContext) interface{}

	// Visit a parse tree produced by scopepascalParser#constSection.
	VisitConstSection(ctx *ConstSectionContext) interface{}

	// Visit a parse tree produced by scopepascalParser#resourceSection.
	VisitResourceSection(ctx *ResourceSectionContext) interface{}

	// Visit a parse tree produced by scopepascalParser#typeSection.
	VisitTypeSection(ctx *TypeSectionContext) interface{}

	// Visit a parse tree produced by scopepascalParser#typeBlock.
	VisitTypeBlock(ctx *TypeBlockContext) interface{}

	// Visit a parse tree produced by scopepascalParser#type.
	VisitType(ctx *TypeContext) interface{}

	// Visit a parse tree produced by scopepascalParser#accessSpecifier.
	VisitAccessSpecifier(ctx *AccessSpecifierContext) interface{}

	// Visit a parse tree produced by scopepascalParser#procedureOrFunctionModifiers.
	VisitProcedureOrFunctionModifiers(ctx *ProcedureOrFunctionModifiersContext) interface{}

	// Visit a parse tree produced by scopepascalParser#classForwardDeclaration.
	VisitClassForwardDeclaration(ctx *ClassForwardDeclarationContext) interface{}

	// Visit a parse tree produced by scopepascalParser#classType.
	VisitClassType(ctx *ClassTypeContext) interface{}

	// Visit a parse tree produced by scopepascalParser#recordType.
	VisitRecordType(ctx *RecordTypeContext) interface{}

	// Visit a parse tree produced by scopepascalParser#propertyDeclaration.
	VisitPropertyDeclaration(ctx *PropertyDeclarationContext) interface{}

	// Visit a parse tree produced by scopepascalParser#propertyReadDeclaration.
	VisitPropertyReadDeclaration(ctx *PropertyReadDeclarationContext) interface{}

	// Visit a parse tree produced by scopepascalParser#propertyWriteDeclaration.
	VisitPropertyWriteDeclaration(ctx *PropertyWriteDeclarationContext) interface{}

	// Visit a parse tree produced by scopepascalParser#propertyDefaultValueDeclaration.
	VisitPropertyDefaultValueDeclaration(ctx *PropertyDefaultValueDeclarationContext) interface{}

	// Visit a parse tree produced by scopepascalParser#propertyIndexDeclaration.
	VisitPropertyIndexDeclaration(ctx *PropertyIndexDeclarationContext) interface{}

	// Visit a parse tree produced by scopepascalParser#propertyIndexParameters.
	VisitPropertyIndexParameters(ctx *PropertyIndexParametersContext) interface{}

	// Visit a parse tree produced by scopepascalParser#propertyIndexParametersList.
	VisitPropertyIndexParametersList(ctx *PropertyIndexParametersListContext) interface{}

	// Visit a parse tree produced by scopepascalParser#arrayType.
	VisitArrayType(ctx *ArrayTypeContext) interface{}

	// Visit a parse tree produced by scopepascalParser#pointerType.
	VisitPointerType(ctx *PointerTypeContext) interface{}

	// Visit a parse tree produced by scopepascalParser#setType.
	VisitSetType(ctx *SetTypeContext) interface{}

	// Visit a parse tree produced by scopepascalParser#fileType.
	VisitFileType(ctx *FileTypeContext) interface{}

	// Visit a parse tree produced by scopepascalParser#scalarType.
	VisitScalarType(ctx *ScalarTypeContext) interface{}

	// Visit a parse tree produced by scopepascalParser#subrangeType.
	VisitSubrangeType(ctx *SubrangeTypeContext) interface{}

	// Visit a parse tree produced by scopepascalParser#blockDeclaration.
	VisitBlockDeclaration(ctx *BlockDeclarationContext) interface{}

	// Visit a parse tree produced by scopepascalParser#functionDeclaration.
	VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{}

	// Visit a parse tree produced by scopepascalParser#procedureDeclaration.
	VisitProcedureDeclaration(ctx *ProcedureDeclarationContext) interface{}

	// Visit a parse tree produced by scopepascalParser#functionOrProcedureDeclaration.
	VisitFunctionOrProcedureDeclaration(ctx *FunctionOrProcedureDeclarationContext) interface{}

	// Visit a parse tree produced by scopepascalParser#functionOrProcedure.
	VisitFunctionOrProcedure(ctx *FunctionOrProcedureContext) interface{}

	// Visit a parse tree produced by scopepascalParser#blockStatement.
	VisitBlockStatement(ctx *BlockStatementContext) interface{}

	// Visit a parse tree produced by scopepascalParser#recordVariantDeclaration.
	VisitRecordVariantDeclaration(ctx *RecordVariantDeclarationContext) interface{}

	// Visit a parse tree produced by scopepascalParser#recordVariant.
	VisitRecordVariant(ctx *RecordVariantContext) interface{}

	// Visit a parse tree produced by scopepascalParser#statementError.
	VisitStatementError(ctx *StatementErrorContext) interface{}

	// Visit a parse tree produced by scopepascalParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by scopepascalParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}

	// Visit a parse tree produced by scopepascalParser#identifierList.
	VisitIdentifierList(ctx *IdentifierListContext) interface{}

	// Visit a parse tree produced by scopepascalParser#expressionList.
	VisitExpressionList(ctx *ExpressionListContext) interface{}

	// Visit a parse tree produced by scopepascalParser#paramsDeclaration.
	VisitParamsDeclaration(ctx *ParamsDeclarationContext) interface{}

	// Visit a parse tree produced by scopepascalParser#paramsDeclarationSection.
	VisitParamsDeclarationSection(ctx *ParamsDeclarationSectionContext) interface{}

	// Visit a parse tree produced by scopepascalParser#paramSpecifier.
	VisitParamSpecifier(ctx *ParamSpecifierContext) interface{}

	// Visit a parse tree produced by scopepascalParser#varSection.
	VisitVarSection(ctx *VarSectionContext) interface{}

	// Visit a parse tree produced by scopepascalParser#varDeclaration.
	VisitVarDeclaration(ctx *VarDeclarationContext) interface{}

	// Visit a parse tree produced by scopepascalParser#inlinedVarDeclaration.
	VisitInlinedVarDeclaration(ctx *InlinedVarDeclarationContext) interface{}

	// Visit a parse tree produced by scopepascalParser#relationaloperator.
	VisitRelationaloperator(ctx *RelationaloperatorContext) interface{}

	// Visit a parse tree produced by scopepascalParser#additiveoperator.
	VisitAdditiveoperator(ctx *AdditiveoperatorContext) interface{}

	// Visit a parse tree produced by scopepascalParser#multiplicativeoperator.
	VisitMultiplicativeoperator(ctx *MultiplicativeoperatorContext) interface{}

	// Visit a parse tree produced by scopepascalParser#operator.
	VisitOperator(ctx *OperatorContext) interface{}

	// Visit a parse tree produced by scopepascalParser#string.
	VisitString(ctx *StringContext) interface{}

	// Visit a parse tree produced by scopepascalParser#number.
	VisitNumber(ctx *NumberContext) interface{}

	// Visit a parse tree produced by scopepascalParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by scopepascalParser#term.
	VisitTerm(ctx *TermContext) interface{}

	// Visit a parse tree produced by scopepascalParser#functionExpression.
	VisitFunctionExpression(ctx *FunctionExpressionContext) interface{}

	// Visit a parse tree produced by scopepascalParser#procedureExpression.
	VisitProcedureExpression(ctx *ProcedureExpressionContext) interface{}

	// Visit a parse tree produced by scopepascalParser#functionDesignator.
	VisitFunctionDesignator(ctx *FunctionDesignatorContext) interface{}

	// Visit a parse tree produced by scopepascalParser#errorExpression.
	VisitErrorExpression(ctx *ErrorExpressionContext) interface{}
}
