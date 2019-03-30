package emperror

import (
	"errors"
	"testing"
)

func TestHandlerFunc(t *testing.T) {
	var actual error
	log := func(err error) {
		actual = err
	}

	fn := HandlerFunc(log)

	expected := errors.New("error")

	fn.Handle(expected)

	if got, want := actual, expected; got != want {
		t.Errorf("error does not match the expected one\nactual:   %s\nexpected: %s", got, want)
	}
}
