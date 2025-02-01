package discover

import "palsp/internal/parser"

type PublicSymbolListener struct {
	parser.BasepascalListener
}

func (s *PublicSymbolListener) EnterInterfaceSection(ctx *parser.InterfaceSectionContext) {
}

func (s *PublicSymbolListener) EnterUsesUnits() {}
func (s *PublicSymbolListener) ExitUsesUnits()  {}

func (s *PublicSymbolListener) ExitInterfaceSection(ctx *parser.InterfaceSectionContext) {
}

func (s *PublicSymbolListener) EnterImplementationSection(ctx *parser.ImplementationSectionContext) {

}
