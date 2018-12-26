package logrushandler

import (
	"github.com/sirupsen/logrus"
)

func ExampleNew() {
	logger := logrus.New()
	_ = New(logger)

	// Output:
}
