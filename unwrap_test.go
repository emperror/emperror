package emperror

import (
	"testing"

	"github.com/pkg/errors"
)

func TestUnwrapNext(t *testing.T) {
	nextErr := errors.New("level 0")
	err := Wrap(
		nextErr,
		"level 1",
	)

	actualNextErr := UnwrapNext(err)

	if got, want := actualNextErr, nextErr; got != want {
		t.Errorf("next error does not match the expected one\nactual:   %s\nexpected: %s", got, want)
	}
}

func TestUnwrapNext_Cause(t *testing.T) {
	nextErr := errors.New("level 0")
	err := errors.WithMessage(
		nextErr,
		"level 1",
	)

	actualNextErr := UnwrapNext(err)

	if got, want := actualNextErr, nextErr; got != want {
		t.Errorf("next error does not match the expected one\nactual:   %s\nexpected: %s", got, want)
	}
}

func TestUnwrap(t *testing.T) {
	lastErr := errors.New("level 0")
	err := Wrap(
		errors.WithMessage(
			Wrap(
				lastErr,
				"level 1",
			),
			"level 2",
		),
		"level 3",
	)

	actualLastErr := Unwrap(err)

	if got, want := actualLastErr, lastErr; got != want {
		t.Errorf("last error does not match the expected one\nactual:   %s\nexpected: %s", got, want)
	}
}

func TestUnwrap_NoChain(t *testing.T) {
	lastErr := errors.New("level 0")
	err := lastErr

	actualLastErr := Unwrap(err)

	if got, want := actualLastErr, lastErr; got != want {
		t.Errorf("last error does not match the expected one\nactual:   %s\nexpected: %s", got, want)
	}
}

func TestUnwrapEach(t *testing.T) {
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

	UnwrapEach(err, fn)

	if got, want := i, 4; got != want {
		t.Errorf("error chain length does not match the expected one\nactual:   %d\nexpected: %d", got, want)
	}
}

func TestUnwrapEach_BreakTheLoop(t *testing.T) {
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

	UnwrapEach(err, fn)

	if got, want := i, 3; got != want {
		t.Errorf("error chain length does not match the expected one\nactual:   %d\nexpected: %d", got, want)
	}
}

func TestUnwrapEach_NilError(t *testing.T) {
	var i int
	fn := func(err error) bool {
		i++

		return !(i > 2)
	}

	UnwrapEach(nil, fn)

	if got, want := i, 0; got != want {
		t.Errorf("error chain length does not match the expected one\nactual:   %d\nexpected: %d", got, want)
	}
}

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
