package errors

import "github.com/pkg/errors"

// ErrorCollection holds a list of errors.
type ErrorCollection interface {
	Errors() []error
}

// Contextor represents an error which holds a context.
type Contextor interface {
	Context() []interface{}
}

// Causer is the interface defined in github.com/pkg/errors for specifying a parent error.
type Causer interface {
	Cause() error
}

// StackTracer is the interface defined in github.com/pkg/errors for exposing stack trace from an error.
type StackTracer interface {
	StackTrace() errors.StackTrace
}

// Handler is responsible for handling an error.
//
// This interface allows libraries to decouple from logging and error handling solutions.
type Handler interface {
	// Handle processes an error.
	Handle(err error)
}
