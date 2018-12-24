// Package airbrakehandler provides Airbrake/Errbit integration.
package airbrakehandler

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
	l.sendAsynchronously = bool(o)
}

// Handler is responsible for sending errors to Airbrake/Errbit.
type Handler struct {
	notifier *gobrake.Notifier

	sendAsynchronously bool
}

// New creates a new Airbrake handler.
func New(projectID int64, projectKey string, opts ...Option) *Handler {
	return NewFromNotifier(gobrake.NewNotifier(projectID, projectKey), opts...)
}

// NewAsync creates a new Airbrake handler that sends errors asynchronously.
func NewAsync(projectID int64, projectKey string, opts ...Option) *Handler {
	h := New(projectID, projectKey, opts...)

	h.sendAsynchronously = true

	return h
}

// NewFromNotifier creates a new Airbrake handler from a notifier instance.
func NewFromNotifier(notifier *gobrake.Notifier, opts ...Option) *Handler {
	h := &Handler{
		notifier: notifier,
	}

	for _, o := range opts {
		o.apply(h)
	}

	return h
}

// NewAsyncFromNotifier creates a new Airbrake handler from a notifier instance that sends errors asynchronously.
func NewAsyncFromNotifier(notifier *gobrake.Notifier, opts ...Option) *Handler {
	h := NewFromNotifier(notifier, opts...)

	h.sendAsynchronously = true

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

	if h.sendAsynchronously {
		h.notifier.SendNoticeAsync(notice)
	} else {
		_, _ = h.notifier.SendNotice(notice)
	}
}

// Close closes the underlying Airbrake instance.
func (h *Handler) Close() error {
	return h.notifier.Close()
}
