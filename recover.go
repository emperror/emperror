package emperror

import (
	"errors"
	"fmt"
)

// Recover accepts a recovered panic (if any) and converts it to an error (if necessary).
func Recover(r interface{}) (err error) {
	if r != nil {
		switch x := r.(type) {
		case string:
			err = errors.New(x)
		case error:
			err = x
		default:
			err = fmt.Errorf("unknown panic, received: %v", r)
		}

		if _, ok := StackTrace(err); !ok {
			err = &wrappedError{
				err:   err,
				stack: callers()[2:], // TODO: improve callers?
			}
		}
	}

	return err
}
