package sentryhandler

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/pkg/errors"

	"emperror.dev/emperror"
	"emperror.dev/emperror/httperr"
)

func newHandler(t *testing.T) *Handler {
	dsn := os.Getenv("SENTRY_DSN")

	if dsn == "" {
		t.Skip("missing sentry credentials")
	}

	handler, err := New(dsn)
	if err != nil {
		t.Fatal(err)
	}

	return handler
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
