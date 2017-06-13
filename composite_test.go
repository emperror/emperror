package emperror

import (
	"errors"
	"testing"
)

func TestCompositeHandler_Handle(t *testing.T) {
	handler1 := NewTestHandler()
	handler2 := NewTestHandler()

	handler := NewCompositeHandler(handler1, handler2)

	err := errors.New("error")

	handler.Handle(err)

	if handler1.Last() != err || handler2.Last() != err {
		t.Error("error is not handled by all handlers")
	}
}
