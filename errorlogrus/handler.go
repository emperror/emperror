/*
Package errorlogrus logs an error using Logrus.

Start by creating a logger instance and inject it into the the handler:

	import (
		"github.com/emperror/errorlogrus"
		"github.com/sirupsen/logrus"
	)

	// ...

	logger := logrus.New()
	handler := errorlogrus.NewHandler(logger)
*/
package errorlogrus

import (
	"github.com/goph/emperror"
	"github.com/goph/emperror/internal/keyvals"
	"github.com/sirupsen/logrus"
)

// Handler handles an error by passing it to a logrus logger.
type Handler struct {
	logger logrus.FieldLogger
}

// NewHandler returns a handler which logs errors using logrus.
func NewHandler(logger logrus.FieldLogger) *Handler {
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
