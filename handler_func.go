package emperror

// HandlerFunc wraps a function and turns it into an error handler.
type HandlerFunc func(err error)

// Handle calls the underlying log function.
func (h HandlerFunc) Handle(err error) {
	h(err)
}

// HandlerLogFunc wraps a log-like function and turns it into an error handler.
type HandlerLogFunc func(args ...interface{})

// Handle calls the underlying log function.
func (h HandlerLogFunc) Handle(err error) {
	h(err)
}
