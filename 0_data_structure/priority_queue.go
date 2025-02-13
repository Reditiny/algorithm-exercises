package ds

type PriorityQueue[T any] struct {
	data    []T
	compare func(a, b T) int
	size    int
}

func NewPriorityQueue[T any](compare func(a, b T) int) *PriorityQueue[T] {
	return &PriorityQueue[T]{
		data:    make([]T, 1),
		compare: compare,
		size:    0,
	}
}

func (pq *PriorityQueue[T]) Push(x T) {
	if pq.size == len(pq.data) {
		temp := make([]T, len(pq.data)*2)
		copy(temp, pq.data)
		pq.data = temp
	}
	pq.data[pq.size] = x
	pq.up(pq.size)
	pq.size++
}

func (pq *PriorityQueue[T]) Top() T {
	if pq.size == 0 {
		panic("empty priority queue")
	}
	return pq.data[0]
}

func (pq *PriorityQueue[T]) Pop() {
	if pq.size != 0 {
		pq.size--
		temp := pq.data[pq.size]
		pq.data[0] = temp
		pq.down(0)
	}
}

func (pq *PriorityQueue[T]) Size() int {
	return pq.size
}

func (pq *PriorityQueue[T]) Empty() bool {
	return pq.size == 0
}

func (pq *PriorityQueue[T]) Clear() {
	pq.data = make([]T, 0)
	pq.size = 0
}

func (pq *PriorityQueue[T]) down(index int) {
	left := index*2 + 1
	right := index*2 + 2
	if left >= pq.size {
		return
	} else if right >= pq.size {
		if pq.compare(pq.data[index], pq.data[left]) < 0 {
			pq.data[index], pq.data[left] = pq.data[left], pq.data[index]
			pq.down(left)
		}
	} else {
		if pq.compare(pq.data[left], pq.data[right]) > 0 {
			if pq.compare(pq.data[index], pq.data[left]) < 0 {
				pq.data[index], pq.data[left] = pq.data[left], pq.data[index]
				pq.down(left)
			}
		} else {
			if pq.compare(pq.data[index], pq.data[right]) < 0 {
				pq.data[index], pq.data[right] = pq.data[right], pq.data[index]
				pq.down(right)
			}
		}
	}
}

func (pq *PriorityQueue[T]) up(index int) {
	if index == 0 {
		return
	}
	parent := (index - 1) / 2
	if pq.compare(pq.data[parent], pq.data[index]) < 0 {
		pq.data[parent], pq.data[index] = pq.data[index], pq.data[parent]
		pq.up(parent)
	}
}

func (pq *PriorityQueue[T]) Range(f func(any T)) {
	for i := 0; i < pq.size; i++ {
		f(pq.data[i])
	}
}
