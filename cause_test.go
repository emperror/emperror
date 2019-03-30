package emperror

import (
	"testing"

	"github.com/pkg/errors"
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

	if got, want := i, 4; got != want {
		t.Errorf("error chain length does not match the expected one\nactual:   %d\nexpected: %d", got, want)
	}
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

	if got, want := i, 3; got != want {
		t.Errorf("error chain length does not match the expected one\nactual:   %d\nexpected: %d", got, want)
	}
}

func TestForEachCause_NilError(t *testing.T) {
	var i int
	fn := func(err error) bool {
		i++

		return !(i > 2)
	}

	ForEachCause(nil, fn)

	if got, want := i, 0; got != want {
		t.Errorf("error chain length does not match the expected one\nactual:   %d\nexpected: %d", got, want)
	}
}
