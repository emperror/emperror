package emperror

import (
	"reflect"
	"testing"

	"emperror.dev/errors"
)

func TestWith(t *testing.T) {
	err := errors.New("error")

	kvs := []interface{}{"a", 123}
	err = With(err, kvs...)
	kvs[1] = 0 // With should copy its key values

	ctx := Context(err)

	if got, want := ctx[0], "a"; got != want {
		t.Errorf("context value does not match the expected one\nactual:   %s\nexpected: %s", got, want)
	}

	if got, want := ctx[1], 123; got != want {
		t.Errorf("context value does not match the expected one\nactual:   %s\nexpected: %d", got, want)
	}

	if got, want := err.Error(), "error"; got != want {
		t.Errorf("error does not match the expected one\nactual:   %s\nexpected: %s", got, want)
	}
}

func TestWith_Multiple(t *testing.T) {
	err := errors.New("")

	err = With(With(err, "a", 123), "b", 321)

	ctx := Context(err)

	if got, want := ctx[0], "a"; got != want {
		t.Errorf("context value does not match the expected one\nactual:   %s\nexpected: %s", got, want)
	}

	if got, want := ctx[1], 123; got != want {
		t.Errorf("context value does not match the expected one\nactual:   %s\nexpected: %d", got, want)
	}

	if got, want := ctx[2], "b"; got != want {
		t.Errorf("context value does not match the expected one\nactual:   %s\nexpected: %s", got, want)
	}

	if got, want := ctx[3], 321; got != want {
		t.Errorf("context value does not match the expected one\nactual:   %s\nexpected: %d", got, want)
	}
}

func TestContextor_MissingValue(t *testing.T) {
	err := errors.New("")

	err = With(With(err, "k0"), "k1")

	ctx := Context(err)

	if got, want := len(ctx), 4; got != want {
		t.Fatalf("context does not have the required length\nactual:   %d\nexpected: %d", got, want)
	}

	for i := 1; i < 4; i += 2 {
		if ctx[i] != nil {
			t.Errorf("context value %d is expected to be nil\nactual: %v", i, ctx[i])
		}
	}
}

func TestContext(t *testing.T) {
	err := With(
		errors.WithMessage(
			With(
				errors.Wrap(
					With(
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

	actual := Context(err)

	if got, want := actual, expected; !reflect.DeepEqual(got, want) {
		t.Errorf("context does not match the expected one\nactual:   %v\nexpected: %v", got, want)
	}
}

func TestWith_NilError(t *testing.T) {
	err := With(nil)

	if err != nil {
		t.Errorf("error is expected to be nil\nactual: %v", err)
	}
}
