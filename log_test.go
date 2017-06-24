package emperror_test

import (
	"testing"

	"fmt"

	"github.com/goph/emperror"
)

type testLogger struct {
	last string
}

func (l *testLogger) Log(keyvals ...interface{}) error {
	l.last = keyvals[3].(string)

	return nil
}

func TestLogHandler_Handle(t *testing.T) {
	logger := &testLogger{}
	handler := emperror.NewLogHandler(logger)

	err := fmt.Errorf("internal error")

	handler.Handle(err)

	if got, want := logger.last, err.Error(); got != want {
		t.Fatalf("expected to log a specific error, received: %v", got)
	}
}
