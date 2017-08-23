package emperror

import "github.com/goph/stdlib/errors"

// HandleRecover recovers from a panic and handles the error.
//
// 		go emperror.HandleRecover(errorHandler)
func HandleRecover(handler errors.Handler) {
	err := errors.Recover(recover())
	if err != nil {
		handler.Handle(err)
	}
}

// HandleIfErr handles an error whenever it occures.
func HandleIfErr(handler errors.Handler, err error) {
	if err != nil {
		handler.Handle(err)
	}
}
