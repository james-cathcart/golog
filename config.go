package golog

// This package is a simple wrapper around the native go logger to provide
// some additional functionality

const (
	// DebugLevel provides information that may be over sharing for typical users but also
	// may be useful for developers.
	DebugLevel = 1
	// InfoLevel general information. This is similar to a "verbose" flag and is the
	// most granular option
	InfoLevel = 2
	// WarnLevel provides warnings on events that may not be critical failures, but are unexpected
	// and potentially undesirable. This is usually the preferred level for production
	WarnLevel = 3
	// ErrorLevel provides information on errors within the application or its dependencies. This
	// is also a common option for production, especially when there is limited space for log storage
	ErrorLevel = 4
	// Disabled will disable all logging output from this golog
	Disabled = 5
)

var (
	logLevel int
	levelMap = map[int]string{
		DebugLevel: `Debug`,
		InfoLevel:  `Info`,
		WarnLevel:  `Warn`,
		ErrorLevel: `Error`,
	}
)

func init() {
	logLevel = DebugLevel
}

// SetLoggingLevel sets the overall logging level for the golog.
func SetLoggingLevel(level int) {
	logLevel = level
}

// GetLoggingLevel
func GetLoggingLevel() int {
	return logLevel
}
