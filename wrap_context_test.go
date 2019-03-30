package emperror

import (
	"fmt"
	"io"
	"testing"

	"github.com/pkg/errors"
)

func TestWrapWith_Format(t *testing.T) {
	testError := WrapWith(io.EOF, "error1", "key", "value")

	tests := []struct {
		error
		format string
		want   string
	}{{
		WrapWith(errors.New("error"), "error2", "key", "value"),
		"%s",
		"error2: error",
	}, {
		WrapWith(errors.New("error"), "error2", "key", "value"),
		"%v",
		"error2: error",
	}, {
		WrapWith(errors.New("error"), "error2", "key", "value"),
		"%+v",
		"error\n" +
			"github.com/goph/emperror.TestWrapWith_Format\n" +
			"\t.+/wrap_context_test.go:27",
	}, {
		WrapWith(io.EOF, "error", "key", "value"),
		"%s",
		"error: EOF",
	}, {
		WrapWith(io.EOF, "error", "key", "value"),
		"%v",
		"error: EOF",
	}, {
		WrapWith(io.EOF, "error", "key", "value"),
		"%+v",
		"EOF\n" +
			"error\n" +
			"github.com/goph/emperror.TestWrapWith_Format\n" +
			"\t.+/wrap_context_test.go:41",
	}, {
		WrapWith(WrapWith(io.EOF, "error1"), "error2", "key", "value"),
		"%+v",
		"EOF\n" +
			"error1\n" +
			"github.com/goph/emperror.TestWrapWith_Format\n" +
			"\t.+/wrap_context_test.go:48\n",
	}, {
		WrapWith(fmt.Errorf("error with space"), "context", "key", "value"),
		"%q",
		`"context: error with space"`,
	}, {
		WrapWith(testError, "error2", "key", "value"),
		"%+v",
		"EOF\n" +
			"error1\n" +
			"github.com/goph/emperror.TestWrapWith_Format\n" +
			"\t.+/wrap_context_test.go:12\n",
	}}

	for i, tt := range tests {
		testFormatRegexp(t, i, tt.error, tt.format, tt.want)
	}
}

func TestWrapWith_Context(t *testing.T) {
	err := errors.New("error")

	kvs := []interface{}{"a", 123}
	err = WrapWith(err, "error2", kvs...)
	kvs[1] = 0 // WrapWith should copy its key values

	ctx := Context(err)

	if got, want := ctx[0], "a"; got != want {
		t.Errorf("context value does not match the expected one\nactual:   %s\nexpected: %s", got, want)
	}

	if got, want := ctx[1], 123; got != want {
		t.Errorf("context value does not match the expected one\nactual:   %s\nexpected: %d", got, want)
	}

	if got, want := err.Error(), "error2: error"; got != want {
		t.Errorf("error does not match the expected one\nactual:   %s\nexpected: %s", got, want)
	}
}
