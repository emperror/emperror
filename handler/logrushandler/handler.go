// Package logrushandler provides Logrus integration.
package logrushandler

import (
	"github.com/goph/emperror"
	"github.com/goph/emperror/internal/keyvals"
	"github.com/sirupsen/logrus"
)

// Handler logs errors using Logrus.
type Handler struct {
	logger logrus.FieldLogger
}

// New creates a new handler.
func New(logger logrus.FieldLogger) *Handler {
	return &Handler{
		logger: logger,
	}
}

// Handle logs an error.
func (h *Handler) Handle(err error) {
	var ctx map[string]interface{}

	// Extract context from the error and attach it to the log
	if kvs := emperror.Context(err); len(kvs) > 0 {
		ctx = keyvals.ToMap(kvs)
	}

	logger := h.logger.WithFields(logrus.Fields(ctx))

	type errorCollection interface {
		Errors() []error
	}

	if errs, ok := err.(errorCollection); ok {
		for _, e := range errs.Errors() {
			logger := logger.WithField("parent", err.Error())

			logger.Error(e.Error())
		}
	} else {
		logger.Error(err.Error())
	}
}
