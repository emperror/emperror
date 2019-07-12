package emperror

import "emperror.dev/errors"

// WrapWith returns an error annotating err with a stack trace
// at the point Wrap is called (if there is none attached to the error yet), the supplied message,
// and the supplied context.
// If err is nil, Wrap returns nil.
//
// Note: do not use this method when passing errors between goroutines.
func WrapWith(err error, message string, keyvals ...interface{}) error {
	if err == nil {
		return nil
	}

	err = errors.WithStackDepthIf(errors.WithMessage(err, message), 1)

	// Attach context to the error
	if len(keyvals) > 0 {
		err = With(err, keyvals...)
	}

	return err
}
