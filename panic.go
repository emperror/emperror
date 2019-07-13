package emperror

import (
	"fmt"

	"emperror.dev/errors"
)

// Panic panics if the passed error is not nil.
// If the error does not contain any stack trace, the function attaches one, starting from the frame of the
// "Panic" function call.
//
// This function is useful with HandleRecover when panic is used as a flow control tool to stop the application.
func Panic(err error) {
	if err != nil {
		var st stackTracer
		if !errors.As(err, &st) {
			err = errors.WithStackDepth(err, 1)
		}

		panic(err)
	}
}

// Recover accepts a recovered panic (if any) and converts it to an error (if necessary).
func Recover(r interface{}) (err error) {
	if r != nil {
		switch x := r.(type) {
		case string:
			err = errors.NewPlain(x)
		case error:
			err = x
		default:
			err = fmt.Errorf("unknown panic, received: %v", r)
		}

		var st stackTracer
		if !errors.As(err, &st) {
			err = errors.WithStackDepth(err, 3)
		}
	}

	return err
}

// HandleRecover recovers from a panic and handles the error.
//
// 		defer emperror.HandleRecover(errorHandler)
func HandleRecover(handler Handler) {
	err := Recover(recover())
	if err != nil {
		handler.Handle(err)
	}
}
