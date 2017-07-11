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

	if want, have := err, f(); want != have {
		t.Errorf("\nwant: %v\nhave: %v", want, have)
	}
}

func TestRecover_StringPanic(t *testing.T) {
	f := createRecoverFunc("internal error")

	if want, have := "internal error", f().Error(); want != have {
		t.Errorf("\nwant: %v\nhave: %v", want, have)
	}
}

func TestRecover_AnyPanic(t *testing.T) {
	f := createRecoverFunc(123)

	if want, have := "Unknown panic, received: 123", f().Error(); want != have {
		t.Errorf("\nwant: %v\nhave: %v", want, have)
	}
}
