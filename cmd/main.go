package main

import (
	"flag"

	"palsp/internal/log"
	"palsp/internal/lsp"
)

func main() {
	port := flag.String("port", "", "Port to run the LSP server on (leave empty for stdio)")
	logLevel := flag.String("log-level", "none", "Log level (debug, info, warn, error, none)")
	logFile := flag.String("log-file", "", "Log file path (defaults to stderr)")
	logLevelAntlrError := flag.String("log-level-antlr-error", "none", "Log level (debug, info, warn, error, none)")
	logFileAntlrError := flag.String("log-file-antlr-error", "", "Log file path (defaults to stderr)")
	logLevelAntlrTrace := flag.String("log-level-antlr-trace", "none", "Log level (debug, info, warn, error, none)")
	logFileAntlrTrace := flag.String("log-file-antlr-trace", "", "Log file path (defaults to stderr)")
	flag.Parse()

	log.Logger = log.NewLogger(logLevel, logFile)
	log.AntlrErrorLogger = log.NewAntlrErrorLogger(logLevelAntlrError, logFileAntlrError)
	log.AntlrTraceLogger = log.NewAntlrTraceLogger(logLevelAntlrTrace, logFileAntlrTrace)

	log.Logger.Info().Msg("Application started")
	lsp.StartServer(*port)
}
