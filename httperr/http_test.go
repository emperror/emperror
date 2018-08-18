package httperr_test

import (
	"testing"

	"errors"
	"net/http"
	"net/url"

	. "github.com/goph/emperror/httperr"
	"github.com/stretchr/testify/assert"
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
		t.Run(name, func(t *testing.T) {
			nerr := WithHTTPRequest(test.err, test.request)

			assert.EqualError(t, nerr, test.err.Error())

			req, ok := HTTPRequest(nerr)
			assert.True(t, ok)
			assert.Equal(t, test.request, req)
		})
	}
}
