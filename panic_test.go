package emperror

import (
	"errors"
	"testing"
)

func panicErrorFunc() error {
	return errors.New("error")
}

func TestPanic(t *testing.T) {
	err := panicErrorFunc()

	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("expected to recover from a panic, but nothing panicked")
		}

		err, ok := r.(error)
		if !ok {
			t.Fatal("expected to recover an error from a panic")
		}

		if err == nil {
			t.Fatal("expected to the recovered error to be an error, received nil")
		}
	}()

	Panic(err)
}

func TestPanic_NoError(t *testing.T) {
	defer func() {
		r := recover()
		if r != nil {
			t.Fatal("unexpected panic")
		}
	}()

	Panic(nil)
}
