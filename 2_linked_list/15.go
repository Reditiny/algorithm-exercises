package main

import ds "algorithm-exercises/0_data_structure"

/*
	现在有一棵搜索二叉树，请将其转换为一个有序的双向链表
*/

func convertToBiLinkList(root *ds.BTNode[int]) *ds.BTNode[int] {
	if root == nil {
		return nil
	}
	dummyHead := &ds.BTNode[int]{}
	stack := ds.NewStack[*ds.BTNode[int]](10)
	work := root
	last := dummyHead
	// 中序遍历的过程中构造链表
	for !stack.Empty() || work != nil {
		for work != nil {
			stack.Push(work)
			work = work.Left
		}
		cur := stack.Top()
		stack.Pop()
		// 每次构建前驱节点与当前节点的关系
		last.Right = cur
		cur.Left = last

		work = cur.Right
		last = cur
	}
	last.Right = dummyHead.Right
	dummyHead.Right.Left = last
	return dummyHead.Right
}
