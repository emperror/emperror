package emperror

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestForEachCause(t *testing.T) {
	err := errors.WithMessage(
		errors.WithMessage(
			errors.WithMessage(
				errors.New("level 0"),
				"level 1",
			),
			"level 2",
		),
		"level 3",
	)

	var i int
	fn := func(err error) bool {
		i++

		return true
	}

	ForEachCause(err, fn)

	assert.Equal(t, 4, i)
}

func TestForEachCause_BreakTheLoop(t *testing.T) {
	err := errors.WithMessage(
		errors.WithMessage(
			errors.WithMessage(
				errors.New("level 0"),
				"level 1",
			),
			"level 2",
		),
		"level 3",
	)

	var i int
	fn := func(err error) bool {
		i++

		return !(i > 2)
	}

	ForEachCause(err, fn)

	assert.Equal(t, 3, i)
}

func TestForEachCause_NilError(t *testing.T) {
	var i int
	fn := func(err error) bool {
		i++

		return !(i > 2)
	}

	ForEachCause(nil, fn)

	assert.Equal(t, 0, i)
}
