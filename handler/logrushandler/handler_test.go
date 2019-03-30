package logrushandler_test

import (
	"github.com/sirupsen/logrus"

	"github.com/goph/emperror/handler/logrushandler"
)

func ExampleNew() {
	logger := logrus.New()
	_ = logrushandler.New(logger)

	// Output:
}
