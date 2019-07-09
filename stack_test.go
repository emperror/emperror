// Copyright 2015 Dave Cheney <dave@cheney.net>. All rights reserved.
// Use of this source code (or at least parts of it) is governed by a BSD-style
// license that can be found in the LICENSE_THIRD_PARTY file.

package emperror

import (
	"runtime"
	"testing"
)

func TestStackTrace(t *testing.T) {
	tests := []struct {
		stack *stack
		want  []string
	}{
		{
			callers(0),
			[]string{
				"emperror.dev/emperror.TestStackTrace\n" +
					"\t.+/stack_test.go:18",
			},
		},
	}
	for i, tt := range tests {
		st := tt.stack.StackTrace()
		for j, want := range tt.want {
			testFormatRegexp(t, i, st[j], "%+v", want)
		}
	}
}

func stackTrace() StackTrace {
	const depth = 8
	var pcs [depth]uintptr
	n := runtime.Callers(1, pcs[:])
	var st stack = pcs[0:n]
	return st.StackTrace()
}

func TestStackTraceFormat(t *testing.T) {
	tests := []struct {
		StackTrace
		format string
		want   string
	}{{
		nil,
		"%s",
		`\[\]`,
	}, {
		nil,
		"%v",
		`\[\]`,
	}, {
		nil,
		"%+v",
		"",
	}, {
		nil,
		"%#v",
		`\[\]errors.Frame\(nil\)`,
	}, {
		make(StackTrace, 0),
		"%s",
		`\[\]`,
	}, {
		make(StackTrace, 0),
		"%v",
		`\[\]`,
	}, {
		make(StackTrace, 0),
		"%+v",
		"",
	}, {
		make(StackTrace, 0),
		"%#v",
		`\[\]errors.Frame{}`,
	}, {
		stackTrace()[:2],
		"%s",
		`\[stack_test.go stack_test.go\]`,
	}, {
		stackTrace()[:2],
		"%v",
		`\[stack_test.go:36 stack_test.go:83\]`,
	}, {
		stackTrace()[:2],
		"%+v",
		"\n" +
			"emperror.dev/emperror.stackTrace\n" +
			"\t.+/stack_test.go:36\n" +
			"emperror.dev/emperror.TestStackTraceFormat\n" +
			"\t.+/stack_test.go:87",
	}, {
		stackTrace()[:2],
		"%#v",
		`\[\]errors.Frame{stack_test.go:36, stack_test.go:95}`,
	}}

	for i, tt := range tests {
		testFormatRegexp(t, i, tt.StackTrace, tt.format, tt.want)
	}
}
