package emperror

import (
	"errors"
	"fmt"

	"emperror.dev/errors/match"
)

func ExampleWithFilter() {
	err := errors.New("no more errors")
	isErr := errors.New("is")

	handler := WithFilter(
		HandlerFunc(func(err error) { fmt.Println(err) }),
		match.Is(isErr),
	)

	handler.Handle(err)
	handler.Handle(isErr)

	// Output: no more errors
}
