package emperror

import (
	"context"
	"reflect"
	"testing"

	"emperror.dev/errors"
)

func TestWithDetails(t *testing.T) {
	testHandler := &TestErrorHandlerFacade{}

	details := []interface{}{"a", 123}
	handler := WithDetails(testHandler, details...)

	handler.Handle(errors.NewPlain("error"))

	cerr := errors.WithDetails(errors.NewPlain("error"), "a", 123)

	testHandleDetails(t, handler, testHandler, cerr)
}

func TestWithDetails_Multi(t *testing.T) {
	testHandler := &TestErrorHandlerFacade{}

	handler := WithDetails(WithDetails(testHandler, "a", 123), "b", 321)

	handler.Handle(errors.NewPlain("error"))

	cerr := errors.WithDetails(errors.NewPlain("error"), "a", 123, "b", 321)

	testHandleDetails(t, handler, testHandler, cerr)
}

func TestWithDetails_MultiPrefix(t *testing.T) {
	testHandler := &TestErrorHandlerFacade{}

	handler := HandlerWithPrefix(WithDetails(testHandler, "a", 123), "b", 321)

	handler.Handle(errors.NewPlain("error"))

	cerr := errors.WithDetails(errors.NewPlain("error"), "b", 321, "a", 123)

	testHandleDetails(t, handler, testHandler, cerr)
}

func TestWithDetails_MissingValue(t *testing.T) {
	testHandler := &TestErrorHandlerFacade{}

	handler := WithDetails(testHandler, "k0")

	handler.Handle(errors.NewPlain("error"))

	cerr := errors.WithDetails(errors.NewPlain("error"), "k0", nil)

	testHandleDetails(t, handler, testHandler, cerr)
}

func testHandleDetails(
	t *testing.T,
	handler ErrorHandlerFacade,
	testHandler *TestErrorHandlerFacade,
	expectedError error,
) {
	handler.Handle(errors.NewPlain("error"))

	if want, have := expectedError, testHandler.LastError(); !reflect.DeepEqual(want, have) {
		t.Errorf("unexpected error\nexpected: %v\nactual:   %v", want, have)
	}

	ctx := context.Background()

	handler.HandleContext(ctx, errors.NewPlain("error"))

	if want, have := expectedError, testHandler.LastError(); !reflect.DeepEqual(want, have) {
		t.Errorf("unexpected error\nexpected: %v\nactual:   %v", want, have)
	}

	if want, have := testHandler.LastContext(), ctx; want != have {
		t.Errorf("unexpected context\nexpected: %v\nactual:   %v", want, have)
	}
}
