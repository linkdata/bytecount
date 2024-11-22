[![build](https://github.com/linkdata/bytecount/actions/workflows/go.yml/badge.svg)](https://github.com/linkdata/bytecount/actions/workflows/go.yml)
[![coverage](https://coveralls.io/repos/github/linkdata/bytecount/badge.svg?branch=main)](https://coveralls.io/github/linkdata/bytecount?branch=main)
[![goreport](https://goreportcard.com/badge/github.com/linkdata/bytecount)](https://goreportcard.com/report/github.com/linkdata/bytecount)
[![Docs](https://godoc.org/github.com/linkdata/bytecount?status.svg)](https://godoc.org/github.com/linkdata/bytecount)

# bytecount

Go pretty printer for byte counts.

Only depends on the standard library.

## Usage

`go get github.com/linkdata/bytecount`

When printing byte counts using the `fmt` package, pass the values using `bytecount.N(n)`.

You may pass width and precision if you wish. The default is to keep the output to
a maximum of five characters while still showing as much precision as possible.

If the formatting verb is `b` (e.g. `"%b"`), the value is multiplied by 8 and the
suffix is changed from `B` (bytes) to `b` (bits).

## Example

```go
import (
	"fmt"

	"github.com/linkdata/bytecount"
)

func main() {
	fmt.Print(bytecount.N(53667551))
	// Output:
	// 51.2MB
}
```