// Code generated from /home/tomas/development/tomas303/projects/palsp/internal/scopeparser/scopepascal.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // scopepascal

import "github.com/antlr4-go/antlr/v4"

type BasescopepascalVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasescopepascalVisitor) VisitSource(ctx *SourceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitUnit(ctx *UnitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitInterfaceSection(ctx *InterfaceSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitImplementationSection(ctx *ImplementationSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitInitializationSection(ctx *InitializationSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitFinalizationSection(ctx *FinalizationSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitInterfaceBlock(ctx *InterfaceBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitImplementationBlock(ctx *ImplementationBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitUnitList(ctx *UnitListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitLabelDeclaration(ctx *LabelDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitConstSection(ctx *ConstSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitResourceSection(ctx *ResourceSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitTypeSection(ctx *TypeSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitTypeBlock(ctx *TypeBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitType(ctx *TypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitAccessSpecifier(ctx *AccessSpecifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitProcedureOrFunctionModifiers(ctx *ProcedureOrFunctionModifiersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitClassForwardDeclaration(ctx *ClassForwardDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitClassType(ctx *ClassTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitRecordType(ctx *RecordTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitPropertyDeclaration(ctx *PropertyDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitPropertyReadDeclaration(ctx *PropertyReadDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitPropertyWriteDeclaration(ctx *PropertyWriteDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitPropertyDefaultValueDeclaration(ctx *PropertyDefaultValueDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitPropertyIndexDeclaration(ctx *PropertyIndexDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitPropertyIndexParameters(ctx *PropertyIndexParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitPropertyIndexParametersList(ctx *PropertyIndexParametersListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitArrayType(ctx *ArrayTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitPointerType(ctx *PointerTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitSetType(ctx *SetTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitFileType(ctx *FileTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitScalarType(ctx *ScalarTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitSubrangeType(ctx *SubrangeTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitBlockDeclaration(ctx *BlockDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitProcedureDeclaration(ctx *ProcedureDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitFunctionOrProcedureDeclaration(ctx *FunctionOrProcedureDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitFunctionOrProcedure(ctx *FunctionOrProcedureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitBlockStatement(ctx *BlockStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitRecordVariantDeclaration(ctx *RecordVariantDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitRecordVariant(ctx *RecordVariantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitStatementError(ctx *StatementErrorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitIdentifierList(ctx *IdentifierListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitExpressionList(ctx *ExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitParamsDeclaration(ctx *ParamsDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitParamsDeclarationSection(ctx *ParamsDeclarationSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitParamSpecifier(ctx *ParamSpecifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitVarSection(ctx *VarSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitVarDeclaration(ctx *VarDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitInlinedVarDeclaration(ctx *InlinedVarDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitRelationaloperator(ctx *RelationaloperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitAdditiveoperator(ctx *AdditiveoperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitMultiplicativeoperator(ctx *MultiplicativeoperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitOperator(ctx *OperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitString(ctx *StringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitNumber(ctx *NumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitTerm(ctx *TermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitFunctionExpression(ctx *FunctionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitProcedureExpression(ctx *ProcedureExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitFunctionDesignator(ctx *FunctionDesignatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasescopepascalVisitor) VisitErrorExpression(ctx *ErrorExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}
