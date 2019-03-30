package httperr

import (
	"errors"
	"net/http"
	"net/url"
	"testing"

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
		name, test := name, test

		t.Run(name, func(t *testing.T) {
			err := WithHTTPRequest(test.err, test.request)

			assert.EqualError(t, err, test.err.Error())

			req, ok := HTTPRequest(err)
			assert.True(t, ok)
			assert.Equal(t, test.request, req)
		})
	}
}
