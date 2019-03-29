package emperror

// Panic panics if the passed error is not nil.
//
// This function is useful with HandleRecover when panic is used as a flow control tool to stop the application.
func Panic(err error) {
	if err != nil {
		panic(err)
	}
}
