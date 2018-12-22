package emperror

import "sync"

// TestHandler is a simple stub for the handler interface recording every error.
//
// The TestHandler is safe for concurrent use.
type TestHandler struct {
	errors []error

	mu sync.RWMutex
}

// NewTestHandler returns a new TestHandler.
func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

// Handle saves the error in a list.
func (h *TestHandler) Handle(err error) {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.errors = append(h.errors, err)
}

// Errors returns all the handled errors.
func (h *TestHandler) Errors() []error {
	h.mu.RLock()
	defer h.mu.RUnlock()

	return h.errors
}

// Last returns the last handled error.
func (h *TestHandler) Last() error {
	h.mu.RLock()
	defer h.mu.RUnlock()

	if len(h.errors) < 1 {
		return nil
	}

	return h.errors[len(h.errors)-1]
}
