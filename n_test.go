package bytecount

import "testing"

func TestN(t *testing.T) {
	if n := N(int(1)); n != Value(1) {
		t.Error(n)
	}
	if n := N(int16(1)); n != Value(1) {
		t.Error(n)
	}
	if n := N(int32(1)); n != Value(1) {
		t.Error(n)
	}
	if n := N(int64(1)); n != Value(1) {
		t.Error(n)
	}
	if n := N(uint(1)); n != Value(1) {
		t.Error(n)
	}
	if n := N(float32(1)); n != Value(1) {
		t.Error(n)
	}
	if n := N(float64(1)); n != Value(1) {
		t.Error(n)
	}
}
