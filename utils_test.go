package emperror_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/goph/emperror"
)

func TestHandleRecovery(t *testing.T) {
	handler := NewTestHandler()
	err := errors.New("error")

	defer func() {
		assert.EqualError(t, handler.LastError(), "error")
	}()
	defer HandleRecover(handler)

	panic(err)
}

func TestHandleIfErr(t *testing.T) {
	handler := NewTestHandler()
	err := errors.New("error")

	Handle(handler, err)

	assert.Equal(t, err, handler.LastError())
}

func TestHandleIfErr_Nil(t *testing.T) {
	handler := NewTestHandler()

	Handle(handler, nil)

	assert.NoError(t, handler.LastError())
}
