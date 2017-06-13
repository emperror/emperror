package airbrake

import (
	"net/http"
)

// httpError interface provides a way to attach an http.Request to the error.
type httpError interface {
	Request() *http.Request
}
