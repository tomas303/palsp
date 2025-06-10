package main

import (
	"flag"
	stdlog "log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"runtime/pprof"
	"syscall"

	"palsp/internal/log"
	"palsp/internal/lsp"
)

func init() {
	go func() {
		stdlog.Println(http.ListenAndServe("localhost:6060", nil))
	}()
}

func main() {
	port := flag.String("port", "", "Port to run the LSP server on (leave empty for stdio)")
	logLevelMain := flag.String("log-level-main", "none", "Log level (debug, info, warn, error, none)")
	logFileMain := flag.String("log-file-main", "", "Log file path (defaults to stderr)")
	logLevelAntlrError := flag.String("log-level-antlr-error", "none", "Log level (debug, info, warn, error, none)")
	logFileAntlrError := flag.String("log-file-antlr-error", "", "Log file path (defaults to stderr)")
	logLevelAntlrTrace := flag.String("log-level-antlr-trace", "none", "Log level (debug, info, warn, error, none)")
	logFileAntlrTrace := flag.String("log-file-antlr-trace", "", "Log file path (defaults to stderr)")

	// Add profiling flags
	cpuProfile := flag.String("cpuprofile", "", "write cpu profile to file")
	memProfile := flag.String("memprofile", "", "write memory profile to file")

	flag.Parse()

	// Start CPU profiling if requested
	if *cpuProfile != "" {
		f, err := os.Create(*cpuProfile)
		if err != nil {
			stdlog.Fatalf("could not create CPU profile: %v", err)
		}
		defer f.Close()

		if err := pprof.StartCPUProfile(f); err != nil {
			stdlog.Fatalf("could not start CPU profile: %v", err)
		}
		defer pprof.StopCPUProfile()
	}

	log.Main = log.NewLogger(logLevelMain, logFileMain)
	log.AntlrError = log.NewAntlrErrorLogger(logLevelAntlrError, logFileAntlrError)
	log.AntlrTrace = log.NewAntlrTraceLogger(logLevelAntlrTrace, logFileAntlrTrace)

	log.Main.Info().Msg("Application started")

	// Set up signal handling for graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Start the LSP server in a goroutine
	done := make(chan error, 1)
	go func() {
		lsp.StartServer(*port)
		done <- nil
	}()

	// Wait for either server to finish or signal
	select {
	case err := <-done:
		if err != nil {
			log.Main.Error().Err(err).Msg("LSP server error")
		}
	case <-c:
		log.Main.Info().Msg("Shutting down...")
	}

	// Write memory profile if requested
	if *memProfile != "" {
		f, err := os.Create(*memProfile)
		if err != nil {
			stdlog.Fatalf("could not create memory profile: %v", err)
		}
		defer f.Close()

		if err := pprof.WriteHeapProfile(f); err != nil {
			stdlog.Fatalf("could not write memory profile: %v", err)
		}
	}

	log.Main.Info().Msg("Application stopped")
}
