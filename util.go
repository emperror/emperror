package emperror

// HandleRecover returns a function that can be deferred.
//
// 		go emperror.HandleRecover(errorHandler)()
func HandleRecover(handler Handler) func() {
	return func() {
		err := Recover(recover())
		if err != nil {
			handler.Handle(err)
		}
	}
}

// HandleIfErr handles an error whenever it occures.
func HandleIfErr(handler Handler, err error) {
	if err != nil {
		handler.Handle(err)
	}
}
