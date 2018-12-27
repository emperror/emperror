// nolint: goconst
package airbrakehandler

import "github.com/airbrake/gobrake"

func ExampleNew() {
	projectID := int64(1)
	projectKey := "key"

	handler := New(projectID, projectKey)
	defer handler.Close() // Make sure to close the handler to flush all error reporting in progress

	// Output:
}

func ExampleNewFromNotifier() {
	projectID := int64(1)
	projectKey := "key"

	notifier := gobrake.NewNotifier(projectID, projectKey)
	handler := NewFromNotifier(notifier)
	defer handler.Close() // Make sure to close the handler to flush all error reporting in progress

	// Output:
}

func ExampleNewSync() {
	projectID := int64(1)
	projectKey := "key"

	handler := NewSync(projectID, projectKey)
	defer handler.Close()

	// Output:
}

func ExampleNewSyncFromNotifier() {
	projectID := int64(1)
	projectKey := "key"

	notifier := gobrake.NewNotifier(projectID, projectKey)
	handler := NewSyncFromNotifier(notifier)
	defer handler.Close()

	// Output:
}
