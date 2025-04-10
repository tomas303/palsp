package log

import (
	"fmt"
	"io"
	"os"
	"strings"
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
	consoleWriter := zerolog.ConsoleWriter{
		Out:        output,
		TimeFormat: time.RFC3339,
		NoColor:    false,
	}

	// Color codes for different parts of the log
	var (
		colorReset  = "\x1b[0m"
		colorDebug  = "\x1b[36m" // Cyan
		colorInfo   = "\x1b[32m" // Green
		colorWarn   = "\x1b[33m" // Yellow
		colorError  = "\x1b[31m" // Red
		colorFields = "\x1b[34m" // Blue for field names
	)

	// Format level with color
	consoleWriter.FormatLevel = func(i interface{}) string {
		level := i.(string)
		var levelColor string

		switch level {
		case "debug":
			levelColor = colorDebug
		case "info":
			levelColor = colorInfo
		case "warn":
			levelColor = colorWarn
		case "error":
			levelColor = colorError
		default:
			return level
		}

		return levelColor + level + colorReset
	}

	// Format field names with color
	consoleWriter.FormatFieldName = func(i interface{}) string {
		return colorFields + i.(string) + colorReset + ":"
	}

	// Format message with same color as level
	consoleWriter.FormatMessage = func(i interface{}) string {
		if i == nil {
			return ""
		}

		msg := fmt.Sprintf("%s", i)
		if msg == "" {
			return ""
		}

		// Get current log level from context (this is tricky in zerolog)
		// We'll use a simpler approach by setting colors based on a prefix check

		// For demonstration, let's color based on message content
		var msgColor string

		if strings.HasPrefix(msg, "ANTLR syntax error") {
			msgColor = colorError // Use error color for ANTLR syntax errors
		} else if strings.Contains(msg, "warning") || strings.Contains(msg, "warn") {
			msgColor = colorWarn
		} else if strings.Contains(msg, "debug") {
			msgColor = colorDebug
		} else {
			msgColor = colorInfo // Default to info color
		}

		return msgColor + msg + colorReset
	}

	// Format timestamp
	consoleWriter.FormatTimestamp = func(i interface{}) string {
		return i.(string)
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
