package testing

import "testing"

// Handler wraps a test state and outputs an error.
type Handler struct {
	T *testing.T
}

// NewHandler returns a new Handler.
func NewHandler(t *testing.T) *Handler {
	return &Handler{t}
}

// Handle calls the underlying test state.
func (h *Handler) Handle(err error) {
	h.T.Error(err)
}
