package log

const (
	// InfoColor used for logging
	InfoColor = "\033[1;34m%s\033[0m"
	// ErrorColor used for logging
	ErrorColor = "\033[1;31m%s\033[0m"
	// WarningColor used for logging
	WarningColor = "\033[1;33m%s\033[0m"
	// DebugColor used for logging
	DebugColor = "\033[0;36m%s\033[0m"
)

// Logger is an interface to represent any log lib
type Logger interface {
	Info(string)
	Infof(string, ...interface{})
	Debug(string)
	Debugf(string, ...interface{})
	Warn(string)
	Warnf(string, ...interface{})
	Error(error)
}
