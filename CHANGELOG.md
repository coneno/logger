# Changelog

## [v1.2.1] - 2021-10-20

### Changed:

- Improved documentation
- Added license to the repository

## [v1.2.0] - 2021-10-05

### New:

- Example usage of the logger available [here](example/main.go).

### Changed:

- Change prefix order, so that log level tag comes directly before the message.


## [v1.1.0] - 2021-10-05

### New:

- Added log level "WARNING" also directed towards stderr.
- Unit tests to check log levels.

### Changed:

- Log levels are: ERROR <  WARNING <  INFO < DEBUG. Log levels include all previous levels from left to right (WARNING includes ERROR, INFO includes WARNING and ERROR, etc.)

## [v1.0.0]

### New:

- Initial release to test integration and usage with other go packages.