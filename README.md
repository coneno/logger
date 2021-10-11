# Logger

Logger package implements logging methods with formatted output and more detailed logging information than provided in the standard logging package of golang. It is possible to specify the level of logging information to be written as output.


## Usage


```go
setLevel(level LogLevel)
```
sets the level of logging information written as output.

There are four possible values for argument `level` of type `LogLevel` in ascending info order:

> 1. `LEVEL_ERROR`: Only Error logs are written to `stderr`. 
> 2. `LEVEL_WARNING`: Warnings and errors are written to `stderr`.
> 3. `LEVEL_INFO`:    Errors and warnings are written to `stderr`, infos to `stdout`. 
> 4. `LEVEL_DEBUG`:   Errors and warnings are written to `stderr`, infos and debugs to `stdout`.

To print logging information you can choose one of the following `*log.Logger` variables chaining to common I/O writers:

```go
Error 
Warning
Info
Debug
```

The following Log flags are set:
> `log.Ldate`: the date in the local time zone \
> `log.Ltime`: the time in the local time zone \
> `log.Lshortfile`: final file name element and line number \
> `log.Lmsgprefix`: move the "prefix" log level tag from the beginning of the line to before the message


## Examples

```go
logger.SetLevel(logger.LEVEL_WARNING)
suffix := 2
logger.Error.Printf("error %d", suffix)
logger.Warning.Printf("warning  %d", suffix)
logger.Info.Printf("info  %d", suffix)
logger.Debug.Printf("debug  %d", suffix)
```
has the following output:

```go
2021/10/07 12:18:12 main.go:16: [ERROR]: error 2
2021/10/07 12:18:12 main.go:17: [WARNING]: warning  2
```
```go
logger.SetLevel(logger.LEVEL_DEBUG)
suffix := 4
logger.Error.Printf("error %d", suffix)
logger.Warning.Printf("warning  %d", suffix)
logger.Info.Printf("info  %d", suffix)
logger.Debug.Printf("debug  %d", suffix)
```

prints all logging information:

```go
2021/10/07 12:18:12 main.go:32: [ERROR]: error 4
2021/10/07 12:18:12 main.go:33: [WARNING]: warning  4
2021/10/07 12:18:12 main.go:34: [INFO]: info  4
2021/10/07 12:18:12 main.go:35: [DEBUG]: debug  4
```


## Tests

To perform the unit tests, you can call from the root folder:

```go
go test
```

To run the example:

```go
go run example/main.go
```
