package emperror_test

import (
	"testing"

	"errors"

	"github.com/goph/emperror"
)

func TestHandlerContext(t *testing.T) {
	mockHandler := new(HandlerMock)
	cerr := emperror.With(errors.New("error"), "a", 123)
	mockHandler.On("Handle", cerr).Return()

	kvs := []interface{}{"a", 123}
	handler := emperror.HandlerWith(mockHandler, kvs...)

	handler.Handle(errors.New("error"))

	mockHandler.AssertExpectations(t)
}

func TestHandlerContext_Multi(t *testing.T) {
	mockHandler := new(HandlerMock)
	cerr := emperror.With(errors.New("error"), "a", 123, "b", 321)
	mockHandler.On("Handle", cerr).Return()

	handler := emperror.HandlerWith(emperror.HandlerWith(mockHandler, "a", 123), "b", 321)

	handler.Handle(errors.New("error"))

	mockHandler.AssertExpectations(t)
}

func TestHandlerContext_MultiPrefix(t *testing.T) {
	mockHandler := new(HandlerMock)
	cerr := emperror.With(errors.New("error"), "b", 321, "a", 123)
	mockHandler.On("Handle", cerr).Return()

	handler := emperror.HandlerWithPrefix(emperror.HandlerWith(mockHandler, "a", 123), "b", 321)

	handler.Handle(errors.New("error"))

	mockHandler.AssertExpectations(t)
}

func TestHandlerContext_MissingValue(t *testing.T) {
	mockHandler := new(HandlerMock)
	cerr := emperror.With(errors.New("error"), "k1", nil, "k0", nil)
	mockHandler.On("Handle", cerr).Return()

	handler := emperror.HandlerWithPrefix(emperror.HandlerWith(mockHandler, "k0"), "k1")

	handler.Handle(errors.New("error"))

	mockHandler.AssertExpectations(t)
}
