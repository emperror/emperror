package bugsnag

import "fmt"

// logger is a simple logger interface based on go-kit's logger.
type logger interface {
	Log(keyvals ...interface{}) error
}

// LoggerOption configures a logger instance.
type LoggerOption interface {
	apply(*handlerLogger)
}

// LogMessageField configures the message field of the logger.
type LogMessageField string

func (o LogMessageField) apply(l *handlerLogger) {
	l.msg = string(o)
}

type handlerLogger struct {
	logger logger

	msg string
}

// NewLogger returns a new logger instance.
func NewLogger(logger logger, opts ...LoggerOption) *handlerLogger {
	l := &handlerLogger{logger: logger}

	for _, opt := range opts {
		opt.apply(l)
	}

	// Default message field
	if l.msg == "" {
		l.msg = "msg"
	}

	return l
}

// Printf implements a bugsnag logger.
func (l *handlerLogger) Printf(format string, v ...interface{}) {
	_ = l.logger.Log(l.msg, fmt.Sprintf(format, v...))
}
