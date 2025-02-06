package discover

import (
	"fmt"
	"palsp/internal/parser" // Ensure this import is correct
	"strings"               // added

	"github.com/antlr4-go/antlr/v4"
)

// type PascalListener interface {
// 	EnterUnit(ctx *parser.UnitContext)
// 	EnterEveryRule(ctx antlr.ParserRuleContext)
// 	ExitEveryRule(ctx antlr.ParserRuleContext)
// }

type finishError struct {
	Message string
}

func (e *finishError) Error() string {
	return e.Message
}

func newFinishError(message string) *finishError {
	return &finishError{Message: message}
}

type listenerFactory func() antlr.ParseTreeListener

type listenerHandler func(listener antlr.ParseTreeListener, path string)

type unitNameListener struct {
	parser.BasepascalListener

	unitName string
	isUnit   bool
}

func (l *unitNameListener) ExitUnit(ctx *parser.UnitContext) {
	identifiers := ctx.AllIdentifier() // get all identifiers
	if len(identifiers) > 0 {          // if there is at least one identifier
		var parts []string
		for i := 0; i < len(identifiers); i++ {
			parts = append(parts, identifiers[i].GetText())
		}
		unitName := strings.Join(parts, ".") // join with dot delimiter
		fmt.Println("Unit identified:", unitName)
		l.unitName = unitName
		l.isUnit = true
		panic(newFinishError("Unit ID rule hit"))
	}
}

// GetUnitName returns the unit name identified by the listener
func (l *unitNameListener) UnitName() string {
	return l.unitName
}

// IsUnit returns whether the listener has identified a unit
func (l *unitNameListener) IsUnit() bool {
	return l.isUnit
}
