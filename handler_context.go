package emperror

import (
	"context"

	"emperror.dev/errors"
)

type errorHandlerContext struct {
	handler   ErrorHandlerFacade
	extractor ContextExtractor
}

// NewErrorHandlerContext returns an error handler that extracts details from the provided context (if any)
// and annotates the handled error with them.
func NewErrorHandlerContext(handler ErrorHandler, extractor ContextExtractor) ErrorHandlerFacade {
	return errorHandlerContext{
		handler:   ensureErrorHandlerFacade(handler),
		extractor: extractor,
	}
}

func (e errorHandlerContext) Handle(err error) {
	e.handler.Handle(err)
}

func (e errorHandlerContext) HandleContext(ctx context.Context, err error) {
	fields := e.extractor(ctx)

	details := make([]interface{}, 0, len(fields)*2)

	for key, value := range fields {
		details = append(details, key, value)
	}

	e.handler.HandleContext(ctx, errors.WithDetails(err, details...))
}

// ContextExtractor extracts a map of details from a context.
type ContextExtractor func(ctx context.Context) map[string]interface{}

// ContextExtractors combines a list of ContextExtractor.
// The returned extractor aggregates the result of the underlying extractors.
func ContextExtractors(extractors ...ContextExtractor) ContextExtractor {
	return func(ctx context.Context) map[string]interface{} {
		fields := make(map[string]interface{})

		for _, extractor := range extractors {
			for key, value := range extractor(ctx) {
				fields[key] = value
			}
		}

		return fields
	}
}

// The implementation bellow is heavily influenced by go-kit's log context.

// HandlerWith returns a new error handler with keyvals context appended to it.
// If the wrapped error handler is already a contextual error handler created by HandlerWith or HandlerWithPrefix
// keyvals is appended to the existing context, but a new error handler is returned.
//
// The created handler will prepend it's own context to the handled errors.
// Deprecated: use WithDetails instead.
func HandlerWith(handler Handler, keyvals ...interface{}) Handler {
	return WithDetails(handler, keyvals...)
}

// HandlerWithPrefix returns a new error handler with keyvals context prepended to it.
// If the wrapped error handler is already a contextual error handler created by HandlerWith or HandlerWithPrefix
// keyvals is prepended to the existing context, but a new error handler is returned.
//
// The created handler will prepend it's own context to the handled errors.
// Deprecated: no replacement at this time.
func HandlerWithPrefix(handler Handler, keyvals ...interface{}) ErrorHandlerFacade {
	handlerFacade := ensureErrorHandlerFacade(handler)

	if len(keyvals) == 0 {
		return handlerFacade
	}

	prevkvs, handlerFacade := extractHandlerContext(handlerFacade)

	n := len(prevkvs) + len(keyvals)
	if len(keyvals)%2 != 0 {
		n++
	}

	kvs := make([]interface{}, 0, n)
	kvs = append(kvs, keyvals...)

	if len(kvs)%2 != 0 {
		kvs = append(kvs, nil)
	}

	kvs = append(kvs, prevkvs...)

	return newContextualHandler(handlerFacade, kvs)
}

// extractHandlerContext extracts the context and optionally the wrapped handler when it's the same container.
func extractHandlerContext(handler ErrorHandlerFacade) ([]interface{}, ErrorHandlerFacade) {
	var kvs []interface{}

	if c, ok := handler.(*contextualHandler); ok {
		handler = c.handler
		kvs = c.keyvals
	}

	return kvs, handler
}

// contextualHandler is a Handler implementation returned by HandlerWith or HandlerWithPrefix.
//
// It wraps an error handler and a holds keyvals as the context.
type contextualHandler struct {
	handler ErrorHandlerFacade
	keyvals []interface{}
}

// newContextualHandler creates a new *contextualHandler or a struct which is contextual and holds a stack trace.
func newContextualHandler(handler ErrorHandlerFacade, kvs []interface{}) ErrorHandlerFacade {
	chandler := &contextualHandler{
		handler: handler,
		keyvals: kvs,
	}

	return chandler
}

// Handle prepends the handler's context to the error's (if any) and delegates the call to the underlying handler.
func (h *contextualHandler) Handle(err error) {
	err = errors.WithDetails(err, h.keyvals...)

	h.handler.Handle(err)
}

// HandleContext prepends the handler's context to the error's (if any)
// and delegates the call to the underlying handler.
func (h *contextualHandler) HandleContext(ctx context.Context, err error) {
	err = errors.WithDetails(err, h.keyvals...)

	h.handler.HandleContext(ctx, err)
}
