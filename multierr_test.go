package emperror

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type errorCollection interface {
	Errors() []error
}

func TestMultiErrorBuilder_ErrOrNil(t *testing.T) {
	builder := NewMultiErrorBuilder()

	err := fmt.Errorf("error")

	builder.Add(err)

	merr := builder.ErrOrNil()

	assert.Equal(t, err, merr.(errorCollection).Errors()[0])
}

func TestMultiErrorBuilder_ErrOrNil_NilWhenEmpty(t *testing.T) {
	builder := NewMultiErrorBuilder()

	assert.NoError(t, builder.ErrOrNil())
}

func TestMultiErrorBuilder_ErrOrNil_Single(t *testing.T) {
	builder := &MultiErrorBuilder{
		SingleWrapMode: ReturnSingle,
	}

	err := fmt.Errorf("error")

	builder.Add(err)

	assert.Equal(t, err, builder.ErrOrNil())
}

func TestMultiErrorBuilder_Message(t *testing.T) {
	want := "Multiple errors happened during action"

	builder := &MultiErrorBuilder{
		Message: want,
	}

	err := fmt.Errorf("error")

	builder.Add(err)

	assert.Equal(t, want, builder.ErrOrNil().Error())
}
