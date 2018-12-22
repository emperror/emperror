package emperror_test

import (
	"errors"
	"testing"

	. "github.com/goph/emperror"
	"github.com/stretchr/testify/assert"
)

func TestHandlerFunc(t *testing.T) {
	var actual error
	log := func(err error) {
		actual = err
	}

	fn := HandlerFunc(log)

	expected := errors.New("error")

	fn.Handle(expected)

	assert.Equal(t, expected, actual)
}

func TestHandlerLogFunc(t *testing.T) {
	var actual error
	log := func(args ...interface{}) {
		actual = args[0].(error)
	}

	fn := HandlerLogFunc(log)

	expected := errors.New("error")

	fn.Handle(expected)

	assert.Equal(t, expected, actual)
}
