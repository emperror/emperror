// Package bugsnaghandler provides Bugsnag integration.
package bugsnaghandler

import (
	"reflect"

	"github.com/bugsnag/bugsnag-go"
	"github.com/goph/emperror"
	"github.com/goph/emperror/internal/keyvals"
	"github.com/pkg/errors"
)

// Handler is responsible for sending errors to Bugsnag.
type Handler struct {
	notifier *bugsnag.Notifier
}

// New creates a new handler.
func New(APIKey string) *Handler {
	return NewFromNotifier(bugsnag.New(bugsnag.Configuration{
		APIKey: APIKey,
	}))
}

// NewFromNotifier creates a new handler from an existing notifier instance.
func NewFromNotifier(notifier *bugsnag.Notifier) *Handler {
	return &Handler{
		notifier: notifier,
	}
}

// Handle sends the error to Bugsnag.
func (h *Handler) Handle(err error) {
	err = emperror.ExposeStackTrace(err)

	if e, ok := err.(stackTracer); ok {
		err = newErrorWithStackFrames(e)
	}

	var rawData []interface{}

	cause := errors.Cause(err)
	if name := reflect.TypeOf(cause).String(); len(name) > 0 {
		errorClass := bugsnag.ErrorClass{Name: name}

		rawData = append(rawData, errorClass)
	}

	if ctx := emperror.Context(err); len(ctx) > 0 {
		rawData = append(rawData, bugsnag.MetaData{
			"Params": keyvals.ToMap(ctx),
		})
	}

	_ = h.notifier.Notify(err, rawData...)
}