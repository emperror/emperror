// nolint: goconst
package airbrakehandler_test

import (
	"github.com/airbrake/gobrake"

	"emperror.dev/emperror/handler/airbrakehandler"
)

func ExampleNew() {
	projectID := int64(1)
	projectKey := "key"

	handler := airbrakehandler.New(projectID, projectKey)
	defer handler.Close() // Make sure to close the handler to flush all error reporting in progress

	// Output:
}

func ExampleNewFromNotifier() {
	projectID := int64(1)
	projectKey := "key"

	notifier := gobrake.NewNotifier(projectID, projectKey)
	handler := airbrakehandler.NewFromNotifier(notifier)
	defer handler.Close() // Make sure to close the handler to flush all error reporting in progress

	// Output:
}

func ExampleNewSync() {
	projectID := int64(1)
	projectKey := "key"

	handler := airbrakehandler.NewSync(projectID, projectKey)
	defer handler.Close()

	// Output:
}

func ExampleNewSyncFromNotifier() {
	projectID := int64(1)
	projectKey := "key"

	notifier := gobrake.NewNotifier(projectID, projectKey)
	handler := airbrakehandler.NewSyncFromNotifier(notifier)
	defer handler.Close()

	// Output:
}
