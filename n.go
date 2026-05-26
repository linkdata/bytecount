package bytecount

// numeric is the set of numeric types accepted by [N] and [Sprint].
type numeric interface {
	~float32 | ~float64 |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// N converts any numeric value to a [Value] suitable for passing to the fmt
// package. Conversion goes through float32, so very large or very precise
// integers may lose precision.
func N[T numeric](n T) Value {
	return Value(n)
}
