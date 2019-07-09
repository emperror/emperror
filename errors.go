package emperror

import (
	"fmt"
	"io"
)

// WithStack annotates err with a stack trace at the point WithStack was called.
// If err is nil, WithStack returns nil.
func WithStack(err error) error {
	return WithStackDepth(err, 1)
}

// WithStackDepth annotates err with a stack trace at the point WithStackDepth was called.
// If err is nil, WithStackDepth returns nil.
func WithStackDepth(err error, depth int) error {
	if err == nil {
		return nil
	}

	return &withStack{
		err,
		callers(depth + 1),
	}
}

type withStack struct {
	error
	*stack
}

func (w *withStack) Cause() error { return w.error }

func (w *withStack) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "%+v", w.Cause())
			w.stack.Format(s, verb)
			return
		}
		fallthrough
	case 's':
		io.WriteString(s, w.Error())
	case 'q':
		fmt.Fprintf(s, "%q", w.Error())
	}
}
