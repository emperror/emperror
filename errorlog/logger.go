package errorlog

import "github.com/InVisionApp/go-logger"

// InvisionLogger wraps an invision logger and exposes it under a custom interface.
type InvisionLogger struct {
	logger log.Logger
}

// NewInvisionLogger returns a new InvisionLogger instance.
func NewInvisionLogger(logger log.Logger) *InvisionLogger {
	return &InvisionLogger{
		logger: logger,
	}
}

// Error logs an error event.
func (l *InvisionLogger) Error(msg ...interface{}) {
	l.logger.Error(msg...)
}

// WithFields annotates a logger with some context.
func (l *InvisionLogger) WithFields(fields map[string]interface{}) Logger {
	return &InvisionLogger{logger: l.logger.WithFields(fields)}
}
