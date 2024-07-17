package main

import (
	ds "algorithm-exercises/0_data_structure"
	"math"
)

/*
	给定一棵二叉树的头节点 head, 已知其中所有节点的值都不一样，
	找到含有节点最多的搜索二叉子树， 并返回这棵子树的头节点。
	时间复杂度为O(N)，额外空间复杂度为O(h)， h 为二叉树的高度
*/

func biggestSubBST(head *ds.BTNode[int]) *ds.BTNode[int] {
	var ans *ds.BTNode[int]
	maxCount := math.MinInt

	var postOrder func(root *ds.BTNode[int]) (bool, int, int, int)
	postOrder = func(root *ds.BTNode[int]) (isBST bool, minVal, maxVal, nodeCount int) {
		if root == nil {
			// 空节点认为是 BST
			return true, 0, 0, 0
		}

		leftIsBST, leftMin, leftMax, leftCount := postOrder(root.Left)
		rightIsBST, rightMin, rightMax, rightCount := postOrder(root.Right)
		// 左右子树均是 BST 且 左子树最大值 < 当前节点值 < 右子树最小值
		// 当前树才是 BST
		if leftIsBST && rightIsBST && (leftCount == 0 || leftMax < root.Val) && (rightCount == 0 || root.Val < rightMin) {
			curMin := leftMin
			curMax := rightMax
			// 额外处理空节点返回的值
			if leftCount == 0 {
				curMin = root.Val
			}
			if rightCount == 0 {
				curMax = root.Val
			}
			// 当前 BST 确定是否更新答案
			curCount := leftCount + rightCount + 1
			if curCount > maxCount {
				maxCount = curCount
				ans = root
			}
			return true, curMin, curMax, curCount
		}
		return false, 0, 0, 0
	}

	postOrder(head)
	return ans
}
