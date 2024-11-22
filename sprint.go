package bytecount

import (
	"fmt"
)

func Sprint[T numeric](v T) string {
	return fmt.Sprint(N(v))
}
