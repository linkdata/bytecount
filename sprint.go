package bytecount

import "bytes"

func Sprint[T numeric](v T) string {
	var buf bytes.Buffer
	_, _ = Write(&buf, float32(v), -1, -1, 'B')
	return buf.String()
}
