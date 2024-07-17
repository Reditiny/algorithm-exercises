package main

import (
	ds "algorithm-exercises/0_data_structure"
)

/*
	分别用递归和非递归方式实现二叉树先序、中序和后序遍历
*/

func preOrderRecur[T any](root *ds.BTNode[T], f func(*ds.BTNode[T])) {
	if root == nil {
		return
	}
	f(root)
	preOrderRecur(root.Left, f)
	preOrderRecur(root.Right, f)
}

func inOrderRecur[T any](root *ds.BTNode[T], f func(*ds.BTNode[T])) {
	if root == nil {
		return
	}
	inOrderRecur(root.Left, f)
	f(root)
	inOrderRecur(root.Right, f)
}

func postOrderRecur[T any](root *ds.BTNode[T], f func(*ds.BTNode[T])) {
	if root == nil {
		return
	}
	postOrderRecur(root.Left, f)
	postOrderRecur(root.Right, f)
	f(root)
}

func preOrderIter[T any](root *ds.BTNode[T], f func(*ds.BTNode[T])) {
	if root == nil {
		return
	}
	stack := ds.NewStack[*ds.BTNode[T]](10)
	stack.Push(root)
	for !stack.Empty() {
		cur := stack.Top()
		stack.Pop()
		f(cur)
		if cur.Right != nil {
			stack.Push(cur.Right)
		}
		if cur.Left != nil {
			stack.Push(cur.Left)
		}
	}
}

func inOrderIter[T any](root *ds.BTNode[T], f func(*ds.BTNode[T])) {
	if root == nil {
		return
	}
	stack := ds.NewStack[*ds.BTNode[T]](10)
	work := root
	for !stack.Empty() || work != nil {
		for work != nil {
			stack.Push(work)
			work = work.Left
		}
		cur := stack.Top()
		stack.Pop()
		f(cur)
		work = cur.Right
	}
}

func postOrderIter[T any](root *ds.BTNode[T], f func(*ds.BTNode[T])) {
	if root == nil {
		return
	}
	if root == nil {
		return
	}
	stack := ds.NewStack[*ds.BTNode[T]](10)
	stack2 := ds.NewStack[*ds.BTNode[T]](10)
	stack.Push(root)
	for !stack.Empty() {
		cur := stack.Top()
		stack.Pop()
		stack2.Push(cur)
		if cur.Left != nil {
			stack.Push(cur.Left)
		}
		if cur.Right != nil {
			stack.Push(cur.Right)
		}
	}
	for !stack2.Empty() {
		f(stack2.Top())
		stack2.Pop()
	}
}
