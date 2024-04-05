package ds

type Stack[T any] struct {
	top  int
	data []T
}

func NewStack[T any](size int) *Stack[T] {
	return &Stack[T]{
		top:  -1,
		data: make([]T, size),
	}
}

func (s *Stack[T]) Push(x T) {
	s.top++
	if s.top == len(s.data) {
		s.data = append(s.data, x)
		return
	}
	s.data[s.top] = x
}

func (s *Stack[T]) Pop() {
	if s.top >= 0 {
		s.top--
	}
}

func (s *Stack[T]) Top() T {
	if s.top == -1 {
		panic("empty stack")
	}
	return s.data[s.top]
}

func (s *Stack[T]) Empty() bool {
	return s.top == -1
}

func (s *Stack[T]) Size() int {
	return s.top + 1
}

func (s *Stack[T]) Clear() {
	s.data = make([]T, 0)
	s.top = -1
}
