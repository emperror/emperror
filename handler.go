package emperror

// Handler is responsible for handling an error.
//
// This interface allows libraries to decouple from logging solutions.
// In most cases the implementation will provide some log functionalities though.
type Handler interface {
	// Handle processes an error.
	Handle(err error)
}
