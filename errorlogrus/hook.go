package errorlogrus

import (
	"errors"

	"github.com/goph/emperror"
	"github.com/sirupsen/logrus"
)

// hook implements the logrus.Hook interface.
type hook struct {
	handler emperror.Handler
}

// NewHook returns a new logrus hook.
func NewHook(handler emperror.Handler) logrus.Hook {
	return &hook{
		handler: handler,
	}
}

func (h *hook) Fire(entry *logrus.Entry) error {
	err, ok := entry.Data["error"].(error)
	if !ok {
		err = errors.New(entry.Message)
	}

	h.handler.Handle(err)

	return nil
}

func (*hook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.ErrorLevel,
		logrus.FatalLevel,
		logrus.PanicLevel,
	}
}
