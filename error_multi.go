package emperror

import (
	"emperror.dev/errors"
)

// Errors is responsible for listing multiple errors.
// Deprecated: use multi error tools from emperror.dev/errors instead.
type Errors interface {
	// Errors returns the list of wrapped errors.
	Errors() []error
}

// multiError implements Errors and aggregates multiple errors into a single value.
// Also implements the error interface so it can be returned as an error.
type multiError struct {
	errors []error
	msg    string
}

// Error implements the error interface.
func (e *multiError) Error() string {
	if e.msg != "" {
		return e.msg
	}

	return "Multiple errors happened"
}

// Errors returns the list of wrapped errors.
func (e *multiError) Errors() []error {
	return e.errors
}

// SingleWrapMode defines how MultiErrorBuilder behaves when there is only one error in the list.
type SingleWrapMode int

// These constants cause MultiErrorBuilder to behave as described if there is only one error in the list.
const (
	AlwaysWrap   SingleWrapMode = iota // Always return a multiError.
	ReturnSingle                       // Return the single error.
)

// MultiErrorBuilder provides an interface for aggregating errors and exposing them as a single value.
// Deprecated: use multi error tools from emperror.dev/errors instead.
type MultiErrorBuilder struct {
	errors []error

	Message        string
	SingleWrapMode SingleWrapMode
}

// NewMultiErrorBuilder returns a new MultiErrorBuilder.
func NewMultiErrorBuilder() *MultiErrorBuilder {
	return &MultiErrorBuilder{
		SingleWrapMode: AlwaysWrap,
	}
}

// Add adds an error to the list.
//
// Calling this method concurrently is not safe.
func (b *MultiErrorBuilder) Add(err error) {
	b.errors = append(b.errors, err)
}

// ErrOrNil returns a multiError the builder aggregates a list of errors,
// or returns nil if the list of errors is empty.
//
// It is useful to avoid checking if there are any errors added to the list.
func (b *MultiErrorBuilder) ErrOrNil() error {
	err := errors.Combine(b.errors...)
	if err == nil {
		return nil
	}

	errs := errors.GetErrors(err)

	// Return a single error when there is only one and the builder is told to do so.
	if len(errs) == 1 && b.SingleWrapMode == ReturnSingle {
		return errs[0]
	}

	return &multiError{errs, b.Message}
}
