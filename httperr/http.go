package httperr

import (
	"fmt"
	"io"
	"net/http"

	"emperror.dev/errors"
)

// WithHTTPRequest attaches an HTTP request to the error.
// Deprecated: no replacement at this time.
func WithHTTPRequest(err error, r *http.Request) error {
	return &withHTTPRequest{
		req: r,
		err: err,
	}
}

// HTTPRequest extracts an HTTP request from an error (if any).
//
// It loops through the whole error chain (if any).
// Deprecated: no replacement at this time.
func HTTPRequest(err error) (*http.Request, bool) {
	type httpError interface {
		HTTPRequest() *http.Request
	}

	var req *http.Request

	errors.UnwrapEach(err, func(err error) bool {
		if httpErr, ok := err.(httpError); ok {
			req = httpErr.HTTPRequest()

			return false
		}

		return true
	})

	return req, req != nil
}

type withHTTPRequest struct {
	req *http.Request
	err error
}

func (w *withHTTPRequest) Error() string {
	return w.err.Error()
}

func (w *withHTTPRequest) Cause() error  { return w.err }
func (w *withHTTPRequest) Unwrap() error { return w.err }

func (w *withHTTPRequest) HTTPRequest() *http.Request {
	return w.req
}

func (w *withHTTPRequest) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			_, _ = fmt.Fprintf(s, "%+v", w.Cause())

			return
		}

		fallthrough

	case 's':
		_, _ = io.WriteString(s, w.Error())

	case 'q':
		_, _ = fmt.Fprintf(s, "%q", w.Error())
	}
}
