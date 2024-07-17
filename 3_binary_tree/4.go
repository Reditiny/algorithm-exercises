package main

import ds "algorithm-exercises/0_data_structure"

/*
	Morris 遍历
	给定 一棵二叉树的头节点 head, 完成二叉树的先序、中序和后序遍历。
	时间复杂度为O(N), 空间复杂度为O(1)
*/

func morrisIn(head *ds.BTNode[int], f func(*ds.BTNode[int])) {
	if head == nil {
		return
	}
	var cur1, cur2 *ds.BTNode[int]
	cur1 = head
	for cur1 != nil {
		// 当前节点(cur1)左子树的最右节点(cur2)
		cur2 = cur1.Left
		if cur2 != nil {
			for cur2.Right != nil && cur2.Right != cur1 {
				cur2 = cur2.Right
			}
			if cur2.Right == nil {
				// cur2 的右节点为空说明是第一次遍历到
				// 构建回溯指针并向左子树迭代
				cur2.Right = cur1
				cur1 = cur1.Left
				continue
			} else {
				// 非空说明第二遍历到，前面的节点均已遍历过
				// 还原指针
				cur2.Right = nil
			}
		}
		f(cur1)
		cur1 = cur1.Right
	}
}

func morrisPre(head *ds.BTNode[int], f func(*ds.BTNode[int])) {
	if head == nil {
		return
	}
	var cur1, cur2 *ds.BTNode[int]
	cur1 = head
	for cur1 != nil {
		// 当前节点(cur1)左子树的最右节点(cur2)
		cur2 = cur1.Left
		if cur2 != nil {
			for cur2.Right != nil && cur2.Right != cur1 {
				cur2 = cur2.Right
			}
			if cur2.Right == nil {
				f(cur1)
				// cur2 的右节点为空说明是第一次遍历到
				// 构建回溯指针并向左子树迭代，先打印再迭代
				cur2.Right = cur1
				cur1 = cur1.Left
				continue
			} else {
				// 非空说明第二遍历到，前面的节点均以遍历过
				// 还原指针
				cur2.Right = nil
			}
		} else {
			f(cur1)
		}
		cur1 = cur1.Right
	}
}

func morrisPost(head *ds.BTNode[int], f func(*ds.BTNode[int])) {
	if head == nil {
		return
	}
	var cur1, cur2 *ds.BTNode[int]
	cur1 = head
	for cur1 != nil {
		cur2 = cur1.Left
		if cur2 != nil {
			for cur2.Right != nil && cur2.Right != cur1 {
				cur2 = cur2.Right
			}
			if cur2.Right == nil {
				cur2.Right = cur1
				cur1 = cur1.Left
				continue
			} else {
				cur2.Right = nil
				onEdges(cur1.Left, f)
			}
		}
		cur1 = cur1.Right
	}
	onEdges(head, f)
}

// 逆序处理右边界
func onEdges(root *ds.BTNode[int], f func(*ds.BTNode[int])) {
	tail := reverseEdge(root)
	work := tail
	for work != nil {
		f(work)
		work = work.Right
	}
	reverseEdge(tail)
}

func reverseEdge(root *ds.BTNode[int]) *ds.BTNode[int] {
	var pre, next *ds.BTNode[int]
	for root != nil {
		next = root.Right
		root.Right = pre
		pre = root
		root = next
	}
	return pre
}
