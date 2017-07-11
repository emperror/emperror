package internal

// ContextualError exposes the same interface as the one in github.com/goph/stdlib/errors.
type ContextualError interface {
	Context() []interface{}
}
