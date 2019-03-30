package emperror

import (
	stderrors "errors"
	"testing"

	"github.com/pkg/errors"
)

func TestStackTrace(t *testing.T) {
	err := errors.WithMessage(errors.New("error"), "wrapper")

	stack, ok := StackTrace(err)
	if !ok {
		t.Fatal("stack trace is missing")
	}

	if stack == nil {
		t.Error("empty stack trace")
	}
}

func TestExposeStackTrace(t *testing.T) {
	err := errors.WithMessage(errors.New("error"), "wrapper")

	err = ExposeStackTrace(err)

	stack := err.(stackTracer).StackTrace()

	if len(stack) < 1 {
		t.Error("empty stack trace")
	}
}

func TestExposeStackTrace_NoStackTrace(t *testing.T) {
	err := stderrors.New("error")

	serr := ExposeStackTrace(err)

	if got, want := serr, err; got != want {
		t.Errorf("error does not match the expected one\nactual:   %s\nexpected: %s", got, want)
	}
}
