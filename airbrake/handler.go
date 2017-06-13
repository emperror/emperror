package airbrake

import (
	"net/http"

	"github.com/airbrake/gobrake"
)

// Handler is responsible for sending errors to Airbrake/Errbit.
type Handler struct {
	Gobrake           *gobrake.Notifier
	SendSynchronously bool
}

// NewHandler creates a new Airbrake handler.
func NewHandler(projectID int64, projectKey string) *Handler {
	return &Handler{
		Gobrake: gobrake.NewNotifier(projectID, projectKey),
	}
}

// Handle calls the underlying Airbrake notifier.
func (h *Handler) Handle(err error) {
	var req *http.Request

	if err, ok := err.(httpError); ok {
		req = err.Request()
	}

	notice := h.Gobrake.Notice(err, req, 1)

	if h.SendSynchronously {
		h.Gobrake.SendNotice(notice)
	} else {
		h.Gobrake.SendNoticeAsync(notice)
	}

}

// Close closes the underlying Airbrake instance.
func (h *Handler) Close() error {
	return h.Gobrake.Close()
}
