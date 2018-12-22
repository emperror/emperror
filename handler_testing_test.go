package emperror_test

import (
	"testing"
	"errors"

	. "github.com/goph/emperror"
	"github.com/stretchr/testify/assert"
)

func TestTestHandler_Handle(t *testing.T) {
	handler := NewTestHandler()

	err1 := errors.New("error 1")
	err2 := errors.New("error 2")

	handler.Handle(err1)
	handler.Handle(err2)

	errors := handler.Errors()

	assert.Equal(t, err1, errors[0])
	assert.Equal(t, err2, errors[1])
}

func TestTestHandler_Last(t *testing.T) {
	handler := NewTestHandler()

	err1 := errors.New("error 1")
	err2 := errors.New("error 2")

	handler.Handle(err1)
	handler.Handle(err2)

	assert.Equal(t, err2, handler.Last())
}

func TestTestHandler_Last_Empty(t *testing.T) {
	handler := NewTestHandler()

	assert.NoError(t, handler.Last())
}
