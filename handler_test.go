package emperror

import (
	"context"
	"errors"
	"testing"
)

func TestHandlers(t *testing.T) {
	handler1 := NewTestHandler()
	handler2 := NewTestHandler()

	handler := Handlers{handler1, handler2}

	err := errors.New("error")

	handler.Handle(err)

	if got, want := handler1.LastError(), err; got != want {
		t.Errorf("error does not match the expected one\nactual:   %s\nexpected: %s", got, want)
	}

	if got, want := handler1.LastError(), err; got != want {
		t.Errorf("error does not match the expected one\nactual:   %s\nexpected: %s", got, want)
	}
}

func TestHandlers_NoHandlers(t *testing.T) {
	var handler Handlers

	err := errors.New("error")

	handler.Handle(err)
}

type closableHandler struct {
	handler    Handler
	closeError error
}

func (h *closableHandler) Handle(err error) {
	h.handler.Handle(err)
}

func (h *closableHandler) Close() error {
	return h.closeError
}

func TestHandlers_Close(t *testing.T) {
	t.Run("no_handlers", func(t *testing.T) {
		var handler Handlers

		err := handler.Close()

		if err != nil {
			t.Errorf("unexpected error when closing handlers\nactual:   %+v", err)
		}
	})

	t.Run("no_closers", func(t *testing.T) {
		handler1 := NewTestHandler()
		handler2 := NewTestHandler()

		handler := Handlers{handler1, handler2}

		err := handler.Close()

		if err != nil {
			t.Errorf("unexpected error when closing handlers\nactual:   %+v", err)
		}
	})

	t.Run("no_errors", func(t *testing.T) {
		handler1 := &closableHandler{handler: NewTestHandler()}
		handler2 := NewTestHandler()

		handler := Handlers{handler1, handler2}

		err := handler.Close()

		if err != nil {
			t.Errorf("unexpected error when closing handlers\nactual:   %+v", err)
		}
	})

	t.Run("error", func(t *testing.T) {
		closeErr := errors.New("close error")

		handler1 := &closableHandler{handler: NewTestHandler(), closeError: closeErr}
		handler2 := NewTestHandler()

		handler := Handlers{handler1, handler2}

		err := handler.Close()

		if err != closeErr {
			t.Errorf("unexpected error when closing handlers\nactual:   %#v\nexpected: %#v", err, closeErr)
		}
	})
}

func TestHandlerFunc(t *testing.T) {
	var actual error
	log := func(err error) {
		actual = err
	}

	fn := HandlerFunc(log)

	expected := errors.New("error")

	fn.Handle(expected)

	if got, want := actual, expected; got != want {
		t.Errorf("error does not match the expected one\nactual:   %s\nexpected: %s", got, want)
	}
}

func TestHandle(t *testing.T) {
	handler := NewTestHandler()
	err := errors.New("error")

	Handle(handler, err)

	if got, want := handler.LastError(), err; got != want {
		t.Errorf("error does not match the expected one\nactual:   %s\nexpected: %s", got, want)
	}
}

func TestHandle_Nil(t *testing.T) {
	handler := NewTestHandler()

	Handle(handler, nil)

	if got := handler.LastError(); got != nil {
		t.Errorf("unexpected error, received: %s", got)
	}
}

func TestMakeContextAware(t *testing.T) {
	testHandler := NewTestHandler()
	handler := MakeContextAware(testHandler)
	err := errors.New("error")

	handler.Handle(context.Background(), err)

	if got, want := testHandler.LastError(), err; got != want {
		t.Errorf("error does not match the expected one\nactual:   %s\nexpected: %s", got, want)
	}
}
