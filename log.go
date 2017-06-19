package emperror

// LogHandler accepts an logger instance and logs an error.
//
// Compatible with most level-based loggers.
type LogHandler struct {
	l logger
}

// logger covers most of the level-based logging solutions.
type logger interface {
	Error(args ...interface{})
}

// NewLogHandler returns a new LogHandler.
func NewLogHandler(l logger) Handler {
	return &LogHandler{l}
}

// Handle takes care of an error by logging it.
func (h *LogHandler) Handle(err error) {
	h.l.Error(err)
}
