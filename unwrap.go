package emperror

// Unwrap returns the next error in an error chain (if any).
// Otherwise it returns nil.
//
// It supports both Go 1.13 errors.Wrapper and github.com/pkg/errors.Causer interfaces
// (the former takes precedence).
func Unwrap(err error) error {
	switch e := err.(type) {
	case interface{ Unwrap() error }:
		return e.Unwrap()

	case interface{ Cause() error }:
		return e.Cause()
	}

	return nil
}

// UnwrapAll returns the last error (root cause) in an error chain.
// If the error has no cause, it is returned directly.
//
// It supports both Go 1.13 errors.Wrapper and github.com/pkg/errors.Causer interfaces
// (the former takes precedence).
func UnwrapAll(err error) error {
	for {
		cause := Unwrap(err)
		if cause == nil {
			break
		}

		err = cause
	}

	return err
}

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

		err = Unwrap(err)
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
