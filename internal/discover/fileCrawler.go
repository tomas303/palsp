package discover

import (
	"io"
	"net/url"
	"os"
	"palsp/internal/log"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
)

type fileCrawler struct{}

// processPasFiles processes Pascal files without loading the entire directory structure first
func (fc *fileCrawler) processPasFiles(rootPath string, processor func(filepath string)) {
	// Convert URI to file path if needed
	root := rootPath

	// Try to parse as URL - this is more robust than just checking prefix
	if uri, err := url.Parse(rootPath); err == nil && uri.Scheme == "file" {
		// Valid file URI, convert to file path
		root = uri.Path

		// On Windows, remove leading slash if present
		if runtime.GOOS == "windows" && strings.HasPrefix(root, "/") {
			root = root[1:]
		}
	}

	// Use a worker pool pattern with bounded concurrency
	numWorkers := runtime.NumCPU()
	pathChan := make(chan string, 100) // Buffer for paths to process
	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for path := range pathChan {
				processor(path)
			}
		}()
	}

	// Start a single goroutine for directory scanning
	go func() {
		defer close(pathChan)
		scanDirTree(root, pathChan)
	}()

	// Wait for all processing to complete
	wg.Wait()
}

// scanDirTree scans directories without sorting and minimal memory usage
func scanDirTree(root string, pathChan chan<- string) {
	// Process the current directory
	dir, err := os.Open(root)
	if err != nil {
		log.Logger.Error().Str("path", root).Msg("Error opening directory")
		return
	}
	defer dir.Close()

	for {
		// Read directory entries in chunks without loading all at once
		entries, err := dir.ReadDir(100)
		if err != nil && err != io.EOF {
			log.Logger.Error().Str("path", root).Msg("Error reading directory")
			return
		}
		if len(entries) == 0 {
			break // No more entries
		}

		// Process each entry
		for _, entry := range entries {
			path := filepath.Join(root, entry.Name())

			if entry.IsDir() {
				// Recursively scan subdirectories
				scanDirTree(path, pathChan)
			} else if strings.HasSuffix(strings.ToLower(path), ".pas") {
				// Send Pascal files for processing
				pathChan <- path
			}
		}
	}
}
