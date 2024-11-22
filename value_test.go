package bytecount

import (
	"fmt"
	"testing"
)

func TestValue_Format(t *testing.T) {
	n := Value(1)
	got := fmt.Sprintf("%v:%3v:%.2v:%5.1v:%0.0v:%b", n, n, n, n, Value(1024+511), n)
	want := "1B: 1B:1.00B: 1.0B:1KB:8b"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
