package emperror

import "github.com/goph/stdlib/errors"

// nullHandler throws every error away.
// Can be used as a fallback.
type nullHandler struct{}

// NewNullHandler returns a new nullHandler.
func NewNullHandler() errors.Handler {
	return &nullHandler{}
}

// Handle throws the error away.
func (h *nullHandler) Handle(err error) {}
