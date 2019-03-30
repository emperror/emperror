package emperror

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanic(t *testing.T) {
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

		st, ok := StackTrace(err)
		if !ok {
			t.Fatal("error is expected to carry a stack trace")
		}

		if got, want := fmt.Sprintf("%n", st[0]), "TestPanic"; got != want {
			t.Errorf("function name does not match the expected one\nactual:   %s\nexpected: %s", got, want)
		}

		if got, want := fmt.Sprintf("%s", st[0]), "panic_test.go"; got != want {
			t.Errorf("file name does not match the expected one\nactual:   %s\nexpected: %s", got, want)
		}

		if got, want := fmt.Sprintf("%d", st[0]), "43"; got != want {
			t.Errorf("line number does not match the expected one\nactual:   %s\nexpected: %s", got, want)
		}
	}()

	Panic(errors.New("error"))
}

func TestPanic_NoError(t *testing.T) {
	defer func() {
		r := recover()
		if r != nil {
			t.Fatalf("unexpected panic, received: %v", r)
		}
	}()

	Panic(nil)
}
