package emperror_test

import (
	"testing"

	"errors"

	"github.com/goph/emperror"
	"github.com/goph/emperror/internal/mocks"
)

func TestCompositeHandler_Handle(t *testing.T) {
	handler1 := &mocks.Handler{}
	handler2 := &mocks.Handler{}

	handler := emperror.NewCompositeHandler(handler1, handler2)

	err := errors.New("error")

	handler1.On("Handle", err).Once().Return()
	handler2.On("Handle", err).Once().Return()

	handler.Handle(err)

	handler1.AssertExpectations(t)
	handler2.AssertExpectations(t)
}
