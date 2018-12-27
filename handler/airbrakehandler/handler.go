// Package airbrakehandler provides Airbrake/Errbit integration.
package airbrakehandler

import (
	"github.com/airbrake/gobrake"
	"github.com/goph/emperror"
	"github.com/goph/emperror/httperr"
	"github.com/goph/emperror/internal/keyvals"
)

// Handler is responsible for sending errors to Airbrake/Errbit.
type Handler struct {
	notifier *gobrake.Notifier

	sendSynchronously bool
}

// New creates a new handler.
func New(projectID int64, projectKey string) *Handler {
	return NewFromNotifier(gobrake.NewNotifier(projectID, projectKey))
}

// NewSync creates a new handler that sends errors synchronously.
func NewSync(projectID int64, projectKey string) *Handler {
	handler := New(projectID, projectKey)

	handler.sendSynchronously = true

	return handler
}

// NewFromNotifier creates a new handler from a notifier instance.
func NewFromNotifier(notifier *gobrake.Notifier) *Handler {
	handler := &Handler{
		notifier: notifier,
	}

	return handler
}

// NewSyncFromNotifier creates a new handler from a notifier instance that sends errors synchronously.
func NewSyncFromNotifier(notifier *gobrake.Notifier) *Handler {
	handler := NewFromNotifier(notifier)

	handler.sendSynchronously = true

	return handler
}

// Handle sends the error to Airbrake/Errbit.
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

// Close closes the underlying notifier and waits for asynchronous reports to finish.
func (h *Handler) Close() error {
	return h.notifier.Close()
}
