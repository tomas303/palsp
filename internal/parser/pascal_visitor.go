// Code generated from /home/tomas/development/tomas303/projects/pascallsp/palsp/internal/pascal.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // pascal

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by pascalParser.
type pascalVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by pascalParser#source.
	VisitSource(ctx *SourceContext) interface{}

	// Visit a parse tree produced by pascalParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by pascalParser#unit.
	VisitUnit(ctx *UnitContext) interface{}

	// Visit a parse tree produced by pascalParser#interfaceSection.
	VisitInterfaceSection(ctx *InterfaceSectionContext) interface{}

	// Visit a parse tree produced by pascalParser#implementationSection.
	VisitImplementationSection(ctx *ImplementationSectionContext) interface{}

	// Visit a parse tree produced by pascalParser#initializationSection.
	VisitInitializationSection(ctx *InitializationSectionContext) interface{}

	// Visit a parse tree produced by pascalParser#finalizationSection.
	VisitFinalizationSection(ctx *FinalizationSectionContext) interface{}

	// Visit a parse tree produced by pascalParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}

	// Visit a parse tree produced by pascalParser#identifierPart.
	VisitIdentifierPart(ctx *IdentifierPartContext) interface{}

	// Visit a parse tree produced by pascalParser#interfaceBlock.
	VisitInterfaceBlock(ctx *InterfaceBlockContext) interface{}

	// Visit a parse tree produced by pascalParser#errorInterfaceBlockPart.
	VisitErrorInterfaceBlockPart(ctx *ErrorInterfaceBlockPartContext) interface{}

	// Visit a parse tree produced by pascalParser#implementationBlock.
	VisitImplementationBlock(ctx *ImplementationBlockContext) interface{}

	// Visit a parse tree produced by pascalParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by pascalParser#usesUnits.
	VisitUsesUnits(ctx *UsesUnitsContext) interface{}

	// Visit a parse tree produced by pascalParser#labelDeclarationPart.
	VisitLabelDeclarationPart(ctx *LabelDeclarationPartContext) interface{}

	// Visit a parse tree produced by pascalParser#label.
	VisitLabel(ctx *LabelContext) interface{}

	// Visit a parse tree produced by pascalParser#constantDefinitionPart.
	VisitConstantDefinitionPart(ctx *ConstantDefinitionPartContext) interface{}

	// Visit a parse tree produced by pascalParser#constantDefinition.
	VisitConstantDefinition(ctx *ConstantDefinitionContext) interface{}

	// Visit a parse tree produced by pascalParser#constantChr.
	VisitConstantChr(ctx *ConstantChrContext) interface{}

	// Visit a parse tree produced by pascalParser#hexConstant.
	VisitHexConstant(ctx *HexConstantContext) interface{}

	// Visit a parse tree produced by pascalParser#constant.
	VisitConstant(ctx *ConstantContext) interface{}

	// Visit a parse tree produced by pascalParser#arrayConstant.
	VisitArrayConstant(ctx *ArrayConstantContext) interface{}

	// Visit a parse tree produced by pascalParser#recordConstant.
	VisitRecordConstant(ctx *RecordConstantContext) interface{}

	// Visit a parse tree produced by pascalParser#recordField.
	VisitRecordField(ctx *RecordFieldContext) interface{}

	// Visit a parse tree produced by pascalParser#unsignedNumber.
	VisitUnsignedNumber(ctx *UnsignedNumberContext) interface{}

	// Visit a parse tree produced by pascalParser#unsignedInteger.
	VisitUnsignedInteger(ctx *UnsignedIntegerContext) interface{}

	// Visit a parse tree produced by pascalParser#unsignedReal.
	VisitUnsignedReal(ctx *UnsignedRealContext) interface{}

	// Visit a parse tree produced by pascalParser#sign.
	VisitSign(ctx *SignContext) interface{}

	// Visit a parse tree produced by pascalParser#bool_.
	VisitBool_(ctx *Bool_Context) interface{}

	// Visit a parse tree produced by pascalParser#string.
	VisitString(ctx *StringContext) interface{}

	// Visit a parse tree produced by pascalParser#stringExpression.
	VisitStringExpression(ctx *StringExpressionContext) interface{}

	// Visit a parse tree produced by pascalParser#resourceDefinitionPart.
	VisitResourceDefinitionPart(ctx *ResourceDefinitionPartContext) interface{}

	// Visit a parse tree produced by pascalParser#resourceDefinition.
	VisitResourceDefinition(ctx *ResourceDefinitionContext) interface{}

	// Visit a parse tree produced by pascalParser#deprecatedHint.
	VisitDeprecatedHint(ctx *DeprecatedHintContext) interface{}

	// Visit a parse tree produced by pascalParser#typeDefinitionPart.
	VisitTypeDefinitionPart(ctx *TypeDefinitionPartContext) interface{}

	// Visit a parse tree produced by pascalParser#typeDefinition.
	VisitTypeDefinition(ctx *TypeDefinitionContext) interface{}

	// Visit a parse tree produced by pascalParser#classTypeOrForward.
	VisitClassTypeOrForward(ctx *ClassTypeOrForwardContext) interface{}

	// Visit a parse tree produced by pascalParser#interfaceTypeOrForward.
	VisitInterfaceTypeOrForward(ctx *InterfaceTypeOrForwardContext) interface{}

	// Visit a parse tree produced by pascalParser#forwardDeclaration.
	VisitForwardDeclaration(ctx *ForwardDeclarationContext) interface{}

	// Visit a parse tree produced by pascalParser#classType.
	VisitClassType(ctx *ClassTypeContext) interface{}

	// Visit a parse tree produced by pascalParser#interfaceType.
	VisitInterfaceType(ctx *InterfaceTypeContext) interface{}

	// Visit a parse tree produced by pascalParser#functionType.
	VisitFunctionType(ctx *FunctionTypeContext) interface{}

	// Visit a parse tree produced by pascalParser#procedureType.
	VisitProcedureType(ctx *ProcedureTypeContext) interface{}

	// Visit a parse tree produced by pascalParser#metaClassType.
	VisitMetaClassType(ctx *MetaClassTypeContext) interface{}

	// Visit a parse tree produced by pascalParser#aliasDistinctType.
	VisitAliasDistinctType(ctx *AliasDistinctTypeContext) interface{}

	// Visit a parse tree produced by pascalParser#aliasType.
	VisitAliasType(ctx *AliasTypeContext) interface{}

	// Visit a parse tree produced by pascalParser#classImplementsInterfaces.
	VisitClassImplementsInterfaces(ctx *ClassImplementsInterfacesContext) interface{}

	// Visit a parse tree produced by pascalParser#accessSpecifier.
	VisitAccessSpecifier(ctx *AccessSpecifierContext) interface{}

	// Visit a parse tree produced by pascalParser#classDeclaration.
	VisitClassDeclaration(ctx *ClassDeclarationContext) interface{}

	// Visit a parse tree produced by pascalParser#classImplicitPublishedDeclaration.
	VisitClassImplicitPublishedDeclaration(ctx *ClassImplicitPublishedDeclarationContext) interface{}

	// Visit a parse tree produced by pascalParser#classDeclarationPart.
	VisitClassDeclarationPart(ctx *ClassDeclarationPartContext) interface{}

	// Visit a parse tree produced by pascalParser#interfaceGuidConst.
	VisitInterfaceGuidConst(ctx *InterfaceGuidConstContext) interface{}

	// Visit a parse tree produced by pascalParser#interfaceDeclaration.
	VisitInterfaceDeclaration(ctx *InterfaceDeclarationContext) interface{}

	// Visit a parse tree produced by pascalParser#interfaceDeclarationPart.
	VisitInterfaceDeclarationPart(ctx *InterfaceDeclarationPartContext) interface{}

	// Visit a parse tree produced by pascalParser#errorInterfaceDeclarationPart.
	VisitErrorInterfaceDeclarationPart(ctx *ErrorInterfaceDeclarationPartContext) interface{}

	// Visit a parse tree produced by pascalParser#errorClassDeclarationPart.
	VisitErrorClassDeclarationPart(ctx *ErrorClassDeclarationPartContext) interface{}

	// Visit a parse tree produced by pascalParser#propertyDeclaration.
	VisitPropertyDeclaration(ctx *PropertyDeclarationContext) interface{}

	// Visit a parse tree produced by pascalParser#propertyReadDeclaration.
	VisitPropertyReadDeclaration(ctx *PropertyReadDeclarationContext) interface{}

	// Visit a parse tree produced by pascalParser#propertyWriteDeclaration.
	VisitPropertyWriteDeclaration(ctx *PropertyWriteDeclarationContext) interface{}

	// Visit a parse tree produced by pascalParser#propertyDefaultValueDeclaration.
	VisitPropertyDefaultValueDeclaration(ctx *PropertyDefaultValueDeclarationContext) interface{}

	// Visit a parse tree produced by pascalParser#propertyIndexDeclaration.
	VisitPropertyIndexDeclaration(ctx *PropertyIndexDeclarationContext) interface{}

	// Visit a parse tree produced by pascalParser#propertyIndexParameters.
	VisitPropertyIndexParameters(ctx *PropertyIndexParametersContext) interface{}

	// Visit a parse tree produced by pascalParser#propertyIndexParametersList.
	VisitPropertyIndexParametersList(ctx *PropertyIndexParametersListContext) interface{}

	// Visit a parse tree produced by pascalParser#genericTemplate.
	VisitGenericTemplate(ctx *GenericTemplateContext) interface{}

	// Visit a parse tree produced by pascalParser#genericTemplateList.
	VisitGenericTemplateList(ctx *GenericTemplateListContext) interface{}

	// Visit a parse tree produced by pascalParser#genericTypeParameter.
	VisitGenericTypeParameter(ctx *GenericTypeParameterContext) interface{}

	// Visit a parse tree produced by pascalParser#genericConstraints.
	VisitGenericConstraints(ctx *GenericConstraintsContext) interface{}

	// Visit a parse tree produced by pascalParser#genericConstraint.
	VisitGenericConstraint(ctx *GenericConstraintContext) interface{}

	// Visit a parse tree produced by pascalParser#type_.
	VisitType_(ctx *Type_Context) interface{}

	// Visit a parse tree produced by pascalParser#simpleType.
	VisitSimpleType(ctx *SimpleTypeContext) interface{}

	// Visit a parse tree produced by pascalParser#scalarType.
	VisitScalarType(ctx *ScalarTypeContext) interface{}

	// Visit a parse tree produced by pascalParser#scalerList.
	VisitScalerList(ctx *ScalerListContext) interface{}

	// Visit a parse tree produced by pascalParser#scalerMember.
	VisitScalerMember(ctx *ScalerMemberContext) interface{}

	// Visit a parse tree produced by pascalParser#subrangeType.
	VisitSubrangeType(ctx *SubrangeTypeContext) interface{}

	// Visit a parse tree produced by pascalParser#typeIdentifier.
	VisitTypeIdentifier(ctx *TypeIdentifierContext) interface{}

	// Visit a parse tree produced by pascalParser#structuredType.
	VisitStructuredType(ctx *StructuredTypeContext) interface{}

	// Visit a parse tree produced by pascalParser#unpackedStructuredType.
	VisitUnpackedStructuredType(ctx *UnpackedStructuredTypeContext) interface{}

	// Visit a parse tree produced by pascalParser#stringtype.
	VisitStringtype(ctx *StringtypeContext) interface{}

	// Visit a parse tree produced by pascalParser#arrayType.
	VisitArrayType(ctx *ArrayTypeContext) interface{}

	// Visit a parse tree produced by pascalParser#typeList.
	VisitTypeList(ctx *TypeListContext) interface{}

	// Visit a parse tree produced by pascalParser#indexType.
	VisitIndexType(ctx *IndexTypeContext) interface{}

	// Visit a parse tree produced by pascalParser#recordType.
	VisitRecordType(ctx *RecordTypeContext) interface{}

	// Visit a parse tree produced by pascalParser#recordDeclaration.
	VisitRecordDeclaration(ctx *RecordDeclarationContext) interface{}

	// Visit a parse tree produced by pascalParser#recordImplicitPublishedDeclaration.
	VisitRecordImplicitPublishedDeclaration(ctx *RecordImplicitPublishedDeclarationContext) interface{}

	// Visit a parse tree produced by pascalParser#recordDeclarationPart.
	VisitRecordDeclarationPart(ctx *RecordDeclarationPartContext) interface{}

	// Visit a parse tree produced by pascalParser#errorRecordDeclarationPart.
	VisitErrorRecordDeclarationPart(ctx *ErrorRecordDeclarationPartContext) interface{}

	// Visit a parse tree produced by pascalParser#recordParts.
	VisitRecordParts(ctx *RecordPartsContext) interface{}

	// Visit a parse tree produced by pascalParser#recordFixedPart.
	VisitRecordFixedPart(ctx *RecordFixedPartContext) interface{}

	// Visit a parse tree produced by pascalParser#recordVariantPart.
	VisitRecordVariantPart(ctx *RecordVariantPartContext) interface{}

	// Visit a parse tree produced by pascalParser#tag.
	VisitTag(ctx *TagContext) interface{}

	// Visit a parse tree produced by pascalParser#recordVariant.
	VisitRecordVariant(ctx *RecordVariantContext) interface{}

	// Visit a parse tree produced by pascalParser#helperType.
	VisitHelperType(ctx *HelperTypeContext) interface{}

	// Visit a parse tree produced by pascalParser#helperDeclaration.
	VisitHelperDeclaration(ctx *HelperDeclarationContext) interface{}

	// Visit a parse tree produced by pascalParser#helperImplicitPublishedDeclaration.
	VisitHelperImplicitPublishedDeclaration(ctx *HelperImplicitPublishedDeclarationContext) interface{}

	// Visit a parse tree produced by pascalParser#helperDeclarationPart.
	VisitHelperDeclarationPart(ctx *HelperDeclarationPartContext) interface{}

	// Visit a parse tree produced by pascalParser#errorHelperDeclarationPart.
	VisitErrorHelperDeclarationPart(ctx *ErrorHelperDeclarationPartContext) interface{}

	// Visit a parse tree produced by pascalParser#setType.
	VisitSetType(ctx *SetTypeContext) interface{}

	// Visit a parse tree produced by pascalParser#fileType.
	VisitFileType(ctx *FileTypeContext) interface{}

	// Visit a parse tree produced by pascalParser#pointerType.
	VisitPointerType(ctx *PointerTypeContext) interface{}

	// Visit a parse tree produced by pascalParser#variableDeclarationPart.
	VisitVariableDeclarationPart(ctx *VariableDeclarationPartContext) interface{}

	// Visit a parse tree produced by pascalParser#variableDeclaration.
	VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{}

	// Visit a parse tree produced by pascalParser#procedureHeader.
	VisitProcedureHeader(ctx *ProcedureHeaderContext) interface{}

	// Visit a parse tree produced by pascalParser#functionHeader.
	VisitFunctionHeader(ctx *FunctionHeaderContext) interface{}

	// Visit a parse tree produced by pascalParser#procedureOrFunctionHeader.
	VisitProcedureOrFunctionHeader(ctx *ProcedureOrFunctionHeaderContext) interface{}

	// Visit a parse tree produced by pascalParser#procedureOrFunctionHeaderModifiers.
	VisitProcedureOrFunctionHeaderModifiers(ctx *ProcedureOrFunctionHeaderModifiersContext) interface{}

	// Visit a parse tree produced by pascalParser#procedureOrFunctionDeclaration.
	VisitProcedureOrFunctionDeclaration(ctx *ProcedureOrFunctionDeclarationContext) interface{}

	// Visit a parse tree produced by pascalParser#procedureDeclaration.
	VisitProcedureDeclaration(ctx *ProcedureDeclarationContext) interface{}

	// Visit a parse tree produced by pascalParser#functionDeclaration.
	VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{}

	// Visit a parse tree produced by pascalParser#procedureLambdaDeclaration.
	VisitProcedureLambdaDeclaration(ctx *ProcedureLambdaDeclarationContext) interface{}

	// Visit a parse tree produced by pascalParser#functionLambdaDeclaration.
	VisitFunctionLambdaDeclaration(ctx *FunctionLambdaDeclarationContext) interface{}

	// Visit a parse tree produced by pascalParser#resultType.
	VisitResultType(ctx *ResultTypeContext) interface{}

	// Visit a parse tree produced by pascalParser#procedureOrFunctionBody.
	VisitProcedureOrFunctionBody(ctx *ProcedureOrFunctionBodyContext) interface{}

	// Visit a parse tree produced by pascalParser#classOperatorHeader.
	VisitClassOperatorHeader(ctx *ClassOperatorHeaderContext) interface{}

	// Visit a parse tree produced by pascalParser#classOperatorDeclaration.
	VisitClassOperatorDeclaration(ctx *ClassOperatorDeclarationContext) interface{}

	// Visit a parse tree produced by pascalParser#formalParameterList.
	VisitFormalParameterList(ctx *FormalParameterListContext) interface{}

	// Visit a parse tree produced by pascalParser#formalParameterSection.
	VisitFormalParameterSection(ctx *FormalParameterSectionContext) interface{}

	// Visit a parse tree produced by pascalParser#parameterGroup.
	VisitParameterGroup(ctx *ParameterGroupContext) interface{}

	// Visit a parse tree produced by pascalParser#identifierList.
	VisitIdentifierList(ctx *IdentifierListContext) interface{}

	// Visit a parse tree produced by pascalParser#constList.
	VisitConstList(ctx *ConstListContext) interface{}

	// Visit a parse tree produced by pascalParser#defaultValue.
	VisitDefaultValue(ctx *DefaultValueContext) interface{}

	// Visit a parse tree produced by pascalParser#typedIdentifierList.
	VisitTypedIdentifierList(ctx *TypedIdentifierListContext) interface{}

	// Visit a parse tree produced by pascalParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by pascalParser#errorStatement.
	VisitErrorStatement(ctx *ErrorStatementContext) interface{}

	// Visit a parse tree produced by pascalParser#unlabelledStatement.
	VisitUnlabelledStatement(ctx *UnlabelledStatementContext) interface{}

	// Visit a parse tree produced by pascalParser#simpleStatement.
	VisitSimpleStatement(ctx *SimpleStatementContext) interface{}

	// Visit a parse tree produced by pascalParser#assignmentStatement.
	VisitAssignmentStatement(ctx *AssignmentStatementContext) interface{}

	// Visit a parse tree produced by pascalParser#raiseExceptionStatement.
	VisitRaiseExceptionStatement(ctx *RaiseExceptionStatementContext) interface{}

	// Visit a parse tree produced by pascalParser#variableDeclarationStatement.
	VisitVariableDeclarationStatement(ctx *VariableDeclarationStatementContext) interface{}

	// Visit a parse tree produced by pascalParser#variableDesignator.
	VisitVariableDesignator(ctx *VariableDesignatorContext) interface{}

	// Visit a parse tree produced by pascalParser#typeCast.
	VisitTypeCast(ctx *TypeCastContext) interface{}

	// Visit a parse tree produced by pascalParser#propertyDesignator.
	VisitPropertyDesignator(ctx *PropertyDesignatorContext) interface{}

	// Visit a parse tree produced by pascalParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by pascalParser#relationaloperator.
	VisitRelationaloperator(ctx *RelationaloperatorContext) interface{}

	// Visit a parse tree produced by pascalParser#simpleExpression.
	VisitSimpleExpression(ctx *SimpleExpressionContext) interface{}

	// Visit a parse tree produced by pascalParser#additiveoperator.
	VisitAdditiveoperator(ctx *AdditiveoperatorContext) interface{}

	// Visit a parse tree produced by pascalParser#term.
	VisitTerm(ctx *TermContext) interface{}

	// Visit a parse tree produced by pascalParser#multiplicativeoperator.
	VisitMultiplicativeoperator(ctx *MultiplicativeoperatorContext) interface{}

	// Visit a parse tree produced by pascalParser#signedFactor.
	VisitSignedFactor(ctx *SignedFactorContext) interface{}

	// Visit a parse tree produced by pascalParser#factor.
	VisitFactor(ctx *FactorContext) interface{}

	// Visit a parse tree produced by pascalParser#unsignedConstant.
	VisitUnsignedConstant(ctx *UnsignedConstantContext) interface{}

	// Visit a parse tree produced by pascalParser#functionDesignator.
	VisitFunctionDesignator(ctx *FunctionDesignatorContext) interface{}

	// Visit a parse tree produced by pascalParser#defaultDesignator.
	VisitDefaultDesignator(ctx *DefaultDesignatorContext) interface{}

	// Visit a parse tree produced by pascalParser#parameterList.
	VisitParameterList(ctx *ParameterListContext) interface{}

	// Visit a parse tree produced by pascalParser#set_.
	VisitSet_(ctx *Set_Context) interface{}

	// Visit a parse tree produced by pascalParser#elementList.
	VisitElementList(ctx *ElementListContext) interface{}

	// Visit a parse tree produced by pascalParser#element.
	VisitElement(ctx *ElementContext) interface{}

	// Visit a parse tree produced by pascalParser#procedureStatement.
	VisitProcedureStatement(ctx *ProcedureStatementContext) interface{}

	// Visit a parse tree produced by pascalParser#methodCallStatement.
	VisitMethodCallStatement(ctx *MethodCallStatementContext) interface{}

	// Visit a parse tree produced by pascalParser#actualParameter.
	VisitActualParameter(ctx *ActualParameterContext) interface{}

	// Visit a parse tree produced by pascalParser#parameterwidth.
	VisitParameterwidth(ctx *ParameterwidthContext) interface{}

	// Visit a parse tree produced by pascalParser#gotoStatement.
	VisitGotoStatement(ctx *GotoStatementContext) interface{}

	// Visit a parse tree produced by pascalParser#inheritedStatement.
	VisitInheritedStatement(ctx *InheritedStatementContext) interface{}

	// Visit a parse tree produced by pascalParser#emptyStatement_.
	VisitEmptyStatement_(ctx *EmptyStatement_Context) interface{}

	// Visit a parse tree produced by pascalParser#empty_.
	VisitEmpty_(ctx *Empty_Context) interface{}

	// Visit a parse tree produced by pascalParser#structuredStatement.
	VisitStructuredStatement(ctx *StructuredStatementContext) interface{}

	// Visit a parse tree produced by pascalParser#compoundStatement.
	VisitCompoundStatement(ctx *CompoundStatementContext) interface{}

	// Visit a parse tree produced by pascalParser#statements.
	VisitStatements(ctx *StatementsContext) interface{}

	// Visit a parse tree produced by pascalParser#conditionalStatement.
	VisitConditionalStatement(ctx *ConditionalStatementContext) interface{}

	// Visit a parse tree produced by pascalParser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by pascalParser#caseStatement.
	VisitCaseStatement(ctx *CaseStatementContext) interface{}

	// Visit a parse tree produced by pascalParser#caseListElement.
	VisitCaseListElement(ctx *CaseListElementContext) interface{}

	// Visit a parse tree produced by pascalParser#repetetiveStatement.
	VisitRepetetiveStatement(ctx *RepetetiveStatementContext) interface{}

	// Visit a parse tree produced by pascalParser#whileStatement.
	VisitWhileStatement(ctx *WhileStatementContext) interface{}

	// Visit a parse tree produced by pascalParser#repeatStatement.
	VisitRepeatStatement(ctx *RepeatStatementContext) interface{}

	// Visit a parse tree produced by pascalParser#forStatement.
	VisitForStatement(ctx *ForStatementContext) interface{}

	// Visit a parse tree produced by pascalParser#forList.
	VisitForList(ctx *ForListContext) interface{}

	// Visit a parse tree produced by pascalParser#initialValue.
	VisitInitialValue(ctx *InitialValueContext) interface{}

	// Visit a parse tree produced by pascalParser#finalValue.
	VisitFinalValue(ctx *FinalValueContext) interface{}

	// Visit a parse tree produced by pascalParser#withStatement.
	VisitWithStatement(ctx *WithStatementContext) interface{}

	// Visit a parse tree produced by pascalParser#tryExceptStatement.
	VisitTryExceptStatement(ctx *TryExceptStatementContext) interface{}

	// Visit a parse tree produced by pascalParser#exceptionCase.
	VisitExceptionCase(ctx *ExceptionCaseContext) interface{}

	// Visit a parse tree produced by pascalParser#exceptionElse.
	VisitExceptionElse(ctx *ExceptionElseContext) interface{}

	// Visit a parse tree produced by pascalParser#tryFinallyStatement.
	VisitTryFinallyStatement(ctx *TryFinallyStatementContext) interface{}

	// Visit a parse tree produced by pascalParser#withStatementVariableList.
	VisitWithStatementVariableList(ctx *WithStatementVariableListContext) interface{}

	// Visit a parse tree produced by pascalParser#attributeSection.
	VisitAttributeSection(ctx *AttributeSectionContext) interface{}

	// Visit a parse tree produced by pascalParser#attributeList.
	VisitAttributeList(ctx *AttributeListContext) interface{}

	// Visit a parse tree produced by pascalParser#attributeItem.
	VisitAttributeItem(ctx *AttributeItemContext) interface{}
}
