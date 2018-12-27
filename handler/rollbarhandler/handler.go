package rollbarhandler

import (
	"github.com/goph/emperror"
	"github.com/goph/emperror/httperr"
	"github.com/goph/emperror/internal/keyvals"
	"github.com/rollbar/rollbar-go"
)

type Handler struct {
	client *rollbar.Client
}

func New(token, environment, codeVersion, serverHost, serverRoot string) *Handler {
	return NewFromClient(rollbar.New(token, environment, codeVersion, serverHost, serverRoot))
}

func NewFromClient(client *rollbar.Client) *Handler {
	return &Handler{
		client: client,
	}
}

func (h *Handler) Handle(err error) {
	// Get HTTP request (if any)
	req, httpok := httperr.HTTPRequest(err)

	ctx := keyvals.ToMap(emperror.Context(err))

	// Expose the stackTracer interface on the outer error (if there is stack trace in the error)
	err = emperror.ExposeStackTrace(err)

	// Convert error with stack trace to an internal error type
	if e, ok := err.(stackTracer); ok {
		err = newCauseStacker(e)
	}

	if httpok {
		h.client.RequestErrorWithStackSkipWithExtras(rollbar.ERR, req, err, 3, ctx)

		return
	}

	h.client.ErrorWithStackSkipWithExtras(rollbar.ERR, err, 3, ctx)
}

func (h *Handler) Close() error {
	return h.client.Close()
}
