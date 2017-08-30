# Change Log


## Unreleased

### Added

- All error related code from [github.com/goph/stdlib](github.com/goph/stdlib)

### Changed

- Moved log handler to separate package
- Import subject package in tests to allow [using the exported identifiers without a qualifier](https://golang.org/ref/spec#Import_declarations)


## 0.4.0 - 2017-08-23

### Changed

- `compositeHandler` not exported anymore
- `nullHandler` not exported anymore

### Removed

- Handler interface (use the one in stdlib)
- Recovery (use the one in stdlib)


## 0.3.0 - 2017-07-11

### Added

- Contextual logging of errors
- Contextual error support to Airbrake handler

### Changed

- Make error level default in Log handler
- Do not export `LogHandler`


## 0.2.2 - 2017-07-07

### Added

- Testing handler wrapping test state


## 0.2.1 - 2017-06-24

### Changed

- Errors are passed as messages to loggers


## 0.2.0 - 2017-06-24

### Changed

- Use go-kit log interface


## 0.1.2 - 2017-06-22

### Changed

- `HandlerRecover` does not return a function anymore


## 0.1.1 - 2017-06-22

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
