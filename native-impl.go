package golog

import (
	nativeLogger "log"
)

type NativeLogger struct{}

func NewNativeLogger() AbstractLogger {
	return &NativeLogger{}
}

// Info log info level message
func (l *NativeLogger) Info(message interface{}) {
	nativeLogger.Print(message)
}

// Debug log debug level message
func (l *NativeLogger) Debug(message interface{}) {
	nativeLogger.Print(message)
}

// Warn log warning level message
func (l *NativeLogger) Warn(message interface{}) {
	nativeLogger.Print(message)
}

// Error log error level message
func (l *NativeLogger) Error(message interface{}) {
	nativeLogger.Print(message)
}
