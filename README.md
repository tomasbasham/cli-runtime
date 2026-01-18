# cli-runtime

A Go module providing a framework for building command-line tools.

## Prerequisites

You will need the following things properly installed on your computer:

- [Go](https://golang.org/): any one of the **three latest major**
  [releases](https://golang.org/doc/devel/release.html)

## Installation

With [Go module](https://go.dev/wiki/Modules) support (Go 1.11+), simply add the
following import

```go
import "github.com/tomasbasham/cli-runtime"
```

to your code, and then `go [build|run|test]` will automatically fetch the
necessary dependencies.

Otherwise, to install the `cli-runtime` package, run the following command:

```bash
go get -u github.com/tomasbasham/cli-runtime
```

## License

This project is licensed under the [MIT License](LICENSE).
