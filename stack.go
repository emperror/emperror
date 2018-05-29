package emperror

import "github.com/pkg/errors"

type stackTracer interface {
	StackTrace() errors.StackTrace
}

// StackTrace returns the stack trace from an error (if any).
func StackTrace(err error) (errors.StackTrace, bool) {
	var stack errors.StackTrace

	ForEachCause(err, func(err error) bool {
		if serr, ok := err.(stackTracer); ok {
			stack = serr.StackTrace()

			return false
		}

		return true
	})

	return stack, stack != nil
}

type withStack struct {
	err   error
	stack errors.StackTrace
}

func (e *withStack) Error() string {
	return e.err.Error()
}

func (e *withStack) Cause() error {
	return e.err
}

func (e *withStack) StackTrace() errors.StackTrace {
	return e.stack
}

// ExposeStackTrace exposes the stack trace (if any) in the outer error.
func ExposeStackTrace(err error) error {
	if err == nil {
		return err
	}

	stack, ok := StackTrace(err)
	if !ok {
		return err
	}

	return &withStack{
		err:   err,
		stack: stack,
	}
}
