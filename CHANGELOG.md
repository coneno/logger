# Changelog

## [v1.1.0] - 2021-10-05

### New:

- Added log level "WARNING" also directed towards stderr.
- Unit tests to check log levels.

### Changed:

- Log levels are: ERROR <  WARNING <  INFO < DEBUG. Log levels include all previous levels from left to right (WARNING includes ERROR, INFO includes WARNING and ERROR, etc.)

## [v1.0.0]

### New:

- Initial release to test integration and usage with other go packages.