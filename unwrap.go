package emperror

import (
	"emperror.dev/errors"
)

// UnwrapEach loops through an error chain and calls a function for each of them.
//
// The provided function can return false to break the loop before it reaches the end of the chain.
//
// It supports both Go 1.13 errors.Wrapper and github.com/pkg/errors.Causer interfaces
// (the former takes precedence).
func UnwrapEach(err error, fn func(err error) bool) {
	for err != nil {
		continueLoop := fn(err)
		if !continueLoop {
			break
		}

		err = errors.Unwrap(err)
	}
}

// ForEachCause loops through an error chain and calls a function for each of them,
// starting with the topmost one.
//
// The function can return false to break the loop before it ends.
//
// Deprecated: use UnwrapEach instead.
func ForEachCause(err error, fn func(err error) bool) {
	// causer is the interface defined in github.com/pkg/errors for specifying a parent error.
	type causer interface {
		Cause() error
	}

	for err != nil {
		continueLoop := fn(err)
		if !continueLoop {
			break
		}

		cause, ok := err.(causer)
		if !ok {
			break
		}

		err = cause.Cause()
	}
}
