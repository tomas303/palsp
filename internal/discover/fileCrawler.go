package discover

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

type listenerData struct {
	Listener antlr.ParseTreeListener
	Path     string
}

type fileCrawler struct{}

func (c *fileCrawler) processPasFiles(rootDir string, factory func() antlr.ParseTreeListener, handler listenerHandler) {
	dataChan := make(chan listenerData)
	var wg sync.WaitGroup

	// Start a single goroutine to process listeners
	go c.processListeners(dataChan, handler)

	// Ensure the channel is closed when all processing is done
	defer close(dataChan)

	// Create a semaphore with a limit on the number of concurrent goroutines
	const maxConcurrentGoroutines = 1
	semaphore := make(chan struct{}, maxConcurrentGoroutines)

	err := c.walk(rootDir, factory, dataChan, &wg, semaphore)
	if err != nil {
		log.Fatalf("Failed to process files: %v", err)
	}

	// Wait for all file processing goroutines to finish
	wg.Wait()
}

func (c *fileCrawler) walk(rootDir string, factory func() antlr.ParseTreeListener, dataChan chan<- listenerData, wg *sync.WaitGroup, semaphore chan struct{}) error {
	return filepath.Walk(rootDir, func(path string, info os.FileInfo, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}

		if !info.IsDir() && strings.HasSuffix(strings.ToLower(info.Name()), ".pas") {
			wg.Add(1)
			semaphore <- struct{}{} // Acquire a slot in the semaphore
			go func(path string, dataChan chan<- listenerData) {
				defer wg.Done()
				defer func() { <-semaphore }() // Release the slot in the semaphore

				file, err := os.Open(path)
				if err != nil {
					log.Printf("Failed to open file %s: %v", path, err)
					return
				}
				defer file.Close()

				listener := factory()
				defer func() {
					if r := recover(); r != nil {
						if finishErr, ok := r.(*finishError); ok {
							log.Printf("Listener finished extraction: %v", finishErr)
							dataChan <- listenerData{Listener: listener, Path: path}
						} else {
							log.Printf("Error parsing file %s: %v", path, r)
						}
					} else {
						log.Printf("Listener finished extraction")
						dataChan <- listenerData{Listener: listener, Path: path}
					}
				}()
				parseFromReader(file, listener, defaultOptions())
			}(path, dataChan)
		}
		return nil
	})
}

func (c *fileCrawler) processListeners(dataChan <-chan listenerData, handler listenerHandler) {
	for data := range dataChan {
		handler(data.Listener, data.Path)
	}
}
