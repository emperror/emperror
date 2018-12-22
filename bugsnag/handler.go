/*
Package bugsnag provides Bugsnag integration.

Bugsnag recommends to create a global scope error handler.
Although that works, makes writing software harder, so this package recommends creating a separate instance:

	APIKey := "key"

	handler := bugsnag.NewHandler(APIKey)

If you need more control over th underlying Notifier instance (eg. more advanced construction),
you can create a custom one and then create a handler using it:

	import (
		"github.com/bugsnag/bugsnag-go"
		emperror_bugsnag "github.com/goph/emperror/bugsnag"
	)

	// ...

	handler := &emperror_bugsnag.NewHandlerFromNotifier(bugsnag.New(bugsnag.Configuration{
		APIKey:      APIKey,
		AppVersion:  "1.0.0",
		Synchronous: true,
	}))
*/
package bugsnag

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

// NewHandler creates a new Handler.
func NewHandler(APIKey string) *Handler {
	return &Handler{
		bugsnag.New(bugsnag.Configuration{
			APIKey: APIKey,
		}),
	}
}

// NewHandlerFromNotifier creates a new Handler from an existing bugsnag notifier.
func NewHandlerFromNotifier(notifier *bugsnag.Notifier) *Handler {
	return &Handler{notifier}
}

// Handle passes the error to the underlying Bugsnag notifier.
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
