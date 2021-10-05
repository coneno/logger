package main

import "github.com/coneno/logger"

func main() {
	// ERROR
	logger.SetLevel(logger.LEVEL_ERROR)
	logger.Error.Println("error 1")
	logger.Warning.Println("warning 1")
	logger.Info.Println("info 1")
	logger.Debug.Println("debug 1")

	// + WARNING
	logger.SetLevel(logger.LEVEL_WARNING)
	suffix := 2
	logger.Error.Printf("error %d", suffix)
	logger.Warning.Printf("warning  %d", suffix)
	logger.Info.Printf("info  %d", suffix)
	logger.Debug.Printf("debug  %d", suffix)

	// + INFO
	logger.SetLevel(logger.LEVEL_INFO)
	suffix = 3
	logger.Error.Printf("error %d", suffix)
	logger.Warning.Printf("warning  %d", suffix)
	logger.Info.Printf("info  %d", suffix)
	logger.Debug.Printf("debug  %d", suffix)

	// + DEBUG
	logger.SetLevel(logger.LEVEL_DEBUG)
	suffix = 4
	logger.Error.Printf("error %d", suffix)
	logger.Warning.Printf("warning  %d", suffix)
	logger.Info.Printf("info  %d", suffix)
	logger.Debug.Printf("debug  %d", suffix)
}
