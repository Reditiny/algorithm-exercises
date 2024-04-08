package ds

type number interface {
	int64 | int32 | int16 | int8 |
		uint64 | uint32 | uint16 |
		uint8 | float64 | float32 |
		int
}

func Max[T number](a, b T) T {
	if a < b {
		return b
	}
	return a
}

func Min[T number](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Abs[T number](a T) T {
	if a < 0 {
		return -a
	}
	return a
}
