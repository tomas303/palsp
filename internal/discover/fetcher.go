package discover

import (
	"runtime"
	"sync"
	"sync/atomic"
	"time"

	"palsp/internal/log"
)

// Fetcher is responsible for scanning unscanned units in parallel
type Fetcher struct {
	jobChan      chan string    // Channel for units to process
	normalChan   chan string    // Channel for priority units
	priorityChan chan string    // Channel for priority units
	running      atomic.Bool    // Atomic flag to indicate if fetcher is running
	wg           sync.WaitGroup // WaitGroup to manage worker goroutines
	maxWorkers   int            // Maximum number of worker goroutines
}

var fetcher *Fetcher
var fetcherOnce sync.Once

// GetFetcher returns the singleton instance of the Fetcher
func GetFetcher() *Fetcher {
	fetcherOnce.Do(func() {
		// Use half the available CPU cores for workers
		workers := runtime.NumCPU() / 4
		if workers < 1 {
			workers = 1
		}
		log.Main.Debug().Int("workers", workers).Msg("Initializing Fetcher with workers")

		fetcher = &Fetcher{
			jobChan:      make(chan string, 10000),
			normalChan:   make(chan string, 10000),
			priorityChan: make(chan string, 1000),
			maxWorkers:   workers,
		}
	})
	return fetcher
}

// Start begins processing unscanned units in the background
// Returns immediately, processing happens in background
func (f *Fetcher) Start() {
	// Only start if not already running
	if f.running.CompareAndSwap(false, true) {
		log.Main.Info().Msg("Starting symbol fetcher")
		// Start the dispatcher
		go f.dispatcher()
	}
}

// Stop gracefully stops the fetcher
func (f *Fetcher) Stop() {
	if f.running.CompareAndSwap(true, false) {
		log.Main.Info().Msg("Stopping symbol fetcher")
		// Wait for all workers to finish
		f.wg.Wait()
		log.Main.Info().Msg("Symbol fetcher stopped")
	}
}

// AddPrioritized adds a unit to be processed with high priority
func (f *Fetcher) AddPrioritized(unitName string) {
	if !f.running.Load() {
		f.Start()
	}
	go func() {
		select {
		case f.priorityChan <- unitName:
			log.Main.Debug().Str("unit", unitName).Msg("Added priority unit to fetch queue")
		case <-time.After(10 * time.Second):
			log.Main.Warn().Str("unit", unitName).Msg("Timed out trying to add unit to priority queue")
		}
	}()
}

// Add adds a unit to be processed with normal priority
func (f *Fetcher) AddNormal(unitName string) {
	if !f.running.Load() {
		f.Start()
	}
	// Spawn goroutine to handle queuing - this returns immediately
	go func() {
		select {
		case f.normalChan <- unitName:
			log.Main.Debug().Str("unit", unitName).Msg("Added unit to fetch queue")
		case <-time.After(10 * time.Second):
			log.Main.Warn().Str("unit", unitName).Msg("Timed out trying to add unit to normal queue")
		}
	}()
}

// dispatcher manages the worker pool and distributes units to process
func (f *Fetcher) dispatcher() {
	// Start worker goroutines
	f.wg.Add(f.maxWorkers)
	for i := 0; i < f.maxWorkers; i++ {
		go f.worker(i)
	}

	// When dispatcher exits, close job channel to signal workers to stop
	defer func() {
		close(f.jobChan)
		log.Main.Debug().Msg("Fetcher dispatcher stopped")
	}()

	for f.running.Load() {
		select {
		// Check high priority channel first (non-blocking)
		case unit := <-f.priorityChan:
			f.jobChan <- unit
			continue
		default:
			// No items in priority channel, check normal channel
			select {
			case unit := <-f.normalChan:
				f.jobChan <- unit
				continue
			case <-time.After(100 * time.Millisecond):
				// Nothing in either channel, just wait
			}
		}
	}
}

// worker processes units from the job channel
func (f *Fetcher) worker(id int) {
	defer f.wg.Done()

	log.Main.Debug().Int("worker", id).Msg("Fetcher worker started")

	for unit := range f.jobChan {
		if !f.running.Load() {
			break
		}

		log.Main.Debug().Int("worker", id).Str("unit", unit).Msg("Processing unit")

		// Start timing
		start := time.Now()

		// Retrieve the unit to parse and index it
		_, _, err := SymDB().RetriveUnit(unit)

		// Calculate duration
		duration := time.Since(start)

		if err != nil {
			log.Main.Error().Err(err).Str("unit", unit).Str("duration", duration.String()).Msg("Failed to process unit")
		} else {
			log.Main.Info().Int("worker", id).Str("unit", unit).Str("duration", duration.String()).Msg("Unit processed")
		}
	}

	log.Main.Debug().Int("worker", id).Msg("Fetcher worker stopped")
}
