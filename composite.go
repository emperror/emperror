package emperror

import "github.com/goph/stdlib/errors"

// compositeHandler allows an error to be processed by multiple handlers.
type compositeHandler struct {
	handlers []errors.Handler
}

// NewCompositeHandler returns a new compositeHandler.
func NewCompositeHandler(handlers ...errors.Handler) errors.Handler {
	return &compositeHandler{handlers}
}

// Handle goes through the handlers and call each of them for the error.
func (h *compositeHandler) Handle(err error) {
	for _, handler := range h.handlers {
		handler.Handle(err)
	}
}
