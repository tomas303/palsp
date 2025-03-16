package discover

import (
	"os"
	"path/filepath"
	"strings"
)

type fileCrawler struct{}

func (c *fileCrawler) processPasFiles(rootDir string, handler func(path string)) error {
	return filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(strings.ToLower(info.Name()), ".pas") {
			handler(path)
		}
		return nil
	})
}
