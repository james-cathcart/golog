package main

import (
	"github.com/james-cathcart/golog"
)

func main() {

	logPrefix := `main`

	// Create a new logger instance
	log := golog.NewLogger(logPrefix)

	// Set logging level for application
	golog.SetLoggingLevel(golog.DebugLevel)

	log.Info("info message")
	log.Debug("debug message")
	log.Warn("warn message")
	log.Error("error message")

}
