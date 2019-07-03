![Emperror](/.github/logo.png?raw=true)

[![CircleCI](https://circleci.com/gh/emperror/emperror.svg?style=svg)](https://circleci.com/gh/emperror/emperror)
[![Go Report Card](https://goreportcard.com/badge/emperror.dev/emperror?style=flat-square)](https://goreportcard.com/report/emperror.dev/emperror)
[![GolangCI](https://golangci.com/badges/github.com/emperror/emperror.svg)](https://golangci.com/r/github.com/emperror/emperror)
[![GoDoc](http://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/emperror.dev/emperror)

**The Emperor takes care of all errors personally.**

Go's philosophy encourages to gracefully handle errors whenever possible,
but some times recovering from an error is not possible.

In those cases handling the error means making the best effort to record every detail
for later inspection, doing that as high in the application stack as possible.

This project provides tools (building on the well-known [pkg/errors](https://github.com/pkg/errors) package)
to make error handling easier.

Read more about the topic here:

- https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully
- https://8thlight.com/blog/kyle-krull/2018/08/13/exploring-error-handling-patterns-in-go.html
- https://banzaicloud.com/blog/error-handling-go/


## Features

- Various error handling strategies (eg. logging, third-party error services) using a simple interface
- Annotate errors with context (key-value pairs, HTTP request, etc)
- Various helpers related to error handling (recovery from panics, etc)
- Integrations with well-known error catchers and libraries:
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

The reference implementation for logging with Emperror can be found in package [logur](https://github.com/goph/logur).
Logur is an opinionated logging toolkit supporting multiple logging libraries (like [logrus](https://github.com/sirupsen/logrus)).


### Attach context to an error

Following [go-kit's logger](https://github.com/go-kit/kit/tree/master/log) context pattern
Emperror gives you tools to attach context (eg. key-value pairs) to an error:

```go
package main

import (
	"emperror.dev/emperror"
	"github.com/pkg/errors"
)

func foo() error { return errors.New("error") }

func bar() error {
	err := foo()
	if err != nil {
	    return emperror.With(err, "key", "value")
	}

	return nil
}
```

Note that (just like with go-kit's logger) the context is *NOT* a set of key-value pairs per se,
but most tools will convert the slice to key-value pairs.
This is to provide flexibility in error handling implementations.


## Development

When all coding and testing is done, please run the test suite:

``` bash
$ make check
```


## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.
