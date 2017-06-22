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
			err = fmt.Errorf("Unknown panic, received: %v", r)
		}
	}

	return err
}
