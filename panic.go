package emperror

// Panic panics if the passed error is not nil.
// If the error does not contain any stack trace, the function attaches one, starting from the frame of the
// "Panic" function call.
//
// This function is useful with HandleRecover when panic is used as a flow control tool to stop the application.
func Panic(err error) {
	if err != nil {
		if _, ok := StackTrace(err); !ok {
			err = &wrappedError{
				err:   err,
				stack: callers(),
			}
		}

		panic(err)
	}
}
