package emperror

import (
	"errors"
	"testing"
)

func TestTestHandler_Count(t *testing.T) {
	handler := NewTestHandler()

	err1 := errors.New("error 1")
	err2 := errors.New("error 2")

	handler.Handle(err1)
	handler.Handle(err2)

	if got, want := handler.Count(), 2; got != want {
		t.Errorf("error count not match the expected one\nactual:   %d\nexpected: %d", got, want)
	}
}

func TestTestHandler_LastError(t *testing.T) {
	handler := NewTestHandler()

	err1 := errors.New("error 1")
	err2 := errors.New("error 2")

	handler.Handle(err1)
	handler.Handle(err2)

	if got, want := handler.LastError(), err2; got != want {
		t.Errorf("error does not match the expected one\nactual:   %s\nexpected: %s", got, want)
	}
}

func TestTestHandler_Errors(t *testing.T) {
	handler := NewTestHandler()

	err1 := errors.New("error 1")
	err2 := errors.New("error 2")

	handler.Handle(err1)
	handler.Handle(err2)

	errs := handler.Errors()

	if got, want := errs[0], err1; got != want {
		t.Errorf("error 1 does not match the expected one\nactual:   %s\nexpected: %s", got, want)
	}

	if got, want := errs[1], err2; got != want {
		t.Errorf("error 2 does not match the expected one\nactual:   %s\nexpected: %s", got, want)
	}
}
