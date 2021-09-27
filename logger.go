package logger

import (
	"fmt"
	"io"
	"log"
	"os"
)

type LogLevel int

const (
	LEVEL_INFO LogLevel = iota
	LEVEL_DEBUG
	LEVEL_ERROR
)

var Info = getInfoLogger()
var Debug = getDebugLogger()
var Error = getErrorLogger()

func SetLevel(level LogLevel) {
	switch level {
	case LEVEL_INFO:
		Info = getInfoLogger()
		Debug = getDebugLogger()
		Error = getErrorLogger()
	case LEVEL_DEBUG:
		Info = getEmptyLogger()
		Debug = getDebugLogger()
		Error = getErrorLogger()
	case LEVEL_ERROR:
		Info = getEmptyLogger()
		Debug = getEmptyLogger()
		Error = getErrorLogger()
	}
}

func getInfoLogger() *log.Logger {
	return getLogger("INFO", false)
}

func getDebugLogger() *log.Logger {
	return getLogger("DEBUG", false)
}

func getErrorLogger() *log.Logger {
	return getLogger("ERROR", true)
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
