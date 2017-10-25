package emperror_test

import (
	"testing"

	stderrors "errors"
	"net/http"
	"net/url"

	"github.com/goph/emperror"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
			err: stderrors.New("error"),
		},
		"error with stacktrace": {
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
			nerr := emperror.WithHttpRequest(test.err, test.request)

			assert.Implements(t, new(emperror.HttpError), nerr)
			assert.EqualError(t, nerr, test.err.Error())

			if serr, ok := test.err.(emperror.StackTracer); ok {
				require.Implements(t, new(emperror.StackTracer), nerr)

				snerr := nerr.(emperror.StackTracer)
				assert.Equal(t, serr.StackTrace(), snerr.StackTrace())
			}
		})
	}
}
