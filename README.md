[![build](https://github.com/linkdata/bytecount/actions/workflows/go.yml/badge.svg)](https://github.com/linkdata/bytecount/actions/workflows/go.yml)
[![coverage](https://github.com/linkdata/bytecount/blob/coverage/main/badge.svg)](https://html-preview.github.io/?url=https://github.com/linkdata/bytecount/blob/coverage/main/report.html)
[![Docs](https://godoc.org/github.com/linkdata/bytecount?status.svg)](https://godoc.org/github.com/linkdata/bytecount)

# bytecount

Go pretty printer for byte counts.

Only depends on the standard library.

## Usage

`go get github.com/linkdata/bytecount`

When printing byte counts using the `fmt` package, pass the values using `bytecount.N(n)`.

You may pass width and precision if you wish. The default is to keep the output to
a maximum of six characters while still showing as much precision as possible.

`bytecount.Value` stores values as `float32`, so the largest representable byte
count is `math.MaxFloat32` (~3.4&#x00D7;10&#x00B3;&#x2078;), which formats as `268435440QB`
(Quetta-bytes use the SI prefix `Q` = 10&#x00B3;&#x2070;). Anything larger overflows
the underlying `float32` and prints as `+InfQB`. Exact integer precision is also
limited to 24 bits (16,777,216); above that, neighboring integer byte counts may
format identically.

If the formatting verb is `d` (e.g. `"%d"`), the divisor is 1000 rather than 1024.

If the formatting verb is `b` (e.g. `"%b"`), the value is multiplied by 8 and the
unit suffix is changed from `B` (bytes) to `b` (bits). The divisor stays at 1024;
combining bit output with an SI divisor is not supported.

If the `"#"` flag is given, no unit suffix is written.

If the `" "` flag is given, a space is written between the digits and the suffix.

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
