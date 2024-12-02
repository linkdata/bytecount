package bytecount_test

import (
	"fmt"
	"math"

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
	fmt.Printf("%v %v\n", math.MaxFloat32, bytecount.N(math.MaxFloat32))
	// Output:
	// 455 455B
	// 3192 3.12kB
	// 22351 21.8kB
	// 156464 153kB
	// 1095255 1.04MB
	// 7666792 7.31MB
	// 53667551 51.2MB
	// 375672864 358MB
	// 2629710055 2.45GB
	// 18407970392 17.1GB
	// 3.4028234663852886e+38 268435440QB
}
