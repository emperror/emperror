package bugsnag

import "fmt"

// Logger is a simple logger interface based on github.com/goph/logur.
type Logger interface {
	// Error logs an error event.
	Error(msg string)
}

// NotifierLogger is a bugsnag error notifier compatible logger.
type NotifierLogger struct {
	logger Logger
}

// NewLogger returns a new logger instance.
func NewLogger(logger Logger) *NotifierLogger {
	return &NotifierLogger{
		logger: logger,
	}
}

// Printf implements a bugsnag logger.
func (l *NotifierLogger) Printf(format string, v ...interface{}) {
	l.logger.Error(fmt.Sprintf(format, v...))
}
