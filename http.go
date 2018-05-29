package emperror

import "net/http"

// WithHttpRequest attaches an HTTP request to the error.
func WithHttpRequest(err error, r *http.Request) error {
	return &httpError{
		req: r,
		err: err,
	}
}

// HttpRequest extracts an HTTP request from an error (if any).
//
// It loops through the whole error chain (if any).
func HttpRequest(err error) (*http.Request, bool) {
	type httpError interface {
		HttpRequest() *http.Request
	}

	var req *http.Request

	// Get the request from the error chain
	ForEachCause(err, func(err error) bool {
		if httpErr, ok := err.(httpError); ok {
			req = httpErr.HttpRequest()

			return false
		}

		return true
	})

	return req, req != nil
}

type httpError struct {
	req *http.Request
	err error
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
	return e.req
}
