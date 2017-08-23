package errors

import "fmt"

// Recover accepts a recovered panic (if any) and converts it to an error (if necessary).
func Recover(r interface{}) (err error) {
	if r != nil {
		switch x := r.(type) {
		case string:
			err = NewWithStackTrace(x)
		case error:
			if _, ok := x.(StackTracer); !ok {
				x = WithStack(x)
			}

			err = x
		default:
			err = NewWithStackTrace(fmt.Sprintf("Unknown panic, received: %v", r))
		}
	}

	return err
}
