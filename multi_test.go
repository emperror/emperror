package emperror_test

import (
	"testing"

	"fmt"

	"github.com/goph/emperror"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMultiErrorBuilder_ErrOrNil(t *testing.T) {
	builder := emperror.NewMultiErrorBuilder()

	err := fmt.Errorf("error")

	builder.Add(err)

	merr := builder.ErrOrNil()

	require.Implements(t, (*emperror.ErrorCollection)(nil), merr)
	assert.Equal(t, err, merr.(emperror.ErrorCollection).Errors()[0])
}

func TestMultiErrorBuilder_ErrOrNil_NilWhenEmpty(t *testing.T) {
	builder := emperror.NewMultiErrorBuilder()

	assert.NoError(t, builder.ErrOrNil())
}

func TestMultiErrorBuilder_ErrOrNil_Single(t *testing.T) {
	builder := &emperror.MultiErrorBuilder{
		SingleWrapMode: emperror.ReturnSingle,
	}

	err := fmt.Errorf("error")

	builder.Add(err)

	assert.Equal(t, err, builder.ErrOrNil())
}

func TestMultiErrorBuilder_Message(t *testing.T) {
	want := "Multiple errors happened during action"

	builder := &emperror.MultiErrorBuilder{
		Message: want,
	}

	err := fmt.Errorf("error")

	builder.Add(err)

	assert.Equal(t, want, builder.ErrOrNil().Error())
}
