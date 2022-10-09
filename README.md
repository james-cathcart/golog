# GoLog
This is a simple wrapper around the native Go logger that provides some additional features, primarily logging levels and the ability to disable the logs if desired (such as for running tests).

## Install
```
go get github.com/james-cathcart/golog
```

## Usage
```
package main

import (
	"github.com/james-cathcart/golog"
)

func main() {

    logPrefix := `main`
    
    // Create a new logger instance
    log := golog.NewLogger(logPrefix)
    
    // Set logging level for application
    golog.SetLoggingLevel(golog.InfoLevel)
    
    // Log a basic info message
    log.Info("initializing application")
    
    // Log a debug message with a string template and arguments
    log.Debugf("Hello %s", `world`)
}
```
