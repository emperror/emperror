package emperror

import (
	"github.com/go-kit/kit/log/level"
	"github.com/goph/emperror/internal/errors"
)

// logger covers most of the level-based logging solutions.
type logger interface {
	Log(keyvals ...interface{}) error
}

// logHandler accepts a logger instance and logs an error.
//
// Compatible with most level-based loggers.
type logHandler struct {
	l logger
}

// NewLogHandler returns a new logHandler.
func NewLogHandler(l logger) errors.Handler {
	return &logHandler{level.Error(l)}
}

// Handle logs an error.
func (h *logHandler) Handle(err error) {
	keyvals := []interface{}{"msg", err.Error()}

	if cerr, ok := err.(errors.Contextor); ok {
		keyvals = append(keyvals, cerr.Context()...)
	}

	h.l.Log(keyvals...)
}
