package golog

import (
	"fmt"
	nativeLogger "log"
)

// Log is the instance of the
type Log struct {
	loggingLevel string
	logPrefix    string
}

// NewLogger returns a new golog instance via the Logger interface
func NewLogger(logPrefix string) Logger {
	return &Log{
		logPrefix: logPrefix,
	}
}

// Info log info level message
func (l *Log) Info(message interface{}) {

	if Disabled == logLevel {
		return
	} else if InfoLevel != logLevel {
		return
	}

	nativeLogger.Printf("[ %s ] %s -> %v", InfoLevel, l.logPrefix, message)
}

// Infof log info level message with arguments
func (l *Log) Infof(template string, args ...interface{}) {

	if Disabled == logLevel {
		return
	} else if InfoLevel != logLevel {
		return
	}
	message := fmt.Sprintf(template, args...)
	nativeLogger.Printf("[ %s ] %s -> %v", InfoLevel, l.logPrefix, message)
}

// Debug log debug level message
func (l *Log) Debug(message interface{}) {

	if Disabled == logLevel {
		return
	} else if DebugLevel != logLevel && WarnLevel != logLevel && ErrorLevel != logLevel {
		return
	}
	nativeLogger.Printf("[ %s ] %s -> %v", DebugLevel, l.logPrefix, message)
}

// Debugf log debug level message with arguments
func (l *Log) Debugf(template string, args ...interface{}) {

	if Disabled == logLevel {
		return
	} else if DebugLevel != logLevel && WarnLevel != logLevel && ErrorLevel != logLevel {
		return
	}
	message := fmt.Sprintf(template, args...)
	nativeLogger.Printf("[ %s ] %s -> %v", DebugLevel, l.logPrefix, message)
}

// Warn log warning level message
func (l *Log) Warn(message interface{}) {
	if Disabled == logLevel {
		return
	} else if WarnLevel != logLevel && ErrorLevel != logLevel {
		return
	}
	nativeLogger.Printf("[ %s ] %s -> %v", WarnLevel, l.logPrefix, message)
}

// Warnf log warning level message with arguments
func (l *Log) Warnf(template string, args ...interface{}) {

	if Disabled == logLevel {
		return
	} else if WarnLevel != logLevel && ErrorLevel != logLevel {
		return
	}
	message := fmt.Sprintf(template, args...)
	nativeLogger.Printf("[ %s ] %s -> %v", WarnLevel, l.logPrefix, message)
}

// Error log error level message
func (l *Log) Error(message interface{}) {

	if Disabled == logLevel {
		return
	}

	nativeLogger.Printf("[ %s ] %s -> %v", ErrorLevel, l.logPrefix, message)
}

// Errorf log error level message with arguments
func (l *Log) Errorf(template string, args ...interface{}) {

	if Disabled == logLevel {
		return
	}

	message := fmt.Sprintf(template, args...)
	nativeLogger.Printf("[ %s ] %s -> %v", ErrorLevel, l.logPrefix, message)
}
