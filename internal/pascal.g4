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
    : UNIT identifier (DOT identifier)* SEMI 
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
    ;

constantChr
    : CHR LPAREN unsignedInteger RPAREN
    ;

constant
    : unsignedNumber
    | sign unsignedNumber
    | identifier
    | sign identifier
    | string
    | constantChr
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
    : STRING_LITERAL
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
    : identifier EQUAL (type_ | functionType | procedureType | forwardClassType)
    ;

functionType
    : FUNCTION (formalParameterList)? COLON resultType OFOBJECT?
    ;

procedureType
    : PROCEDURE (formalParameterList)? OFOBJECT?
    ;

forwardClassType
    : CLASS
    ;

classType
    : CLASS LPAREN identifier classImplementsInterfaces RPAREN classImplicitPublishedDeclaration (classDeclaration)* END
    ;

classImplementsInterfaces
    : (COMMA typeIdentifier)*
    ;

classDeclaration
    : classPrivateDeclaration
    | classStrictPrivateDeclaration
    | classProtectedDeclaration
    | classStrictProtectedDeclaration
    | classPublicDeclaration
    | classPublishedDeclaration
    ;    

classPrivateDeclaration
    : PRIVATE classDeclarationPart*
    ;

classStrictPrivateDeclaration
    : STRICTPRIVATE classDeclarationPart*
    ;

classProtectedDeclaration
    : PROTECTED classDeclarationPart*
    ;

classStrictProtectedDeclaration
    : STRICTPROTECTED classDeclarationPart*
    ;

classPublicDeclaration
    : PUBLIC classDeclarationPart*
    ;

classPublishedDeclaration
    : PUBLISHED classDeclarationPart*
    ;

classImplicitPublishedDeclaration
    : classDeclarationPart*
    ;

classDeclarationPart
    : identifierList COLON typeIdentifier SEMI
    | typeDefinitionPart
    | constantDefinitionPart
    | FUNCTION identifier (formalParameterList)? COLON resultType SEMI procedureOrFunctionHeaderModifiers
    | (PROCEDURE| CONSTRUCTOR | DESTRUCTOR) identifier (formalParameterList)? SEMI procedureOrFunctionHeaderModifiers
    | propertyDeclaration SEMI (DEFAULT SEMI)?
    ;

propertyDeclaration
    : PROPERTY identifier propertyIndexParameters? COLON typeIdentifier propertyReadDeclaration? propertyWriteDeclaration? propertyDefaultValueDeclaration? propertyIndexDeclaration?
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
    : LBRACK propertyIndexParametersList RBRACK
    ;

propertyIndexParametersList
    : identifierList COLON indexType (SEMI identifierList COLON indexType)*
    ;

methodIdentifier
    : (typeIdentifier DOT)+ identifier
    ;

type_
    : simpleType
    | structuredType
    | pointerType
    ;

simpleType
    : scalarType
    | subrangeType
    | typeIdentifier
    | stringtype
    ;

scalarType
    : LPAREN identifierList RPAREN
    ;

subrangeType
    : constant DOTDOT constant
    ;

typeIdentifier
    : identifier
    | (CHAR | BOOLEAN | INTEGER | REAL | STRING | LONGBOOL | CARDINAL)
    | identifier LT typeIdentifier GT
    | arrayType
    ;

structuredType
    : PACKED unpackedStructuredType
    | unpackedStructuredType
    | classType
    ;

unpackedStructuredType
    : recordType
    | setType
    | fileType
    ;

stringtype
    : STRING LBRACK (identifier | unsignedNumber) RBRACK
    ;

arrayType
    : ARRAY LBRACK typeList RBRACK OF componentType
    | ARRAY LBRACK2 typeList RBRACK2 OF componentType
    | ARRAY OF CONST
    | ARRAY OF componentType
    ;

typeList
    : indexType (COMMA indexType)*
    ;

indexType
    : simpleType
    ;

componentType
    : type_
    ;

recordType
    : RECORD fieldList? END
    ;

fieldList
    : fixedPart (SEMI variantPart)?
    | variantPart
    ;

fixedPart
    : recordSection (SEMI recordSection)* SEMI?
    ;

recordSection
    : identifierList COLON type_
    ;

variantPart
    : CASE tag OF variant (SEMI variant)*
    ;

tag
    : identifier COLON typeIdentifier
    | typeIdentifier
    ;

variant
    : constList COLON LPAREN fieldList RPAREN
    ;

setType
    : SET OF baseType
    ;

baseType
    : simpleType
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
    : identifierList COLON type_
    ;

procedureHeader
    : (PROCEDURE| CONSTRUCTOR | DESTRUCTOR) (identifier|methodIdentifier) (formalParameterList)? SEMI procedureOrFunctionHeaderModifiers
    ;

functionHeader
    : FUNCTION (identifier|methodIdentifier) (formalParameterList)? COLON resultType SEMI procedureOrFunctionHeaderModifiers
    ;

procedureOrFunctionHeader
    : procedureHeader
    | functionHeader
    ;

procedureOrFunctionHeaderModifiers: (
		(ABSTRACT | VIRTUAL | OVERRIDE | OVERLOAD | INLINE) SEMI
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
    : parameterGroup
    | VAR parameterGroup
    | CONST parameterGroup
    | OUT parameterGroup
    | FUNCTION parameterGroup
    | PROCEDURE parameterGroup
    ;

parameterGroup
    : identifierList COLON typeIdentifier defaultValue?
	| (VAR | CONST) identifierList;

identifierList
    : identifier (COMMA identifier)*
    ;

constList
    : constant (COMMA constant)*
    ;

defaultValue
    : EQUAL expression
    ;

statement
    : label COLON unlabelledStatement
    | unlabelledStatement
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
    | emptyStatement_
    | raiseExceptionStatement
    ;

assignmentStatement
    : variable ASSIGN expression
    ;

raiseExceptionStatement
    : RAISE expression?
    ;

variable
    : (typeCast | AT identifier | identifier) (
        LBRACK expression (COMMA expression)* RBRACK
        | LBRACK2 expression (COMMA expression)* RBRACK2
        | DOT identifier
        | DEREFERENCE+
    )*
    ;

typeCast
    : typeIdentifier LPAREN expression RPAREN 
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
    ;

signedFactor
    : (PLUS | MINUS)? factor
    ;

factor
    : variable (AS identifier)?
    | LPAREN expression RPAREN
    | INHERITED? functionDesignator
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
    | string
    | NIL
    ;

functionDesignator
    : (identifier|methodIdentifier) (LT typeIdentifier GT)? (LPAREN parameterList RPAREN)?
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
    | methodIdentifier (LPAREN parameterList RPAREN)?
    ;

methodCallStatement
    : variable (DOT variable)* DOT identifier (LPAREN parameterList RPAREN)?
    ;

actualParameter
    : expression parameterwidth*
    ;

parameterwidth
    : COLON expression
    ;

gotoStatement
    : GOTO label
    ;

inheritedStatement
    : INHERITED (identifier|methodIdentifier) (LPAREN parameterList RPAREN)?
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
    : IF expression THEN statement (ELSE statement)?
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
    : WITH recordVariableList DO statement
    | WITH expression DO statement
    ;

tryExceptStatement
    : TRY statements EXCEPT statements END
    ;

tryFinallyStatement
    : TRY statements FINALLY statements END
    ;

recordVariableList
    : variable (COMMA variable)*
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

LONGBOOL
    : 'LONGBOOL'
    ;

CARDINAL
    : 'Cardinal'
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
    : ('A' .. 'Z') ('A' .. 'Z' | '0' .. '9' | '_')*
    ;

STRING_LITERAL
    : '\'' ('\'\'' | ~ ('\''))* '\''
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
