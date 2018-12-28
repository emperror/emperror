package rollbarhandler

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/goph/emperror"
	"github.com/goph/emperror/httperr"
	"github.com/pkg/errors"
)

func newHandler(t *testing.T) *Handler {
	token := os.Getenv("ROLLBAR_TOKEN")

	if token == "" {
		t.Skip("missing rollbar credentials")
	}

	return New(token, "test", "latest", "localhost", "github.com/goph/emperror")
}

func TestIntegration_Handler(t *testing.T) {
	handler := newHandler(t)
	defer handler.Close()

	err := errors.New("error")

	handler.Handle(err)

	// Wait for the notice to reach the queue before closing
	time.Sleep(500 * time.Millisecond)
}

func TestIntegration_WithContext(t *testing.T) {
	handler := newHandler(t)
	defer handler.Close()

	err := emperror.With(errors.New("error with context"), "key", "value")

	handler.Handle(err)

	// Wait for the notice to reach the queue before closing
	time.Sleep(500 * time.Millisecond)
}

func TestIntegration_WithHTTPRequest(t *testing.T) {
	handler := newHandler(t)
	defer handler.Close()

	req, e := http.NewRequest("GET", "https://google.com", nil)
	if e != nil {
		t.Fatal(e)
	}

	err := httperr.WithHTTPRequest(errors.New("error with http request"), req)

	handler.Handle(err)

	// Wait for the notice to reach the queue before closing
	time.Sleep(500 * time.Millisecond)
}
