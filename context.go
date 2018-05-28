package emperror

// The implementation bellow is heavily influenced by go-kit's log context.

// With returns a new error with keyvals context appended to it.
// If the wrapped error is already a contextual error created by With
// keyvals is appended to the existing context, but a new error is returned.
func With(err error, keyvals ...interface{}) error {
	if len(keyvals) == 0 {
		return err
	}

	var kvs []interface{}

	// extract context from previous error
	if c, ok := err.(*contextualError); ok {
		err = c.error

		kvs = append(kvs, c.keyvals...)

		if len(kvs)%2 != 0 {
			kvs = append(kvs, nil)
		}
	}

	kvs = append(kvs, keyvals...)

	if len(kvs)%2 != 0 {
		kvs = append(kvs, nil)
	}

	// Limiting the capacity of the stored keyvals ensures that a new
	// backing array is created if the slice must grow in With.
	// Using the extra capacity without copying risks a data race.
	return &contextualError{
		error:   err,
		keyvals: kvs[:len(kvs):len(kvs)],
	}
}

// Context extracts the context key-value pairs from an error (or error chain).
func Context(err error) []interface{} {
	type contextor interface {
		Context() []interface{}
	}

	var kvs []interface{}

	ForEachCause(err, func(err error) bool {
		if cerr, ok := err.(contextor); ok {
			kvs = append(cerr.Context(), kvs...)
		}

		return true
	})

	return kvs
}

// contextualError is the ContextualError implementation returned by With.
//
// It wraps an error and a holds keyvals as the context.
type contextualError struct {
	error
	keyvals []interface{}
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
