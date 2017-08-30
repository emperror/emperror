package emperror_test

import (
	"testing"

	"bytes"
	"errors"

	"github.com/go-kit/kit/log"
	"github.com/goph/emperror"
	"github.com/goph/emperror/internal"
	"github.com/stretchr/testify/assert"
)

func TestLogHandler_Handle(t *testing.T) {
	buf := &bytes.Buffer{}
	logger := log.NewLogfmtLogger(buf)
	handler := emperror.NewLogHandler(logger)

	err := errors.New("internal error")

	handler.Handle(err)

	assert.Equal(t, "level=error msg=\"internal error\"\n", buf.String())
}

func TestLogHandler_HandleContext(t *testing.T) {
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

	assert.Equal(t, "level=error msg=\"internal error\" a=123 previous=\"previous error\"\n", buf.String())
}
