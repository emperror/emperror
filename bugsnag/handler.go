/*
Package bugsnag provides Bugsnag integration.

Bugsnag recommends to create a global scope error handler.
Although that works, makes writing software harder, so this package recommends creating a separate instance:

	APIKey := "key"

	handler := bugsnag.NewHandler(APIKey)

If you need more control over th underlying Notifier instance (eg. more advanced construction),
you can create a custom one and then create a handler using it:

	// Note: there is a conflict in package names which is resolved here with the "_" (underline) prefix.
	handler := &_bugsnag.NewHandlerFromNotifier(bugsnag.New(bugsnag.Configuration{
		APIKey:      APIKey,
		AppVersion:  "1.0.0",
		Synchronous: true,
	}))
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

// NewHandlerFromNotifier creates a new Bugsnag handler from a notifier instance.
func NewHandlerFromNotifier(notifier *bugsnag.Notifier) *handler {
	return &handler{notifier}
}

// Handle calls the underlying Bugsnag notifier.
func (h *handler) Handle(err error) {
	h.notifier.Notify(err)
}
