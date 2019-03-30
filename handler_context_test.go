package emperror

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandlerContext(t *testing.T) {
	testHandler := NewTestHandler()

	kvs := []interface{}{"a", 123}
	handler := HandlerWith(testHandler, kvs...)

	handler.Handle(errors.New("error"))

	cerr := With(errors.New("error"), "a", 123)

	assert.Equal(t, cerr, testHandler.LastError())
}

func TestHandlerContext_Multi(t *testing.T) {
	testHandler := NewTestHandler()

	handler := HandlerWith(HandlerWith(testHandler, "a", 123), "b", 321)

	handler.Handle(errors.New("error"))

	cerr := With(errors.New("error"), "a", 123, "b", 321)

	assert.Equal(t, cerr, testHandler.LastError())
}

func TestHandlerContext_MultiPrefix(t *testing.T) {
	testHandler := NewTestHandler()

	handler := HandlerWithPrefix(HandlerWith(testHandler, "a", 123), "b", 321)

	handler.Handle(errors.New("error"))

	cerr := With(errors.New("error"), "b", 321, "a", 123)

	assert.Equal(t, cerr, testHandler.LastError())
}

func TestHandlerContext_MissingValue(t *testing.T) {
	testHandler := NewTestHandler()

	handler := HandlerWithPrefix(HandlerWith(testHandler, "k0"), "k1")

	handler.Handle(errors.New("error"))

	cerr := With(errors.New("error"), "k1", nil, "k0", nil)

	assert.Equal(t, cerr, testHandler.LastError())
}
