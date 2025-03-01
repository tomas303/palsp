package edit

import (
	"palsp/internal/parser"

	"github.com/antlr4-go/antlr/v4"
)

type lsnFindOnPos struct {
	parser.BasepascalListener
	line      int
	character int
	found     antlr.ParseTree
	foundSpan int
}

func newLsnFindOnPos(line, character int) *lsnFindOnPos {
	return &lsnFindOnPos{
		line:      line,
		character: character,
	}
}

// EnterEveryRule checks each node; if the token span covers the given position, it stores the node.
func (l *lsnFindOnPos) EnterEveryRule(ctx antlr.ParserRuleContext) {
	startToken := ctx.GetStart()
	stopToken := ctx.GetStop()
	if startToken == nil || stopToken == nil {
		return
	}

	startLine := startToken.GetLine()
	startCol := startToken.GetColumn()
	endLine := stopToken.GetLine()
	endCol := stopToken.GetColumn() + len(stopToken.GetText())

	// Check if the position (l.line, l.character) is within the token's range
	positionInRange := (startLine < l.line || (startLine == l.line && startCol <= l.character)) &&
		(endLine > l.line || (endLine == l.line && endCol >= l.character))

	if positionInRange {
		newSpan := stopToken.GetTokenIndex() - startToken.GetTokenIndex() + 1
		if newSpan < l.foundSpan || l.found == nil {
			l.found = ctx
			l.foundSpan = newSpan
		}
	}
}

func (l *lsnFindOnPos) GetFound() antlr.ParseTree {
	return l.found
}
