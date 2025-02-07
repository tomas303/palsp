package discover

import (
	"path/filepath"
	"strings"
)

type Discover struct{}

func (d *Discover) Units(rootDir string) {
	fc := fileCrawler{}
	// fc.processPasListeners(rootDir,
	// 	func() antlr.ParseTreeListener {
	// 		return &unitNameListener{}
	// 	},
	// 	func(listener antlr.ParseTreeListener, path string) {
	// 		unitNameListener := listener.(*unitNameListener)
	// 		if unitNameListener.IsUnit() {
	// 			fmt.Println("Unit found:", unitNameListener.UnitName())
	// 			SymDB().insertUnit(unitNameListener.UnitName(), path)
	// 		}
	// 	})
	fc.processPasFiles(rootDir,
		func(path string) {
			filename := filepath.Base(path)
			ext := filepath.Ext(path)
			unitName := strings.TrimSuffix(filename, ext)
			SymDB().insertUnit(unitName, path)
		})

}
