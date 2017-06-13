package emperror

// CompositeHandler allows an error to be processed by multiple handlers.
type CompositeHandler struct {
	handlers []Handler
}

// NewCompositeHandler returns a new CompositeHandler.
func NewCompositeHandler(handlers ...Handler) Handler {
	return &CompositeHandler{handlers}
}

// Handle goes through the handlers and call each of them for the error.
func (h *CompositeHandler) Handle(err error) {
	for _, handler := range h.handlers {
		handler.Handle(err)
	}
}
