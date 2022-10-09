package golog

type Logger interface {
	// Info logs an info level message
	Info(message interface{})
	// Infof logs an info level message with arguments
	Infof(template string, args ...interface{})
	// Debug logs a debug level message
	Debug(message interface{})
	// Debugf logs a debug level message with arguments
	Debugf(template string, args ...interface{})
	// Warn logs a warning level message
	Warn(message interface{})
	// Warnf logs a warning level message with arguments
	Warnf(template string, args ...interface{})
	// Error logs an error level message
	Error(message interface{})
	// Errorf logs an error level message with arguments
	Errorf(template string, args ...interface{})
}
