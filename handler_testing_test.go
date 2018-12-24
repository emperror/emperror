package emperror_test

import (
	"errors"
	"testing"

	. "github.com/goph/emperror"
	"github.com/stretchr/testify/assert"
)

func TestTestHandler_Count(t *testing.T) {
	handler := NewTestHandler()

	err1 := errors.New("error 1")
	err2 := errors.New("error 2")

	handler.Handle(err1)
	handler.Handle(err2)

	assert.Equal(t, 2, handler.Count())
}

func TestTestHandler_LastError(t *testing.T) {
	handler := NewTestHandler()

	err1 := errors.New("error 1")
	err2 := errors.New("error 2")

	handler.Handle(err1)
	handler.Handle(err2)

	assert.Equal(t, err2, handler.LastError())
}

func TestTestHandler_Errors(t *testing.T) {
	handler := NewTestHandler()

	err1 := errors.New("error 1")
	err2 := errors.New("error 2")

	handler.Handle(err1)
	handler.Handle(err2)

	errs := handler.Errors()

	assert.Equal(t, err1, errs[0])
	assert.Equal(t, err2, errs[1])
}
