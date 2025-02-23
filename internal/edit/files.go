package edit

import (
	"fmt"
	"net/url"
	"palsp/internal/discover"
	dsc "palsp/internal/discover"
	"path/filepath"
)

var Lspi *lsp

func init() {
	// Create a new lsp instance.
	Lspi = &lsp{
		fls: &files{
			fileDict: make(map[string]file),
		},
	}
}

type file struct {
	scope *discover.UnitScope
	path  string
}

type files struct {
	// Path to the file.
	fileDict map[string]file
}

type lsp struct {
	fls *files
}

// splitURI splits a URI into its path, name (without extension), and extension.
func splitURI(uri string) (dir, name, ext string, err error) {
	parsed, err := url.Parse(uri)
	if err != nil {
		return "", "", "", err
	}
	dir = parsed.Path
	base := filepath.Base(dir)
	ext = filepath.Ext(base)
	name = base[:len(base)-len(ext)]
	return dir, name, ext, nil
}

func (l *lsp) DidOpen(uri string, text string) OpResult {
	dir, name, _, err := splitURI(uri)
	if err != nil {
		return NewOpResultFail(fmt.Sprintf("Invalid URI: %s", uri), err)
	}
	delete(l.fls.fileDict, dir)
	fileName := filepath.Base(dir)
	unit := fileName[:len(fileName)-len(filepath.Ext(fileName))]
	d := &dsc.Discover{}
	d.Units(name)
	sc := d.ScopeSymbols(unit)
	l.fls.fileDict[dir] = file{
		scope: sc,
		path:  uri,
	}
	return NewOpResultSuccess()
}

func (l *lsp) DidChange(uri string, text string) OpResult {
	return l.DidOpen(uri, text)
}

func (l *lsp) DidClose(uri string) OpResult {
	dir, _, _, err := splitURI(uri)
	if err != nil {
		return NewOpResultFail(fmt.Sprintf("Invalid URI: %s", uri), err)
	}
	delete(l.fls.fileDict, dir)
	return NewOpResultSuccess()
}

func (l *lsp) Hover(uri string, line int, character int) OpResult {
	hoverResp := Hover{
		Contents: MarkupContent{
			Kind:  "plaintext",
			Value: fmt.Sprintf("Hover info for %s", uri),
		},
		Range: &Range{
			Start: Position{Line: line, Character: character},
			End:   Position{Line: line, Character: character},
		},
	}
	return NewOpResultSuccessWithResult(hoverResp)
}
