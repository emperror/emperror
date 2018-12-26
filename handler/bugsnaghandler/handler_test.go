// nolint: goconst
package bugsnaghandler

import "github.com/bugsnag/bugsnag-go"

func ExampleNew() {
	apiKey := "key"

	_ = New(apiKey)

	// Output:
}

func ExampleNewFromNotifier() {
	apiKey := "key"

	_ = NewFromNotifier(bugsnag.New(bugsnag.Configuration{
		APIKey: apiKey,
	}))
}
