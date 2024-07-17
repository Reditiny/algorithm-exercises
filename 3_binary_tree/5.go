package main

import (
	ds "algorithm-exercises/0_data_structure"
	"math"
)

/*
	在二叉树中找到累加和为指定值的最长路径长度
	给定一棵二叉树的头节点 head 和一个 32 位整数 sum, 二叉树节点值类型为整型求累加和为 sum 的最长路径长度。
	路径是指从某个节点往下， 每次最多选择一个孩子节点或者不选所形成的节点链
*/

func getMaxLenWithSum(head *ds.BTNode[int], sum int) int {
	maxLen := math.MinInt
	// 记录 sum 到 level 的映射
	sumToLevel := make(map[int]int)

	var preOrder func(*ds.BTNode[int], int, int)
	preOrder = func(root *ds.BTNode[int], preSum int, level int) {
		if root == nil {
			return
		}
		// 记录 curSum 第一次出现的 level
		curSum := preSum + root.Val
		if _, ok := sumToLevel[curSum]; !ok {
			sumToLevel[curSum] = level
		}
		// 如果 curSum-sum 出现过，两层之间的路径和为 sum
		if l, ok := sumToLevel[curSum-sum]; ok {
			maxLen = ds.Max[int](maxLen, level-l)
		}

		preOrder(root.Left, curSum, level+1)
		preOrder(root.Right, curSum, level+1)
		// 当本节点的子节点均递归完成后，本节点存入的信息就无效了
		// 不删除会影响后续的计算，导致两层之间的和为 sum 但是却无路径
		if level == sumToLevel[curSum] {
			delete(sumToLevel, curSum)
		}
	}

	preOrder(head, 0, 0)

	return maxLen
}
