package main

import ds "algorithm-exercises/0_data_structure"

/*
	设计一个有 getMin 功能的栈
	pop push top getMin 操作的时间复杂度都是 0(1)
	空则返回 -1
*/

type StackWithMin struct {
	mainStack ds.Stack[int] // 数据栈
	minStack  ds.Stack[int] // 辅助栈，高度与数据栈一致，保存当前最小值
}

func (s *StackWithMin) Push(x int) {
	s.mainStack.Push(x)
	// 入栈时辅助栈保存 min(x,minStack.Top)
	if s.minStack.Empty() {
		s.minStack.Push(x)
		return
	}
	s.minStack.Push(ds.Min(x, s.minStack.Top()))
}

func (s *StackWithMin) Pop() int {
	if s.mainStack.Empty() {
		return -1
	}
	ans := s.mainStack.Top()
	s.minStack.Pop()
	s.mainStack.Pop()
	return ans
}

func (s *StackWithMin) Top() int {
	if s.mainStack.Empty() {
		return -1
	}
	return s.mainStack.Top()
}

func (s *StackWithMin) GetMin() int {
	if s.minStack.Empty() {
		return -1
	}
	return s.minStack.Top()
}
