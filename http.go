package emperror

import "net/http"

// WithHTTPRequest attaches an HTTP request to the error.
func WithHTTPRequest(err error, r *http.Request) error {
	return &httpError{
		req: r,
		err: err,
	}
}

// HTTPRequest extracts an HTTP request from an error (if any).
//
// It loops through the whole error chain (if any).
func HTTPRequest(err error) (*http.Request, bool) {
	type httpError interface {
		HTTPRequest() *http.Request
	}

	var req *http.Request

	ForEachCause(err, func(err error) bool {
		if httpErr, ok := err.(httpError); ok {
			req = httpErr.HTTPRequest()

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

// HTTPRequest returns the current HTTP request.
func (e *httpError) HTTPRequest() *http.Request {
	return e.req
}
