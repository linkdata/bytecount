package bytecount

import (
	"fmt"
	"io"
)

func Write(w io.Writer, v float32, wid, prec int, unit rune) (int, error) {
	defs := func(w, p int) {
		if wid < 0 {
			wid = w
		}
		if prec < 0 {
			prec = p
		}
	}
	var scale rune
	negative := v < 0
	if negative {
		v = -v
	}
	if v > 1000 {
		for _, scale = range "KMGTPEZY" {
			v /= 1024.0
			if v < 10 {
				defs(1, 2)
				break
			}
			if v < 100 {
				defs(2, 1)
				break
			}
			if v < 1000 {
				defs(3, 0)
				break
			}
		}
	}
	if negative {
		v = -v
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
