package emperror_test

import (
	"testing"

	"fmt"

	"github.com/goph/emperror"
)

func createRecoverFunc(p interface{}) func() error {
	return func() (err error) {
		defer func() {
			err = emperror.Recover(recover())
		}()

		panic(p)
	}
}

func TestRecover_ErrorPanic(t *testing.T) {
	err := fmt.Errorf("internal error")

	f := createRecoverFunc(err)

	if got, want := f(), err; got != want {
		t.Fatalf("expected to recover a specific error, received: %v", got)
	}
}

func TestRecover_StringPanic(t *testing.T) {
	f := createRecoverFunc("internal error")

	if got, want := f().Error(), "internal error"; got != want {
		t.Fatalf("expected to recover a specific error, received: %v", got)
	}
}

func TestRecover_AnyPanic(t *testing.T) {
	f := createRecoverFunc(123)

	if got, want := f().Error(), "Unknown panic, received: 123"; got != want {
		t.Fatalf("expected to recover a specific error, received: %v", got)
	}
}
