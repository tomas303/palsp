package log

import (
	"io"
	"os"

	"github.com/rs/zerolog"
)

var (
	// Global logger instance
	Logger zerolog.Logger
)

// LogLevel represents log severity levels
type LogLevel string

const (
	DebugLevel LogLevel = "debug"
	InfoLevel  LogLevel = "info"
	WarnLevel  LogLevel = "warn"
	ErrorLevel LogLevel = "error"
	NoneLevel  LogLevel = "none"
)

// Initialize sets up the global logger with specified level and output
func Initialize(level LogLevel, output io.Writer) {
	// Default to stderr if no output specified
	if output == nil {
		output = os.Stderr
	}

	// Set global log level
	var logLevel zerolog.Level
	switch level {
	case DebugLevel:
		logLevel = zerolog.DebugLevel
	case InfoLevel:
		logLevel = zerolog.InfoLevel
	case WarnLevel:
		logLevel = zerolog.WarnLevel
	case ErrorLevel:
		logLevel = zerolog.ErrorLevel
	case NoneLevel:
		logLevel = zerolog.Disabled
	default:
		logLevel = zerolog.InfoLevel
	}

	// Create console writer for pretty output
	output = zerolog.ConsoleWriter{Out: output, TimeFormat: "15:04:05"}

	// Initialize the global logger
	Logger = zerolog.New(output).
		Level(logLevel).
		With().
		Timestamp().
		// Caller().
		Logger()
}

func IsDebugEnabled() bool {
	return Logger.GetLevel() == zerolog.DebugLevel
}
