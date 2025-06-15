// Code generated from /home/tomas/development/tomas303/projects/pascallsp/palsp/internal/pascal.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // pascal

import "github.com/antlr4-go/antlr/v4"

type BasepascalVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasepascalVisitor) VisitSource(ctx *SourceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitUnit(ctx *UnitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitInterfaceSection(ctx *InterfaceSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitImplementationSection(ctx *ImplementationSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitInitializationSection(ctx *InitializationSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitFinalizationSection(ctx *FinalizationSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitIdentifierPart(ctx *IdentifierPartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitInterfaceBlock(ctx *InterfaceBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitErrorInterfaceBlockPart(ctx *ErrorInterfaceBlockPartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitImplementationBlock(ctx *ImplementationBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitUsesUnits(ctx *UsesUnitsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitLabelDeclarationPart(ctx *LabelDeclarationPartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitLabel(ctx *LabelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitConstantDefinitionPart(ctx *ConstantDefinitionPartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitConstantDefinition(ctx *ConstantDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitConstantChr(ctx *ConstantChrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitHexConstant(ctx *HexConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitConstant(ctx *ConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitArrayConstant(ctx *ArrayConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitRecordConstant(ctx *RecordConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitRecordField(ctx *RecordFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitUnsignedNumber(ctx *UnsignedNumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitUnsignedInteger(ctx *UnsignedIntegerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitUnsignedReal(ctx *UnsignedRealContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitSign(ctx *SignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitBool_(ctx *Bool_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitString(ctx *StringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitStringExpression(ctx *StringExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitResourceDefinitionPart(ctx *ResourceDefinitionPartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitResourceDefinition(ctx *ResourceDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitDeprecatedHint(ctx *DeprecatedHintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitTypeDefinitionPart(ctx *TypeDefinitionPartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitTypeDefinition(ctx *TypeDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitForwardDeclaration(ctx *ForwardDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitClassType(ctx *ClassTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitInterfaceType(ctx *InterfaceTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitFunctionType(ctx *FunctionTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitProcedureType(ctx *ProcedureTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitMetaClassType(ctx *MetaClassTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitAliasDistinctType(ctx *AliasDistinctTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitAliasType(ctx *AliasTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitClassImplementsInterfaces(ctx *ClassImplementsInterfacesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitAccessSpecifier(ctx *AccessSpecifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitClassDeclaration(ctx *ClassDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitClassImplicitPublishedDeclaration(ctx *ClassImplicitPublishedDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitClassDeclarationPart(ctx *ClassDeclarationPartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitInterfaceGuidConst(ctx *InterfaceGuidConstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitInterfaceDeclaration(ctx *InterfaceDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitInterfaceDeclarationPart(ctx *InterfaceDeclarationPartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitErrorInterfaceDeclarationPart(ctx *ErrorInterfaceDeclarationPartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitErrorClassDeclarationPart(ctx *ErrorClassDeclarationPartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitPropertyDeclaration(ctx *PropertyDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitPropertyReadDeclaration(ctx *PropertyReadDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitPropertyWriteDeclaration(ctx *PropertyWriteDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitPropertyDefaultValueDeclaration(ctx *PropertyDefaultValueDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitPropertyIndexDeclaration(ctx *PropertyIndexDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitPropertyIndexParameters(ctx *PropertyIndexParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitPropertyIndexParametersList(ctx *PropertyIndexParametersListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitGenericTemplate(ctx *GenericTemplateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitGenericTemplateList(ctx *GenericTemplateListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitGenericTypeParameter(ctx *GenericTypeParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitGenericConstraints(ctx *GenericConstraintsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitGenericConstraint(ctx *GenericConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitType_(ctx *Type_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitSimpleType(ctx *SimpleTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitScalarType(ctx *ScalarTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitScalerList(ctx *ScalerListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitScalerMember(ctx *ScalerMemberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitSubrangeType(ctx *SubrangeTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitTypeIdentifier(ctx *TypeIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitStructuredType(ctx *StructuredTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitUnpackedStructuredType(ctx *UnpackedStructuredTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitStringtype(ctx *StringtypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitArrayType(ctx *ArrayTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitTypeList(ctx *TypeListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitIndexType(ctx *IndexTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitRecordType(ctx *RecordTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitRecordDeclaration(ctx *RecordDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitRecordImplicitPublishedDeclaration(ctx *RecordImplicitPublishedDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitRecordDeclarationPart(ctx *RecordDeclarationPartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitErrorRecordDeclarationPart(ctx *ErrorRecordDeclarationPartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitRecordParts(ctx *RecordPartsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitRecordFixedPart(ctx *RecordFixedPartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitRecordVariantPart(ctx *RecordVariantPartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitTag(ctx *TagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitRecordVariant(ctx *RecordVariantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitHelperType(ctx *HelperTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitHelperDeclaration(ctx *HelperDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitHelperImplicitPublishedDeclaration(ctx *HelperImplicitPublishedDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitHelperDeclarationPart(ctx *HelperDeclarationPartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitErrorHelperDeclarationPart(ctx *ErrorHelperDeclarationPartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitSetType(ctx *SetTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitFileType(ctx *FileTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitPointerType(ctx *PointerTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitVariableDeclarationPart(ctx *VariableDeclarationPartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitProcedureHeader(ctx *ProcedureHeaderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitFunctionHeader(ctx *FunctionHeaderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitProcedureOrFunctionHeader(ctx *ProcedureOrFunctionHeaderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitProcedureOrFunctionHeaderModifiers(ctx *ProcedureOrFunctionHeaderModifiersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitProcedureOrFunctionDeclaration(ctx *ProcedureOrFunctionDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitProcedureDeclaration(ctx *ProcedureDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitProcedureLambdaDeclaration(ctx *ProcedureLambdaDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitFunctionLambdaDeclaration(ctx *FunctionLambdaDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitResultType(ctx *ResultTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitProcedureOrFunctionBody(ctx *ProcedureOrFunctionBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitClassOperatorHeader(ctx *ClassOperatorHeaderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitClassOperatorDeclaration(ctx *ClassOperatorDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitFormalParameterList(ctx *FormalParameterListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitFormalParameterSection(ctx *FormalParameterSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitParameterGroup(ctx *ParameterGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitIdentifierList(ctx *IdentifierListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitConstList(ctx *ConstListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitDefaultValue(ctx *DefaultValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitTypedIdentifierList(ctx *TypedIdentifierListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitErrorStatement(ctx *ErrorStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitUnlabelledStatement(ctx *UnlabelledStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitSimpleStatement(ctx *SimpleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitAssignmentStatement(ctx *AssignmentStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitRaiseExceptionStatement(ctx *RaiseExceptionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitVariableDeclarationStatement(ctx *VariableDeclarationStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitVariableDesignator(ctx *VariableDesignatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitTypeCast(ctx *TypeCastContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitPropertyDesignator(ctx *PropertyDesignatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitRelationaloperator(ctx *RelationaloperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitSimpleExpression(ctx *SimpleExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitAdditiveoperator(ctx *AdditiveoperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitTerm(ctx *TermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitMultiplicativeoperator(ctx *MultiplicativeoperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitSignedFactor(ctx *SignedFactorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitFactor(ctx *FactorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitUnsignedConstant(ctx *UnsignedConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitFunctionDesignator(ctx *FunctionDesignatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitDefaultDesignator(ctx *DefaultDesignatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitParameterList(ctx *ParameterListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitSet_(ctx *Set_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitElementList(ctx *ElementListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitElement(ctx *ElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitProcedureStatement(ctx *ProcedureStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitMethodCallStatement(ctx *MethodCallStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitActualParameter(ctx *ActualParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitParameterwidth(ctx *ParameterwidthContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitGotoStatement(ctx *GotoStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitInheritedStatement(ctx *InheritedStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitEmptyStatement_(ctx *EmptyStatement_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitEmpty_(ctx *Empty_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitStructuredStatement(ctx *StructuredStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitCompoundStatement(ctx *CompoundStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitStatements(ctx *StatementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitConditionalStatement(ctx *ConditionalStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitIfStatement(ctx *IfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitCaseStatement(ctx *CaseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitCaseListElement(ctx *CaseListElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitRepetetiveStatement(ctx *RepetetiveStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitWhileStatement(ctx *WhileStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitRepeatStatement(ctx *RepeatStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitForStatement(ctx *ForStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitForList(ctx *ForListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitInitialValue(ctx *InitialValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitFinalValue(ctx *FinalValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitWithStatement(ctx *WithStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitTryExceptStatement(ctx *TryExceptStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitExceptionCase(ctx *ExceptionCaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitExceptionElse(ctx *ExceptionElseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitTryFinallyStatement(ctx *TryFinallyStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitWithStatementVariableList(ctx *WithStatementVariableListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitAttributeSection(ctx *AttributeSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitAttributeList(ctx *AttributeListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasepascalVisitor) VisitAttributeItem(ctx *AttributeItemContext) interface{} {
	return v.VisitChildren(ctx)
}
