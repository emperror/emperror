package bugsnaghandler

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/pkg/errors"

	"github.com/goph/emperror"
	"github.com/goph/emperror/httperr"
)

func newHandler(t *testing.T) *Handler {
	apiKey := os.Getenv("BUGSNAG_API_KEY")

	if apiKey == "" {
		t.Skip("missing bugsnag credentials")
	}

	return New(apiKey)
}

func TestIntegration_Handler(t *testing.T) {
	handler := newHandler(t)

	err := errors.New("error")

	handler.Handle(err)

	// Wait for the notice to reach the queue before closing
	time.Sleep(500 * time.Millisecond)
}

func TestIntegration_WithContext(t *testing.T) {
	handler := newHandler(t)

	err := emperror.With(errors.New("error with context"), "key", "value")

	handler.Handle(err)

	// Wait for the notice to reach the queue before closing
	time.Sleep(500 * time.Millisecond)
}

func TestIntegration_WithHTTPRequest(t *testing.T) {
	handler := newHandler(t)

	req, e := http.NewRequest("GET", "https://google.com", nil)
	if e != nil {
		t.Fatal(e)
	}

	err := httperr.WithHTTPRequest(errors.New("error with http request"), req)

	handler.Handle(err)

	// Wait for the notice to reach the queue before closing
	time.Sleep(500 * time.Millisecond)
}
