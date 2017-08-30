package emperror

// TestHandler throws every error away.
type TestHandler struct {
	errors []error
}

// Handle saves the error in a list.
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
