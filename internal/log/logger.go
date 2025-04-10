package log

import (
	"io"
	"os"
	"time"

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
	consoleWriter := zerolog.ConsoleWriter{Out: output, TimeFormat: time.RFC3339, NoColor: false}

	// Custom colors for different log levels (only if colors are enabled)
	consoleWriter.FormatLevel = func(i interface{}) string {
		level := i.(string)
		switch level {
		case "debug":
			return "\x1b[36m" + level + "\x1b[0m" // Cyan
		case "info":
			return "\x1b[32m" + level + "\x1b[0m" // Green
		case "warn":
			return "\x1b[33m" + level + "\x1b[0m" // Yellow
		case "error":
			return "\x1b[31m" + level + "\x1b[0m" // Red
		default:
			return level
		}
	}

	// Optional: Customize field name colors
	consoleWriter.FormatFieldName = func(i interface{}) string {
		return "\x1b[34m" + i.(string) + "\x1b[0m:" // Blu field names
	}

	// Initialize the global logger
	Logger = zerolog.New(consoleWriter).
		Level(logLevel).
		With().
		Timestamp().
		// Caller().
		Logger()
}

func IsDebugEnabled() bool {
	return Logger.GetLevel() == zerolog.DebugLevel
}
