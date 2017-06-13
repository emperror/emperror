package emperror

// NullHandler throws every error away.
// Can be used as a fallback.
type NullHandler struct{}

// NewNullHandler returns a new NullHandler.
func NewNullHandler() Handler {
	return &NullHandler{}
}

// Handle does the actual throwing away.
func (h *NullHandler) Handle(err error) {
}
