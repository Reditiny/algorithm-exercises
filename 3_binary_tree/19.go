package main

import ds "algorithm-exercises/0_data_structure"

/**
从二叉树的节点 A 出发，可以向上或者向下走，但沿途的节点只能经过一次，当到达节点 B 时，路径上的节点数叫作 A 到 B 的距离
给定一棵二叉树的头节点 head, 求整棵树上节点间的最大距离。如果二叉树的节点数为N 时间复杂度要求为O(N)
*/

func maxDistance(head *ds.BTNode[int]) int {
	ans := 0

	var getMaxDeep func(root *ds.BTNode[int]) int

	getMaxDeep = func(root *ds.BTNode[int]) int {
		if root == nil {
			return 0
		}

		leftDeep := getMaxDeep(root.Left)
		RightDeep := getMaxDeep(root.Right)

		curDeep := ds.Max(leftDeep, RightDeep) + 1

		ans = ds.Max(ans, leftDeep+RightDeep+1)

		return curDeep
	}

	getMaxDeep(head)

	return ans
}
