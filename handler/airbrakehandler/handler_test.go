package airbrakehandler

import "github.com/airbrake/gobrake"

func ExampleNew() {
	projectID := int64(1)
	projectKey := "key"

	_ = New(projectID, projectKey)

	// Output:
}

func ExampleNewFromNotifier() {
	projectID := int64(1)
	projectKey := "key"

	notifier := gobrake.NewNotifier(projectID, projectKey)
	_ = NewFromNotifier(notifier)

	// Output:
}

func ExampleNewAsync() {
	projectID := int64(1)
	projectKey := "key"

	handler := NewAsync(projectID, projectKey)
	defer handler.Close()

	// Output:
}

func ExampleNewAsyncFromNotifier() {
	projectID := int64(1)
	projectKey := "key"

	notifier := gobrake.NewNotifier(projectID, projectKey)
	handler := NewAsyncFromNotifier(notifier)
	defer handler.Close()

	// Output:
}
