/*
Package airbrake provides Airbrake integration.

Create a new handler as you would create a gobrake.Notifier:

	projectID  := int64(1)
	projectKey := "key"

	handler := airbrake.NewHandler(projectID, projectKey)

If you need access to the underlying Notifier instance (or need more advanced construction),
you can create the handler from a notifier directly:

	projectID  := int64(1)
	projectKey := "key"

	notifier := gobrake.NewNotifier(projectID, projectKey)
	handler := airbrake.NewHandlerFromNotifier(notifier)

By default Gobrake sends errors asynchronously and expects to be closed before the program finishes:

	func main() {
		defer handler.Close()
	}

If you want to Flush notices you can do it as you would with Gobrake's notifier
or you can configure the handler to send notices synchronously:

	handler := airbrake.NewHandlerFromNotifier(notifier, airbrake.SendSynchronously(true))
*/
package airbrake

import (
	"github.com/airbrake/gobrake"
	"github.com/goph/emperror"
	"github.com/goph/emperror/httperr"
	"github.com/goph/emperror/internal/keyvals"
)

// Option configures a logger instance.
type Option interface {
	apply(*Handler)
}

// SendSynchronously configures the handler to send notices synchronously.
type SendSynchronously bool

func (o SendSynchronously) apply(l *Handler) {
	l.sendSynchronously = bool(o)
}

// Handler is responsible for sending errors to Airbrake/Errbit.
type Handler struct {
	notifier *gobrake.Notifier

	sendSynchronously bool
}

// NewHandler creates a new Airbrake handler.
func NewHandler(projectID int64, projectKey string, opts ...Option) *Handler {
	h := &Handler{
		notifier: gobrake.NewNotifier(projectID, projectKey),
	}

	for _, o := range opts {
		o.apply(h)
	}

	return h
}

// NewHandlerFromNotifier creates a new Airbrake handler from a notifier instance.
func NewHandlerFromNotifier(notifier *gobrake.Notifier, opts ...Option) *Handler {
	h := &Handler{
		notifier: notifier,
	}

	for _, o := range opts {
		o.apply(h)
	}

	return h
}

// Handle calls the underlying Airbrake notifier.
func (h *Handler) Handle(err error) {
	// Get HTTP request (if any)
	req, _ := httperr.HTTPRequest(err)

	// Expose the stackTracer interface on the outer error (if there is stack trace in the error)
	err = emperror.ExposeStackTrace(err)

	notice := h.notifier.Notice(err, req, 1)

	// Extract context from the error and attach it to the notice
	if kvs := emperror.Context(err); len(kvs) > 0 {
		notice.Params = keyvals.ToMap(kvs)
	}

	if h.sendSynchronously {
		_, _ = h.notifier.SendNotice(notice)
	} else {
		h.notifier.SendNoticeAsync(notice)
	}
}

// Close closes the underlying Airbrake instance.
func (h *Handler) Close() error {
	return h.notifier.Close()
}
