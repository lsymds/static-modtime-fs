# staticmodtimefs

A Go package that wraps a filesystem and provides a constant modification time of any files and/or directories
contained within.

This package was created to allow Go's default `http.FileServerFS` handler to correctly apply `Last-Modified` headers
for embedded filesystem files.

## License

This package is licensed under the MIT license. More information is available here: [LICENSE](/LICENSE).

## Installation

`go get github.com/lsymds/staticmodtimefs`

## Usage

```go
import "github.com/lsymds/staticmodtimefs"

var myOtherFS fs.FS

// ... set it

wrappedFS := staticmodtimefs.NewStaticModTimeFS(myOtherFS, time.Now())

// ... use it however you wish
```
