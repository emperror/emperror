package errorlogrus

import (
	"errors"

	"github.com/goph/emperror"
	"github.com/sirupsen/logrus"
)

// HookOption configures a logger instance.
type HookOption interface {
	apply(*hook)
}

// AttachContext tells to hook to attach the logrus context to the error.
type AttachContext bool

func (o AttachContext) apply(l *hook) {
	l.attachContext = bool(o)
}

// hook implements the logrus.Hook interface.
type hook struct {
	handler emperror.Handler

	attachContext bool
}

// NewHook returns a new logrus hook.
func NewHook(handler emperror.Handler, opts ...HookOption) logrus.Hook {
	h := &hook{handler: handler}

	for _, o := range opts {
		o.apply(h)
	}

	return h
}

func (h *hook) Fire(entry *logrus.Entry) error {
	err, ok := entry.Data["error"].(error)
	if !ok {
		err = errors.New(entry.Message)
	}

	if h.attachContext {
		var ctx []interface{}

		for key, value := range entry.Data {
			if key == "error" {
				continue
			}

			ctx = append(ctx, key, value)
		}

		err = emperror.With(err, ctx...)
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
