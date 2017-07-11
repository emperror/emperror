package emperror

import "github.com/go-kit/kit/log/level"

// LogHandler accepts an logger instance and logs an error.
//
// Compatible with most level-based loggers.
type LogHandler struct {
	l logger
}

// logger covers most of the level-based logging solutions.
type logger interface {
	Log(keyvals ...interface{}) error
}

// NewLogHandler returns a new LogHandler.
func NewLogHandler(l logger) Handler {
	return &LogHandler{level.Error(l)}
}

// Handle takes care of an error by logging it.
func (h *LogHandler) Handle(err error) {
	h.l.Log(
		"msg", err.Error(),
	)
}
