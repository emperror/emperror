package emperror

import (
	"context"
	"errors"
	"testing"
)

func TestTestErrorHandler(t *testing.T) {
	handler := &TestErrorHandler{}

	err1 := errors.New("error 1")
	err2 := errors.New("error 2")

	handler.Handle(err1)
	handler.Handle(nil)
	handler.Handle(err2)
	handler.Handle(nil)

	if want, have := 2, handler.Count(); want != have {
		t.Errorf("unexpected error count\nexpected: %d\nactual:   %d", want, have)
	}

	if want, have := err2, handler.LastError(); want != have {
		t.Errorf("unexpected last error\nexpected: %s\nactual:   %s", want, have)
	}

	errs := handler.Errors()

	if want, have := err1, errs[0]; want != have {
		t.Errorf("unexpected error [1]\nexpected: %s\nactual:   %s", want, have)
	}

	if want, have := err2, errs[1]; want != have {
		t.Errorf("unexpected error [2]\nexpected: %s\nactual:   %s", want, have)
	}
}

func TestTestErrorHandlerContext(t *testing.T) {
	handler := &TestErrorHandlerContext{}

	err1 := errors.New("error 1")
	ctx1 := context.WithValue(context.Background(), "key1", "value1") // nolint: golint
	err2 := errors.New("error 2")
	ctx2 := context.WithValue(context.Background(), "key2", "value2") // nolint: golint

	handler.HandleContext(ctx1, err1)
	handler.HandleContext(context.TODO(), nil)
	handler.HandleContext(ctx2, err2)
	handler.HandleContext(context.TODO(), nil)

	if want, have := 2, handler.Count(); want != have {
		t.Errorf("unexpected error count\nexpected: %d\nactual:   %d", want, have)
	}

	if want, have := err2, handler.LastError(); want != have {
		t.Errorf("unexpected last error\nexpected: %s\nactual:   %s", want, have)
	}

	if want, have := ctx2, handler.LastContext(); want != have {
		t.Errorf("unexpected last context\nexpected: %v\nactual:   %v", want, have)
	}

	errs := handler.Errors()

	if want, have := err1, errs[0]; want != have {
		t.Errorf("unexpected error [1]\nexpected: %s\nactual:   %s", want, have)
	}

	if want, have := err2, errs[1]; want != have {
		t.Errorf("unexpected error [2]\nexpected: %s\nactual:   %s", want, have)
	}

	ctxs := handler.Contexts()

	if want, have := ctx1, ctxs[0]; want != have {
		t.Errorf("unexpected context [1]\nexpected: %v\nactual:   %v", want, have)
	}

	if want, have := ctx2, ctxs[1]; want != have {
		t.Errorf("unexpected context [2]\nexpected: %v\nactual:   %v", want, have)
	}
}

// nolint: dupl
func TestTestErrorHandlerFacade(t *testing.T) {
	handler := &TestErrorHandlerFacade{}

	err1 := errors.New("error 1")
	ctx1 := context.WithValue(context.Background(), "key1", "value1") // nolint: golint
	err2 := errors.New("error 2")
	err3 := errors.New("error 3")
	ctx3 := context.WithValue(context.Background(), "key3", "value3") // nolint: golint

	handler.HandleContext(ctx1, err1)
	handler.HandleContext(context.TODO(), nil)
	handler.Handle(err2)
	handler.HandleContext(ctx3, err3)
	handler.HandleContext(context.TODO(), nil)

	if want, have := 3, handler.Count(); want != have {
		t.Errorf("unexpected error count\nexpected: %d\nactual:   %d", want, have)
	}

	if want, have := err3, handler.LastError(); want != have {
		t.Errorf("unexpected last error\nexpected: %s\nactual:   %s", want, have)
	}

	if want, have := ctx3, handler.LastContext(); want != have {
		t.Errorf("unexpected last context\nexpected: %v\nactual:   %v", want, have)
	}

	errs := handler.Errors()

	if want, have := err1, errs[0]; want != have {
		t.Errorf("unexpected error [1]\nexpected: %s\nactual:   %s", want, have)
	}

	if want, have := err2, errs[1]; want != have {
		t.Errorf("unexpected error [2]\nexpected: %s\nactual:   %s", want, have)
	}

	if want, have := err3, errs[2]; want != have {
		t.Errorf("unexpected error [3]\nexpected: %s\nactual:   %s", want, have)
	}

	ctxs := handler.Contexts()

	if want, have := ctx1, ctxs[0]; want != have {
		t.Errorf("unexpected context [1]\nexpected: %v\nactual:   %v", want, have)
	}

	if ctxs[1] != nil {
		t.Errorf("unexpected context [2]\nactual:   %v\nexpected: %v", ctxs[1], nil)
	}

	if want, have := ctx3, ctxs[2]; want != have {
		t.Errorf("unexpected context [3]\nexpected: %v\nactual:   %v", want, have)
	}
}

// nolint: dupl
func TestTestErrorHandlerSet(t *testing.T) {
	handler := &TestErrorHandlerSet{}

	err1 := errors.New("error 1")
	ctx1 := context.WithValue(context.Background(), "key1", "value1") // nolint: golint
	err2 := errors.New("error 2")
	err3 := errors.New("error 3")
	ctx3 := context.WithValue(context.Background(), "key3", "value3") // nolint: golint

	handler.HandleContext(ctx1, err1)
	handler.HandleContext(context.TODO(), nil)
	handler.Handle(err2)
	handler.HandleContext(ctx3, err3)
	handler.HandleContext(context.TODO(), nil)

	if want, have := 3, handler.Count(); want != have {
		t.Errorf("unexpected error count\nexpected: %d\nactual:   %d", want, have)
	}

	if want, have := err3, handler.LastError(); want != have {
		t.Errorf("unexpected last error\nexpected: %s\nactual:   %s", want, have)
	}

	if want, have := ctx3, handler.LastContext(); want != have {
		t.Errorf("unexpected last context\nexpected: %v\nactual:   %v", want, have)
	}

	errs := handler.Errors()

	if want, have := err1, errs[0]; want != have {
		t.Errorf("unexpected error [1]\nexpected: %s\nactual:   %s", want, have)
	}

	if want, have := err2, errs[1]; want != have {
		t.Errorf("unexpected error [2]\nexpected: %s\nactual:   %s", want, have)
	}

	if want, have := err3, errs[2]; want != have {
		t.Errorf("unexpected error [3]\nexpected: %s\nactual:   %s", want, have)
	}

	ctxs := handler.Contexts()

	if want, have := ctx1, ctxs[0]; want != have {
		t.Errorf("unexpected context [1]\nexpected: %v\nactual:   %v", want, have)
	}

	if ctxs[1] != nil {
		t.Errorf("unexpected context [2]\nactual:   %v\nexpected: %v", ctxs[1], nil)
	}

	if want, have := ctx3, ctxs[2]; want != have {
		t.Errorf("unexpected context [3]\nexpected: %v\nactual:   %v", want, have)
	}
}

func TestTestHandler_Count(t *testing.T) {
	handler := NewTestHandler()

	err1 := errors.New("error 1")
	err2 := errors.New("error 2")

	handler.Handle(err1)
	handler.Handle(err2)

	if got, want := handler.Count(), 2; got != want {
		t.Errorf("error count not match the expected one\nactual:   %d\nexpected: %d", got, want)
	}
}

func TestTestHandler_LastError(t *testing.T) {
	handler := NewTestHandler()

	err1 := errors.New("error 1")
	err2 := errors.New("error 2")

	handler.Handle(err1)
	handler.Handle(err2)

	if got, want := handler.LastError(), err2; got != want {
		t.Errorf("error does not match the expected one\nactual:   %s\nexpected: %s", got, want)
	}
}

func TestTestHandler_Errors(t *testing.T) {
	handler := NewTestHandler()

	err1 := errors.New("error 1")
	err2 := errors.New("error 2")

	handler.Handle(err1)
	handler.Handle(err2)

	errs := handler.Errors()

	if got, want := errs[0], err1; got != want {
		t.Errorf("error 1 does not match the expected one\nactual:   %s\nexpected: %s", got, want)
	}

	if got, want := errs[1], err2; got != want {
		t.Errorf("error 2 does not match the expected one\nactual:   %s\nexpected: %s", got, want)
	}
}
