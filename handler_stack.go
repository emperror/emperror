package emperror

import (
	"fmt"

	"emperror.dev/errors"
)

// WithStack annotates every error passing through the handler with the
// function name and file line of the stack trace's top frame (if one is found).
func WithStack(handler Handler) Handler {
	if handler == nil {
		return NewNoopHandler()
	}

	return &stackHandler{handler}
}

type stackHandler struct {
	handler Handler
}

func (h *stackHandler) Handle(err error) {
	if err == nil {
		return
	}

	var stackTracer interface{ StackTrace() errors.StackTrace }
	if errors.As(err, &stackTracer) {
		stackTrace := stackTracer.StackTrace()

		if len(stackTrace) > 0 {
			frame := stackTrace[0]

			err = errors.WithDetails(
				err,
				"func", fmt.Sprintf("%n", frame), // nolint: govet
				"file", fmt.Sprintf("%v", frame), // nolint: govet
			)
		}
	}

	h.handler.Handle(err)
}
