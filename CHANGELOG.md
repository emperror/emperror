# Change Log


## 0.2.1 - 2017-07-24

### Changed

- Errors are passed as messages to loggers


## 0.2.0 - 2017-07-24

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
