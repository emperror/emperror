package emperror_test

import (
	"testing"

	stderrors "errors"

	"github.com/goph/emperror"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func TestStackTrace(t *testing.T) {
	err := errors.WithMessage(errors.New("error"), "wrapper")

	stack, ok := emperror.StackTrace(err)

	assert.True(t, ok)
	assert.NotNil(t, stack)
}

func TestExposeStackTrace(t *testing.T) {
	err := errors.WithMessage(errors.New("error"), "wrapper")

	err = emperror.ExposeStackTrace(err)

	stack := err.(stackTracer).StackTrace()

	assert.NotEmpty(t, stack)
}

func TestExposeStackTrace_NoStackTrace(t *testing.T) {
	err := stderrors.New("error")

	serr := emperror.ExposeStackTrace(err)

	assert.Equal(t, err, serr)
}
