package rollbarhandler_test

import (
	"github.com/rollbar/rollbar-go"

	"emperror.dev/emperror/handler/rollbarhandler"
)

func ExampleNew() {
	token := "token"

	_ = rollbarhandler.New(token, "env", "version", "host", "serverRoot")

	// Output:
}

func ExampleNewFromClient() {
	token := "token"

	_ = rollbarhandler.NewFromClient(rollbar.New(
		token,
		"env",
		"version",
		"host",
		"serverRoot",
	))

	// Output:
}
