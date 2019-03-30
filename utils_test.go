package emperror

import (
	"errors"
	"testing"
)

func TestHandleRecovery(t *testing.T) {
	err := errors.New("error")

	defer HandleRecover(HandlerFunc(func(err error) {
		if got, want := err.Error(), "error"; got != want {
			t.Errorf("error does not match the expected one\nactual:   %s\nexpected: %s", got, want)
		}
	}))

	panic(err)
}

func TestHandle(t *testing.T) {
	handler := NewTestHandler()
	err := errors.New("error")

	Handle(handler, err)

	if got, want := handler.LastError(), err; got != want {
		t.Errorf("error does not match the expected one\nactual:   %s\nexpected: %s", got, want)
	}
}

func TestHandle_Nil(t *testing.T) {
	handler := NewTestHandler()

	Handle(handler, nil)

	if got := handler.LastError(); got != nil {
		t.Errorf("unexpected error, received: %s", got)
	}
}
