// Package errors extends the errors package in the stdlib.
//
// Despite the implicit nature of interface satisfication in Go
// this package exports a number of interfaces to avoid defining them over and over again.
// Although it means coupling between the consumer code and this package,
// the purpose of this library (being a stdlib extension) justifies that.
package errors

import (
	std__errors "errors"

	"github.com/pkg/errors"
)

// New returns an error that formats as the given text.
//
// This is an alias to the stdlib errors.New function.
func New(text string) error {
	return std__errors.New(text)
}

// NewWithStackTrace returns an error that formats as the given text and contains stack trace.
//
// This is an alias to github.com/pkg/errors.New function.
func NewWithStackTrace(text string) error {
	return errors.New(text)
}

// WithStack annotates err with a stack trace at the point WithStack was called.
//
// If err is nil, WithStack returns nil.
//
// This is an alias to github.com/pkg/errors.WithStack function.
func WithStack(err error) error {
	return errors.WithStack(err)
}

// WithMessage annotates err with a new message.
//
// If err is nil, WithMessage returns nil.
//
// This is an alias to github.com/pkg/errors.WithMessage function.
func WithMessage(err error, message string) error {
	return errors.WithMessage(err, message)
}

// Cause returns the underlying cause of the error, if possible.
// An error value has a cause if it implements the Causer interface.
//
// If the error does not implement Cause, the original error will
// be returned. If the error is nil, nil will be returned without further
// investigation.
//
// This is an alias to github.com/pkg/errors.Cause function.
func Cause(err error) error {
	return errors.Cause(err)
}
