package discover

import (
	"fmt"
	"log"
	"palsp/internal/scopeparser/parser"
)

type scopeSymbolsListener struct {
	parser.BasescopepascalListener

	unitName              string
	unit_id               int
	scopeStack            stack[string]
	identExitActionStack  stack[func(identifier string)]
	accessSpecifiersStack stack[AccessSpec]
}

func (s *scopeSymbolsListener) insertSymbol(symbol string, kind int, definition string) {
	err := SymDB().insertSymbol(s.unit_id, symbol, s.scope(), kind, definition)
	if err != nil {
		log.Printf("Non-fatal error encountered: %v", err)
	}
}

func (s *scopeSymbolsListener) scope() string {
	return s.scopeStack.joinByDot()
}

func (s *scopeSymbolsListener) ExitIdentifier(ctx *parser.IdentifierContext) {
	action := s.identExitActionStack.pop()
	if action != nil {
		identifier := ctx.GetText()
		action(identifier)
	}
}

func (s *scopeSymbolsListener) EnterTypeBlock(ctx *parser.TypeBlockContext) {
	s.identExitActionStack.push(func(identifier string) {
		fmt.Printf("Entering type: %s\n", identifier)
		s.scopeStack.push(identifier)
	})

}

func (s *scopeSymbolsListener) ExitTypeBlock(ctx *parser.TypeBlockContext) {
	identifier := s.scopeStack.pop()
	fmt.Printf("Exiting type: %s\n", identifier)
	s.insertSymbol(identifier, int(TypeSymbol), identifier)

}
