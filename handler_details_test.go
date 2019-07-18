package emperror

import (
	"reflect"
	"testing"

	"emperror.dev/errors"
)

func TestWithDetails(t *testing.T) {
	testHandler := NewTestHandler()

	details := []interface{}{"a", 123}
	handler := WithDetails(testHandler, details...)

	handler.Handle(errors.NewPlain("error"))

	cerr := errors.WithDetails(errors.NewPlain("error"), "a", 123)

	if got, want := testHandler.LastError(), cerr; !reflect.DeepEqual(got, want) {
		t.Errorf("error does not match the expected one\nactual:   %s\nexpected: %s", got, want)
	}
}

func TestWithDetails_Multi(t *testing.T) {
	testHandler := NewTestHandler()

	handler := WithDetails(WithDetails(testHandler, "a", 123), "b", 321)

	handler.Handle(errors.NewPlain("error"))

	cerr := errors.WithDetails(errors.NewPlain("error"), "a", 123, "b", 321)

	if got, want := testHandler.LastError(), cerr; !reflect.DeepEqual(got, want) {
		t.Errorf("error does not match the expected one\nactual:   %s\nexpected: %s", got, want)
	}
}

func TestWithDetails_MultiPrefix(t *testing.T) {
	testHandler := NewTestHandler()

	handler := HandlerWithPrefix(WithDetails(testHandler, "a", 123), "b", 321)

	handler.Handle(errors.NewPlain("error"))

	cerr := errors.WithDetails(errors.NewPlain("error"), "b", 321, "a", 123)

	if got, want := testHandler.LastError(), cerr; !reflect.DeepEqual(got, want) {
		t.Errorf("error does not match the expected one\nactual:   %s\nexpected: %s", got, want)
	}
}

func TestWithDetails_MissingValue(t *testing.T) {
	testHandler := NewTestHandler()

	handler := WithDetails(testHandler, "k0")

	handler.Handle(errors.NewPlain("error"))

	cerr := errors.WithDetails(errors.NewPlain("error"), "k0", nil)

	if got, want := testHandler.LastError(), cerr; !reflect.DeepEqual(got, want) {
		t.Errorf("error does not match the expected one\nactual:   %#v\nexpected: %#v", got, want)
	}
}
