package emperror_test

import (
	"testing"

	"errors"

	"github.com/goph/emperror"
)

func TestCompositeHandler_Handle(t *testing.T) {
	t.Parallel()

	handler1 := emperror.NewTestHandler()
	handler2 := emperror.NewTestHandler()

	handler := emperror.NewCompositeHandler(handler1, handler2)

	err := errors.New("error")

	handler.Handle(err)

	if want, have := err, handler1.Last(); want != have {
		t.Errorf("\nwant: %v\nhave: %v", want, have)
	}

	if want, have := err, handler2.Last(); want != have {
		t.Errorf("\nwant: %v\nhave: %v", want, have)
	}
}
