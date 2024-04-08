package main

import ds "algorithm-exercises/0_data_structure"

/*
	用一个栈实现另一个栈的排序
	一个栈中元素的类型为整型， 现在想将该栈从顶到底按从大到小的顺序排序， 只许申
	请一个栈。除此之外， 可以申请新的变量， 但不能申请额外的数据结构。如何完成排序？
*/

func SortStack(stack *ds.Stack[int]) *ds.Stack[int] {
	ans := ds.NewStack[int](stack.Size())
	for !stack.Empty() {
		top := stack.Top()
		stack.Pop()
		insert(ans, top)
	}
	return ans
}

func insert(stack *ds.Stack[int], cur int) {
	if stack.Empty() || cur >= stack.Top() {
		stack.Push(cur)
		return
	}
	top := stack.Top()
	stack.Pop()
	insert(stack, cur)
	stack.Push(top)
}

func SortStack2(stack *ds.Stack[int]) *ds.Stack[int] {
	// 类似插入排序
	ans := ds.NewStack[int](stack.Size())
	for !stack.Empty() {
		top := stack.Top()
		stack.Pop()
		for !ans.Empty() && ans.Top() > top {
			temp := ans.Top()
			ans.Pop()
			stack.Push(temp)
		}
		ans.Push(top)
	}
	return ans
}
