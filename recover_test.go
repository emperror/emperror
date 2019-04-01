package emperror

import (
	"fmt"
	"testing"
)

func createRecoverFunc(p interface{}) func() error {
	return func() (err error) {
		defer func() {
			err = Recover(recover())
		}()

		panic(p)
	}
}

func assertRecoveredError(t *testing.T, err error, msg string) {
	t.Helper()

	st, ok := StackTrace(err)
	if !ok {
		t.Fatal("error is expected to carry a stack trace")
	}

	if got, want := fmt.Sprintf("%n", st[0]), "createRecoverFunc.func1"; got != want { // nolint: govet
		t.Errorf("function name does not match the expected one\nactual:   %s\nexpected: %s", got, want)
	}

	if got, want := fmt.Sprintf("%s", st[0]), "recover_test.go"; got != want { // nolint: govet
		t.Errorf("file name does not match the expected one\nactual:   %s\nexpected: %s", got, want)
	}

	if got, want := fmt.Sprintf("%d", st[0]), "14"; got != want { // nolint: govet
		t.Errorf("line number does not match the expected one\nactual:   %s\nexpected: %s", got, want)
	}

	if got, want := err.Error(), msg; got != want {
		t.Errorf("error message does not match the expected one\nactual:   %s\nexpected: %s", got, want)
	}
}

func TestRecover_ErrorPanic(t *testing.T) {
	err := fmt.Errorf("internal error")

	f := createRecoverFunc(err)

	v := f()

	assertRecoveredError(t, v, "internal error")
}

func TestRecover_StringPanic(t *testing.T) {
	f := createRecoverFunc("internal error")

	v := f()

	assertRecoveredError(t, v, "internal error")
}

func TestRecover_AnyPanic(t *testing.T) {
	f := createRecoverFunc(123)

	v := f()

	assertRecoveredError(t, v, "unknown panic, received: 123")
}

func TestRecover_Nil(t *testing.T) {
	f := createRecoverFunc(nil)

	v := f()

	if got, want := v, error(nil); got != want { // nolint: govet
		t.Errorf("the recovered value is expected to be nil\nactual: %v", got)
	}
}
