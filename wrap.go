package emperror

import (
	"emperror.dev/errors"
)

// Wrap returns an error annotating err with a stack trace
// at the point Wrap is called (if there is none attached to the error yet), and the supplied message.
// If err is nil, Wrap returns nil.
//
// Note: do not use this method when passing errors between goroutines.
// Deprecated: use emperror.dev/errors.WrapIf instead.
func Wrap(err error, message string) error {
	return errors.WithStackDepthIf(errors.WithMessage(err, message), 1)
}

// Wrapf returns an error annotating err with a stack trace
// at the point Wrapf is call (if there is none attached to the error yet), and the format specifier.
// If err is nil, Wrapf returns nil.
//
// Note: do not use this method when passing errors between goroutines.
// Deprecated: use emperror.dev/errors.WrapIff instead.
func Wrapf(err error, format string, args ...interface{}) error {
	return errors.WithStackDepthIf(errors.WithMessagef(err, format, args...), 1)
}

// WrapWith returns an error annotating err with a stack trace
// at the point Wrap is called (if there is none attached to the error yet), the supplied message,
// and the supplied context.
// If err is nil, Wrap returns nil.
//
// Note: do not use this method when passing errors between goroutines.
// Deprecated: use emperror.dev/errors.WrapIfWithDetails instead.
func WrapWith(err error, message string, keyvals ...interface{}) error {
	return errors.WithDetails(errors.WithStackDepthIf(errors.WithMessage(err, message), 1), keyvals...)
}
