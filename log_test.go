package emperror_test

import (
	"testing"

	"bytes"
	"errors"

	"github.com/go-kit/kit/log"
	"github.com/goph/emperror"
	"github.com/goph/emperror/internal"
)

func TestLogHandler_Handle(t *testing.T) {
	t.Parallel()
	buf := &bytes.Buffer{}
	logger := log.NewLogfmtLogger(buf)
	handler := emperror.NewLogHandler(logger)

	err := errors.New("internal error")

	handler.Handle(err)

	if want, have := "level=error msg=\"internal error\"\n", buf.String(); want != have {
		t.Errorf("\nwant: %shave: %s", want, have)
	}
}

func TestLogHandler_HandleContext(t *testing.T) {
	t.Parallel()
	buf := &bytes.Buffer{}
	logger := log.NewLogfmtLogger(buf)
	handler := emperror.NewLogHandler(logger)

	err := internal.ErrorWithContext{
		Msg: "internal error",
		Keyvals: []interface{}{
			"a", 123,
			"previous", errors.New("previous error"),
		},
	}

	handler.Handle(err)

	if want, have := "level=error msg=\"internal error\" a=123 previous=\"previous error\"\n", buf.String(); want != have {
		t.Errorf("\nwant: %shave: %s", want, have)
	}
}
