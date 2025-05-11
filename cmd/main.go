package main

import (
	"flag"

	"palsp/internal/log"
	"palsp/internal/lsp"
)

func main() {
	port := flag.String("port", "", "Port to run the LSP server on (leave empty for stdio)")
	logLevelMain := flag.String("log-level-main", "none", "Log level (debug, info, warn, error, none)")
	logFileMain := flag.String("log-file-main", "", "Log file path (defaults to stderr)")
	logLevelAntlrError := flag.String("log-level-antlr-error", "none", "Log level (debug, info, warn, error, none)")
	logFileAntlrError := flag.String("log-file-antlr-error", "", "Log file path (defaults to stderr)")
	logLevelAntlrTrace := flag.String("log-level-antlr-trace", "none", "Log level (debug, info, warn, error, none)")
	logFileAntlrTrace := flag.String("log-file-antlr-trace", "", "Log file path (defaults to stderr)")
	logLevelStructure := flag.String("log-level-structure", "none", "Log level (debug, info, warn, error, none)")
	logFileStructure := flag.String("log-file-antlr-structure", "", "Log file path (defaults to stderr)")
	flag.Parse()

	log.Main = log.NewLogger(logLevelMain, logFileMain)
	log.AntlrError = log.NewAntlrErrorLogger(logLevelAntlrError, logFileAntlrError)
	log.AntlrTrace = log.NewAntlrTraceLogger(logLevelAntlrTrace, logFileAntlrTrace)
	log.Structure = log.NewLogger(logLevelStructure, logFileStructure)

	log.Main.Info().Msg("Application started")
	lsp.StartServer(*port)
}
