package ds

type BTNode[T any] struct {
	Val   T
	Left  *BTNode[T]
	Right *BTNode[T]
}

type LinkNode[T any] struct {
	Val  T
	Next *LinkNode[T]
}
