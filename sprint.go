package bytecount

import "bytes"

func Sprint(v float64) string {
	var buf bytes.Buffer
	_, _ = Write(&buf, v, -1, -1, 'B')
	return buf.String()
}
