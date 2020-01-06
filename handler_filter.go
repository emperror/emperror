package emperror

type filterHandler struct {
	matcher ErrorMatcher
	handler Handler
}

func (h filterHandler) Handle(err error) {
	if h.matcher.MatchError(err) {
		return
	}

	h.handler.Handle(err)
}

// WithDetails returns a new error handler that discards errors matching any of the specified filters.
// Otherwise it passes errors to the next handler.
func WithFilter(handler Handler, matcher ErrorMatcher) Handler {
	return filterHandler{
		matcher: matcher,
		handler: handler,
	}
}

// ErrorMatcher checks if an error matches a certain condition.
type ErrorMatcher interface {
	// MatchError checks if err matches a certain condition.
	MatchError(err error) bool
}
