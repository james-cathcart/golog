package main

import (
	"github.com/james-cathcart/golog"
)

func main() {

	// Create a new AbstractLogger interface and inject the NativeLogger implementation
	var nativeLogger golog.AbstractLogger = golog.NewNativeLogger()
	log := golog.NewLogger(nativeLogger)

	// Set logging level for application
	golog.SetLoggingLevel(golog.DebugLevel)

	log.Debug("debug message")
	log.Info("info message")
	log.Warn("warn message")
	log.Error("error message")

}
