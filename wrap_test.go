package emperror_test

import (
	"fmt"
	"io"
	"testing"

	. "github.com/goph/emperror"
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
			"github.com/goph/emperror_test.TestWrap_Format\n" +
			"\t.+/github.com/goph/emperror/wrap_test.go:28",
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
			"github.com/goph/emperror_test.TestWrap_Format\n" +
			"\t.+/github.com/goph/emperror/wrap_test.go:42",
	}, {
		Wrap(Wrap(io.EOF, "error1"), "error2"),
		"%+v",
		"EOF\n" +
			"error1\n" +
			"github.com/goph/emperror_test.TestWrap_Format\n" +
			"\t.+/github.com/goph/emperror/wrap_test.go:49\n",
	}, {
		Wrap(fmt.Errorf("error with space"), "context"),
		"%q",
		`"context: error with space"`,
	}, {
		Wrap(testError, "error2"),
		"%+v",
		"EOF\n" +
			"error1\n" +
			"github.com/goph/emperror_test.TestWrap_Format\n" +
			"\t.+/github.com/goph/emperror/wrap_test.go:13\n",
	}}

	for i, tt := range tests {
		testFormatRegexp(t, i, tt.error, tt.format, tt.want)
	}
}
