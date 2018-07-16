/*
Package bugsnag provides Bugsnag integration.

Bugsnag recommends to create a global scope error handler.
Although that works, makes writing software harder, so this package recommends creating a separate instance:

	APIKey := "key"

	handler := bugsnag.NewHandler(APIKey)
*/
package bugsnag

import "github.com/bugsnag/bugsnag-go"

// handler is responsible for sending errors to Bugsnag.
type handler struct {
	notifier *bugsnag.Notifier
}

// NewHandler creates a new Bugsnag handler.
func NewHandler(APIKey string) *handler {
	return &handler{
		bugsnag.New(bugsnag.Configuration{
			APIKey: APIKey,
		}),
	}
}

// Handle calls the underlying Bugsnag notifier.
func (h *handler) Handle(err error) {
	h.notifier.Notify(err)
}
