package emperror_test

import (
	"testing"

	"errors"

	"github.com/goph/emperror"
	"github.com/stretchr/testify/assert"
)

func TestHandlerContext(t *testing.T) {
	testHandler := emperror.NewTestHandler()

	kvs := []interface{}{"a", 123}
	handler := emperror.HandlerWith(testHandler, kvs...)

	handler.Handle(errors.New("error"))

	cerr := emperror.With(errors.New("error"), "a", 123)

	assert.Equal(t, cerr, testHandler.Last())
}

func TestHandlerContext_Multi(t *testing.T) {
	testHandler := emperror.NewTestHandler()

	handler := emperror.HandlerWith(emperror.HandlerWith(testHandler, "a", 123), "b", 321)

	handler.Handle(errors.New("error"))

	cerr := emperror.With(errors.New("error"), "a", 123, "b", 321)

	assert.Equal(t, cerr, testHandler.Last())
}

func TestHandlerContext_MultiPrefix(t *testing.T) {
	testHandler := emperror.NewTestHandler()

	handler := emperror.HandlerWithPrefix(emperror.HandlerWith(testHandler, "a", 123), "b", 321)

	handler.Handle(errors.New("error"))

	cerr := emperror.With(errors.New("error"), "b", 321, "a", 123)

	assert.Equal(t, cerr, testHandler.Last())
}

func TestHandlerContext_MissingValue(t *testing.T) {
	testHandler := emperror.NewTestHandler()

	handler := emperror.HandlerWithPrefix(emperror.HandlerWith(testHandler, "k0"), "k1")

	handler.Handle(errors.New("error"))

	cerr := emperror.With(errors.New("error"), "k1", nil, "k0", nil)

	assert.Equal(t, cerr, testHandler.Last())
}
