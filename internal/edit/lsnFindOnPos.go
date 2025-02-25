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
	// Check if the token's span covers the target position.
	// Condition: token starts before or at the target position and ends after or at the target.
	if startToken.GetLine() < l.line ||
		(startToken.GetLine() == l.line && startToken.GetColumn() <= l.character) {
		// Compute end column from stop token text length.
		text := stopToken.GetText()
		endCol := stopToken.GetColumn() + len(text) - 1
		if stopToken.GetLine() > l.line ||
			(stopToken.GetLine() == l.line && endCol >= l.character) {
			l.found = ctx
			// Optionally print debug info:
			// fmt.Println("Found node covering position", l.line, strconv.Itoa(l.character))
			panic("found")
		}
	}
}

func (l *lsnFindOnPos) GetFound() antlr.ParseTree {
	return l.found
}
