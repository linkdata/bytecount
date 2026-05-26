package bytecount

import (
	"fmt"
	"math"
	"strconv"
)

// Value is a byte count expressed as a float32, formatted via [Value.Format]
// as a short human-readable string with an SI-style scale prefix and a unit
// suffix (`B` for bytes, `b` for bits). The zero value formats as "0B".
type Value float32

// String returns the default formatting of v, equivalent to fmt.Sprint(v).
func (v Value) String() string {
	return fmt.Sprint(v)
}

// Format implements [fmt.Formatter]. It supports the verbs `v`, `d`, and `b`:
//
//   - `v`: bytes with a 1024 divisor and `B` suffix.
//   - `d`: bytes with a 1000 (SI) divisor and `B` suffix.
//   - `b`: bits (value multiplied by 8) with a 1024 divisor and `b` suffix.
//
// The `#` flag omits the unit suffix. The ` ` (space) flag inserts a space
// between the digits and the scale/unit. The `+`, `-`, `0`, width, and
// precision flags behave as in the standard fmt package. Unsupported verbs
// produce the usual `%!verb(type=value)` error string.
func (v Value) Format(f fmt.State, verb rune) {
	var wid, prec int
	userPrec, hasUserPrec := f.Precision()
	var scale rune
	divisor := Value(1024.0)
	sign := byte(0)
	unit := byte('B')

	switch verb {
	case 'v':
		// default byte formatter
	case 'b':
		v *= 8
		unit = 'b'
	case 'd':
		divisor = 1000.0
	default:
		_, _ = fmt.Fprintf(f, "%%!%c(%T=%v)", verb, v, float32(v))
		return
	}

	if f.Flag('#') {
		unit = 0
	}

	negative := math.Signbit(float64(v))
	if negative {
		v = Value(math.Abs(float64(v)))
		sign = '-'
	} else if f.Flag('+') {
		sign = '+'
	}

	scaleAt := Value(1000)
	if !hasUserPrec {
		// With default precision (0), 999.5 rounds to 1000.
		scaleAt = 999.5
	}

	if v >= scaleAt {
		for _, scale = range "kMGTPEZYRQ" {
			v /= divisor
			// Thresholds are nudged below the next power of ten so that
			// values that would round up across a tier (e.g. 9.999 → "10.00")
			// fall into the next branch and keep the 6-character budget.
			if v < 9.995 {
				wid = 1
				prec = 2
				break
			}
			if v < 99.95 {
				wid = 2
				prec = 1
				break
			}
			if v < 999.5 {
				wid = 3
				prec = 0
				break
			}
		}
	}
	if n, ok := f.Width(); ok {
		wid = n
	}
	if hasUserPrec {
		prec = userPrec
	}

	var buf []byte
	buf = strconv.AppendFloat(buf, float64(v), 'f', prec, 32)
	if len(buf) > 0 && (buf[0] == '+' || buf[0] == '-') {
		if sign == 0 {
			sign = buf[0]
		}
		buf = buf[1:]
	}
	if f.Flag(' ') && (scale != 0 || unit != 0) {
		buf = append(buf, ' ')
	}
	if scale != 0 {
		buf = append(buf, byte(scale))
	}
	if unit != 0 {
		buf = append(buf, unit)
	}

	buflen := len(buf)
	if sign != 0 {
		buflen++
	}
	padchar := byte(' ')
	if f.Flag('0') && !f.Flag('-') {
		padchar = '0'
	}
	padlen := max(wid-buflen, 0)

	var out []byte
	if f.Flag('-') {
		if sign != 0 {
			out = append(out, sign)
		}
		out = append(out, buf...)
		for i := 0; i < padlen; i++ {
			out = append(out, ' ')
		}
	} else {
		if padchar == '0' && sign != 0 {
			out = append(out, sign)
		}
		for i := 0; i < padlen; i++ {
			out = append(out, padchar)
		}
		if padchar != '0' && sign != 0 {
			out = append(out, sign)
		}
		out = append(out, buf...)
	}

	_, _ = f.Write(out)
}
