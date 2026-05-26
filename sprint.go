package bytecount

import (
	"fmt"
)

// Sprint returns the default human-readable formatting of v, equivalent to
// fmt.Sprint(N(v)).
func Sprint[T numeric](v T) string {
	return fmt.Sprint(N(v))
}
