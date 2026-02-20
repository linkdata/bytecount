package bytecount

import (
	"fmt"
	"math"
	"testing"
)

func TestValue_Format(t *testing.T) {
	overflow := N(math.MaxFloat64)
	underflow := N(-math.MaxFloat64)
	negzero := N(math.Copysign(0, -1))
	tests := []struct {
		name string
		f    string
		v    Value
	}{
		{
			name: "1B",
			f:    "%v",
			v:    1,
		},
		{
			name: " 1B",
			f:    "%3v",
			v:    1,
		},
		{
			name: "1.00B",
			f:    "%.2v",
			v:    1,
		},
		{
			name: " 1.0B",
			f:    "%5.1v",
			v:    1,
		},
		{
			name: "01kB",
			f:    "%04.0v",
			v:    1024 + 511,
		},
		{
			name: "1.54 k",
			f:    "% #d",
			v:    1536,
		},
		{
			name: "1.00kB",
			f:    "%d",
			v:    1000,
		},
		{
			name: "8b",
			f:    "%b",
			v:    1,
		},
		{
			name: "1.0  ",
			f:    "%#-5.1v",
			v:    1,
		},
		{
			name: "1",
			f:    "% #v",
			v:    1,
		},
		{
			name: "-01B",
			f:    "%04v",
			v:    -1,
		},
		{
			name: "-1B ",
			f:    "%-04v",
			v:    -1,
		},
		{
			name: "     -1B",
			f:    "%8v",
			v:    -1,
		},
		{
			name: "-1 B",
			f:    "%- 04v",
			v:    -1,
		},
		{
			name: "+1B",
			f:    "%+v",
			v:    1,
		},
		{
			name: "     +1B",
			f:    "%+8v",
			v:    1,
		},
		{
			name: "+InfQB",
			f:    "%v",
			v:    overflow,
		},
		{
			name: "-InfQB",
			f:    "%v",
			v:    underflow,
		},
		{
			name: "-0B",
			f:    "%v",
			v:    negzero,
		},
		{
			name: "-000000B",
			f:    "%08v",
			v:    negzero,
		},
		{
			name: "%!x(bytecount.Value=31)",
			f:    "%x",
			v:    31,
		},
		{
			name: "%!q(bytecount.Value=31)",
			f:    "%q",
			v:    31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fmt.Sprintf(tt.f, tt.v)
			if got != tt.name {
				t.Errorf("want %q, got %q", tt.name, got)
			}
		})
	}
}

func TestValue_String(t *testing.T) {
	want := "1.95kB"
	got := N(2000).String()
	if got != want {
		t.Errorf("want %q got %q", want, got)
	}
}
