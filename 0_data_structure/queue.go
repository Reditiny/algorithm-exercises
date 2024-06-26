package ds

// todo: 扩容逻辑有问题

type Queue[T any] struct {
	data  []T
	front int
	rear  int
}

func NewQueue[T any](size int) *Queue[T] {
	return &Queue[T]{
		data:  make([]T, size),
		front: 0,
		rear:  0,
	}
}

func (q *Queue[T]) EnQueue(x T) {
	if (q.rear+1)%len(q.data) == q.front {
		temp := make([]T, len(q.data)*2)
		for i := 0; i < len(q.data); i++ {
			temp[i] = q.data[(q.front+i)%len(q.data)]
		}
		q.data = temp
		q.front = 0
		q.rear = len(q.data)/2 - 1
	}
	q.data[q.rear] = x
	q.rear = (q.rear + 1) % len(q.data)
}

func (q *Queue[T]) DeQueue() {
	if q.front == q.rear {
	} else {
		q.front = (q.front + 1) % len(q.data)
	}
}

func (q *Queue[T]) Rear() T {
	if q.Empty() {
		panic("queue is empty")
	}
	return q.data[(q.rear-1+len(q.data))%len(q.data)]
}

func (q *Queue[T]) Front() T {
	if q.Empty() {
		panic("queue is empty")
	}
	return q.data[q.front]
}

func (q *Queue[T]) Empty() bool {
	return q.front == q.rear
}

func (q *Queue[T]) Size() int {
	return (q.rear - q.front + len(q.data)) % len(q.data)
}

func (q *Queue[T]) Clear() {
	q.data = make([]T, 0)
	q.front = 0
	q.rear = 0
}
