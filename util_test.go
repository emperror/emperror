package emperror_test

import (
	"errors"
	"testing"

	"github.com/goph/emperror"
)

func TestHandleRecovery(t *testing.T) {
	handler := emperror.NewTestHandler()
	err := errors.New("error")

	defer func() {
		if want, have := err, handler.Last(); want != have {
			t.Errorf("\nwant: %v\nhave: %v", want, have)
		}
	}()
	defer emperror.HandleRecover(handler)

	panic(err)
}

func TestHandleIfErr(t *testing.T) {
	t.Parallel()

	handler := emperror.NewTestHandler()
	err := errors.New("error")

	emperror.HandleIfErr(handler, err)

	if want, have := err, handler.Last(); want != have {
		t.Errorf("\nwant: %v\nhave: %v", want, have)
	}
}

func TestHandleIfErr_Nil(t *testing.T) {
	t.Parallel()

	handler := emperror.NewTestHandler()

	emperror.HandleIfErr(handler, nil)

	var want, have error
	if want, have = nil, handler.Last(); want != have {
		t.Errorf("\nwant: %v\nhave: %v", want, have)
	}
}
