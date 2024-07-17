package main

import (
	ds "algorithm-exercises/0_data_structure"
	"math"
)

/**
判断一棵二叉树是否为搜索二叉树和完全二叉树
*/

func isBST(root *ds.BTNode[int]) bool {
	if root == nil {
		return true
	}
	last := math.MinInt
	ans := true
	// morris 中序遍历
	morrisIn(root, func(node *ds.BTNode[int]) {
		if node.Val <= last {
			ans = false
		}
		last = node.Val
	})
	return ans
}

func isCBT(root *ds.BTNode[int]) bool {
	if root == nil {
		return true
	}
	queue := ds.NewQueue[*ds.BTNode[int]](10)
	queue.EnQueue(root)
	for !queue.Empty() {
		curNode := queue.Front()
		queue.DeQueue()
		// 有右孩子但是没有左孩子，不是完全二叉树
		if curNode.Left == nil && curNode.Right != nil {
			return false
		}
		if curNode.Left != nil {
			queue.EnQueue(curNode.Left)
		}
		if curNode.Right != nil {
			queue.EnQueue(curNode.Right)
		}
	}
	return true
}
