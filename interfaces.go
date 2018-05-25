package emperror

import "github.com/pkg/errors"

// ErrorCollection holds a list of errors.
type ErrorCollection interface {
	Errors() []error
}

// Contextor represents an error which holds a context.
type Contextor interface {
	Context() []interface{}
}

// StackTracer is the interface defined in github.com/pkg/errors for exposing stack trace from an error.
type StackTracer interface {
	StackTrace() errors.StackTrace
}
