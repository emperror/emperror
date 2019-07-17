package emperror

import (
	"emperror.dev/errors"
)

// The implementation bellow is heavily influenced by go-kit's log context.

// With returns a new error with keyvals context appended to it.
// If the wrapped error is already a contextual error created by With
// keyvals is appended to the existing context, but a new error is returned.
// Deprecated: use emperror.dev/errors.WithDetails instead.
func With(err error, keyvals ...interface{}) error {
	return errors.WithDetails(err, keyvals...)
}

// Context extracts the context key-value pairs from an error (or error chain).
// Deprecated: use emperror.dev/errors.GetDetails instead.
func Context(err error) []interface{} {
	type contextor interface {
		Context() []interface{}
	}

	var kvs []interface{}

	errors.UnwrapEach(err, func(err error) bool {
		if cerr, ok := err.(contextor); ok {
			kvs = append(cerr.Context(), kvs...)
		}

		return true
	})

	kvs = append(kvs, errors.GetDetails(err)...)

	return kvs
}
