package emperror

import (
	"regexp"
	"testing"

	"emperror.dev/errors"
)

func TestWithStack(t *testing.T) {
	testHandler := NewTestHandler()
	handler := WithStack(testHandler)

	err := errors.New("error")

	handler.Handle(err)

	lastErr := testHandler.LastError()
	details := []interface{}{"func", "TestWithStack", "file", "handler_stack_test.go:14"}
	for i, value := range errors.GetDetails(lastErr) {
		r := regexp.MustCompile(details[i].(string))

		if !r.MatchString(value.(string)) {
			t.Errorf("stack trace does not match the expected one\nactual:   %s\nexpected: %s", value, details[i])
		}
	}
}

func TestWithStack_NoStack(t *testing.T) {
	testHandler := NewTestHandler()
	handler := WithStack(testHandler)

	err := errors.NewPlain("error")

	handler.Handle(err)

	if details := errors.GetDetails(testHandler.LastError()); len(details) > 0 {
		t.Errorf("unexpected trace information: %+v", details)
	}
}

func TestWithStack_NoError(t *testing.T) {
	testHandler := NewTestHandler()
	handler := WithStack(testHandler)

	handler.Handle(nil)

	if testHandler.Count() > 0 {
		t.Errorf("unexpected error: %+v", testHandler.LastError())
	}
}
