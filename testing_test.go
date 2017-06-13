package emperror

import (
	"errors"
	"testing"
)

func TestTestHandler_Handle(t *testing.T) {
	handler := &TestHandler{}

	err1 := errors.New("error 1")
	err2 := errors.New("error 2")

	handler.Handle(err1)
	handler.Handle(err2)

	errors := handler.Errors()

	if errors[0] != err1 || errors[1] != err2 {
		t.Error("errors do not match with the handled ones")
	}
}

func TestTestHandler_Last(t *testing.T) {
	handler := &TestHandler{}

	err1 := errors.New("error 1")
	err2 := errors.New("error 2")

	handler.Handle(err1)
	handler.Handle(err2)

	if handler.Last() != err2 {
		t.Error("last error does not match with the expected one")
	}
}

func TestTestHandler_Last_Empty(t *testing.T) {
	handler := &TestHandler{}

	if handler.Last() != nil {
		t.Error("empty handler, expected nil")
	}
}
