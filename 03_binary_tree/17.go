package main

import ds "algorithm-exercises/0_data_structure"

/**
在二叉树中找到两个节点的最近公共祖先
*/

func findLowestCommonAncestor(root, p, q *ds.BTNode[int]) *ds.BTNode[int] {
	if root == nil || root == p || root == q {
		return root
	}
	l := findLowestCommonAncestor(root.Left, p, q)
	r := findLowestCommonAncestor(root.Right, p, q)
	if l != nil && r != nil {
		return root
	}
	if l != nil {
		return l
	}
	return r
}
