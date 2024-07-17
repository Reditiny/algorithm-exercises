package main

import ds "algorithm-exercises/0_data_structure"

/**
给定一棵完全二叉树的头节点 head, 返回这棵树的节点个数
如果完全二叉树的节点数为 N, 请实现时间复杂度低于 O(N)的解法
*/

func countNodes(head *ds.BTNode[int]) int {
	rightDeep := maxRightDeep(head)
	leftDeep := maxLeftDeep(head)
	if rightDeep == leftDeep {
		return 1<<rightDeep - 1
	} else {
		return 1 + countNodes(head.Left) + countNodes(head.Right)
	}
}

func maxRightDeep(head *ds.BTNode[int]) int {
	deep := 0
	for head != nil {
		deep++
		head = head.Right
	}
	return deep
}

func maxLeftDeep(head *ds.BTNode[int]) int {
	deep := 0
	for head != nil {
		deep++
		head = head.Left
	}
	return deep
}
