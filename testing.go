package emperror

// TestHandler throws every error away.
type TestHandler struct {
	errors []error
}

// NewTestHandler returns a new TestHandler.
func NewTestHandler() Handler {
	return &TestHandler{}
}

// Handle does the actual throwing away.
func (h *TestHandler) Handle(err error) {
	h.errors = append(h.errors, err)
}

// Errors returns all the handled errors.
func (h *TestHandler) Errors() []error {
	return h.errors
}

// Last returns the last handled error.
func (h *TestHandler) Last() error {
	if len(h.errors) < 1 {
		return nil
	}

	return h.errors[len(h.errors)-1]
}
