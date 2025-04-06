package discover

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
)

type fileCrawler struct{}

func (fc *fileCrawler) processPasFiles(root string, processor func(filepath string)) {
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

	// Walk the directory tree and send paths to workers
	// This part still needs to be sequential
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Only process .pas files
		if strings.ToLower(filepath.Ext(path)) == ".pas" {
			pathChan <- path
		}
		return nil
	})

	close(pathChan) // Signal workers that no more paths are coming
	wg.Wait()       // Wait for all workers to finish

	if err != nil {
		log.Printf("Error walking directory: %v", err)
	}
}
