# Change Log


All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).


## [Unreleased]


## [0.17.1] - 2019-04-01

### Fixed

- Nil pointer in `Recover`


## [0.17.0] - 2019-03-30

## Changed

- Attach stack trace to panicked and recovered errors and skip unnecessary frames
- Switch to Go modules
- Replace testify assertions with manual checks, drop testify dependency


## [0.16.0] - 2018-12-29

### Added

- [Rollbar](https://rollbar.com) handler
- [Sentry](https://sentry.io) handler
- Integration tests for handlers

### Changed

- Refactored and renamed nopHandler to noopHandler
- Refactored the test handler
- Refactored and relocated the airbrake handler
- Refactored and relocated the bugsnag handler
- Refactored and relocated the logrus handler
- Rename `HandleIfErr` to `Handle`
- Make handlers async by default

### Removed

- **bugsnag:** Logger (use [github.com/goph/logur](https://github.com/goph/logur) instead)
- Logrus hook
- Handler log func (unused)


## [0.15.0] - 2018-12-22

### Added

- `Panic` function to only panic if an error is not nil

### Changed

- TestHandler is now concurrent safe
- **bugsnag:** Completely rewritten bugsnag logger
- **bugsnag:** Unexport `NewErrorWithStackFrames`
- **bugsnag:** Export `handler`
- **airbrake:** Export `handler`

### Removed

- Handler mock


## [0.14.0] - 2018-12-11

### Removed

- **errorlog:** implementation moved to [github.com/goph/logur](https://github.com/goph/logur)


## [0.13.0] - 2018-12-07

### Changed

- Replaced go-kit errorlog with a custom interface


## [0.12.1] - 2018-12-07

### Added

- Return nil from `With` when error is nil


## [0.12.0] - 2018-09-24

### Added

- `WrapWith` function to wrap an error with message, stack trace and context at the same time
- Release scripts


## [0.11.0] - 2018-08-30

### Fixed

- **httperr:** Fix wrapped HTTP error formatting
- Fix stack expose wrapper error formatting
- Add Wrap and Wrapf functions


## [0.10.0] - 2018-08-21

### Added

- **errorlogrus:** Add `AttachContext` option to the Hook so that the entry data is appended to the error
- **errorlogrus:** Add an error handler logging with Logrus

### Changed

- **httperr:** Moved HTTP related code to separate package
- **bugsnag:** Improve logger
- **errorlog:** Renamed log package
- **errorlog:** Improved package
- **errorlogrus:** Rename logrus package
- **airbrake:** Improve Airbrake package


## [0.9.1] - 2018-07-27

### Added

- **bugsnag:** logger


## [0.9.0] - 2018-07-27

### Added

- **bugsnag:** support stack trace
- **bugsnag:** context and error name

### Changed

- **bugsnag:** notifier struct is not exported anymore
- **bugsnag:** `NewNotifierFromHandler` constructor to create a handler from a custom notifier instance


## [0.8.0] - 2018-06-24

### Added

- `ForEachCause` function to be able to loop through all errors in a chain
- `Context` function to get the context from an error (and all parent errors)
- `StackTrace` and `ExposeStackTrace` for working with stack trace
- logrus hook

### Changed

- **airbrake:** Use `ForEachCause` to find an HTTP request embedded into an error
- HTTP Request handling
- Refactor converting key-value pairs to maps
- Rename HTTP related function names to upper case (according to golint)

### Fixed

- **airbrake:** Make sure the stack trace is available from the topmost error

### Removed

- `Causer` interface
- `WithPrefix` didn't really make sense with the decorator pattern
- `Contextor` interface
- `StackTracer` interface
- `ErrorCollection` interface


## [0.7.1] - 2018-04-27

### Changed

- `ErrorCollection` errors are handled as separate lines in the log handler


## [0.7.0] - 2018-04-26

### Added

- `HandlerWith` and `HandlerWithPrefix` to attach context to a handler

### Changed

- Append nil instead of `ErrMissingValue` to the context when a value is missing


## [0.6.0] - 2017-10-26

### Added

- HttpError interface for representing errors with an HTTP error context

### Removed

- Aliases to functions in [github.com/pkg/errors](https://github.com/pkg/errors)


## [0.5.0] - 2017-08-30

### Added

- All error related code from [github.com/goph/stdlib](github.com/goph/stdlib)

### Changed

- Moved log handler to separate package
- Import subject package in tests to allow [using the exported identifiers without a qualifier](https://golang.org/ref/spec#Import_declarations)


## [0.4.0] - 2017-08-23

### Changed

- `compositeHandler` not exported anymore
- `nullHandler` not exported anymore

### Removed

- Handler interface (use the one in stdlib)
- Recovery (use the one in stdlib)


## [0.3.0] - 2017-07-11

### Added

- Contextual logging of errors
- Contextual error support to Airbrake handler

### Changed

- Make error level default in Log handler
- Do not export `LogHandler`


## [0.2.2] - 2017-07-07

### Added

- Testing handler wrapping test state


## [0.2.1] - 2017-06-24

### Changed

- Errors are passed as messages to loggers


## [0.2.0] - 2017-06-24

### Changed

- Use go-kit log interface


## [0.1.2] - 2017-06-22

### Changed

- `HandlerRecover` does not return a function anymore


## [0.1.1] - 2017-06-22

### Added

- `Recover` function (from [github.com/goph/stdlib](https://github.com/goph/stdlib))
- `HandlerRecover` to make recovering from a panic easier
- `HandleIfErr` to spare ifs in code if the only handling logic is passing to an error handler


## 0.1.0 - 2017-06-19

### Added

- `Handler` interface
- `NullHandler` serving as a fallback
- `LogHandler` to send errors to log collectors
- `TestHandler` to test code using error handlers
- `CompositeHandler` to handle errors in multiple handlers
- [Airbrake](https://airbrake.io) handler
- [Bugsnag](https://bugsnag.com) handler


[Unreleased]: https://github.com/goph/emperror/compare/v0.17.1...HEAD
[0.17.1]: https://github.com/goph/emperror/compare/v0.17.0...v0.17.1
[0.17.0]: https://github.com/goph/emperror/compare/v0.16.0...v0.17.0
[0.16.0]: https://github.com/goph/emperror/compare/v0.15.0...v0.16.0
[0.15.0]: https://github.com/goph/emperror/compare/v0.14.0...v0.15.0
[0.14.0]: https://github.com/goph/emperror/compare/v0.13.0...v0.14.0
[0.13.0]: https://github.com/goph/emperror/compare/v0.12.1...v0.13.0
[0.12.1]: https://github.com/goph/emperror/compare/v0.12.0...v0.12.1
[0.12.0]: https://github.com/goph/emperror/compare/v0.11.0...v0.12.0
[0.11.0]: https://github.com/goph/emperror/compare/v0.10.0...v0.11.0
[0.10.0]: https://github.com/goph/emperror/compare/v0.9.1...v0.10.0
[0.9.1]: https://github.com/goph/emperror/compare/v0.9.0...v0.9.1
[0.9.0]: https://github.com/goph/emperror/compare/v0.8.0...v0.9.0
[0.8.0]: https://github.com/goph/emperror/compare/v0.7.1...v0.8.0
[0.7.1]: https://github.com/goph/emperror/compare/v0.7.0...v0.7.1
[0.7.0]: https://github.com/goph/emperror/compare/v0.6.0...v0.7.0
[0.6.0]: https://github.com/goph/emperror/compare/v0.5.0...v0.6.0
[0.5.0]: https://github.com/goph/emperror/compare/v0.4.0...v0.5.0
[0.4.0]: https://github.com/goph/emperror/compare/v0.3.0...v0.4.0
[0.3.0]: https://github.com/goph/emperror/compare/v0.2.2...v0.3.0
[0.2.2]: https://github.com/goph/emperror/compare/v0.2.1...v0.2.2
[0.2.1]: https://github.com/goph/emperror/compare/v0.2.0...v0.2.1
[0.2.0]: https://github.com/goph/emperror/compare/v0.1.2...v0.2.0
[0.1.2]: https://github.com/goph/emperror/compare/v0.1.1...v0.1.2
[0.1.1]: https://github.com/goph/emperror/compare/v0.1.0...v0.1.1
