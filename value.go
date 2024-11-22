package bytecount

import (
	"fmt"
	"strconv"
)

type Value float32

func (v Value) Format(f fmt.State, verb rune) {
	var wid, prec int
	var scale rune

	negative := v < 0
	if negative {
		v = -v
	}
	if v > 1000 {
		for _, scale = range "KMGTPEZY" {
			v /= 1024.0
			if v < 10 {
				wid = 1
				prec = 2
				break
			}
			if v < 100 {
				wid = 2
				prec = 1
				break
			}
			if v < 1000 {
				wid = 3
				prec = 0
				break
			}
		}
	}
	if n, ok := f.Width(); ok {
		wid = n
	}
	if n, ok := f.Precision(); ok {
		prec = n
	}

	unit := byte('B')
	if verb == 'b' {
		v *= 8
		unit = 'b'
	}
	if f.Flag('#') {
		unit = 0
	}

	var buf []byte
	buf = strconv.AppendFloat(buf, float64(v), 'f', prec, 32)
	if scale != 0 {
		buf = append(buf, byte(scale))
	}
	if unit != 0 {
		buf = append(buf, unit)
	}

	buflen := len(buf)
	if negative {
		buflen++
	}
	if buflen < wid {
		var pad []byte
		padchar := byte(' ')
		if f.Flag('0') && !f.Flag('-') {
			padchar = '0'
		}
		for i := buflen; i < wid; i++ {
			pad = append(pad, padchar)
		}
		if f.Flag('-') {
			buf = append(buf, pad...)
		} else {
			pad = append(pad, buf...)
			buf = pad
		}
	}
	if negative {
		buf = append([]byte{'-'}, buf...)
	}

	_, _ = f.Write(buf)
}
