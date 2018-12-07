package errorlog_test

import (
	"errors"
	"testing"

	"github.com/InVisionApp/go-logger/shims/testlog"
	"github.com/goph/emperror"
	. "github.com/goph/emperror/errorlog"
	"github.com/goph/emperror/internal"
	"github.com/stretchr/testify/assert"
)

func TestHandler_Handle(t *testing.T) {
	logger := testlog.New()
	handler := NewHandler(NewInvisionLogger(logger))

	err := errors.New("internal error")

	handler.Handle(err)

	assert.Equal(t, 1, logger.CallCount())
	assert.Equal(t, "[ERROR] internal error \n", string(logger.Bytes()))
}

func TestHandler_Handle_Context(t *testing.T) {
	logger := testlog.New()
	handler := NewHandler(NewInvisionLogger(logger))

	err := internal.ErrorWithContext{
		Msg: "internal error",
		Keyvals: []interface{}{
			"a", 123,
			"previous", errors.New("previous error"),
		},
	}

	handler.Handle(err)

	assert.Equal(t, 1, logger.CallCount())

	line := string(logger.Bytes())
	assert.Contains(t, line, "[ERROR] internal error")
	assert.Contains(t, line, "a=123")
	assert.Contains(t, line, "previous=previous error")
}

func TestHandler_Handle_MultiError(t *testing.T) {
	logger := testlog.New()
	handler := NewHandler(NewInvisionLogger(logger))

	err := emperror.NewMultiErrorBuilder()
	err.Add(errors.New("internal error"))
	err.Add(errors.New("something else"))

	handler.Handle(err.ErrOrNil())

	line := string(logger.Bytes())
	assert.Contains(t, line, "[ERROR] internal error")
	assert.Contains(t, line, "[ERROR] something else")
	assert.Contains(t, line, "parent=Multiple errors happened")
}
