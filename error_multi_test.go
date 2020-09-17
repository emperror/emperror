package emperror

import (
	"fmt"
	"reflect"
	"testing"
)

// guarantee multiError implements Errors and error.
var (
	_ Errors = new(multiError)
	_ error  = new(multiError)
)

func TestMultiErrorBuilder_ErrOrNil(t *testing.T) {
	builder := NewMultiErrorBuilder()

	err := fmt.Errorf("error")

	builder.Add(err)

	merr := builder.ErrOrNil()

	if got, want := merr.(Errors).Errors()[0], err; got != want {
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

func TestMultiErrorBuilder_MultipleErrors(t *testing.T) {
	want := []error{
		fmt.Errorf("first"),
		fmt.Errorf("second"),
	}

	builder := NewMultiErrorBuilder()

	for _, e := range want {
		builder.Add(e)
	}

	if got := builder.ErrOrNil().(Errors).Errors(); !reflect.DeepEqual(got, want) {
		t.Errorf("error does not match the expected one\nactual:   %s\nexpected: %s", got, want)
	}
}
