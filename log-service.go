package golog

import (
	"fmt"
)

// Log is the instance of the logger
type Log struct {
	logger    AbstractLogger
}

// NewLogger returns a new golog instance via the GoLogger interface
func NewLogger(logger AbstractLogger) GoLogger {
	return &Log{
		logger: logger,
	}
}

// Debug log debug level message
func (l *Log) Debug(message interface{}) {

	if Disabled == logLevel {
		return
	}

	if logLevel == DebugLevel {
		l.logger.Debug(message)
	}
}

// Debugf log debug level message with arguments
func (l *Log) Debugf(template string, args ...interface{}) {

	if Disabled == logLevel {
		return
	}

	if logLevel == DebugLevel {
		message := fmt.Sprintf(template, args...)
		l.logger.Debug(message)
	}
}

// Info log info level message
func (l *Log) Info(message interface{}) {

	if logLevel == Disabled {
		return
	}

	if logLevel <= InfoLevel {
		l.logger.Info(message)
	}
}

// Infof log info level message with arguments
func (l *Log) Infof(template string, args ...interface{}) {

	if logLevel == Disabled {
		return
	}

	if logLevel <= InfoLevel {
		message := fmt.Sprintf(template, args...)
		l.logger.Info(message)
	}
}

// Warn log warning level message
func (l *Log) Warn(message interface{}) {

	if Disabled == logLevel {
		return
	}

	if logLevel <= WarnLevel {
		l.logger.Warn(message)
	}
}

// Warnf log warning level message with arguments
func (l *Log) Warnf(template string, args ...interface{}) {

	if Disabled == logLevel {
		return
	}

	if logLevel <= WarnLevel {
		message := fmt.Sprintf(template, args...)
		l.logger.Warn(message)
	}
}

// Error log error level message
func (l *Log) Error(message interface{}) {

	if Disabled == logLevel {
		return
	}

	if logLevel <= ErrorLevel {
		l.logger.Error(message)
	}
}

// Errorf log error level message with arguments
func (l *Log) Errorf(template string, args ...interface{}) {

	if Disabled == logLevel {
		return
	}

	if logLevel <= ErrorLevel {
		message := fmt.Sprintf(template, args...)
		l.logger.Error(message)
	}
}
