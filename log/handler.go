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
	keyvals := []interface{}{"msg", err.Error()}

	if cerr, ok := err.(emperror.Contextor); ok {
		keyvals = append(keyvals, cerr.Context()...)
	}

	h.l.Log(keyvals...)
}
