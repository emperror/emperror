// nolint: goconst
package bugsnaghandler_test

import (
	"github.com/bugsnag/bugsnag-go"

	"github.com/goph/emperror/handler/bugsnaghandler"
)

func ExampleNew() {
	apiKey := "key"

	_ = bugsnaghandler.New(apiKey)

	// Output:
}

func ExampleNewFromNotifier() {
	apiKey := "key"

	_ = bugsnaghandler.NewFromNotifier(bugsnag.New(bugsnag.Configuration{
		APIKey: apiKey,
	}))

	// Output:
}
