package discover

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

type Discover struct{}

func (d *Discover) Units(rootDir string) {
	fc := fileCrawler{}
	fc.processPasFiles(rootDir,
		func() antlr.ParseTreeListener {
			return &unitNameListener{}
		},
		func(listener antlr.ParseTreeListener, path string) {
			unitNameListener := listener.(*unitNameListener)
			if unitNameListener.IsUnit() {
				fmt.Println("Unit found:", unitNameListener.UnitName())
				SymDB().insertUnit(unitNameListener.UnitName(), path)
			}
		})
}
