package terminal

import (
	"fmt"

	"github.com/lucasvmiguel/task/internal/log"
)

// Logger to use log commands
type Logger struct {
	DebugEnabled bool
}

// Info logs info message
func (l *Logger) Info(message string) {
	fmt.Printf(log.InfoColor, fmt.Sprintf("[INFO]: %s.", message))
	fmt.Println()
}

// Infof logs info formatted message
func (l *Logger) Infof(format string, a ...interface{}) {
	fmt.Printf(log.InfoColor, fmt.Sprintf("[INFO]: %s.", fmt.Sprintf(format, a...)))
	fmt.Println()
}

// Warn logs info message
func (l *Logger) Warn(message string) {
	fmt.Printf(log.WarningColor, fmt.Sprintf("[WARN]: %s.", message))
	fmt.Println()
}

// Warnf logs info formatted message
func (l *Logger) Warnf(format string, a ...interface{}) {
	fmt.Printf(log.WarningColor, fmt.Sprintf("[WARN]: %s.", fmt.Sprintf(format, a...)))
	fmt.Println()
}

// Debug logs info message
func (l *Logger) Debug(message string) {
	if l.DebugEnabled {
		fmt.Printf(log.DebugColor, fmt.Sprintf("[DEBUG]: %s.", message))
		fmt.Println()
	}
}

// Debugf logs info formatted message
func (l *Logger) Debugf(format string, a ...interface{}) {
	if l.DebugEnabled {
		fmt.Printf(log.DebugColor, fmt.Sprintf("[DEBUG]: %s.", fmt.Sprintf(format, a...)))
		fmt.Println()
	}
}

// Error logs error message
func (l *Logger) Error(err error) {
	fmt.Printf(log.ErrorColor, fmt.Sprintf("[ERROR]: %s.", err.Error()))
	fmt.Println()
}
