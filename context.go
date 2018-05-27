package emperror

// The implementation bellow is heavily influenced by go-kit's log context.

// With returns a new error with keyvals context appended to it.
// If the wrapped error is already a contextual error created by With or WithPrefix
// keyvals is appended to the existing context, but a new error is returned.
func With(err error, keyvals ...interface{}) error {
	if len(keyvals) == 0 {
		return err
	}

	kvs, err := extractContext(err)

	kvs = append(kvs, keyvals...)

	if len(kvs)%2 != 0 {
		kvs = append(kvs, nil)
	}

	// Limiting the capacity of the stored keyvals ensures that a new
	// backing array is created if the slice must grow in With.
	// Using the extra capacity without copying risks a data race.
	return newContextualError(err, kvs[:len(kvs):len(kvs)])
}

// WithPrefix returns a new error with keyvals context appended to it.
// If the wrapped error is already a contextual error created by With or WithPrefix
// keyvals is prepended to the existing context, but a new error is returned.
func WithPrefix(err error, keyvals ...interface{}) error {
	if len(keyvals) == 0 {
		return err
	}

	prevkvs, err := extractContext(err)

	n := len(prevkvs) + len(keyvals)
	if len(keyvals)%2 != 0 {
		n++
	}

	kvs := make([]interface{}, 0, n)
	kvs = append(kvs, keyvals...)

	if len(kvs)%2 != 0 {
		kvs = append(kvs, nil)
	}

	kvs = append(kvs, prevkvs...)

	return newContextualError(err, kvs)
}

// Context extracts the context key-value pairs from an error (or error chain).
func Context(err error) []interface{} {
	type contextor interface {
		Context() []interface{}
	}

	var kvs []interface{}

	ForEachCause(err, func(err error) bool {
		if cerr, ok := err.(contextor); ok {
			kvs = append(kvs, cerr.Context()...)
		}

		return true
	})

	return kvs
}

// extractContext extracts the context and optionally the wrapped error when it's the same container.
func extractContext(err error) ([]interface{}, error) {
	var kvs []interface{}

	if c, ok := err.(*contextualError); ok {
		err = c.error
		kvs = c.keyvals
	} else if c, ok := err.(Contextor); ok {
		kvs = c.Context()
	}

	return kvs, err
}

// contextualError is the ContextualError implementation returned by With or WithPrefix.
//
// It wraps an error and a holds keyvals as the context.
type contextualError struct {
	error
	keyvals []interface{}
}

// newContextualError creates a new *contextualError or a struct which is contextual and holds a stack trace.
func newContextualError(err error, kvs []interface{}) error {
	cerr := &contextualError{
		error:   err,
		keyvals: kvs,
	}

	if serr, ok := err.(StackTracer); ok {
		return struct {
			*contextualError
			StackTracer
		}{
			contextualError: cerr,
			StackTracer:     serr,
		}
	}

	return cerr
}

// Context returns the appended keyvals.
func (e *contextualError) Context() []interface{} {
	return e.keyvals
}

// Cause returns the underlying error.
//
// This method fulfills the causer interface described in github.com/pkg/errors.
func (e *contextualError) Cause() error {
	return e.error
}
