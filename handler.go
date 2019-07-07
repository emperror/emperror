package emperror

// Handler is a generic error handler. It allows applications (and libraries) to handle errors
// without worrying about the actual error handling strategy (logging, error tracking service, etc).
type Handler interface {
	// Handle handles an error.
	Handle(err error)
}

// Handlers collects a number of error handlers into a single one.
type Handlers []Handler

func (h Handlers) Handle(err error) {
	if len(h) < 1 {
		return
	}

	for _, handler := range h {
		handler.Handle(err)
	}
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
