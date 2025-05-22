/*
BSD License

Copyright (c) 2013, Tom Everett
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions
are met:

1. Redistributions of source code must retain the above copyright
   notice, this list of conditions and the following disclaimer.
2. Redistributions in binary form must reproduce the above copyright
   notice, this list of conditions and the following disclaimer in the
   documentation and/or other materials provided with the distribution.
3. Neither the name of Tom Everett nor the names of its contributors
   may be used to endorse or promote products derived from this software
   without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
"AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/
/*
Adapted from pascal.g by  Hakki Dogusan, Piet Schoutteten and Marton Papp
*/

// $antlr-format alignTrailingComments true, columnLimit 150, minEmptyLines 1, maxEmptyLinesToKeep 1, reflowComments false, useTab false
// $antlr-format allowShortRulesOnASingleLine false, allowShortBlocksOnASingleLine true, alignSemicolons hanging, alignColons hanging

grammar pascal;

options {
    caseInsensitive = true;
}

source
    : program
    | unit
    ;

program
    : PROGRAM identifier (LPAREN identifierList RPAREN)? SEMI usesUnits? implementationBlock BEGIN statements END DOT EOF
    ;

unit
    : UNIT identifier SEMI 
      interfaceSection
      implementationSection?
      initializationSection?
      finalizationSection?
      END DOT EOF
    ;

interfaceSection
    : INTERFACE usesUnits? interfaceBlock
    ;

implementationSection
    : IMPLEMENTATION usesUnits? implementationBlock
    ;

initializationSection
    : INITIALIZATION statements
    ;

finalizationSection
    : FINALIZATION statements
    ;

identifier
    : identifierPart (DOT identifierPart)*
    ;

identifierPart
    : IDENT
    | INDEX
    | READ
    | WRITE
    ;

interfaceBlock
    : (
        labelDeclarationPart
        | constantDefinitionPart
        | resourceDefinitionPart
        | typeDefinitionPart
        | variableDeclarationPart
        | procedureOrFunctionHeader
    )*
    ;

implementationBlock
    : (
        labelDeclarationPart
        | constantDefinitionPart
        | resourceDefinitionPart
        | typeDefinitionPart
        | variableDeclarationPart
        | procedureOrFunctionDeclaration
        | procedureOrFunctionHeader FORWARD SEMI
    )*
    ;

block
    : (
        labelDeclarationPart
        | constantDefinitionPart
        | variableDeclarationPart
        | procedureOrFunctionDeclaration
    )*
    ;

usesUnits
    : USES identifierList SEMI
    ;

labelDeclarationPart
    : LABEL label (COMMA label)* SEMI
    ;

label
    : unsignedInteger
    ;

constantDefinitionPart
    : CONST (constantDefinition SEMI)+
    ;

constantDefinition
    : identifier (COLON typeIdentifier)? EQUAL constant
    | identifier COLON arrayType EQUAL constant
    ;

constantChr
    : CHR LPAREN unsignedInteger RPAREN
    ;

hexConstant
    : HEX_LITERAL
    ;

constant
    : unsignedNumber
    | sign unsignedNumber
    | identifier
    | sign identifier
    | string
    | constantChr
    | scalarType
    | arrayConstant (PLUS (arrayConstant|identifier))*
    ;

arrayConstant
    : LBRACK constant (COMMA constant)* RBRACK
    ;

unsignedNumber
    : unsignedInteger
    | unsignedReal
    ;

unsignedInteger
    : NUM_INT
    ;

unsignedReal
    : NUM_REAL
    ;

sign
    : PLUS
    | MINUS
    ;

bool_
    : TRUE
    | FALSE
    ;

string
    : (STRING_LITERAL | STRING_CROSSHATCH_LITERAL)+
    ;

resourceDefinitionPart
    : RESOURCESTRING (resourceDefinition)+
    ;

resourceDefinition
    : identifier EQUAL string SEMI
    ;    

typeDefinitionPart
    : TYPE (typeDefinition SEMI)+
    ;

typeDefinition
    : attributeSection? identifier genericTemplate? EQUAL (type_ | functionType | procedureType | forwardClassType | forwardInterfaceType)
    ;

functionType
    : FUNCTION (formalParameterList)? COLON resultType OFOBJECT? procedureOrFunctionHeaderModifiers
    ;

procedureType
    : PROCEDURE (formalParameterList)? OFOBJECT? procedureOrFunctionHeaderModifiers
    ;

forwardClassType
    : CLASS SEMI
    ;

forwardInterfaceType
    : INTERFACE SEMI
    ;

classType
    : CLASS (LPAREN identifier genericTemplate? classImplementsInterfaces RPAREN)? ABSTRACT? classImplicitPublishedDeclaration (classDeclaration)* END
    ;

classImplementsInterfaces
    : (COMMA typeIdentifier)*
    ;

accessSpecifier
    : PRIVATE
    | STRICTPRIVATE
    | PROTECTED
    | STRICTPROTECTED
    | PUBLIC
    | PUBLISHED
    ;

classDeclaration
    : accessSpecifier classDeclarationPart*
    ;

classImplicitPublishedDeclaration
    : classDeclarationPart*
    ;

classDeclarationPart
    : attributeSection? typedIdentifierList SEMI
    | typeDefinitionPart
    | constantDefinitionPart
    // | CLASS? FUNCTION identifier (formalParameterList)? COLON resultType procedureOrFunctionHeaderModifiers SEMI
    // | CLASS? (PROCEDURE| CONSTRUCTOR | DESTRUCTOR) identifier (formalParameterList)? procedureOrFunctionHeaderModifiers SEMI
    | functionHeader
    | procedureHeader
    | propertyDeclaration SEMI (DEFAULT SEMI)?
    | errorClassDeclarationPart SEMI
    ;

GUID_LITERAL
    : '[' '\'' '{' HEX_CHAR_SEQ8 '-' HEX_CHAR_SEQ4 '-' HEX_CHAR_SEQ4 '-' HEX_CHAR_SEQ4 '-' HEX_CHAR_SEQ12 '}' '\'' ']'
    ;

interfaceGuidConst
    : GUID_LITERAL
    ;
   
interfaceType
    : INTERFACE (LPAREN identifier genericTemplate? RPAREN)? GUID_LITERAL? interfaceDeclaration END
    ;

interfaceDeclaration
    : interfaceDeclarationPart*
    ;

interfaceDeclarationPart
    : functionHeader
    | procedureHeader
    | propertyDeclaration SEMI
    | errorInterfaceDeclarationPart SEMI
    ;

errorInterfaceDeclarationPart
    : ~(END)+ // Consume tokens until a likely statement boundary
    ;

errorClassDeclarationPart
    : ~(PRIVATE | STRICTPRIVATE | PROTECTED | STRICTPROTECTED | PUBLIC | PUBLISHED | END)+ // Consume tokens until a likely statement boundary
    ;

propertyDeclaration
    : PROPERTY identifier propertyIndexParameters? COLON typeIdentifier propertyReadDeclaration? propertyWriteDeclaration? propertyDefaultValueDeclaration? propertyIndexDeclaration? (SEMI DEFAULT)?
    | PROPERTY identifier propertyDefaultValueDeclaration?
    ;

propertyReadDeclaration
    : READ identifier
    ;

propertyWriteDeclaration
    : WRITE identifier
    ;

propertyDefaultValueDeclaration 
    : DEFAULT expression
    ;

propertyIndexDeclaration 
    : INDEX unsignedNumber
    ;

propertyIndexParameters
    : LBRACK formalParameterSection (COMMA formalParameterSection)* RBRACK
    ;

propertyIndexParametersList
    : identifierList COLON indexType (SEMI identifierList COLON indexType)*
    ;

genericTemplate
    : '<' genericTemplateList '>'
    ;

genericTemplateList
    : typeIdentifier (COMMA typeIdentifier)*
    ;

type_
    : simpleType
    | structuredType
    | pointerType
    ;

simpleType
    : scalarType
    | subrangeType
    | typeIdentifier genericTemplate?
    | stringtype
    ;

scalarType
    : LPAREN identifierList RPAREN
    ;

subrangeType
    : simpleExpression DOTDOT simpleExpression
    ;

typeIdentifier
    : identifier genericTemplate?
    | (CHAR | BOOLEAN | INTEGER | REAL | STRING | CARDINAL | LONGBOOL | LONGINT)
    // | arrayType
    ;

structuredType
    : PACKED unpackedStructuredType
    | unpackedStructuredType
    | classType
    | interfaceType
    ;

unpackedStructuredType
    : arrayType
    | recordType
    | setType
    | fileType
    ;

stringtype
    : STRING LBRACK (identifier | unsignedNumber) RBRACK
    ;

arrayType
    : ARRAY LBRACK typeList RBRACK OF type_ (EQUAL scalarType)?
    | ARRAY LBRACK2 typeList RBRACK2 OF type_ (EQUAL scalarType)?
    | ARRAY OF CONST
    | ARRAY OF type_
    ;

typeList
    : indexType (COMMA indexType)*
    ;

indexType
    : simpleType
    ;

recordType
    : RECORD recordParts? END
    ;

recordParts
    : recordFixedPart (SEMI recordVariantPart)?
    | recordVariantPart
    ;

recordFixedPart
    : typedIdentifierList (SEMI typedIdentifierList)* SEMI?
    ;

recordVariantPart
    : CASE tag OF recordVariant (SEMI recordVariant)*
    ;

tag
    : identifier COLON typeIdentifier
    | typeIdentifier
    ;

recordVariant
    : constList COLON LPAREN recordParts RPAREN
    ;

setType
    : SET OF simpleType
    ;

fileType
    : FILE OF type_
    | FILE
    ;

pointerType
    : DEREFERENCE typeIdentifier
    ;

variableDeclarationPart
    : VAR variableDeclaration (SEMI variableDeclaration)* SEMI
    ;

variableDeclaration
    : attributeSection? typedIdentifierList (EQUAL simpleExpression)?
    ;

procedureHeader
    : attributeSection? CLASS? (PROCEDURE| CONSTRUCTOR | DESTRUCTOR) identifier (formalParameterList)? procedureOrFunctionHeaderModifiers SEMI
    ;

functionHeader
    : attributeSection? CLASS? FUNCTION identifier (formalParameterList)? COLON resultType procedureOrFunctionHeaderModifiers SEMI
    ;

procedureOrFunctionHeader
    : procedureHeader
    | functionHeader
    ;

procedureOrFunctionHeaderModifiers: (
		SEMI (ABSTRACT | VIRTUAL | OVERRIDE | REINTRODUCE | OVERLOAD | INLINE | STDCALL | CDECL)
	)*;

procedureOrFunctionDeclaration
    : procedureDeclaration
    | functionDeclaration
    ;

procedureDeclaration
    : procedureHeader procedureOrFunctionBody SEMI
    ;

functionDeclaration
    : functionHeader procedureOrFunctionBody SEMI
    ;

procedureLambdaDeclaration
    : PROCEDURE (formalParameterList)? procedureOrFunctionBody
    ;

functionLambdaDeclaration
    : FUNCTION (formalParameterList)? procedureOrFunctionBody COLON resultType
    ;

resultType
    : typeIdentifier
    ;

procedureOrFunctionBody
    : block compoundStatement
    ;

formalParameterList
    : LPAREN formalParameterSection (SEMI formalParameterSection)* RPAREN
    ;

formalParameterSection
    : attributeSection? parameterGroup
    | attributeSection? VAR parameterGroup
    | attributeSection? CONST parameterGroup
    | attributeSection? OUT parameterGroup
    | attributeSection? FUNCTION parameterGroup
    | attributeSection? PROCEDURE parameterGroup
    ;

parameterGroup
    : identifierList (COLON typeIdentifier)? defaultValue?
	// | (VAR | CONST) identifierList
    ;

identifierList
    : identifier (COMMA identifier)*
    ;

constList
    : constant (COMMA constant)*
    ;

defaultValue
    : EQUAL expression
    ;

typedIdentifierList
    : identifierList COLON type_
    ;

statement
    : label COLON unlabelledStatement
    | unlabelledStatement
    | errorStatement
    ;

errorStatement
    : ~('END' | ';')+ // Consume tokens until a likely statement boundary
    ;

unlabelledStatement
    : simpleStatement
    | structuredStatement
    ;

simpleStatement
    : assignmentStatement
    | methodCallStatement
    | procedureStatement
    | gotoStatement
    | inheritedStatement
    | typeCast
    | emptyStatement_
    | raiseExceptionStatement
    | variableDeclarationStatement
    ;

assignmentStatement
    : variableDesignator ASSIGN expression
    | propertyDesignator ASSIGN expression
    ;

raiseExceptionStatement
    : RAISE expression?
    ;

variableDeclarationStatement
    : VAR identifierList (COLON typeDefinition)? (ASSIGN expression)?
    ;

variableDesignator
    : (typeCast | AT identifier | functionDesignator) (
        LBRACK expression (COMMA expression)* RBRACK
        | LBRACK2 expression (COMMA expression)* RBRACK2
        | DOT functionDesignator
        | DEREFERENCE+
    )*
    ;

typeCast
    : typeIdentifier LPAREN expression RPAREN 
    | LPAREN expression AS typeIdentifier  RPAREN 
    ;
  
propertyDesignator
    : functionDesignator (DOT functionDesignator)* DOT identifier (LBRACK expression (COMMA expression)* RBRACK)?
    ;

expression
    : simpleExpression (relationaloperator expression)?
    ;

relationaloperator
    : EQUAL
    | NOT_EQUAL
    | LT
    | LE
    | GE
    | GT
    | IN
    ;

simpleExpression
    : term (additiveoperator simpleExpression)?
    ;

additiveoperator
    : PLUS
    | MINUS
    | OR
    ;

term
    : signedFactor (multiplicativeoperator term)?
    ;

multiplicativeoperator
    : STAR
    | SLASH
    | DIV
    | MOD
    | AND
    | SHR
    | SHL
    | XOR
    ;

signedFactor
    : (PLUS | MINUS)? factor
    ;

factor
    : INHERITED? functionDesignator
    | variableDesignator (AS identifier)?
    | LPAREN expression RPAREN
    | unsignedConstant
    | set_
    | NOT factor
    | bool_
    | factor LBRACK expression (COMMA expression)* RBRACK
    | AT? typeIdentifier (LPAREN expression RPAREN)? (DEREFERENCE)*
    | factor (DOT expression)+
    ;

unsignedConstant
    : unsignedNumber
    | constantChr
    | hexConstant
    | string
    | NIL
    ;

functionDesignator
    : (identifier) (LT typeIdentifier GT)? (LPAREN parameterList RPAREN)?
    ;

parameterList
    : actualParameter (COMMA actualParameter)*
    ;

set_
    : LBRACK elementList RBRACK
    | LBRACK2 elementList RBRACK2
    ;

elementList
    : element (COMMA element)*
    |
    ;

element
    : expression (DOTDOT expression)?
    ;

procedureStatement
    : identifier (LPAREN parameterList RPAREN)?
    ;

methodCallStatement
    : variableDesignator (DOT variableDesignator)* DOT identifier (LPAREN parameterList RPAREN)?
    ;

actualParameter
    : expression parameterwidth*
    | procedureLambdaDeclaration
    | functionLambdaDeclaration
    ;

parameterwidth
    : COLON expression
    ;

gotoStatement
    : GOTO label
    ;

inheritedStatement
    : INHERITED (identifier (LPAREN parameterList RPAREN)?)?
    ;

emptyStatement_
    :
    ;

empty_
    :
    /* empty */
    ;

structuredStatement
    : compoundStatement
    | conditionalStatement
    | repetetiveStatement
    | withStatement
    | tryExceptStatement
    | tryFinallyStatement
    ;

compoundStatement
    : BEGIN statements END
    ;

statements
    : statement (SEMI statement)*
    ;

conditionalStatement
    : ifStatement
    | caseStatement
    ;

ifStatement
    : IF expression THEN statement (ELSE statement)? SEMI?
    ;

caseStatement
    : CASE expression OF caseListElement (SEMI caseListElement)* (SEMI ELSE statements)? SEMI? END
    ;

caseListElement
    : constList COLON statement
    ;

repetetiveStatement
    : whileStatement
    | repeatStatement
    | forStatement
    ;

whileStatement
    : WHILE expression DO statement
    ;

repeatStatement
    : REPEAT statements UNTIL expression
    ;

forStatement
    : FOR identifier ASSIGN forList DO statement
    ;

forList
    : initialValue (TO | DOWNTO) finalValue
    ;

initialValue
    : expression
    ;

finalValue
    : expression
    ;

withStatement
    : WITH withStatementVariableList DO statement
    | WITH expression DO statement
    ;

tryExceptStatement
    : TRY statements EXCEPT statements END
    ;

tryFinallyStatement
    : TRY statements FINALLY statements END
    ;

withStatementVariableList
    : variableDesignator (COMMA variableDesignator)*
    ;

attributeSection
    : LBRACK attributeList RBRACK
    ;

attributeList
    : attributeItem (COMMA attributeItem)*
    ;

attributeItem
    : identifier (LPAREN parameterList RPAREN)?
    ;

AND
    : 'AND'
    ;

ARRAY
    : 'ARRAY'
    ;

BEGIN
    : 'BEGIN'
    ;

BOOLEAN
    : 'BOOLEAN'
    ;

CASE
    : 'CASE'
    ;

CHAR
    : 'CHAR'
    ;

CHR
    : 'CHR'
    ;

CONST
    : 'CONST'
    ;

DIV
    : 'DIV'
    ;

DO
    : 'DO'
    ;

DOWNTO
    : 'DOWNTO'
    ;

ELSE
    : 'ELSE'
    ;

END
    : 'END'
    ;

FILE
    : 'FILE'
    ;

FOR
    : 'FOR'
    ;

FUNCTION
    : 'FUNCTION'
    ;

GOTO
    : 'GOTO'
    ;

IF
    : 'IF'
    ;

IN
    : 'IN'
    ;

INTEGER
    : 'INTEGER'
    ;

LABEL
    : 'LABEL'
    ;

MOD
    : 'MOD'
    ;

NIL
    : 'NIL'
    ;

NOT
    : 'NOT'
    ;

OF
    : 'OF'
    ;

OR
    : 'OR'
    ;

PACKED
    : 'PACKED'
    ;

PROCEDURE
    : 'PROCEDURE'
    ;

PROGRAM
    : 'PROGRAM'
    ;

REAL
    : 'REAL'
    ;

RECORD
    : 'RECORD'
    ;

REPEAT
    : 'REPEAT'
    ;

SET
    : 'SET'
    ;

THEN
    : 'THEN'
    ;

TO
    : 'TO'
    ;

TYPE
    : 'TYPE'
    ;

UNTIL
    : 'UNTIL'
    ;

VAR
    : 'VAR'
    ;

WHILE
    : 'WHILE'
    ;

WITH
    : 'WITH'
    ;

PLUS
    : '+'
    ;

MINUS
    : '-'
    ;

STAR
    : '*'
    ;

SLASH
    : '/'
    ;

ASSIGN
    : ':='
    ;

COMMA
    : ','
    ;

SEMI
    : ';'
    ;

COLON
    : ':'
    ;

EQUAL
    : '='
    ;

NOT_EQUAL
    : '<>'
    ;

LT
    : '<'
    ;

LE
    : '<='
    ;

GE
    : '>='
    ;

GT
    : '>'
    ;

LPAREN
    : '('
    ;

RPAREN
    : ')'
    ;

LBRACK
    : '['
    ;

LBRACK2
    : '(.'
    ;

RBRACK
    : ']'
    ;

RBRACK2
    : '.)'
    ;

DEREFERENCE
    : '^'
    ;

AT
    : '@'
    ;

DOT
    : '.'
    ;

DOTDOT
    : '..'
    ;

LCURLY
    : '{'
    ;

RCURLY
    : '}'
    ;

UNIT
    : 'UNIT'
    ;

INTERFACE
    : 'INTERFACE'
    ;

USES
    : 'USES'
    ;

STRING
    : 'STRING'
    ;

IMPLEMENTATION
    : 'IMPLEMENTATION'
    ;

TRUE
    : 'TRUE'
    ;

FALSE
    : 'FALSE'
    ;

CLASS
    : 'CLASS'
    ;

PRIVATE
    : 'PRIVATE'
    ;

PROTECTED
    : 'PROTECTED'
    ;

PUBLIC
    : 'PUBLIC'
    ;

PUBLISHED
    : 'PUBLISHED'
    ;

STRICTPRIVATE
    : 'STRICT' WHITESPACE 'PRIVATE'
    ;

STRICTPROTECTED
    : 'STRICT' WHITESPACE 'PROTECTED'
    ;

OUT
    : 'OUT'
    ;

PROPERTY
    : 'PROPERTY'
    ;

READ
    : 'READ'
    ;

WRITE
    : 'WRITE'
    ;

DEFAULT
    : 'DEFAULT'
    ;

INDEX
    : 'INDEX'
    ;

AS
    : 'AS'
    ;

TRY
    : 'TRY'
    ;

FINALLY
    : 'FINALLY'
    ;

EXCEPT
    : 'EXCEPT'
    ;

INITIALIZATION
    : 'INITIALIZATION'
    ;

FINALIZATION
    : 'FINALIZATION'
    ;

OFOBJECT
    : 'OF' WHITESPACE 'OBJECT'
    ;

INHERITED
    : 'INHERITED'
    ;

ABSTRACT
    : 'ABSTRACT'
    ;

REINTRODUCE
    : 'REINTRODUCE'
    ;

VIRTUAL
    : 'VIRTUAL'
    ;

OVERRIDE
    : 'OVERRIDE'
    ;    

OVERLOAD
    : 'OVERLOAD'
    ;

INLINE
    : 'INLINE'
    ;

CDECL
    : 'CDECL'
    ;

STDCALL
    : 'stdcall'
    ;  

CONSTRUCTOR
    : 'CONSTRUCTOR'
    ;

DESTRUCTOR
    : 'DESTRUCTOR'
    ;

RESOURCESTRING
    : 'resourcestring'
    ;

FORWARD
    : 'FORWARD'
    ;

RAISE
    : 'RAISE'
    ;

SHR
    : 'SHR'
    ;

SHL
    : 'SHL'
    ;

XOR
    : 'XOR'
    ;

CARDINAL
    : 'Cardinal'
    ;

LONGBOOL
    : 'LONGBOOL'
    ;

LONGINT
    : 'LONGINT'
    ;


fragment WHITESPACE : [ \t\r\n]+ ;

WS
    : [ \t\r\n] -> skip
    ;

COMMENT_1
    : '(*' .*? '*)' -> skip
    ;

COMMENT_2
    : '{' .*? '}' -> skip
    ;

COMMENT_3
    : '//' ~[\r\n]* -> skip
    ;

IDENT
    : ('A' .. 'Z' | '_') ('A' .. 'Z' | '0' .. '9' | '_')*
    ;

HEX_CHAR
    : ('A' .. 'F' | '0' .. '9')
    ;

HEX_LITERAL
    : '$' HEX_CHAR+
    ;

STRING_LITERAL
    : '\'' ('\'\'' | ~ ('\''))* '\''
    ;

STRING_CROSSHATCH_LITERAL
    : '#' ([0-9]+ | HEX_LITERAL)
    ;

NUM_INT
    : ('0' .. '9')+
    ;

NUM_REAL
    : ('0' .. '9')+ (('.' ('0' .. '9')+ (EXPONENT)?)? | EXPONENT)
    ;

fragment EXPONENT
    : ('E') ('+' | '-')? ('0' .. '9')+
    ;

UTF8BOM
    : '\uFEFF' -> skip
    ;

fragment HEX_CHAR_SEQ12
    : HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR
    ;

fragment HEX_CHAR_SEQ8
    : HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR
    ;

fragment HEX_CHAR_SEQ4
    : HEX_CHAR HEX_CHAR HEX_CHAR HEX_CHAR
    ;
