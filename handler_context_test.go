package emperror

import (
	"reflect"
	"testing"

	"emperror.dev/errors"
)

func TestHandlerContext(t *testing.T) {
	testHandler := NewTestHandler()

	kvs := []interface{}{"a", 123}
	handler := HandlerWith(testHandler, kvs...)

	handler.Handle(errors.NewPlain("error"))

	cerr := errors.WithDetails(errors.NewPlain("error"), "a", 123)

	if got, want := testHandler.LastError(), cerr; !reflect.DeepEqual(got, want) {
		t.Errorf("error does not match the expected one\nactual:   %s\nexpected: %s", got, want)
	}
}

func TestHandlerContext_Multi(t *testing.T) {
	testHandler := NewTestHandler()

	handler := HandlerWith(HandlerWith(testHandler, "a", 123), "b", 321)

	handler.Handle(errors.NewPlain("error"))

	cerr := errors.WithDetails(errors.NewPlain("error"), "a", 123, "b", 321)

	if got, want := testHandler.LastError(), cerr; !reflect.DeepEqual(got, want) {
		t.Errorf("error does not match the expected one\nactual:   %s\nexpected: %s", got, want)
	}
}

func TestHandlerContext_MultiPrefix(t *testing.T) {
	testHandler := NewTestHandler()

	handler := HandlerWithPrefix(HandlerWith(testHandler, "a", 123), "b", 321)

	handler.Handle(errors.NewPlain("error"))

	cerr := errors.WithDetails(errors.NewPlain("error"), "b", 321, "a", 123)

	if got, want := testHandler.LastError(), cerr; !reflect.DeepEqual(got, want) {
		t.Errorf("error does not match the expected one\nactual:   %s\nexpected: %s", got, want)
	}
}

func TestHandlerContext_MissingValue(t *testing.T) {
	testHandler := NewTestHandler()

	handler := HandlerWithPrefix(HandlerWith(testHandler, "k0"), "k1")

	handler.Handle(errors.NewPlain("error"))

	cerr := errors.WithDetails(errors.NewPlain("error"), "k1", nil, "k0", nil)

	if got, want := testHandler.LastError(), cerr; !reflect.DeepEqual(got, want) {
		t.Errorf("error does not match the expected one\nactual:   %s\nexpected: %s", got, want)
	}
}
