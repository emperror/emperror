package emperror

import "github.com/pkg/errors"

// ErrorCollection holds a list of errors.
type ErrorCollection interface {
	Errors() []error
}

// StackTracer is the interface defined in github.com/pkg/errors for exposing stack trace from an error.
type StackTracer interface {
	StackTrace() errors.StackTrace
}
