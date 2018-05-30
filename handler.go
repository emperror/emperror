package emperror

//go:generate sh -c "CGO_ENABLED=0 go run internal/bin/mockery.go -name=Handler -output . -outpkg emperror_test -testonly -case underscore"

// Handler is responsible for handling an error.
//
// This interface allows libraries to decouple from logging and error handling solutions.
type Handler interface {
	// Handle takes care of unhandled errors.
	Handle(err error)
}
