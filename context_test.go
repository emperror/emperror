package emperror_test

import (
	"testing"

	"github.com/goph/emperror"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWith(t *testing.T) {
	err := errors.New("error")

	kvs := []interface{}{"a", 123}
	err = emperror.With(err, kvs...)
	kvs[1] = 0 // With should copy its key values

	ctx := emperror.Context(err)

	assert.Equal(t, "a", ctx[0])
	assert.Equal(t, 123, ctx[1])
	assert.EqualError(t, err, "error")
}

func TestWith_Multiple(t *testing.T) {
	err := errors.New("")

	err = emperror.With(emperror.With(err, "a", 123), "b", 321)

	ctx := emperror.Context(err)

	assert.Equal(t, "a", ctx[0])
	assert.Equal(t, 123, ctx[1])
	assert.Equal(t, "b", ctx[2])
	assert.Equal(t, 321, ctx[3])
}

func TestContextor_MissingValue(t *testing.T) {
	err := errors.New("")

	err = emperror.With(emperror.With(err, "k0"), "k1")

	ctx := emperror.Context(err)

	require.Len(t, ctx, 4)

	for i := 1; i < 4; i += 2 {
		assert.Nil(t, ctx[i])
	}
}

func TestContext(t *testing.T) {
	err := emperror.With(
		errors.WithMessage(
			emperror.With(
				errors.Wrap(
					emperror.With(
						errors.New("error"),
						"key", "value",
					),
					"wrapped error",
				),
				"key2", "value2",
			),
			"another wrapped error",
		),
		"key3", "value3",
	)

	expected := []interface{}{
		"key", "value",
		"key2", "value2",
		"key3", "value3",
	}

	actual := emperror.Context(err)

	assert.Equal(t, expected, actual)
}
