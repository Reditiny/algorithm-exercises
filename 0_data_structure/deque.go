package ds

type Deque[T any] struct {
	data       map[int][]T
	size       int
	front      int
	rear       int
	frontIndex int
	rearIndex  int
	bufSize    int
}

func NewDeque[T any](bufSize int) *Deque[T] {
	d := &Deque[T]{
		data:       make(map[int][]T),
		size:       0,
		front:      0,
		rear:       0,
		frontIndex: 0,
		rearIndex:  0,
		bufSize:    bufSize,
	}
	d.data[0] = make([]T, d.bufSize)
	return d
}

func (d *Deque[T]) PushFront(x T) {
	if d.front == 0 {
		d.frontIndex--
		d.data[d.frontIndex] = make([]T, d.bufSize)
		d.front = d.bufSize
	}
	d.front--
	d.data[d.frontIndex][d.front] = x
	d.size++
}

func (d *Deque[T]) PushBack(x T) {
	if d.rear == d.bufSize {
		d.rearIndex++
		d.data[d.rearIndex] = make([]T, d.bufSize)
		d.rear = 0
	}
	d.data[d.rearIndex][d.rear] = x
	d.rear++
	d.size++
}

func (d *Deque[T]) PopFront() {
	if d.size > 0 {
		d.front++
		if d.front == d.bufSize {
			d.frontIndex++
			d.front = 0
			delete(d.data, d.frontIndex-1)
		}
		d.size--
	}
}

func (d *Deque[T]) PopBack() {
	if d.size > 0 {
		d.rear--
		if d.rear == -1 {
			d.rearIndex--
			d.rear = d.bufSize - 1
			delete(d.data, d.rearIndex+1)
		}
		d.size--
	}
}

func (d *Deque[T]) Front() T {
	if d.size == 0 {
		panic("deque is empty")
	}
	return d.data[d.frontIndex][d.front]
}

func (d *Deque[T]) Back() T {
	if d.size == 0 {
		panic("deque is empty")
	}
	return d.data[d.rearIndex][d.rear-1]
}

func (d *Deque[T]) Empty() bool {
	return d.size == 0
}

func (d *Deque[T]) Size() int {
	return d.size
}

func (d *Deque[T]) Clear() {
	d.data = make(map[int][]T)
	d.size = 0
	d.front = 0
	d.rear = 0
	d.frontIndex = 0
	d.rearIndex = 0
	d.data[0] = make([]T, d.bufSize)
}
