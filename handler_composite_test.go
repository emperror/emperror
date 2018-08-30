package emperror_test

import (
	"testing"

	"errors"

	"github.com/goph/emperror"
)

func TestCompositeHandler(t *testing.T) {
	handler1 := new(HandlerMock)
	handler2 := new(HandlerMock)

	handler := emperror.NewCompositeHandler(handler1, handler2)

	err := errors.New("error")

	handler1.On("Handle", err).Once().Return()
	handler2.On("Handle", err).Once().Return()

	handler.Handle(err)

	handler1.AssertExpectations(t)
	handler2.AssertExpectations(t)
}
