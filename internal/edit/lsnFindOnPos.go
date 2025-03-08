package edit

import (
	"palsp/internal/discover"
	"palsp/internal/parser"

	"github.com/antlr4-go/antlr/v4"
)

type lsnFindOnPos struct {
	parser.BasepascalListener
	position  discover.Position
	found     antlr.ParseTree
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

// EnterEveryRule checks each node; if the token span covers the given position, it stores the node.
func (l *lsnFindOnPos) EnterEveryRule(ctx antlr.ParserRuleContext) {
	startToken := ctx.GetStart()
	stopToken := ctx.GetStop()
	if startToken == nil || stopToken == nil {
		return
	}

	var text string
	startLine := startToken.GetLine()
	startCol := startToken.GetColumn()
	endLine := stopToken.GetLine()
	endCol := stopToken.GetColumn() + len(stopToken.GetText())

	// Check if the position (l.position.Line, l.position.Character) is within the token's range
	positionInRange := (startLine < l.position.Line || (startLine == l.position.Line && startCol <= l.position.Character)) &&
		(endLine > l.position.Line || (endLine == l.position.Line && endCol >= l.position.Character))
	if positionInRange {
		newSpan := stopToken.GetTokenIndex() - startToken.GetTokenIndex() + 1
		if newSpan < l.foundSpan || l.found == nil {
			l.found = ctx
			text = ctx.GetText() // Get the text of the current context
			l.debugText = text   // Store it for debugging
			l.foundSpan = newSpan
		}
	}
}

func (l *lsnFindOnPos) GetFound() antlr.ParseTree {
	return l.found
}

// GetDebugText returns the text content of the found node for debugging
func (l *lsnFindOnPos) GetDebugText() string {
	return l.debugText
}
