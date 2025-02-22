// Code generated from /home/tomas/development/tomas303/projects/palsp/internal/scopeparser/scopepascal.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // scopepascal

import "github.com/antlr4-go/antlr/v4"

// scopepascalListener is a complete listener for a parse tree produced by scopepascalParser.
type scopepascalListener interface {
	antlr.ParseTreeListener

	// EnterSource is called when entering the source production.
	EnterSource(c *SourceContext)

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterUnit is called when entering the unit production.
	EnterUnit(c *UnitContext)

	// EnterInterfaceSection is called when entering the interfaceSection production.
	EnterInterfaceSection(c *InterfaceSectionContext)

	// EnterImplementationSection is called when entering the implementationSection production.
	EnterImplementationSection(c *ImplementationSectionContext)

	// EnterInitializationSection is called when entering the initializationSection production.
	EnterInitializationSection(c *InitializationSectionContext)

	// EnterFinalizationSection is called when entering the finalizationSection production.
	EnterFinalizationSection(c *FinalizationSectionContext)

	// EnterInterfaceBlock is called when entering the interfaceBlock production.
	EnterInterfaceBlock(c *InterfaceBlockContext)

	// EnterImplementationBlock is called when entering the implementationBlock production.
	EnterImplementationBlock(c *ImplementationBlockContext)

	// EnterUnitList is called when entering the unitList production.
	EnterUnitList(c *UnitListContext)

	// EnterLabelDeclaration is called when entering the labelDeclaration production.
	EnterLabelDeclaration(c *LabelDeclarationContext)

	// EnterConstSection is called when entering the constSection production.
	EnterConstSection(c *ConstSectionContext)

	// EnterResourceSection is called when entering the resourceSection production.
	EnterResourceSection(c *ResourceSectionContext)

	// EnterTypeSection is called when entering the typeSection production.
	EnterTypeSection(c *TypeSectionContext)

	// EnterTypeBlock is called when entering the typeBlock production.
	EnterTypeBlock(c *TypeBlockContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterAccessSpecifier is called when entering the accessSpecifier production.
	EnterAccessSpecifier(c *AccessSpecifierContext)

	// EnterProcedureOrFunctionModifiers is called when entering the procedureOrFunctionModifiers production.
	EnterProcedureOrFunctionModifiers(c *ProcedureOrFunctionModifiersContext)

	// EnterClassForwardDeclaration is called when entering the classForwardDeclaration production.
	EnterClassForwardDeclaration(c *ClassForwardDeclarationContext)

	// EnterClassType is called when entering the classType production.
	EnterClassType(c *ClassTypeContext)

	// EnterRecordType is called when entering the recordType production.
	EnterRecordType(c *RecordTypeContext)

	// EnterPropertyDeclaration is called when entering the propertyDeclaration production.
	EnterPropertyDeclaration(c *PropertyDeclarationContext)

	// EnterPropertyReadDeclaration is called when entering the propertyReadDeclaration production.
	EnterPropertyReadDeclaration(c *PropertyReadDeclarationContext)

	// EnterPropertyWriteDeclaration is called when entering the propertyWriteDeclaration production.
	EnterPropertyWriteDeclaration(c *PropertyWriteDeclarationContext)

	// EnterPropertyDefaultValueDeclaration is called when entering the propertyDefaultValueDeclaration production.
	EnterPropertyDefaultValueDeclaration(c *PropertyDefaultValueDeclarationContext)

	// EnterPropertyIndexDeclaration is called when entering the propertyIndexDeclaration production.
	EnterPropertyIndexDeclaration(c *PropertyIndexDeclarationContext)

	// EnterPropertyIndexParameters is called when entering the propertyIndexParameters production.
	EnterPropertyIndexParameters(c *PropertyIndexParametersContext)

	// EnterPropertyIndexParametersList is called when entering the propertyIndexParametersList production.
	EnterPropertyIndexParametersList(c *PropertyIndexParametersListContext)

	// EnterArrayType is called when entering the arrayType production.
	EnterArrayType(c *ArrayTypeContext)

	// EnterPointerType is called when entering the pointerType production.
	EnterPointerType(c *PointerTypeContext)

	// EnterSetType is called when entering the setType production.
	EnterSetType(c *SetTypeContext)

	// EnterFileType is called when entering the fileType production.
	EnterFileType(c *FileTypeContext)

	// EnterScalarType is called when entering the scalarType production.
	EnterScalarType(c *ScalarTypeContext)

	// EnterSubrangeType is called when entering the subrangeType production.
	EnterSubrangeType(c *SubrangeTypeContext)

	// EnterBlockDeclaration is called when entering the blockDeclaration production.
	EnterBlockDeclaration(c *BlockDeclarationContext)

	// EnterFunctionDeclaration is called when entering the functionDeclaration production.
	EnterFunctionDeclaration(c *FunctionDeclarationContext)

	// EnterProcedureDeclaration is called when entering the procedureDeclaration production.
	EnterProcedureDeclaration(c *ProcedureDeclarationContext)

	// EnterFunctionOrProcedureDeclaration is called when entering the functionOrProcedureDeclaration production.
	EnterFunctionOrProcedureDeclaration(c *FunctionOrProcedureDeclarationContext)

	// EnterFunctionOrProcedure is called when entering the functionOrProcedure production.
	EnterFunctionOrProcedure(c *FunctionOrProcedureContext)

	// EnterBlockStatement is called when entering the blockStatement production.
	EnterBlockStatement(c *BlockStatementContext)

	// EnterRecordVariantDeclaration is called when entering the recordVariantDeclaration production.
	EnterRecordVariantDeclaration(c *RecordVariantDeclarationContext)

	// EnterRecordVariant is called when entering the recordVariant production.
	EnterRecordVariant(c *RecordVariantContext)

	// EnterStatementError is called when entering the statementError production.
	EnterStatementError(c *StatementErrorContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterIdentifierList is called when entering the identifierList production.
	EnterIdentifierList(c *IdentifierListContext)

	// EnterExpressionList is called when entering the expressionList production.
	EnterExpressionList(c *ExpressionListContext)

	// EnterParamsDeclaration is called when entering the paramsDeclaration production.
	EnterParamsDeclaration(c *ParamsDeclarationContext)

	// EnterParamsDeclarationSection is called when entering the paramsDeclarationSection production.
	EnterParamsDeclarationSection(c *ParamsDeclarationSectionContext)

	// EnterParamSpecifier is called when entering the paramSpecifier production.
	EnterParamSpecifier(c *ParamSpecifierContext)

	// EnterVarSection is called when entering the varSection production.
	EnterVarSection(c *VarSectionContext)

	// EnterVarDeclaration is called when entering the varDeclaration production.
	EnterVarDeclaration(c *VarDeclarationContext)

	// EnterInlinedVarDeclaration is called when entering the inlinedVarDeclaration production.
	EnterInlinedVarDeclaration(c *InlinedVarDeclarationContext)

	// EnterRelationaloperator is called when entering the relationaloperator production.
	EnterRelationaloperator(c *RelationaloperatorContext)

	// EnterAdditiveoperator is called when entering the additiveoperator production.
	EnterAdditiveoperator(c *AdditiveoperatorContext)

	// EnterMultiplicativeoperator is called when entering the multiplicativeoperator production.
	EnterMultiplicativeoperator(c *MultiplicativeoperatorContext)

	// EnterOperator is called when entering the operator production.
	EnterOperator(c *OperatorContext)

	// EnterString is called when entering the string production.
	EnterString(c *StringContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterFunctionExpression is called when entering the functionExpression production.
	EnterFunctionExpression(c *FunctionExpressionContext)

	// EnterProcedureExpression is called when entering the procedureExpression production.
	EnterProcedureExpression(c *ProcedureExpressionContext)

	// EnterFunctionDesignator is called when entering the functionDesignator production.
	EnterFunctionDesignator(c *FunctionDesignatorContext)

	// EnterErrorExpression is called when entering the errorExpression production.
	EnterErrorExpression(c *ErrorExpressionContext)

	// ExitSource is called when exiting the source production.
	ExitSource(c *SourceContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitUnit is called when exiting the unit production.
	ExitUnit(c *UnitContext)

	// ExitInterfaceSection is called when exiting the interfaceSection production.
	ExitInterfaceSection(c *InterfaceSectionContext)

	// ExitImplementationSection is called when exiting the implementationSection production.
	ExitImplementationSection(c *ImplementationSectionContext)

	// ExitInitializationSection is called when exiting the initializationSection production.
	ExitInitializationSection(c *InitializationSectionContext)

	// ExitFinalizationSection is called when exiting the finalizationSection production.
	ExitFinalizationSection(c *FinalizationSectionContext)

	// ExitInterfaceBlock is called when exiting the interfaceBlock production.
	ExitInterfaceBlock(c *InterfaceBlockContext)

	// ExitImplementationBlock is called when exiting the implementationBlock production.
	ExitImplementationBlock(c *ImplementationBlockContext)

	// ExitUnitList is called when exiting the unitList production.
	ExitUnitList(c *UnitListContext)

	// ExitLabelDeclaration is called when exiting the labelDeclaration production.
	ExitLabelDeclaration(c *LabelDeclarationContext)

	// ExitConstSection is called when exiting the constSection production.
	ExitConstSection(c *ConstSectionContext)

	// ExitResourceSection is called when exiting the resourceSection production.
	ExitResourceSection(c *ResourceSectionContext)

	// ExitTypeSection is called when exiting the typeSection production.
	ExitTypeSection(c *TypeSectionContext)

	// ExitTypeBlock is called when exiting the typeBlock production.
	ExitTypeBlock(c *TypeBlockContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitAccessSpecifier is called when exiting the accessSpecifier production.
	ExitAccessSpecifier(c *AccessSpecifierContext)

	// ExitProcedureOrFunctionModifiers is called when exiting the procedureOrFunctionModifiers production.
	ExitProcedureOrFunctionModifiers(c *ProcedureOrFunctionModifiersContext)

	// ExitClassForwardDeclaration is called when exiting the classForwardDeclaration production.
	ExitClassForwardDeclaration(c *ClassForwardDeclarationContext)

	// ExitClassType is called when exiting the classType production.
	ExitClassType(c *ClassTypeContext)

	// ExitRecordType is called when exiting the recordType production.
	ExitRecordType(c *RecordTypeContext)

	// ExitPropertyDeclaration is called when exiting the propertyDeclaration production.
	ExitPropertyDeclaration(c *PropertyDeclarationContext)

	// ExitPropertyReadDeclaration is called when exiting the propertyReadDeclaration production.
	ExitPropertyReadDeclaration(c *PropertyReadDeclarationContext)

	// ExitPropertyWriteDeclaration is called when exiting the propertyWriteDeclaration production.
	ExitPropertyWriteDeclaration(c *PropertyWriteDeclarationContext)

	// ExitPropertyDefaultValueDeclaration is called when exiting the propertyDefaultValueDeclaration production.
	ExitPropertyDefaultValueDeclaration(c *PropertyDefaultValueDeclarationContext)

	// ExitPropertyIndexDeclaration is called when exiting the propertyIndexDeclaration production.
	ExitPropertyIndexDeclaration(c *PropertyIndexDeclarationContext)

	// ExitPropertyIndexParameters is called when exiting the propertyIndexParameters production.
	ExitPropertyIndexParameters(c *PropertyIndexParametersContext)

	// ExitPropertyIndexParametersList is called when exiting the propertyIndexParametersList production.
	ExitPropertyIndexParametersList(c *PropertyIndexParametersListContext)

	// ExitArrayType is called when exiting the arrayType production.
	ExitArrayType(c *ArrayTypeContext)

	// ExitPointerType is called when exiting the pointerType production.
	ExitPointerType(c *PointerTypeContext)

	// ExitSetType is called when exiting the setType production.
	ExitSetType(c *SetTypeContext)

	// ExitFileType is called when exiting the fileType production.
	ExitFileType(c *FileTypeContext)

	// ExitScalarType is called when exiting the scalarType production.
	ExitScalarType(c *ScalarTypeContext)

	// ExitSubrangeType is called when exiting the subrangeType production.
	ExitSubrangeType(c *SubrangeTypeContext)

	// ExitBlockDeclaration is called when exiting the blockDeclaration production.
	ExitBlockDeclaration(c *BlockDeclarationContext)

	// ExitFunctionDeclaration is called when exiting the functionDeclaration production.
	ExitFunctionDeclaration(c *FunctionDeclarationContext)

	// ExitProcedureDeclaration is called when exiting the procedureDeclaration production.
	ExitProcedureDeclaration(c *ProcedureDeclarationContext)

	// ExitFunctionOrProcedureDeclaration is called when exiting the functionOrProcedureDeclaration production.
	ExitFunctionOrProcedureDeclaration(c *FunctionOrProcedureDeclarationContext)

	// ExitFunctionOrProcedure is called when exiting the functionOrProcedure production.
	ExitFunctionOrProcedure(c *FunctionOrProcedureContext)

	// ExitBlockStatement is called when exiting the blockStatement production.
	ExitBlockStatement(c *BlockStatementContext)

	// ExitRecordVariantDeclaration is called when exiting the recordVariantDeclaration production.
	ExitRecordVariantDeclaration(c *RecordVariantDeclarationContext)

	// ExitRecordVariant is called when exiting the recordVariant production.
	ExitRecordVariant(c *RecordVariantContext)

	// ExitStatementError is called when exiting the statementError production.
	ExitStatementError(c *StatementErrorContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitIdentifierList is called when exiting the identifierList production.
	ExitIdentifierList(c *IdentifierListContext)

	// ExitExpressionList is called when exiting the expressionList production.
	ExitExpressionList(c *ExpressionListContext)

	// ExitParamsDeclaration is called when exiting the paramsDeclaration production.
	ExitParamsDeclaration(c *ParamsDeclarationContext)

	// ExitParamsDeclarationSection is called when exiting the paramsDeclarationSection production.
	ExitParamsDeclarationSection(c *ParamsDeclarationSectionContext)

	// ExitParamSpecifier is called when exiting the paramSpecifier production.
	ExitParamSpecifier(c *ParamSpecifierContext)

	// ExitVarSection is called when exiting the varSection production.
	ExitVarSection(c *VarSectionContext)

	// ExitVarDeclaration is called when exiting the varDeclaration production.
	ExitVarDeclaration(c *VarDeclarationContext)

	// ExitInlinedVarDeclaration is called when exiting the inlinedVarDeclaration production.
	ExitInlinedVarDeclaration(c *InlinedVarDeclarationContext)

	// ExitRelationaloperator is called when exiting the relationaloperator production.
	ExitRelationaloperator(c *RelationaloperatorContext)

	// ExitAdditiveoperator is called when exiting the additiveoperator production.
	ExitAdditiveoperator(c *AdditiveoperatorContext)

	// ExitMultiplicativeoperator is called when exiting the multiplicativeoperator production.
	ExitMultiplicativeoperator(c *MultiplicativeoperatorContext)

	// ExitOperator is called when exiting the operator production.
	ExitOperator(c *OperatorContext)

	// ExitString is called when exiting the string production.
	ExitString(c *StringContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitFunctionExpression is called when exiting the functionExpression production.
	ExitFunctionExpression(c *FunctionExpressionContext)

	// ExitProcedureExpression is called when exiting the procedureExpression production.
	ExitProcedureExpression(c *ProcedureExpressionContext)

	// ExitFunctionDesignator is called when exiting the functionDesignator production.
	ExitFunctionDesignator(c *FunctionDesignatorContext)

	// ExitErrorExpression is called when exiting the errorExpression production.
	ExitErrorExpression(c *ErrorExpressionContext)
}
