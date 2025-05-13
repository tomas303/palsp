package log

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
)

var (
	// Global logger instance
	Main       zerolog.Logger
	AntlrError zerolog.Logger
	AntlrTrace zerolog.Logger
	Structure  zerolog.Logger
)

// LogLevel represents log severity levels
const (
	DebugLevel string = "debug"
	InfoLevel  string = "info"
	WarnLevel  string = "warn"
	ErrorLevel string = "error"
	NoneLevel  string = "none"
)

const (
	// Color codes for different parts of the log
	colorCyan        = "\x1b[36m"       // Cyan
	colorWhite       = "\x1b[37m"       // White
	colorBrightBlue  = "\x1b[94m"       // Bright blue
	colorBlue        = "\x1b[34m"       // Blue
	colorGreen       = "\x1b[32m"       // Green
	colorMagenta     = "\x1b[35m"       // Magenta
	colorYellow      = "\x1b[33m"       // Yellow
	colorLightOrange = "\x1b[38;5;214m" // Light orange
	colorRed         = "\x1b[31m"       // Red
)

const (
	// Color codes for different parts of the log
	colorReset  = "\x1b[0m"
	colorDebug  = colorCyan
	colorInfo   = colorGreen
	colorWarn   = colorYellow
	colorError  = colorRed
	colorFields = colorBrightBlue
)

func NewLogger(level *string, file *string) zerolog.Logger {
	logLevel := logLevelToZerologLevel(level)
	consoleWriter := newConsoleWriter(file)
	result := zerolog.New(consoleWriter).
		Level(logLevel).
		With().
		Timestamp().
		// Caller().
		Logger()
	return result
}

func NewAntlrErrorLogger(level *string, file *string) zerolog.Logger {
	logLevel := logLevelToZerologLevel(level)
	consoleWriter := newConsoleWriter(file)
	consoleWriter.FieldsOrder = []string{"di", "line", "column", "msg"}
	consoleWriter.FormatFieldName = func(i interface{}) string {
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
	consoleWriter.FormatFieldValue = func(i interface{}) string {
		return toString(i) + colorReset
	}
	result := zerolog.New(consoleWriter).
		Level(logLevel).
		With().
		Timestamp().
		// Caller().
		Logger()
	return result
}

func NewAntlrTraceLogger(level *string, file *string) zerolog.Logger {
	logLevel := logLevelToZerologLevel(level)
	consoleWriter := newConsoleWriter(file)
	consoleWriter.FieldsOrder = []string{"di", "line", "column", "msg"}
	consoleWriter.FormatFieldName = func(i interface{}) string {
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
	consoleWriter.FormatFieldValue = func(i interface{}) string {
		return toString(i) + colorReset
	}
	result := zerolog.New(consoleWriter).
		Level(logLevel).
		With().
		Timestamp().
		// Caller().
		Logger()
	return result
}

func toString(i interface{}) string {
	if i == nil {
		return ""
	}
	return fmt.Sprintf("%v", i)
}

func logLevelToZerologLevel(level *string) zerolog.Level {
	if level == nil || *level == "" {
		return zerolog.Disabled
	} else {
		switch *level {
		case DebugLevel:
			return zerolog.DebugLevel
		case InfoLevel:
			return zerolog.InfoLevel
		case WarnLevel:
			return zerolog.WarnLevel
		case ErrorLevel:
			return zerolog.ErrorLevel
		case NoneLevel:
			return zerolog.Disabled
		default:
			return zerolog.InfoLevel
		}
	}
}

func newConsoleWriter(file *string) zerolog.ConsoleWriter {

	output := getOutput(file)

	// Create console writer for pretty output
	writer := zerolog.ConsoleWriter{
		Out:        output,
		TimeFormat: time.RFC3339,
		NoColor:    false,
	}

	// Format level with color
	writer.FormatLevel = func(i interface{}) string {
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
	writer.FormatFieldName = func(i interface{}) string {
		return colorFields + i.(string) + colorReset + ":"
	}

	// Format timestamp
	writer.FormatTimestamp = func(i interface{}) string {
		return i.(string)
	}

	return writer
}

func getOutput(file *string) *os.File {

	output := os.Stderr
	if file != nil && *file != "" {
		file, err := os.OpenFile(*file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			output = file
		} else {
			os.Stderr.WriteString("Failed to open log file: " + err.Error() + "\n")
		}
	}
	return output
}
