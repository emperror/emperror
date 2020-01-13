package emperror

import (
	"context"
	"errors"
	"fmt"

	"emperror.dev/errors/match"
)

func ExampleWithFilter() {
	err := errors.New("no more errors")
	err2 := errors.New("one last error")
	isErr := errors.New("is")

	handler := WithFilter(
		ErrorHandlerFunc(func(err error) { fmt.Println(err) }),
		match.Is(isErr),
	)

	handler.Handle(err)
	handler.Handle(isErr)
	handler.HandleContext(context.Background(), err2)
	handler.HandleContext(context.Background(), isErr)

	// Output:
	// no more errors
	// one last error
}
