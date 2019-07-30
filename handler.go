package emperror

import (
	"context"

	"emperror.dev/errors"
)

// Handler is a generic error handler. It allows applications (and libraries) to handle errors
// without worrying about the actual error handling strategy (logging, error tracking service, etc).
type Handler interface {
	// Handle handles an error.
	Handle(err error)
}

// ContextAwareHandler is similar to Handler, except it receives a context as well.
// It is useful in request terminal error handling situations.
// An implementation MAY extract information from the context and annotate err with it.
type ContextAwareHandler interface {
	// Handle handles an error.
	Handle(ctx context.Context, err error)
}

// Handlers collects a number of error handlers into a single one.
type Handlers []Handler

func (h Handlers) Handle(err error) {
	for _, handler := range h {
		handler.Handle(err)
	}
}

// Close calls Close on the underlying handlers (if there is any closable handler).
func (h Handlers) Close() error {
	if len(h) < 1 {
		return nil
	}

	errs := make([]error, len(h))

	for i, handler := range h {
		if closer, ok := handler.(interface{ Close() error }); ok {
			errs[i] = closer.Close()
		}
	}

	return errors.Combine(errs...)
}

// HandlerFunc wraps a function and turns it into an error handler.
type HandlerFunc func(err error)

// Handle calls the underlying function.
func (h HandlerFunc) Handle(err error) {
	h(err)
}

// Handle handles an error (if one occurred).
func Handle(handler Handler, err error) {
	if err != nil {
		handler.Handle(err)
	}
}

type noopHandler struct{}

// NewNoopHandler creates a no-op error handler that discards all received errors.
// Useful in examples and as a fallback error handler.
func NewNoopHandler() Handler {
	return &noopHandler{}
}

func (*noopHandler) Handle(err error) {}

// MakeContextAware wraps an error handler and turns it into a ContextAwareHandler.
func MakeContextAware(handler Handler) ContextAwareHandler {
	return &contextAwareHandler{handler}
}

// contextAwareHandler wraps an error handler and turns it into a ContextAwareHandler.
type contextAwareHandler struct {
	handler Handler
}

// Handle calls the underlying error handler.
func (h contextAwareHandler) Handle(ctx context.Context, err error) {
	h.handler.Handle(err)
}
