package bytecount

func N[T ~int | ~int32 | ~int64 |
	~uint | ~uint32 | ~uint64 |
	~float32 | ~float64](n T) Value {
	return Value(n)
}
