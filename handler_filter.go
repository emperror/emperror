package emperror

import (
	"context"
)

type filterHandler struct {
	errorMatcher ErrorMatcher
	handler      ErrorHandlerFacade
}

func (h filterHandler) Handle(err error) {
	if h.errorMatcher(err) {
		return
	}

	h.handler.Handle(err)
}

func (h filterHandler) HandleContext(ctx context.Context, err error) {
	if h.errorMatcher(err) {
		return
	}

	h.handler.HandleContext(ctx, err)
}

// ErrorMatcher checks if an error matches a certain condition.
type ErrorMatcher func(err error) bool

// WithDetails returns a new error handler that discards errors matching any of the specified filters.
// Otherwise it passes errors to the next handler.
func WithFilter(handler ErrorHandler, matcher ErrorMatcher) ErrorHandlerFacade {
	return filterHandler{
		errorMatcher: matcher,
		handler:      ensureErrorHandlerFacade(handler),
	}
}
