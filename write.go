package bytecount

import (
	"fmt"
	"io"
	"math"
)

func Write(w io.Writer, v float64, wid, prec int, unit rune) (int, error) {
	defs := func(w, p int) {
		if wid < 0 {
			wid = w
		}
		if prec < 0 {
			prec = p
		}
	}
	var scale rune
	if v > 1000 {
		for _, scale = range "KMGTPEZY" {
			v /= 1024.0
			if math.Abs(v) < 10 {
				defs(1, 2)
				break
			}
			if math.Abs(v) < 100 {
				defs(2, 1)
				break
			}
			if math.Abs(v) < 1000 {
				defs(3, 0)
				break
			}
		}
	}
	defs(0, 0)
	var suffix string
	if scale != 0 {
		suffix = string(scale)
	}
	suffix += string(unit)
	if wid > 0 {
		wid -= len(suffix)
		if wid < 0 {
			wid = 0
		}
	}
	return fmt.Fprintf(w, "%*.*f%s", wid, prec, v, suffix)
}
