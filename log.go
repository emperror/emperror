package emperror

// LogHandler accepts an errorLogger instance and logs an error.
//
// Compatible with most level-based loggers.
type LogHandler struct {
	logger errorLogger
}

// errorLogger covers most of the level-based logging solutions.
type errorLogger interface {
	Error(args ...interface{})
}

// NewLogHandler returns a new LogHandler.
func NewLogHandler(logger errorLogger) Handler {
	return &LogHandler{logger}
}

// Handle takes care of an error by logging it.
func (h *LogHandler) Handle(err error) {
	h.logger.Error(err)
}
