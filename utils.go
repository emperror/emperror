package emperror

// HandleRecover recovers from a panic and handles the error.
//
// 		defer emperror.HandleRecover(errorHandler)
func HandleRecover(handler Handler) {
	err := Recover(recover())
	if err != nil {
		handler.Handle(err)
	}
}

// Handle handles an error whenever it occurs.
func Handle(handler Handler, err error) {
	if err != nil {
		handler.Handle(err)
	}
}
