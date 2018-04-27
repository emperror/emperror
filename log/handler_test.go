package log_test

import (
	"io"
	"testing"

	"bytes"
	"errors"

	"github.com/go-kit/kit/log"
	"github.com/goph/emperror"
	"github.com/goph/emperror/internal"
	. "github.com/goph/emperror/log"
	"github.com/stretchr/testify/assert"
)

func TestHandler_Handle(t *testing.T) {
	buf := &bytes.Buffer{}
	logger := log.NewLogfmtLogger(buf)
	handler := NewHandler(logger)

	err := errors.New("internal error")

	handler.Handle(err)

	assert.Equal(t, "level=error msg=\"internal error\"\n", buf.String())
}

func TestHandler_Handle_Context(t *testing.T) {
	buf := &bytes.Buffer{}
	logger := log.NewLogfmtLogger(buf)
	handler := NewHandler(logger)

	err := internal.ErrorWithContext{
		Msg: "internal error",
		Keyvals: []interface{}{
			"a", 123,
			"previous", errors.New("previous error"),
		},
	}

	handler.Handle(err)

	assert.Equal(t, "level=error a=123 previous=\"previous error\" msg=\"internal error\"\n", buf.String())
}

func TestHandler_Handle_MultiError(t *testing.T) {
	tests := map[string]struct{
		logger func(w io.Writer) log.Logger
		expected string
	}{
		"logfmt": {
			log.NewLogfmtLogger,
			"level=error msg=\"internal error\" parent=\"Multiple errors happened\"\nlevel=error msg=\"something else\" parent=\"Multiple errors happened\"\n",
		},
		"json": {
			log.NewJSONLogger,
			"{\"level\":\"error\",\"msg\":\"internal error\",\"parent\":\"Multiple errors happened\"}\n{\"level\":\"error\",\"msg\":\"something else\",\"parent\":\"Multiple errors happened\"}\n",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			buf := &bytes.Buffer{}
			logger := test.logger(buf)
			handler := NewHandler(logger)

			err := emperror.NewMultiErrorBuilder()
			err.Add(errors.New("internal error"))
			err.Add(errors.New("something else"))

			handler.Handle(err.ErrOrNil())

			assert.Equal(t, test.expected, buf.String())
		})
	}
}
