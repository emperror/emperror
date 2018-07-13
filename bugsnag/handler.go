/*
Package bugsnag provides Bugsnag integration.

Bugsnag recommends to create a global scope error handler.
Although that works, makes writing software harder,
so this package recommends creating a separate instance:

	APIKey := "key"

	handler := bugsnag.NewHandler(APIKey)

If you need access to the underlying Notifier instance (or need more advanced construction), you can access it from the handler:

	handler.Notifier.Config.AppVersion = "1.0.0"

Bugsnag provides an extensive set of configuration options, so it might make sense to manually construct it:

	// Note: there is a conflict in package names which we resolved with the "_" (underline) prefix.
	handler := &_bugsnag.Handler {
		Notifier: bugsnag.New(bugsnag.Configuration{
			APIKey:      APIKey,
			Synchronous: true,
		})
	}
*/
package bugsnag

import bugsnag "github.com/bugsnag/bugsnag-go"

// Handler is responsible for sending errors to Bugsnag.
type Handler struct {
	Notifier *bugsnag.Notifier
}

// NewHandler creates a new Bugsnag handler.
func NewHandler(APIKey string) *Handler {
	return &Handler{
		Notifier: bugsnag.New(bugsnag.Configuration{
			APIKey:      APIKey,
			Synchronous: true,
		}),
	}
}

// Handle calls the underlying Bugsnag notifier.
func (h *Handler) Handle(err error) {
	h.Notifier.Notify(err)
}
