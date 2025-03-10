package edit

import (
	"palsp/internal/discover"
	"palsp/internal/parser"

	"github.com/antlr4-go/antlr/v4"
)

type lsnFindOnPos struct {
	parser.BasepascalListener
	position  discover.Position
	found     antlr.TerminalNode
	foundSpan int
	debugText string // Store text for debugging purposes
}

func newLsnFindOnPos(line, character int) *lsnFindOnPos {
	return &lsnFindOnPos{
		position: discover.Position{
			Line:      line,
			Character: character,
		},
		foundSpan: 0,
	}
}

func (l *lsnFindOnPos) VisitTerminal(node antlr.TerminalNode) {
	termPos := discover.Position{
		Line:      node.GetSymbol().GetLine(),
		Character: node.GetSymbol().GetColumn()}
	l.debugText = node.GetText()
	if termPos.Compare(l.position) <= 0 {
		l.found = node
	}
}

func (l *lsnFindOnPos) GetFound() antlr.TerminalNode {
	return l.found
}

// GetDebugText returns the text content of the found node for debugging
func (l *lsnFindOnPos) GetDebugText() string {
	return l.debugText
}
