package emperror

import (
	"emperror.dev/errors"
)

// HandlerWithDetails returns a new error handler annotated with key-value pairs.
//
// The created handler will add it's own details to the handled errors.
func HandlerWithDetails(handler Handler, details ...interface{}) Handler {
	if len(details) == 0 {
		return handler
	}

	d, handler := extractHandlerDetails(handler)

	d = append(d, details...)

	if len(d)%2 != 0 {
		d = append(d, nil)
	}

	// Limiting the capacity of the stored keyvals ensures that a new
	// backing array is created if the slice must grow in HandlerWith.
	// Using the extra capacity without copying risks a data race.
	return newHandlerWithDetails(handler, d[:len(d):len(d)])
}

// extractHandlerDetails extracts the context and optionally the wrapped handler when it's the same container.
func extractHandlerDetails(handler Handler) ([]interface{}, Handler) {
	var d []interface{}

	if c, ok := handler.(*handlerWithDetails); ok {
		handler = c.handler
		d = c.details
	}

	return d, handler
}

// handlerWithDetails is a Handler implementation returned by HandlerWith or HandlerWithPrefix.
//
// It wraps an error handler and a holds keyvals as the context.
type handlerWithDetails struct {
	handler Handler
	details []interface{}
}

// newHandlerWithDetails creates a new handler annotated with a set of key-value pairs.
func newHandlerWithDetails(handler Handler, details []interface{}) Handler {
	chandler := &handlerWithDetails{
		handler: handler,
		details: details,
	}

	return chandler
}

// Handle adds the handler's details to err and delegates the call to the underlying handler.
func (h *handlerWithDetails) Handle(err error) {
	// TODO: prepend details so the error takes higher precedence
	err = errors.WithDetails(err, h.details...)

	h.handler.Handle(err)
}
