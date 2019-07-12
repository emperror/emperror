package emperror

import (
	"emperror.dev/errors"
)

// Wrap returns an error annotating err with a stack trace
// at the point Wrap is called (if there is none attached to the error yet), and the supplied message.
// If err is nil, Wrap returns nil.
//
// Note: do not use this method when passing errors between goroutines.
func Wrap(err error, message string) error {
	if err == nil {
		return nil
	}

	err = errors.WithMessage(err, message)

	// There is no stack trace in the error, so attach it here
	var st stackTracer
	if !errors.As(err, &st) {
		return errors.WithStackDepth(err, 1)
	}

	return err
}

// Wrapf returns an error annotating err with a stack trace
// at the point Wrapf is call (if there is none attached to the error yet), and the format specifier.
// If err is nil, Wrapf returns nil.
//
// Note: do not use this method when passing errors between goroutines.
func Wrapf(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}

	err = errors.WithMessagef(err, format, args...)

	// There is no stack trace in the error, so attach it here
	var st stackTracer
	if !errors.As(err, &st) {
		return errors.WithStackDepth(err, 1)
	}

	return err
}
