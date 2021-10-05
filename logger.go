package logger

import (
	"fmt"
	"io"
	"log"
	"os"
)

type LogLevel int

const (
	LEVEL_ERROR LogLevel = iota
	LEVEL_WARNING
	LEVEL_INFO
	LEVEL_DEBUG
)

var Error = getErrorLogger()
var Warning = getWarningLogger()
var Info = getInfoLogger()
var Debug = getDebugLogger()

func SetLevel(level LogLevel) {
	Error = getErrorLogger()
	Warning = getWarningLogger()
	Info = getInfoLogger()
	Debug = getDebugLogger()

	switch level {
	case LEVEL_ERROR:
		Warning = getEmptyLogger()
		Info = getEmptyLogger()
		Debug = getEmptyLogger()
	case LEVEL_WARNING:
		Info = getEmptyLogger()
		Debug = getEmptyLogger()
	case LEVEL_INFO:
		Debug = getEmptyLogger()
	case LEVEL_DEBUG:
		return
	}
}

func getErrorLogger() *log.Logger {
	return getLogger("ERROR", true)
}

func getWarningLogger() *log.Logger {
	return getLogger("WARNING", true)
}

func getInfoLogger() *log.Logger {
	return getLogger("INFO", false)
}

func getDebugLogger() *log.Logger {
	return getLogger("DEBUG", false)
}

func getLogger(prefix string, error bool) *log.Logger {
	var writer io.Writer

	if error {
		writer = os.Stderr
	} else {
		writer = os.Stdout
	}

	return log.New(writer, fmt.Sprintf("[%s]\t- ", prefix), log.Ldate|log.Ltime|log.Lshortfile)
}

func getEmptyLogger() *log.Logger {
	return log.New(emptyWriter{}, "", 0)
}
