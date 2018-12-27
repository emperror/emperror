package rollbarhandler

import (
	"github.com/rollbar/rollbar-go"
)

func ExampleNew() {
	token := "token"

	_ = New(token, "env", "version", "host", "serverRoot")

	// Output:
}

func ExampleNewFromClient() {
	token := "token"

	_ = NewFromClient(rollbar.New(token, "env", "version", "host", "serverRoot"))

	// Output:
}
