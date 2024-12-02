package bytecount

import (
	"fmt"
	"math"
	"testing"
)

func TestValue_Format(t *testing.T) {
	overflow := N(math.MaxFloat64)
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
			name: "-1 B",
			f:    "%- 04v",
			v:    -1,
		},
		{
			name: "+InfQB",
			f:    "%v",
			v:    overflow,
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
