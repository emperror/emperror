package emperror_test

import (
	"testing"

	"github.com/goph/emperror"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestContextor(t *testing.T) {
	err := errors.New("error")

	kvs := []interface{}{"a", 123}
	err = emperror.With(err, kvs...)
	kvs[1] = 0 // With should copy its key values

	require.Implements(t, (*emperror.Contextor)(nil), err)

	ctx := err.(emperror.Contextor).Context()

	assert.Equal(t, "a", ctx[0])
	assert.Equal(t, 123, ctx[1])
	assert.EqualError(t, err, "error")
}

func TestContextor_Multi(t *testing.T) {
	err := errors.New("")

	err = emperror.With(emperror.With(err, "a", 123), "b", 321)

	require.Implements(t, (*emperror.Contextor)(nil), err)

	ctx := err.(emperror.Contextor).Context()

	assert.Equal(t, "a", ctx[0])
	assert.Equal(t, 123, ctx[1])
	assert.Equal(t, "b", ctx[2])
	assert.Equal(t, 321, ctx[3])
}

func TestContextor_MultiPrefix(t *testing.T) {
	err := errors.New("")

	err = emperror.WithPrefix(emperror.With(err, "a", 123), "b", 321)

	require.Implements(t, (*emperror.Contextor)(nil), err)

	ctx := err.(emperror.Contextor).Context()

	assert.Equal(t, "a", ctx[2])
	assert.Equal(t, 123, ctx[3])
	assert.Equal(t, "b", ctx[0])
	assert.Equal(t, 321, ctx[1])
}

func TestContextor_MissingValue(t *testing.T) {
	err := errors.New("")

	err = emperror.WithPrefix(emperror.With(err, "k0"), "k1")

	require.Implements(t, (*emperror.Contextor)(nil), err)

	ctx := err.(emperror.Contextor).Context()

	require.Len(t, ctx, 4)

	for i := 1; i < 4; i += 2 {
		assert.Nil(t, ctx[i])
	}
}
