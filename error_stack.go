package emperror

import (
	"fmt"
	"io"

	"emperror.dev/errors"
)

type stackTracer interface {
	StackTrace() errors.StackTrace
}

// ExposeStackTrace exposes the stack trace (if any) in the outer error.
func ExposeStackTrace(err error) error {
	if err == nil {
		return err
	}

	var st stackTracer
	if errors.As(err, &st) {
		return &withExposedStack{
			err: err,
			st:  st,
		}
	}

	return err
}

type withExposedStack struct {
	err error
	st  stackTracer
}

func (w *withExposedStack) Error() string {
	return w.err.Error()
}

func (w *withExposedStack) Cause() error  { return w.err }
func (w *withExposedStack) Unwrap() error { return w.err }

func (w *withExposedStack) StackTrace() errors.StackTrace {
	return w.st.StackTrace()
}

// Format implements the fmt.Formatter interface.
func (w *withExposedStack) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			_, _ = fmt.Fprintf(s, "%+v", w.Cause())
			return
		}
		fallthrough

	case 's':
		_, _ = io.WriteString(s, w.Error())

	case 'q':
		_, _ = fmt.Fprintf(s, "%q", w.Error())
	}
}
