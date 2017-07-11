package emperror_test

import (
	"testing"

	"errors"

	"github.com/goph/emperror"
)

func TestTestHandler_Handle(t *testing.T) {
	t.Parallel()

	handler := &emperror.TestHandler{}

	err1 := errors.New("error 1")
	err2 := errors.New("error 2")

	handler.Handle(err1)
	handler.Handle(err2)

	errors := handler.Errors()

	if want, have := err1, errors[0]; want != have {
		t.Errorf("\nwant: %v\nhave: %v", want, have)
	}

	if want, have := err2, errors[1]; want != have {
		t.Errorf("\nwant: %v\nhave: %v", want, have)
	}
}

func TestTestHandler_Last(t *testing.T) {
	t.Parallel()

	handler := &emperror.TestHandler{}

	err1 := errors.New("error 1")
	err2 := errors.New("error 2")

	handler.Handle(err1)
	handler.Handle(err2)

	if want, have := err2, handler.Last(); want != have {
		t.Errorf("\nwant: %v\nhave: %v", want, have)
	}
}

func TestTestHandler_Last_Empty(t *testing.T) {
	t.Parallel()

	handler := &emperror.TestHandler{}

	var want, have error
	if want, have = nil, handler.Last(); want != have {
		t.Errorf("\nwant: %v\nhave: %v", want, have)
	}
}
