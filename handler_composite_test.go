package emperror

import (
	"errors"
	"testing"
)

func TestCompositeHandler(t *testing.T) {
	handler1 := NewTestHandler()
	handler2 := NewTestHandler()

	handler := NewCompositeHandler(handler1, handler2)

	err := errors.New("error")

	handler.Handle(err)

	if got, want := handler1.LastError(), err; got != want {
		t.Errorf("error does not match the expected one\nactual:   %s\nexpected: %s", got, want)
	}

	if got, want := handler1.LastError(), err; got != want {
		t.Errorf("error does not match the expected one\nactual:   %s\nexpected: %s", got, want)
	}
}
