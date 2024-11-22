package bytecount_test

import (
	"fmt"

	"github.com/linkdata/bytecount"
)

func ExampleN() {
	var n int64
	n = 64
	for i := 0; i < 10; i++ {
		n++
		n *= 7
		fmt.Printf("%v %v\n", n, bytecount.N(n))
	}
	// Output:
	// 455 455B
	// 3192 3.12KB
	// 22351 21.8KB
	// 156464 153KB
	// 1095255 1.04MB
	// 7666792 7.31MB
	// 53667551 51.2MB
	// 375672864 358MB
	// 2629710055 2.45GB
	// 18407970392 17.1GB
}
