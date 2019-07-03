package emperror

import (
	"io"
	"testing"

	"github.com/pkg/errors"
)

func TestWrapf_Format(t *testing.T) {
	testError := errors.New("error")

	tests := []struct {
		error
		format string
		want   string
	}{{
		Wrapf(io.EOF, "error%d", 2),
		"%s",
		"error2: EOF",
	}, {
		Wrapf(io.EOF, "error%d", 2),
		"%v",
		"error2: EOF",
	}, {
		Wrapf(io.EOF, "error%d", 2),
		"%+v",
		"EOF\n" +
			"error2\n" +
			"emperror.dev/emperror.TestWrapf_Format\n" +
			"\t.+/wrapf_test.go:26",
	}, {
		Wrapf(errors.New("error"), "error%d", 2),
		"%s",
		"error2: error",
	}, {
		Wrapf(errors.New("error"), "error%d", 2),
		"%v",
		"error2: error",
	}, {
		Wrapf(errors.New("error"), "error%d", 2),
		"%+v",
		"error\n" +
			"emperror.dev/emperror.TestWrapf_Format\n" +
			"\t.+/wrapf_test.go:41",
	}, {
		Wrapf(testError, "error%d", 2),
		"%+v",
		"error\n" +
			"emperror.dev/emperror.TestWrapf_Format\n" +
			"\t.+/wrapf_test.go:11",
	}}

	for i, tt := range tests {
		testFormatRegexp(t, i, tt.error, tt.format, tt.want)
	}
}
