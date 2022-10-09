package golog

import (
	"fmt"
	nativeLogger "log"
)

// Log is the instance of the
type Log struct {
	logPrefix string
}

// NewLogger returns a new golog instance via the Logger interface
func NewLogger(logPrefix string) Logger {
	return &Log{
		logPrefix: logPrefix,
	}
}

// Info log info level message
func (l *Log) Info(message interface{}) {

	if logLevel == Disabled {
		return
	} else if logLevel == InfoLevel {
		nativeLogger.Printf("[ %s ] %s -> %v", levelMap[InfoLevel], l.logPrefix, message)
	}
}

// Infof log info level message with arguments
func (l *Log) Infof(template string, args ...interface{}) {

	if logLevel == Disabled {
		return
	} else if logLevel == InfoLevel {
		message := fmt.Sprintf(template, args...)
		nativeLogger.Printf("[ %s ] %s -> %v", levelMap[InfoLevel], l.logPrefix, message)
	}
}

// Debug log debug level message
func (l *Log) Debug(message interface{}) {

	if Disabled == logLevel {
		return
	} else if logLevel <= DebugLevel {
		nativeLogger.Printf("[ %s ] %s -> %v", levelMap[DebugLevel], l.logPrefix, message)
	}
}

// Debugf log debug level message with arguments
func (l *Log) Debugf(template string, args ...interface{}) {

	if Disabled == logLevel {
		return
	} else if logLevel <= DebugLevel {
		message := fmt.Sprintf(template, args...)
		nativeLogger.Printf("[ %s ] %s -> %v", levelMap[DebugLevel], l.logPrefix, message)
	}
}

// Warn log warning level message
func (l *Log) Warn(message interface{}) {

	if Disabled == logLevel {
		return
	} else if logLevel <= WarnLevel {
		nativeLogger.Printf("[ %s ] %s -> %v", levelMap[WarnLevel], l.logPrefix, message)
	}
}

// Warnf log warning level message with arguments
func (l *Log) Warnf(template string, args ...interface{}) {

	if Disabled == logLevel {
		return
	} else if logLevel <= WarnLevel {
		message := fmt.Sprintf(template, args...)
		nativeLogger.Printf("[ %s ] %s -> %v", levelMap[WarnLevel], l.logPrefix, message)
	}
}

// Error log error level message
func (l *Log) Error(message interface{}) {

	if Disabled == logLevel {
		return
	} else if logLevel <= ErrorLevel {
		nativeLogger.Printf("[ %s ] %s -> %v", levelMap[ErrorLevel], l.logPrefix, message)
	}
}

// Errorf log error level message with arguments
func (l *Log) Errorf(template string, args ...interface{}) {

	if Disabled == logLevel {
		return
	} else if logLevel <= ErrorLevel {
		message := fmt.Sprintf(template, args...)
		nativeLogger.Printf("[ %s ] %s -> %v", levelMap[ErrorLevel], l.logPrefix, message)
	}
}
