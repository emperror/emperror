package emperror_test

import (
	"testing"

	. "github.com/goph/emperror"
	"github.com/goph/emperror/internal/mocks"
)

func TestCompositeHandler(t *testing.T) {
	handler1 := new(mocks.Handler)
	handler2 := new(mocks.Handler)

	handler := NewCompositeHandler(handler1, handler2)

	err := New("error")

	handler1.On("Handle", err).Once().Return()
	handler2.On("Handle", err).Once().Return()

	handler.Handle(err)

	handler1.AssertExpectations(t)
	handler2.AssertExpectations(t)
}
