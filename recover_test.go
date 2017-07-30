package emperror_test

import (
	"testing"

	"fmt"

	"github.com/goph/emperror"
	"github.com/stretchr/testify/assert"
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

	assert.Equal(t, err, f())
}

func TestRecover_StringPanic(t *testing.T) {
	f := createRecoverFunc("internal error")

	assert.Equal(t, "internal error", f().Error())
}

func TestRecover_AnyPanic(t *testing.T) {
	f := createRecoverFunc(123)

	assert.Equal(t, "Unknown panic, received: 123", f().Error())
}
