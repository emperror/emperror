package emperror_test

import (
	stderrors "errors"
	"testing"

	. "github.com/goph/emperror"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func TestStackTrace(t *testing.T) {
	err := errors.WithMessage(errors.New("error"), "wrapper")

	stack, ok := StackTrace(err)

	assert.True(t, ok)
	assert.NotNil(t, stack)
}

func TestExposeStackTrace(t *testing.T) {
	err := errors.WithMessage(errors.New("error"), "wrapper")

	err = ExposeStackTrace(err)

	stack := err.(stackTracer).StackTrace()

	assert.NotEmpty(t, stack)
}

func TestExposeStackTrace_NoStackTrace(t *testing.T) {
	err := stderrors.New("error")

	serr := ExposeStackTrace(err)

	assert.Equal(t, err, serr)
}
