package bytecount

import (
	"fmt"
	"testing"
)

func TestValue_Format(t *testing.T) {
	n := Value(1)
	got := fmt.Sprintf("%v:%3v:%.2v:%5.1v:%04.0v:%b:%#-5.1v:%04v:%-04v", n, n, n, n, Value(1024+511), n, n, -n, -n)
	want := "1B: 1B:1.00B: 1.0B:01KB:8b:1.0  :-01B:-1B "
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
