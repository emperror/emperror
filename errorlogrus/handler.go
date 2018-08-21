package errorlogrus

import (
	"github.com/goph/emperror"
	"github.com/goph/emperror/internal/keyvals"
	"github.com/sirupsen/logrus"
)

type handler struct {
	logger logrus.FieldLogger
}

// NewHandler returns a handler which logs errors using logrus.
func NewHandler(logger logrus.FieldLogger) *handler {
	return &handler{logger: logger}
}

// Handle logs an error.
func (h *handler) Handle(err error) {
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
