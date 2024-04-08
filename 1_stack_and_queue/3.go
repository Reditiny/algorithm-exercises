package main

import ds "algorithm-exercises/0_data_structure"

/*
	仅用递归函数和栈操作逆序一个栈
	一个栈依次压入 1、2、3、4、5, 那么从栈顶到栈底分别为5、4、3、2、1。
	将这个栈转置后， 从栈顶到栈底为 1、2、3、4、5, 也就是实现栈中元素的逆序，
	但是只能用递归函数来实现， 不能用其他数据结构
*/

func ReverseStack(s ds.Stack[int]) {
	if s.Empty() {
		return
	}
	last := getAndRemoveLast(s)
	ReverseStack(s)
	s.Push(last)
}

// 递归的分解 n -> 1 + n-1
// 1、2、3、4、5 分解为 1 和 2、3、4、5
func getAndRemoveLast(s ds.Stack[int]) int {
	cur := s.Top()
	s.Pop()
	if s.Empty() {
		return cur
	} else {
		last := getAndRemoveLast(s)
		s.Push(cur)
		return last
	}
}
