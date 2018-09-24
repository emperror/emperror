package emperror_test

import (
	"fmt"
	"io"
	"testing"

	. "github.com/goph/emperror"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
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
			"github.com/goph/emperror_test.TestWrapWith_Format\n" +
			"\t.+/github.com/goph/emperror/wrap_context_test.go:29",
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
			"github.com/goph/emperror_test.TestWrapWith_Format\n" +
			"\t.+/github.com/goph/emperror/wrap_context_test.go:43",
	}, {
		WrapWith(WrapWith(io.EOF, "error1"), "error2", "key", "value"),
		"%+v",
		"EOF\n" +
			"error1\n" +
			"github.com/goph/emperror_test.TestWrapWith_Format\n" +
			"\t.+/github.com/goph/emperror/wrap_context_test.go:50\n",
	}, {
		WrapWith(fmt.Errorf("error with space"), "context", "key", "value"),
		"%q",
		`"context: error with space"`,
	}, {
		WrapWith(testError, "error2", "key", "value"),
		"%+v",
		"EOF\n" +
			"error1\n" +
			"github.com/goph/emperror_test.TestWrapWith_Format\n" +
			"\t.+/github.com/goph/emperror/wrap_context_test.go:14\n",
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

	assert.Equal(t, "a", ctx[0])
	assert.Equal(t, 123, ctx[1])
	assert.EqualError(t, err, "error2: error")
}
