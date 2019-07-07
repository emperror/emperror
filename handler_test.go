package emperror

import (
	"errors"
	"testing"
)

func TestHandlers(t *testing.T) {
	handler1 := NewTestHandler()
	handler2 := NewTestHandler()

	handler := Handlers{handler1, handler2}

	err := errors.New("error")

	handler.Handle(err)

	if got, want := handler1.LastError(), err; got != want {
		t.Errorf("error does not match the expected one\nactual:   %s\nexpected: %s", got, want)
	}

	if got, want := handler1.LastError(), err; got != want {
		t.Errorf("error does not match the expected one\nactual:   %s\nexpected: %s", got, want)
	}
}

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
