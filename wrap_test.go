package emperror_test

import (
	"fmt"
	"io"
	"regexp"
	"strings"
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
			"\t.+/github.com/goph/emperror/wrap_test.go:30",
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
			"\t.+/github.com/goph/emperror/wrap_test.go:44",
	}, {
		Wrap(Wrap(io.EOF, "error1"), "error2"),
		"%+v",
		"EOF\n" +
			"error1\n" +
			"github.com/goph/emperror_test.TestWrap_Format\n" +
			"\t.+/github.com/goph/emperror/wrap_test.go:51\n",
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
			"\t.+/github.com/goph/emperror/wrap_test.go:15\n",
	}}

	for i, tt := range tests {
		testFormatRegexp(t, i, tt.error, tt.format, tt.want)
	}
}

func testFormatRegexp(t *testing.T, n int, arg interface{}, format, want string) {
	got := fmt.Sprintf(format, arg)
	gotLines := strings.SplitN(got, "\n", -1)
	wantLines := strings.SplitN(want, "\n", -1)

	if len(wantLines) > len(gotLines) {
		t.Errorf("test %d: wantLines(%d) > gotLines(%d):\n got: %q\nwant: %q", n+1, len(wantLines), len(gotLines), got, want)
		return
	}

	for i, w := range wantLines {
		match, err := regexp.MatchString(w, gotLines[i])
		if err != nil {
			t.Fatal(err)
		}
		if !match {
			t.Errorf("test %d: line %d: fmt.Sprintf(%q, err):\n got: %q\nwant: %q", n+1, i+1, format, got, want)
		}
	}
}
