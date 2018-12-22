package emperror_test

import (
	"errors"
	"testing"

	. "github.com/goph/emperror"
	"github.com/stretchr/testify/assert"
)

func TestHandleRecovery(t *testing.T) {
	handler := NewTestHandler()
	err := errors.New("error")

	defer func() {
		assert.EqualError(t, handler.Last(), "error")
	}()
	defer HandleRecover(handler)

	panic(err)
}

func TestPanic(t *testing.T) {
	assert.Panics(t, func() {
		Panic(errors.New("error"))
	})
}

func TestPanic_NoError(t *testing.T) {
	assert.NotPanics(t, func() {
		Panic(nil)
	})
}

func TestHandleIfErr(t *testing.T) {
	handler := NewTestHandler()
	err := errors.New("error")

	HandleIfErr(handler, err)

	assert.Equal(t, err, handler.Last())
}

func TestHandleIfErr_Nil(t *testing.T) {
	handler := NewTestHandler()

	HandleIfErr(handler, nil)

	assert.NoError(t, handler.Last())
}
