package emperror_test

import (
	"testing"

	"errors"

	"github.com/goph/emperror"
)

func TestCompositeHandler_Handle(t *testing.T) {
	handler1 := emperror.NewTestHandler()
	handler2 := emperror.NewTestHandler()

	handler := emperror.NewCompositeHandler(handler1, handler2)

	err := errors.New("error")

	handler.Handle(err)

	if handler1.Last() != err || handler2.Last() != err {
		t.Error("error is not handled by all handlers")
	}
}
