![Emperror](.github/logo.png?raw=true)

[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge-flat.svg)](https://github.com/avelino/awesome-go#error-handling)

[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/emperror/emperror/CI?style=flat-square)](https://github.com/emperror/emperror/actions?query=workflow%3ACI)
[![Codecov](https://img.shields.io/codecov/c/github/emperror/emperror?style=flat-square)](https://codecov.io/gh/emperror/emperror)
[![Go Report Card](https://goreportcard.com/badge/emperror.dev/emperror?style=flat-square)](https://goreportcard.com/report/emperror.dev/emperror)
![Go Version](https://img.shields.io/badge/go%20version-%3E=1.12-61CFDD.svg?style=flat-square)
[![PkgGoDev](https://pkg.go.dev/badge/mod/emperror.dev/emperror)](https://pkg.go.dev/mod/emperror.dev/emperror)
[![FOSSA Status](https://app.fossa.com/api/projects/custom%2B8125%2Femperror.dev%2Femperror.svg?type=shield)](https://app.fossa.com/projects/custom%2B8125%2Femperror.dev%2Femperror?ref=badge_shield)


**The Emperor takes care of all errors personally.**

Go's philosophy encourages to gracefully handle errors whenever possible,
but some times recovering from an error is not.

In those cases handling the error means making the best effort to record every detail
for later inspection, doing that as high in the application stack as possible.

This project provides tools to make error handling easier.

Read more about the topic here:

- https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully
- https://8thlight.com/blog/kyle-krull/2018/08/13/exploring-error-handling-patterns-in-go.html
- https://banzaicloud.com/blog/error-handling-go/


## Features

- Various error handling strategies (eg. logging, third-party error services) using a simple interface
- Various helpers related to error handling (recovery from panics, etc)
- [Integrations](https://github.com/emperror?utf8=%E2%9C%93&q=handler-*&type=&language=) with well-known error catchers and libraries:
    - [Logur](https://github.com/logur/logur)
    - [Logrus](https://github.com/sirupsen/logrus)
    - [Sentry](https://sentry.io) [SDK](https://godoc.org/github.com/getsentry/raven-go) (both hosted and on-premise)
    - [Bugsnag](https://bugsnag.com) [SDK](https://godoc.org/github.com/bugsnag/bugsnag-go)
    - [Airbrake](https://airbrake.com) [SDK](https://godoc.org/github.com/airbrake/gobrake) / [Errbit](https://errbit.com/)
    - [Rollbar](https://rollbar.com) [SDK](https://godoc.org/github.com/rollbar/rollbar-go)


## Installation

```bash
go get emperror.dev/emperror
```


## Usage

### Log errors

Logging is one of the most common target to record error events.

Emperror has two logger integrations by default:
- [Logur handler](https://github.com/emperror/handler-logur)
- [Logrus handler](https://github.com/emperror/handler-logrus)


### Annotate errors passing through an error handler

Emperror can annotate errors with details as defined in [emperror.dev/errors](https://github.com/emperror/errors)

```go
package main

import (
	"emperror.dev/emperror"
	"emperror.dev/errors"
)

func main() {
	handler := emperror.WithDetails(newHandler(), "key", "value")

	err := errors.New("error")

	// handled error will receive the handler details
	handler.Handle(err)
}
```

### Panics and recovers

```go
package main

import (
	"emperror.dev/emperror"
	"emperror.dev/errors"
)

func main() {
	var handler emperror.Handler = newHandler()

	// Recover from panics and handle them as errors
	defer emperror.HandleRecover(handler)

	// nil errors will not panic
	emperror.Panic(nil)

	// this will panic if foo returns with a non-nil error
	// useful in main func for initial setup where "if err != nil" does not make much sense
	emperror.Panic(foo())
}

func foo() error {
	return errors.New("error")
}
```

### Filter errors

Sometimes you might not want to handle certain errors that reach the error handler.
A common example is a catch-all error handler in a server. You want to return business errors to the client.

```go
package main

import (
	"emperror.dev/emperror"
	"emperror.dev/errors/match"
)

func main() {
	var handler emperror.Handler = emperror.WithFilter(newHandler(), match.Any{/*any emperror.ErrorMatcher*/})

    // errors matching the provided matcher will not be handled
	handler.Handle(err)
}
```


## Development

When all coding and testing is done, please run the test suite:

```bash
./pleasew lint
./pleasew gotest
```


## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.

[![FOSSA Status](https://app.fossa.com/api/projects/custom%2B8125%2Femperror.dev%2Femperror.svg?type=large)](https://app.fossa.com/projects/custom%2B8125%2Femperror.dev%2Femperror?ref=badge_large)
