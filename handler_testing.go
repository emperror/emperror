package emperror

import (
	"context"
	"sync"
)

// TestErrorHandler is an ErrorHandler recording every error.
//
// Useful when you want to test behavior with an ErrorHandler, but not with ErrorHandlerContext.
// In every other cases TestErrorHandlerFacade should be the default choice of test handler.
//
// TestErrorHandler is safe for concurrent use.
type TestErrorHandler struct {
	errors []error

	mu sync.RWMutex
}

// Count returns the number of recorded events.
func (h *TestErrorHandler) Count() int {
	h.mu.RLock()
	defer h.mu.RUnlock()

	return len(h.errors)
}

// LastError returns the last handled error (if any).
func (h *TestErrorHandler) LastError() error {
	h.mu.RLock()
	defer h.mu.RUnlock()

	if len(h.errors) < 1 {
		return nil
	}

	return h.errors[len(h.errors)-1]
}

// Errors returns all handled errors.
func (h *TestErrorHandler) Errors() []error {
	h.mu.RLock()
	defer h.mu.RUnlock()

	return h.errors
}

// Handle records an error.
func (h *TestErrorHandler) Handle(err error) {
	if err == nil {
		return
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	h.errors = append(h.errors, err)
}

// TestErrorHandlerContext is an ErrorHandlerContext recording every error.
//
// Useful when you want to test behavior with an ErrorHandlerContext, but not with ErrorHandler.
// In every other cases TestErrorHandlerFacade should be the default choice of test handler.
//
// TestErrorHandlerContext is safe for concurrent use.
type TestErrorHandlerContext struct {
	errors   []error
	contexts []context.Context

	mu sync.RWMutex
}

// Count returns the number of recorded events.
func (h *TestErrorHandlerContext) Count() int {
	h.mu.RLock()
	defer h.mu.RUnlock()

	return len(h.errors)
}

// LastError returns the last handled error (if any).
func (h *TestErrorHandlerContext) LastError() error {
	h.mu.RLock()
	defer h.mu.RUnlock()

	if len(h.errors) < 1 {
		return nil
	}

	return h.errors[len(h.errors)-1]
}

// Errors returns all handled errors.
func (h *TestErrorHandlerContext) Errors() []error {
	h.mu.RLock()
	defer h.mu.RUnlock()

	return h.errors
}

// LastContext returns the context of the last handled error (if any).
func (h *TestErrorHandlerContext) LastContext() context.Context {
	h.mu.RLock()
	defer h.mu.RUnlock()

	if len(h.contexts) < 1 {
		return nil
	}

	return h.contexts[len(h.contexts)-1]
}

// Contexts returns contexts of all handled errors.
func (h *TestErrorHandlerContext) Contexts() []context.Context {
	h.mu.RLock()
	defer h.mu.RUnlock()

	return h.contexts
}

// HandleContext records an error.
func (h *TestErrorHandlerContext) HandleContext(ctx context.Context, err error) {
	if err == nil {
		return
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	h.errors = append(h.errors, err)
	h.contexts = append(h.contexts, ctx)
}

// TestErrorHandlerFacade is an ErrorHandlerFacade recording every error.
//
// TestErrorHandlerFacade is safe for concurrent use.
type TestErrorHandlerFacade struct {
	errors   []error
	contexts []context.Context

	mu sync.RWMutex
}

// TestErrorHandlerSet is an ErrorHandlerSet recording every error.
//
// TestErrorHandlerSet is safe for concurrent use.
//
// Deprecated: use TestErrorHandlerFacade.
type TestErrorHandlerSet = TestErrorHandlerFacade

// Count returns the number of recorded events.
func (h *TestErrorHandlerFacade) Count() int {
	h.mu.RLock()
	defer h.mu.RUnlock()

	return len(h.errors)
}

// LastError returns the last handled error (if any).
func (h *TestErrorHandlerFacade) LastError() error {
	h.mu.RLock()
	defer h.mu.RUnlock()

	if len(h.errors) < 1 {
		return nil
	}

	return h.errors[len(h.errors)-1]
}

// Errors returns all handled errors.
func (h *TestErrorHandlerFacade) Errors() []error {
	h.mu.RLock()
	defer h.mu.RUnlock()

	return h.errors
}

// LastContext returns the context of the last handled error (if any).
func (h *TestErrorHandlerFacade) LastContext() context.Context {
	h.mu.RLock()
	defer h.mu.RUnlock()

	if len(h.contexts) < 1 {
		return nil
	}

	return h.contexts[len(h.contexts)-1]
}

// Contexts returns contexts of all handled errors.
func (h *TestErrorHandlerFacade) Contexts() []context.Context {
	h.mu.RLock()
	defer h.mu.RUnlock()

	return h.contexts
}

// Handle records an error.
func (h *TestErrorHandlerFacade) Handle(err error) {
	if err == nil {
		return
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	h.errors = append(h.errors, err)
	h.contexts = append(h.contexts, nil)
}

// HandleContext records an error.
func (h *TestErrorHandlerFacade) HandleContext(ctx context.Context, err error) {
	if err == nil {
		return
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	h.errors = append(h.errors, err)
	h.contexts = append(h.contexts, ctx)
}

// TestHandler is a simple stub for the handler interface recording every error.
//
// The TestHandler is safe for concurrent use.
//
// Deprecated: use TestErrorHandler.
type TestHandler struct {
	errors []error

	mu sync.RWMutex
}

// NewTestHandler returns a new TestHandler.
func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

// Count returns the number of events recorded in the logger.
func (h *TestHandler) Count() int {
	h.mu.RLock()
	defer h.mu.RUnlock()

	return len(h.errors)
}

// LastError returns the last handled error (if any).
func (h *TestHandler) LastError() error {
	h.mu.RLock()
	defer h.mu.RUnlock()

	if len(h.errors) < 1 {
		return nil
	}

	return h.errors[len(h.errors)-1]
}

// Errors returns all handled errors.
func (h *TestHandler) Errors() []error {
	h.mu.RLock()
	defer h.mu.RUnlock()

	return h.errors
}

// Handle records the error.
func (h *TestHandler) Handle(err error) {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.errors = append(h.errors, err)
}
