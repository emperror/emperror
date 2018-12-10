package errorlog

import "github.com/goph/logur"

// LogurLogger wraps an invision logger and exposes it under a custom interface.
type LogurLogger struct {
	logger logur.Logger
}

// NewLogger returns a new Logger instance based on github.com/goph/logur.
func NewLogger(logger logur.Logger) *LogurLogger {
	return &LogurLogger{
		logger: logger,
	}
}

// Error logs an error event.
func (l *LogurLogger) Error(msg ...interface{}) {
	l.logger.Error(msg...)
}

// WithFields annotates a logger with some context.
func (l *LogurLogger) WithFields(fields map[string]interface{}) Logger {
	return &LogurLogger{logger: l.logger.WithFields(fields)}
}
