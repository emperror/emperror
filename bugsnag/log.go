package bugsnag

import (
	"fmt"

	"github.com/go-kit/kit/log"
)

// KitLoggerOption configures a kit logger instance.
type KitLoggerOption interface {
	apply(*kitLogger)
}

// KitLoggerMessage configures the message field of the logger.
type KitLoggerMessage string

func (o KitLoggerMessage) apply(l *kitLogger) {
	l.msg = string(o)
}

type kitLogger struct {
	logger log.Logger

	msg string
}

// NewKitLogger returns a new kit logger instance.
func NewKitLogger(logger log.Logger, opts ...KitLoggerOption) *kitLogger {
	l := &kitLogger{logger: logger}

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
func (l *kitLogger) Printf(format string, v ...interface{}) {
	l.logger.Log(l.msg, fmt.Sprintf(format, v...))
}
