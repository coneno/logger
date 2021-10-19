//Package logger implements logging methods with formatted output
//and more detailed logging information than provided in the standard
//logging package of golang. It is possible to specify the level of
//logging information to be written as output (Default: all information is written).
//Use the `*log.Logger` variables and chain them to common I/O writers
// for printing logging information.
package logger

import (
	"fmt"
	"io"
	"log"
	"os"
)

//LogLevel indicates the level of logging information passed as
//argument to SetLevel function.
type LogLevel int

//There are four possible values for type `LogLevel`:
//
// 1. `LEVEL_ERROR`: Only Error logs are written to `stderr`.
//
// 2. `LEVEL_WARNING`: Warnings and errors are written to `stderr`.
//
// 3. `LEVEL_INFO`:    Errors and warnings are written to `stderr`, infos to `stdout`.
//
// 4. `LEVEL_DEBUG`:   Errors and warnings are written to `stderr`, infos and debugs to `stdout`.
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

//SetLevel sets the level of logging information written as output.
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

	return log.New(writer, fmt.Sprintf("[%s]: ", prefix), log.Ldate|log.Ltime|log.Lshortfile|log.Lmsgprefix)
}

func getEmptyLogger() *log.Logger {
	return log.New(emptyWriter{}, "", 0)
}
