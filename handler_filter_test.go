package emperror

import (
	"errors"
	"fmt"

	"emperror.dev/errors/match"
)

func ExampleWithFilter() {
	err := errors.New("no more errors")
	isErr := errors.New("is")

	type asError struct {
		error
	}
	asErr := asError{errors.New("as")}

	handler := WithFilter(
		HandlerFunc(func(err error) { fmt.Println(err) }),
		match.Any{
			match.Is(isErr),
			match.As(&asError{}),
		},
	)

	handler.Handle(err)
	handler.Handle(isErr)
	handler.Handle(asErr)

	// Output: no more errors
}
