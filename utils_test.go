package emperror

import (
	"errors"
	"testing"
)

func TestHandleRecovery(t *testing.T) {
	err := errors.New("error")

	defer HandleRecover(HandlerFunc(func(err error) {
		if got, want := err.Error(), "error"; got != want {
			t.Errorf("error does not match the expected one\nactual:   %s\nexpected: %s", got, want)
		}
	}))

	panic(err)
}
