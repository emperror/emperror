package httperr

import (
	"errors"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestWithHttpRequest(t *testing.T) {
	tests := map[string]struct {
		request *http.Request
		err     error
	}{
		"simple error": {
			request: &http.Request{
				Method: "POST",
				URL: &url.URL{
					Scheme: "http",
					Host:   "localhost",
				},
			},
			err: errors.New("error"),
		},
	}

	for name, test := range tests {
		name, test := name, test

		t.Run(name, func(t *testing.T) {
			err := WithHTTPRequest(test.err, test.request)

			if got, want := err.Error(), test.err.Error(); got != want {
				t.Errorf("error does not match the expected one\nactual:   %v\nexpected: %v", got, want)
			}

			req, ok := HTTPRequest(err)
			if !ok {
				t.Error("error is expected to contain an HTTP request")
			}

			if got, want := req, test.request; !reflect.DeepEqual(got, want) {
				t.Errorf("request does not match the expected one\nactual:   %v\nexpected: %v", got, want)
			}
		})
	}
}
