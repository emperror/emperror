package airbrakehandler

import (
	"net/http"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/airbrake/gobrake"
	"github.com/pkg/errors"

	"emperror.dev/emperror"
	"emperror.dev/emperror/httperr"
)

func newHandler(t *testing.T) *Handler {
	host := os.Getenv("AIRBRAKE_HOST")
	projectID, _ := strconv.Atoi(os.Getenv("AIRBRAKE_PROJECT_ID"))
	projectKey := os.Getenv("AIRBRAKE_PROJECT_KEY")

	if host == "" || projectID == 0 || projectKey == "" {
		t.Skip("missing airbrake credentials")
	}

	return NewFromNotifier(gobrake.NewNotifierWithOptions(&gobrake.NotifierOptions{
		ProjectId:  int64(projectID),
		ProjectKey: projectKey,
		Host:       host,
	}))
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
