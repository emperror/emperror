package emperror_test

import (
	"errors"
	"testing"

	"github.com/goph/emperror"
	"github.com/stretchr/testify/assert"
)

func TestHandleRecovery(t *testing.T) {
	handler := emperror.NewTestHandler()
	err := errors.New("error")

	defer func() {
		assert.Equal(t, err, handler.Last())
	}()
	defer emperror.HandleRecover(handler)

	panic(err)
}

func TestHandleIfErr(t *testing.T) {
	handler := emperror.NewTestHandler()
	err := errors.New("error")

	emperror.HandleIfErr(handler, err)

	assert.Equal(t, err, handler.Last())
}

func TestHandleIfErr_Nil(t *testing.T) {
	handler := emperror.NewTestHandler()

	emperror.HandleIfErr(handler, nil)

	assert.NoError(t, handler.Last())
}
