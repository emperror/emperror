package emperror

import (
	"context"
)

type filterHandler struct {
	matcher ErrorMatcher
	handler ErrorHandlerSet
}

func (h filterHandler) Handle(err error) {
	if h.matcher.MatchError(err) {
		return
	}

	h.handler.Handle(err)
}

func (h filterHandler) HandleContext(ctx context.Context, err error) {
	if h.matcher.MatchError(err) {
		return
	}

	h.handler.HandleContext(ctx, err)
}

// ErrorMatcher checks if an error matches a certain condition.
type ErrorMatcher interface {
	// MatchError checks if err matches a certain condition.
	MatchError(err error) bool
}

// WithDetails returns a new error handler that discards errors matching any of the specified filters.
// Otherwise it passes errors to the next handler.
func WithFilter(handler ErrorHandler, matcher ErrorMatcher) ErrorHandlerSet {
	return filterHandler{
		matcher: matcher,
		handler: ensureErrorHandlerSet(handler),
	}
}
