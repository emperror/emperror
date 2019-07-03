package logrushandler_test

import (
	"github.com/sirupsen/logrus"

	"emperror.dev/emperror/handler/logrushandler"
)

func ExampleNew() {
	logger := logrus.New()
	_ = logrushandler.New(logger)

	// Output:
}
