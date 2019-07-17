package emperror

import (
	"reflect"
	"testing"

	"emperror.dev/errors"
)

func TestHandlerWithDetails(t *testing.T) {
	testHandler := NewTestHandler()

	details := []interface{}{"a", 123}
	handler := HandlerWithDetails(testHandler, details...)

	handler.Handle(errors.NewPlain("error"))

	cerr := errors.WithDetails(errors.NewPlain("error"), "a", 123)

	if got, want := testHandler.LastError(), cerr; !reflect.DeepEqual(got, want) {
		t.Errorf("error does not match the expected one\nactual:   %s\nexpected: %s", got, want)
	}
}

func TestHandlerWithDetails_Multi(t *testing.T) {
	testHandler := NewTestHandler()

	handler := HandlerWithDetails(HandlerWithDetails(testHandler, "a", 123), "b", 321)

	handler.Handle(errors.NewPlain("error"))

	cerr := errors.WithDetails(errors.NewPlain("error"), "a", 123, "b", 321)

	if got, want := testHandler.LastError(), cerr; !reflect.DeepEqual(got, want) {
		t.Errorf("error does not match the expected one\nactual:   %s\nexpected: %s", got, want)
	}
}

func TestHandlerWithDetails_MultiPrefix(t *testing.T) {
	testHandler := NewTestHandler()

	handler := HandlerWithPrefix(HandlerWithDetails(testHandler, "a", 123), "b", 321)

	handler.Handle(errors.NewPlain("error"))

	cerr := errors.WithDetails(errors.NewPlain("error"), "b", 321, "a", 123)

	if got, want := testHandler.LastError(), cerr; !reflect.DeepEqual(got, want) {
		t.Errorf("error does not match the expected one\nactual:   %s\nexpected: %s", got, want)
	}
}

func TestHandlerWithDetails_MissingValue(t *testing.T) {
	testHandler := NewTestHandler()

	handler := HandlerWithDetails(testHandler, "k0")

	handler.Handle(errors.NewPlain("error"))

	cerr := errors.WithDetails(errors.NewPlain("error"), "k0", nil)

	if got, want := testHandler.LastError(), cerr; !reflect.DeepEqual(got, want) {
		t.Errorf("error does not match the expected one\nactual:   %#v\nexpected: %#v", got, want)
	}
}
