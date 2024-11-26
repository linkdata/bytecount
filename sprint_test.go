package bytecount

import (
	"fmt"
	"testing"
)

func TestSprint(t *testing.T) {
	tests := []struct {
		arg  float32
		want string
	}{
		{
			0,
			"0B",
		},
		{
			+1,
			"1B",
		},
		{
			-1,
			"-1B",
		},
		{
			1023,
			"1.00KB",
		},
		{
			1024 + 512,
			"1.50KB",
		},
		{
			10 * 1024,
			"10.0KB",
		},
		{
			100 * 1024,
			"100KB",
		},
		{
			999 * 1024,
			"999KB",
		},
		{
			1000 * 1024,
			"0.98MB",
		},
		{
			1010 * 1024,
			"0.99MB",
		},
		{
			1023 * 1024,
			"1.00MB",
		},
		{
			1024 * 1024,
			"1.00MB",
		},
		{
			1025 * 1024,
			"1.00MB",
		},
		{
			1024 * 1024 * 1024,
			"1.00GB",
		},
		{
			1024 * 1024 * 1024 * 1024,
			"1.00TB",
		},
		{
			1024 * 1024 * 1024 * 1024 * 1024,
			"1.00PB",
		},
		{
			1024 * 1024 * 1024 * 1024 * 1024 * 1024,
			"1.00EB",
		},
		{
			1024 * 1024 * 1024 * 1024 * 1024 * 1024 * 1024,
			"1.00ZB",
		},
		{
			1024 * 1024 * 1024 * 1024 * 1024 * 1024 * 1024 * 1024,
			"1.00YB",
		},
		{
			10 * 1024 * 1024 * 1024 * 1024 * 1024 * 1024 * 1024 * 1024,
			"10.0YB",
		},
		{
			1024 * 1024 * 1024 * 1024 * 1024 * 1024 * 1024 * 1024 * 1024,
			"1024YB",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%.0f", tt.arg), func(t *testing.T) {
			if got := Sprint(tt.arg); got != tt.want {
				t.Errorf("Sprint() = %v, want %v", got, tt.want)
			}
		})
	}
}
