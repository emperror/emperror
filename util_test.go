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
		if handler.Last() != err {
			t.Error("handler was expected to handle the error")
		}
	}()
	defer emperror.HandleRecover(handler)

	panic(err)
}

func TestHandleIfErr(t *testing.T) {
	handler := emperror.NewTestHandler()
	err := errors.New("error")

	emperror.HandleIfErr(handler, err)

	if handler.Last() != err {
		t.Error("handler was expected to handle the error")
	}
}

func TestHandleIfErr_Nil(t *testing.T) {
	handler := emperror.NewTestHandler()

	emperror.HandleIfErr(handler, nil)

	if handler.Last() != nil {
		t.Error("handler was not expected to handle anything")
	}
}
