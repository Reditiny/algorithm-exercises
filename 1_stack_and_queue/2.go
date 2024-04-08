package main

import ds "algorithm-exercises/0_data_structure"

/*
编写一个类， 用两个栈实现队列， 支持队列的基本操作（push、pop、peek）
*/
type QueueByStack struct {
	MainStack *ds.Stack[int]
	HelpStack *ds.Stack[int]
}

func (q *QueueByStack) moveToHelp() {
	for !q.MainStack.Empty() {
		top := q.MainStack.Top()
		q.MainStack.Pop()
		q.HelpStack.Push(top)
	}
}

func (q *QueueByStack) moveToMain() {
	for !q.HelpStack.Empty() {
		top := q.HelpStack.Top()
		q.HelpStack.Pop()
		q.MainStack.Push(top)
	}
}

func (q *QueueByStack) Push(x int) {
	q.MainStack.Push(x)
}

func (q *QueueByStack) Pop() int {
	q.moveToHelp()
	x := q.HelpStack.Top()
	q.HelpStack.Pop()
	q.moveToMain()
	return x
}
func (q *QueueByStack) Peek() int {
	q.moveToHelp()
	x := q.HelpStack.Top()
	q.moveToMain()

	return x
}
