package bytecount

import "fmt"

type Value float32

func (v Value) Format(f fmt.State, verb rune) {
	wid := -1
	prec := -1
	if n, ok := f.Width(); ok {
		wid = n
	}
	if n, ok := f.Precision(); ok {
		prec = n
	}
	if verb == 'b' {
		v *= 8
	} else {
		verb = 'B'
	}
	_, _ = Write(f, float32(v), wid, prec, verb)
}
