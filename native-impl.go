package golog

import (
    nativeLogger "log"
)

type NativeLogger struct {
    prefix string
}

func NewNativeLogger(prefix string) AbstractLogger {
    return &NativeLogger{
        prefix: prefix,
    }
}

// Info log info level message
func (l *NativeLogger) Info(message interface{}) {
    nativeLogger.Printf("%s%s", l.prefix, message)
}

// Debug log debug level message
func (l *NativeLogger) Debug(message interface{}) {
    nativeLogger.Printf("%s%s", l.prefix, message)
}

// Warn log warning level message
func (l *NativeLogger) Warn(message interface{}) {
    nativeLogger.Printf("%s%s", l.prefix, message)
}

// Error log error level message
func (l *NativeLogger) Error(message interface{}) {
    nativeLogger.Printf("%s%s", l.prefix, message)
}
