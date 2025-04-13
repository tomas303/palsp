package log

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
)

var (
	// Global logger instance
	Logger           zerolog.Logger
	AntlrErrorLogger zerolog.Logger
	AntlrTraceLogger zerolog.Logger
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

const (
	// Color codes for different parts of the log
	colorReset       = "\x1b[0m"
	colorDebug       = "\x1b[36m"       // Cyan
	colorInfo        = "\x1b[32m"       // Green
	colorWarn        = "\x1b[33m"       // Yellow
	colorError       = "\x1b[31m"       // Red
	colorFields      = "\x1b[34m"       // Blue for field names
	colorCyan        = "\x1b[36m"       // Cyan
	colorWhite       = "\x1b[37m"       // White
	colorBrightBlue  = "\x1b[94m"       // Bright blue
	colorBlue        = "\x1b[34m"       // Blue
	colorGreen       = "\x1b[32m"       // Green
	colorMagenta     = "\x1b[35m"       // Magenta
	colorYellow      = "\x1b[33m"       // Yellow
	colorLightOrange = "\x1b[38;5;214m" // Light orange

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

	AntlrErrorConsoleWriter := consoleWriter
	AntlrErrorConsoleWriter.FieldsOrder = []string{"di", "line", "column", "msg"}
	AntlrErrorConsoleWriter.FormatFieldName = func(i interface{}) string {
		// return colorError + i.(string) + ":" + colorReset
		fieldName := toString(i)
		switch fieldName {
		case "di":
			return colorError + fieldName + ": " + colorCyan
		case "line":
			return colorError + fieldName + ": " + colorYellow
		case "column":
			return colorError + fieldName + ": " + colorYellow
		case "msg":
			return colorError + fieldName + ": " + colorLightOrange
		default:
			return colorError + fieldName + ": " + colorWhite
		}
	}
	AntlrErrorConsoleWriter.FormatFieldValue = func(i interface{}) string {
		return toString(i) + colorReset
	}
	AntlrErrorLogger = Logger.Output(AntlrErrorConsoleWriter)

	AntlrTraceConsoleWriter := consoleWriter
	AntlrTraceConsoleWriter.FieldsOrder = []string{"di", "enter", "exit", "token", "rule", "consume"}
	AntlrTraceConsoleWriter.FormatFieldName = func(i interface{}) string {
		fieldName := toString(i)
		switch fieldName {
		case "di":
			return colorDebug + fieldName + ": " + colorCyan
		case "enter", "exit":
			return colorDebug + fieldName + ": " + colorGreen
		case "token":
			return colorDebug + fieldName + ": " + colorYellow
		case "rule":
			return colorDebug + fieldName + ": " + colorGreen
		case "consume":
			return colorDebug + fieldName + ": " + colorMagenta
		default:
			return colorDebug + fieldName + ": " + colorWhite
		}
	}
	AntlrTraceConsoleWriter.FormatFieldValue = func(i interface{}) string {
		return toString(i) + colorReset
	}

	AntlrTraceLogger = Logger.Output(AntlrTraceConsoleWriter)

}

func IsDebugEnabled() bool {
	return Logger.GetLevel() == zerolog.DebugLevel
}

func toString(i interface{}) string {
	if i == nil {
		return ""
	}
	return fmt.Sprintf("%v", i)
}
