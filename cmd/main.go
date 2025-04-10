package main

import (
	"flag"
	"os"

	"palsp/internal/log"
	"palsp/internal/lsp"
)

func main() {
	port := flag.String("port", "", "Port to run the LSP server on (leave empty for stdio)")
	logLevel := flag.String("log-level", "debug", "Log level (debug, info, warn, error, none)")
	logFile := flag.String("log-file", "", "Log file path (defaults to stderr)")
	flag.Parse()

	var output = os.Stderr
	if *logFile != "" {
		file, err := os.OpenFile(*logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			output = file
		} else {
			os.Stderr.WriteString("Failed to open log file: " + err.Error() + "\n")
		}
	}

	log.Initialize(log.LogLevel(*logLevel), output)
	log.Logger.Info().Msg("Application started")

	if *port == "" {
		log.Logger.Info().Msg("Starting LSP server on stdio")
	} else {
		log.Logger.Info().Str("Starting LSP server on port", *port)
	}

	lsp.StartServer(*port)
}
