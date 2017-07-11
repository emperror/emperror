package emperror

import "github.com/go-kit/kit/log/level"

// logHandler accepts an logger instance and logs an error.
//
// Compatible with most level-based loggers.
type logHandler struct {
	l logger
}

// logger covers most of the level-based logging solutions.
type logger interface {
	Log(keyvals ...interface{}) error
}

// NewLogHandler returns a new logHandler.
func NewLogHandler(l logger) Handler {
	return &logHandler{level.Error(l)}
}

// Handle takes care of an error by logging it.
func (h *logHandler) Handle(err error) {
	h.l.Log(
		"msg", err.Error(),
	)
}
