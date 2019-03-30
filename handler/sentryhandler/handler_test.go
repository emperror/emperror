// nolint: goconst
package sentryhandler_test

import (
	"github.com/getsentry/raven-go"

	"github.com/goph/emperror/handler/sentryhandler"
)

func ExampleNew() {
	dsn := "https://user:password@sentry.io/1234"

	handler, err := sentryhandler.New(dsn)
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

	handler := sentryhandler.NewFromClient(client)
	defer handler.Close() // Make sure to close the handler to flush all error reporting in progress

	// Output:
}

func ExampleNewSync() {
	dsn := "https://user:password@sentry.io/1234"

	handler, err := sentryhandler.NewSync(dsn)
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

	handler := sentryhandler.NewSyncFromClient(client)
	defer handler.Close()

	// Output:
}
