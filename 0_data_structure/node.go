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

type BiLinkNode[T any] struct {
	Val  T
	Next *BiLinkNode[T]
	Pre  *BiLinkNode[T]
}

type LinkNodeWithRand[T any] struct {
	Val  T
	Next *LinkNodeWithRand[T]
	Rand *LinkNodeWithRand[T]
}

type HeapNode[T any] struct {
	Val   T
	Par   *HeapNode[T]
	Left  *HeapNode[T]
	Right *HeapNode[T]
}
