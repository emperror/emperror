package emperror_test

import (
	"testing"

	"github.com/goph/emperror"
	"github.com/stretchr/testify/assert"
)

func TestHandleRecovery(t *testing.T) {
	handler := new(emperror.TestHandler)
	err := emperror.New("error")

	defer func() {
		assert.EqualError(t, handler.Last(), "error")
	}()
	defer emperror.HandleRecover(handler)

	panic(err)
}

func TestHandleIfErr(t *testing.T) {
	handler := new(emperror.TestHandler)
	err := emperror.New("error")

	emperror.HandleIfErr(handler, err)

	assert.Equal(t, err, handler.Last())
}

func TestHandleIfErr_Nil(t *testing.T) {
	handler := new(emperror.TestHandler)

	emperror.HandleIfErr(handler, nil)

	assert.NoError(t, handler.Last())
}
