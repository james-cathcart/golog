package golog

// This package is a simple wrapper around the native go logger to provide
// some additional functionality

const (
	// InfoLevel provides general information. This is similar to a "verbose" flag and is the
	// most granular option
	InfoLevel = "Info"
	// DebugLevel provides information that may be useful for developers, but not relevant
	// for the typical user.
	DebugLevel = "Debug"
	// WarnLevel provides warnings on events that may not be critical failures, but are unexpected
	// and potentially undesirable. This is usually the preferred level for production
	WarnLevel = "Warn"
	// ErrorLevel provides information on errors within the application or its dependencies. This
	// is also a common option for production, especially when there is limited space for log storage
	ErrorLevel = "Error"
	// Disabled will disable all logging output from this golog
	Disabled = "Disabled"
)

var (
	logLevel string
)

func init() {
	logLevel = DebugLevel
}

// SetLoggingLevel sets the overall logging level for the golog.
func SetLoggingLevel(level string) {
	logLevel = level
}
