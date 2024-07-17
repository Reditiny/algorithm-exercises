package main

import ds "algorithm-exercises/0_data_structure"

/**
判断二叉树是否为平衡二叉树
如果二叉树的节点数为M 要求时间复杂度为 O(M)
*/

func isBalancedTree(root *ds.BTNode[int]) (balanced bool, height int) {
	if root == nil {
		return true, 0
	}
	leftBalanced, leftHeight := isBalancedTree(root.Left)
	if !leftBalanced {
		return false, 0
	}
	rightBalanced, rightHeight := isBalancedTree(root.Right)
	if !rightBalanced {
		return false, 0

	}
	if ds.Abs[int](leftHeight-rightHeight) > 1 {
		return false, 0
	}
	return true, ds.Max[int](leftHeight, rightHeight) + 1
}
