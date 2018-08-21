/*
Package errorlog logs an error with a go-kit compatible logger.

Start by creating a logger instance and inject it into the the handler:

	import (
		"os"

		"github.com/emperror/errorlog"
		"github.com/go-kit/kit/log"
		"github.com/go-kit/kit/log/level"
	)

	// ...

	logger := level.Error(log.NewLogfmtLogger(os.Stdout))
	handler := errorlog.NewHandler(logger)
*/
package errorlog

import "github.com/goph/emperror"

// logger is a simple logger interface based on go-kit's logger.
type logger interface {
	Log(keyvals ...interface{}) error
}

// Option configures a handler instance.
type Option interface {
	apply(*handler)
}

// MessageField configures the message field of the logger.
type MessageField string

func (o MessageField) apply(l *handler) {
	l.msg = string(o)
}

// handler accepts a logger instance and logs an error.
type handler struct {
	logger logger

	msg string
}

// NewHandler returns a new handler.
func NewHandler(logger logger, opts ...Option) emperror.Handler {
	h := &handler{logger: logger}

	for _, o := range opts {
		o.apply(h)
	}

	// Default message field
	if h.msg == "" {
		h.msg = "msg"
	}

	return h
}

// Handle logs an error.
func (h *handler) Handle(err error) {
	var keyvals []interface{}

	// Extract context from the error and attach it to the log
	if kvs := emperror.Context(err); len(kvs) > 0 {
		keyvals = append(keyvals, kvs...)
	}

	type errorCollection interface {
		Errors() []error
	}

	if errs, ok := err.(errorCollection); ok {
		for _, e := range errs.Errors() {
			keyvals := append(keyvals, h.msg, e.Error(), "parent", err.Error())

			h.logger.Log(keyvals...)
		}
	} else {
		keyvals = append(keyvals, h.msg, err.Error())

		h.logger.Log(keyvals...)
	}
}
