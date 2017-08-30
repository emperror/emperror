package emperror

// nopHandler doesn't do anything.
type nopHandler struct{}

// NewNopHandler returns an error handler that doesn't do anything.
func NewNopHandler() Handler { return nopHandler{} }

// Handle implements the Handler interface and it doesn't do anything.
func (nopHandler) Handle(error) {}
