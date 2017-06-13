package emperror

import (
	"testing"

	"fmt"
)

type testErrorLogger struct {
	last error
}

func (l *testErrorLogger) Error(args ...interface{}) {
	l.last = args[0].(error)
}

func TestLogHandler_Handle(t *testing.T) {
	logger := &testErrorLogger{}
	handler := NewLogHandler(logger)

	err := fmt.Errorf("internal error")

	handler.Handle(err)

	if got, want := logger.last, err; got != want {
		t.Fatalf("expected to log a specific error, received: %v", got)
	}
}
