// nolint: goconst
package sentryhandler

import (
	"github.com/getsentry/raven-go"
)

func ExampleNew() {
	dsn := "https://user:password@sentry.io/1234"

	handler, err := New(dsn)
	if err != nil {
		panic(err)
	}
	defer handler.Close() // Make sure to close the handler to flush all error reporting in progress

	// Output:
}

func ExampleNewFromClient() {
	dsn := "https://user:password@sentry.io/1234"

	client, err := raven.New(dsn)
	if err != nil {
		panic(err)
	}

	handler := NewFromClient(client)
	defer handler.Close() // Make sure to close the handler to flush all error reporting in progress

	// Output:
}

func ExampleNewSync() {
	dsn := "https://user:password@sentry.io/1234"

	handler, err := NewSync(dsn)
	if err != nil {
		panic(err)
	}
	defer handler.Close()

	// Output:
}

func ExampleNewSyncFromClient() {
	dsn := "https://user:password@sentry.io/1234"

	client, err := raven.New(dsn)
	if err != nil {
		panic(err)
	}

	handler := NewSyncFromClient(client)
	defer handler.Close()

	// Output:
}
