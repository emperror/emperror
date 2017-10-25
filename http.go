package emperror

import "net/http"

// WithHttpRequest attaches an HTTP request to the error.
func WithHttpRequest(err error, r *http.Request) error {
	herr := &httpError{
		httpRequest: r,
		err:         err,
	}

	if serr, ok := err.(StackTracer); ok {
		return struct {
			*httpError
			StackTracer
		}{
			httpError:   herr,
			StackTracer: serr,
		}
	}

	return herr
}

// HttpError interface provides a way to attach an http.Request to the error.
type HttpError interface {
	HttpRequest() *http.Request
}

type httpError struct {
	httpRequest *http.Request
	err         error
}

// Error implements the error interface.
func (e *httpError) Error() string {
	return e.err.Error()
}

// Cause implements the Causer interface.
func (e *httpError) Cause() error {
	return e.err
}

// HttpRequest implements the HttpError interface.
func (e *httpError) HttpRequest() *http.Request {
	return e.httpRequest
}
