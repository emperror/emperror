package emperror_test

import (
	"testing"

	"fmt"

	"github.com/goph/emperror"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

	require.NotPanics(t, func() { f() })

	v := f()

	assert.EqualError(t, v, "internal error")
	assert.Equal(t, err, v.(emperror.Causer).Cause())
	assert.Implements(t, (*emperror.StackTracer)(nil), v)
}

func TestRecover_StringPanic(t *testing.T) {
	f := createRecoverFunc("internal error")

	require.NotPanics(t, func() { f() })

	v := f()

	assert.EqualError(t, v, "internal error")
	assert.Implements(t, (*emperror.StackTracer)(nil), v)
}

func TestRecover_AnyPanic(t *testing.T) {
	f := createRecoverFunc(123)

	require.NotPanics(t, func() { f() })

	v := f()

	assert.EqualError(t, v, "Unknown panic, received: 123")
	assert.Implements(t, (*emperror.StackTracer)(nil), v)
}
