/*
Package errorlog logs an error with a go-kit compatible logger.

Start by creating a logger instance and inject it into the the handler:

	import (
		"os"

		"github.com/emperror/errorlog"
		"github.com/InVisionApp/go-logger"
	)

	// ...

	handler := errorlog.NewHandler(errorlog.NewInvisionLogger(log.NewSimple())
*/
package errorlog

import (
	"github.com/goph/emperror"
	"github.com/goph/emperror/internal/keyvals"
)

// Logger is a simple logger interface for recording errors as log events.
type Logger interface {
	// Error logs an Error event.
	Error(msg ...interface{})

	// WithFields returns a new logger with appended fields to the internal context.
	WithFields(fields map[string]interface{}) Logger
}

// handler accepts a logger instance and logs an error.
type handler struct {
	logger Logger
}

// NewHandler returns a new handler.
func NewHandler(logger Logger) emperror.Handler {
	return &handler{logger: logger}
}

// Handle logs an error.
func (h *handler) Handle(err error) {
	var kvs []interface{}

	// Extract context from the error and attach it to the log
	if kv := emperror.Context(err); len(kv) > 0 {
		kvs = append(kvs, kv...)
	}

	type errorCollection interface {
		Errors() []error
	}

	if errs, ok := err.(errorCollection); ok {
		for _, e := range errs.Errors() {
			kvs := append(kvs, "parent", err.Error())

			h.logger.WithFields(keyvals.ToMap(kvs)).Error(e.Error())
		}
	} else {
		h.logger.WithFields(keyvals.ToMap(kvs)).Error(err.Error())
	}
}
