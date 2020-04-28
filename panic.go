package emperror

// HandleRecover recovers from a panic and handles the error.
//
//		defer emperror.HandleRecover(errorHandler)
func HandleRecover(handler ErrorHandler) {
	err := Recover(recover())
	if err != nil {
		handler.Handle(err)
	}
}
