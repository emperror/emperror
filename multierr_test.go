package emperror

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMultiErrorBuilder_ErrOrNil(t *testing.T) {
	builder := NewMultiErrorBuilder()

	err := fmt.Errorf("error")

	builder.Add(err)

	merr := builder.ErrOrNil()

	if got, want := merr.(*MultiError).Errors()[0], err; got != want {
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

func TestMultiError_MultipleErrors(t *testing.T) {
	want := []error{
		fmt.Errorf("first"),
		fmt.Errorf("second"),
	}

	builder := &MultiErrorBuilder{}

	for _, e := range want {
		builder.Add(e)
	}

	errsType, ok := builder.ErrOrNil().(*MultiError)
	if !ok {
		t.Errorf("error is not of type MultiError")
	}

	if got := errsType.Errors(); !reflect.DeepEqual(got, want) {
		t.Errorf("error does not match the expected one\nactual:   %s\nexpected: %s", got, want)
	}
}
