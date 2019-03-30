package emperror

import (
	"fmt"
	"testing"
)

type errorCollection interface {
	Errors() []error
}

func TestMultiErrorBuilder_ErrOrNil(t *testing.T) {
	builder := NewMultiErrorBuilder()

	err := fmt.Errorf("error")

	builder.Add(err)

	merr := builder.ErrOrNil()

	if got, want := merr.(errorCollection).Errors()[0], err; got != want {
		t.Errorf("error does not match the expected one\nactual:   %s\nexpected: %s", got, want)
	}
}

func TestMultiErrorBuilder_ErrOrNil_NilWhenEmpty(t *testing.T) {
	builder := NewMultiErrorBuilder()

	if got := builder.ErrOrNil(); got != nil {
		t.Errorf("unexpected error, received: %s", got)
	}
}

func TestMultiErrorBuilder_ErrOrNil_Single(t *testing.T) {
	builder := &MultiErrorBuilder{
		SingleWrapMode: ReturnSingle,
	}

	err := fmt.Errorf("error")

	builder.Add(err)

	if got, want := builder.ErrOrNil(), err; got != want {
		t.Errorf("error does not match the expected one\nactual:   %s\nexpected: %s", got, want)
	}
}

func TestMultiErrorBuilder_Message(t *testing.T) {
	want := "Multiple errors happened during action"

	builder := &MultiErrorBuilder{
		Message: want,
	}

	err := fmt.Errorf("error")

	builder.Add(err)

	if got, want := builder.ErrOrNil().Error(), want; got != want {
		t.Errorf("error does not match the expected one\nactual:   %s\nexpected: %s", got, want)
	}
}
