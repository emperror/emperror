package log

import (
	"github.com/go-kit/kit/log/level"
	"github.com/goph/emperror"
)

// logger covers most of the level-based logging solutions.
type logger interface {
	Log(keyvals ...interface{}) error
}

// handler accepts a logger instance and logs an error.
//
// Compatible with most level-based loggers.
type handler struct {
	l logger
}

// NewHandler returns a new handler.
func NewHandler(l logger) emperror.Handler {
	return &handler{level.Error(l)}
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
			keyvals := append(keyvals, "msg", e.Error(), "parent", err.Error())

			h.l.Log(keyvals...)
		}
	} else {
		keyvals = append(keyvals, "msg", err.Error())

		h.l.Log(keyvals...)
	}
}
