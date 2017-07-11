package internal

// ErrorWithContext implements the ContextualError interface.
type ErrorWithContext struct {
	Msg     string
	Keyvals []interface{}
}

// Error implements the error interface.
func (e ErrorWithContext) Error() string {
	return e.Msg
}

// Context implements the ContextualError interface.
func (e ErrorWithContext) Context() []interface{} {
	return e.Keyvals
}
