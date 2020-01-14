package emperror

import (
	"context"
	"reflect"
	"testing"
	"time"

	"emperror.dev/errors"
)

func TestNewErrorHandlerContext(t *testing.T) {
	t.Run("no_details", func(t *testing.T) {
		testHandler := &TestErrorHandlerFacade{}

		handler := NewErrorHandlerContext(testHandler, func(ctx context.Context) map[string]interface{} {
			return nil
		})

		handler.HandleContext(context.Background(), errors.New("error"))

		details := errors.GetDetails(testHandler.LastError())

		if len(details) > 0 {
			t.Error("error is not expected to have any details")
		}
	})

	t.Run("details", func(t *testing.T) {
		testHandler := &TestErrorHandlerFacade{}

		handler := NewErrorHandlerContext(testHandler, func(ctx context.Context) map[string]interface{} {
			return map[string]interface{}{
				"key": "value",
			}
		})

		handler.HandleContext(context.Background(), errors.New("error"))

		details := errors.GetDetails(testHandler.LastError())

		if want, have := []interface{}{"key", "value"}, details; !reflect.DeepEqual(want, have) {
			t.Errorf("unexpexted \nexpected: %v\nactual:   %v", want, have)
		}
	})
}

func TestContextExtractors(t *testing.T) {
	extractor := ContextExtractors(
		func(_ context.Context) map[string]interface{} {
			return nil
		},
		func(_ context.Context) map[string]interface{} {
			return map[string]interface{}{
				"key":  "value",
				"key2": "value2",
			}
		},
		func(_ context.Context) map[string]interface{} {
			return map[string]interface{}{
				"key":  "another_value",
				"key3": "value3",
			}
		},
		func(_ context.Context) map[string]interface{} {
			return map[string]interface{}{
				"key4": time.Minute,
			}
		},
	)

	expected := map[string]interface{}{
		"key":  "another_value",
		"key2": "value2",
		"key3": "value3",
		"key4": time.Minute,
	}

	if want, have := expected, extractor(context.Background()); !reflect.DeepEqual(want, have) {
		t.Errorf("unexpexted details\nexpected: %v\nactual:   %v", want, have)
	}
}

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
