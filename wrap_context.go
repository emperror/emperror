package emperror

import (
	"fmt"

	"github.com/pkg/errors"
)

// WrapWith returns an error annotating err with a stack trace
// at the point WrapWith is called (if there is none attached to the error
// yet), the supplied message, and the supplied context.
// If err is nil, WrapWith returns nil.
//
// Note: do not use this method when passing errors between goroutines.
func WrapWith(err error, message string, keyvals ...interface{}) error {
	if err == nil {
		return nil
	}

	_, ok := getStackTracer(err)

	err = errors.WithMessage(err, message)

	// There is no stack trace in the error, so attach it here
	if !ok {
		err = &wrappedError{
			err:   err,
			stack: callers(),
		}
	}

	// Attach context to the error
	if len(keyvals) > 0 {
		err = With(err, keyvals...)
	}

	return err
}

// WrapfWith returns an error annotating err with a stack trace
// at the point WrapfWith is called (if there is none attached to the error
// yet), the format specifier substituted with the values from the key-value
// pairs, and the supplied context.  If err is nil, WrapfWith returns nil.
//
// Note: do not use this method when passing errors between goroutines.
func WrapfWith(err error, format string, keyvals ...interface{}) error {
	if err == nil {
		return nil
	}

	// Use every second argument (i.e. the values from the key-value pairs) as formatting parameter
	args := make([]interface{}, 0, len(keyvals)/2)

	for i, keyval := range keyvals {
		if i%2 == 1 {
			args = append(args, keyval)
		}
	}

	err = errors.WithMessage(err, fmt.Sprintf(format, args...))

	_, ok := getStackTracer(err)

	// There is no stack trace in the error, so attach it here
	if !ok {
		err = &wrappedError{
			err:   err,
			stack: callers(),
		}
	}

	// Attach context to the error
	if len(keyvals) > 0 {
		err = With(err, keyvals...)
	}

	return err
}
