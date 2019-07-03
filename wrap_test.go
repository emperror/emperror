package emperror

import (
	"fmt"
	"io"
	"testing"

	"github.com/pkg/errors"
)

func TestWrap_Format(t *testing.T) {
	testError := Wrap(io.EOF, "error1")

	tests := []struct {
		error
		format string
		want   string
	}{{
		Wrap(errors.New("error"), "error2"),
		"%s",
		"error2: error",
	}, {
		Wrap(errors.New("error"), "error2"),
		"%v",
		"error2: error",
	}, {
		Wrap(errors.New("error"), "error2"),
		"%+v",
		"error\n" +
			"emperror.dev/emperror.TestWrap_Format\n" +
			"\t.+/wrap_test.go:27",
	}, {
		Wrap(io.EOF, "error"),
		"%s",
		"error: EOF",
	}, {
		Wrap(io.EOF, "error"),
		"%v",
		"error: EOF",
	}, {
		Wrap(io.EOF, "error"),
		"%+v",
		"EOF\n" +
			"error\n" +
			"emperror.dev/emperror.TestWrap_Format\n" +
			"\t.+/wrap_test.go:41",
	}, {
		Wrap(Wrap(io.EOF, "error1"), "error2"),
		"%+v",
		"EOF\n" +
			"error1\n" +
			"emperror.dev/emperror.TestWrap_Format\n" +
			"\t.+/wrap_test.go:48\n",
	}, {
		Wrap(fmt.Errorf("error with space"), "context"),
		"%q",
		`"context: error with space"`,
	}, {
		Wrap(testError, "error2"),
		"%+v",
		"EOF\n" +
			"error1\n" +
			"emperror.dev/emperror.TestWrap_Format\n" +
			"\t.+/wrap_test.go:12\n",
	}}

	for i, tt := range tests {
		testFormatRegexp(t, i, tt.error, tt.format, tt.want)
	}
}
