package errorlog_test

import (
	"errors"
	"testing"

	"github.com/goph/emperror"
	. "github.com/goph/emperror/errorlog"
	"github.com/goph/emperror/internal"
	"github.com/goph/logur"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandler_Handle(t *testing.T) {
	logger := logur.NewTestLogger()
	handler := NewHandler(NewLogger(logger))

	err := errors.New("internal error")

	handler.Handle(err)

	require.Equal(t, 1, logger.Count())

	lastEvent := logger.LastEvent()

	assert.Equal(t, logur.ErrorLevel, lastEvent.Level)
	assert.Equal(t, "internal error", lastEvent.Line)
}

func TestHandler_Handle_Context(t *testing.T) {
	logger := logur.NewTestLogger()
	handler := NewHandler(NewLogger(logger))

	err := internal.ErrorWithContext{
		Msg: "internal error",
		Keyvals: []interface{}{
			"a", 123,
			"previous", errors.New("previous error"),
		},
	}

	handler.Handle(err)

	require.Equal(t, 1, logger.Count())

	lastEvent := logger.LastEvent()

	assert.Equal(t, logur.ErrorLevel, lastEvent.Level)
	assert.Equal(t, "internal error", lastEvent.Line)
	assert.Equal(t, logur.Fields{"a": 123, "previous": "previous error"}, lastEvent.Fields)
}

func TestHandler_Handle_MultiError(t *testing.T) {
	logger := logur.NewTestLogger()
	handler := NewHandler(NewLogger(logger))

	err := emperror.NewMultiErrorBuilder()
	err.Add(errors.New("internal error"))
	err.Add(errors.New("something else"))

	handler.Handle(err.ErrOrNil())

	require.Equal(t, 2, logger.Count())

	events := logger.Events()

	assert.Equal(t, logur.ErrorLevel, events[0].Level)
	assert.Equal(t, "internal error", events[0].Line)
	assert.Equal(t, logur.Fields{"parent": "Multiple errors happened"}, events[0].Fields)

	assert.Equal(t, logur.ErrorLevel, events[1].Level)
	assert.Equal(t, "something else", events[1].Line)
	assert.Equal(t, logur.Fields{"parent": "Multiple errors happened"}, events[1].Fields)
}
