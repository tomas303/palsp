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
		workers := runtime.NumCPU() / 2
		if workers < 1 {
			workers = 1
		}
		log.Main.Debug().Int("workers", workers).Msg("Initializing Fetcher with workers")

		fetcher = &Fetcher{
			jobChan:      make(chan string, 100),
			priorityChan: make(chan string, 10),
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

		// // Queue unscanned units
		// go f.queueUnscannedUnits()
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

// AddUnit adds a unit to be processed with high priority
func (f *Fetcher) AddUnit(unitName string) {
	if !f.running.Load() {
		f.Start()
	}

	select {
	case f.priorityChan <- unitName:
		log.Main.Debug().Str("unit", unitName).Msg("Added priority unit to fetch queue")
	default:
		log.Main.Warn().Str("unit", unitName).Msg("Priority queue full, unable to add unit")
	}
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
		// First check priority channel
		select {
		case unit := <-f.priorityChan:
			f.jobChan <- unit
			continue
		default:
			// No priority units, continue to regular handling
		}

		// Then check regular job channel or wait a bit
		select {
		case unit := <-f.priorityChan:
			f.jobChan <- unit
		case <-time.After(100 * time.Millisecond):
			// Just a small delay to prevent tight looping
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

		// Retrieve the unit to parse and index it
		_, _, err := SymDB().RetriveUnit(unit)
		if err != nil {
			log.Main.Error().Err(err).Str("unit", unit).Msg("Failed to process unit")
		}
	}

	log.Main.Debug().Int("worker", id).Msg("Fetcher worker stopped")
}

// queueUnscannedUnits gets all unscanned units and adds them to the job queue
func (f *Fetcher) queueUnscannedUnits() {
	units := SymDB().UnscannedUnits()
	log.Main.Info().Int("count", len(units)).Msg("Queueing unscanned units")

	for _, unit := range units {
		if !f.running.Load() {
			break
		}

		select {
		case f.jobChan <- unit:
			// Unit added to queue
		default:
			// Queue is full, wait a bit
			time.Sleep(10 * time.Millisecond)
			f.jobChan <- unit // Try again (this will block if still full)
		}
	}
}
