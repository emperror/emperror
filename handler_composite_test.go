package emperror

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompositeHandler(t *testing.T) {
	handler1 := NewTestHandler()
	handler2 := NewTestHandler()

	handler := NewCompositeHandler(handler1, handler2)

	err := errors.New("error")

	handler.Handle(err)

	assert.Equal(t, err, handler1.LastError())
	assert.Equal(t, err, handler2.LastError())
}
