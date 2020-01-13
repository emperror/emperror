package emperror

import (
	"context"

	"emperror.dev/errors"
)

// ErrorHandler is a generic error handler that allows applications (and libraries) to handle errors
// without worrying about the actual error handling strategy (logging, error tracking service, etc).
type ErrorHandler interface {
	// Handle handles an error.
	//
	// If err is nil, Handle should immediately return.
	Handle(err error)
}

// ErrorHandlerContext is an optional interface that MAY be implemented by an ErrorHandler.
// It is similar to ErrorHandler, but it receives a context as the first parameter.
// An implementation MAY extract information from the context and annotate err with it.
//
// ErrorHandlerContext MAY honor the deadline carried by the context, but that's not a hard requirement.
type ErrorHandlerContext interface {
	// HandleContext handles an error.
	//
	// If err is nil, HandleContext should immediately return.
	HandleContext(ctx context.Context, err error)
}

// ErrorHandlerSet is a combination of ErrorHandler and ErrorHandlerContext.
// It's sole purpose is to make the API of the package concise by exposing a common interface type
// for return values. It's not supposed to be used by consumers of this package.
//
// It goes directly against the "Use interfaces, return structs" idiom of Go,
// but at the current phase of the package the smaller API surface makes more sense.
//
// In the future it might get replaced with concrete types.
type ErrorHandlerSet interface {
	ErrorHandler
	ErrorHandlerContext
}

// ErrorHandlers combines a number of error handlers into a single one.
type ErrorHandlers []ErrorHandler

func (h ErrorHandlers) Handle(err error) {
	for _, handler := range h {
		handler.Handle(err)
	}
}

func (h ErrorHandlers) HandleContext(ctx context.Context, err error) {
	for _, handler := range h {
		if handlerCtx, ok := handler.(ErrorHandlerContext); ok {
			handlerCtx.HandleContext(ctx, err)
		} else {
			handler.Handle(err)
		}
	}
}

// Close calls Close on the underlying handlers (if there is any closable handler).
func (h ErrorHandlers) Close() error {
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

// ErrorHandlerFunc wraps a function and turns it into an ErrorHandler
// if the function's definition matches the interface.
type ErrorHandlerFunc func(err error)

func (h ErrorHandlerFunc) Handle(err error) {
	h(err)
}

func (h ErrorHandlerFunc) HandleContext(_ context.Context, err error) {
	h(err)
}

// ErrorHandlerContextFunc wraps a function and turns it into an ErrorHandlerContext
// if the function's definition matches the interface.
type ErrorHandlerContextFunc func(ctx context.Context, err error)

func (h ErrorHandlerContextFunc) Handle(err error) {
	h(context.Background(), err)
}

func (h ErrorHandlerContextFunc) HandleContext(ctx context.Context, err error) {
	h(ctx, err)
}

// NoopHandler is a no-op error handler that discards all received errors.
//
// It implements both ErrorHandler and ErrorHandlerContext interfaces.
type NoopHandler struct{}

func (NoopHandler) Handle(_ error) {}

func (NoopHandler) HandleContext(_ context.Context, _ error) {}

func ensureErrorHandlerSet(handler ErrorHandler) ErrorHandlerSet {
	if handler, ok := handler.(ErrorHandlerSet); ok {
		return handler
	}

	return ErrorHandlerFunc(handler.Handle)
}

// Handler is a generic error handler. It allows applications (and libraries) to handle errors
// without worrying about the actual error handling strategy (logging, error tracking service, etc).
//
// Deprecated: use ErrorHandler instead.
type Handler interface {
	// Handle handles an error.
	Handle(err error)
}

// ContextAwareHandler is similar to Handler, except it receives a context as well.
// It is useful in request terminal error handling situations.
// An implementation MAY extract information from the context and annotate err with it.
//
// Deprecated: user ErrorHandlerContext instead.
type ContextAwareHandler interface {
	// Handle handles an error.
	Handle(ctx context.Context, err error)
}

// Handlers collects a number of error handlers into a single one.
//
// Deprecated: use ErrorHandlers instead.
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
//
// Deprecated: use ErrorHandlerFunc.
type HandlerFunc func(err error)

// Handle calls the underlying function.
func (h HandlerFunc) Handle(err error) {
	h(err)
}

// Handle handles an error (if one occurred).
//
// Deprecated: no replacement. ErrorHandler should return if err is nil.
func Handle(handler Handler, err error) {
	if err != nil {
		handler.Handle(err)
	}
}

// NewNoopHandler creates a no-op error handler that discards all received errors.
// Useful in examples and as a fallback error handler.
//
// Deprecated: use NoopHandler.
func NewNoopHandler() Handler {
	return NoopHandler{}
}

// MakeContextAware wraps an error handler and turns it into a ContextAwareHandler.
//
// Deprecated: no replacement at this time.
func MakeContextAware(handler Handler) ContextAwareHandler {
	return &contextAwareHandler{handler}
}

// contextAwareHandler wraps an error handler and turns it into a ContextAwareHandler.
type contextAwareHandler struct {
	handler Handler
}

// Handle calls the underlying error handler.
func (h contextAwareHandler) Handle(_ context.Context, err error) {
	h.handler.Handle(err)
}
