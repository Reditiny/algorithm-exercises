package main

import ds "algorithm-exercises/0_data_structure"

/*
	给定一棵二叉树的头节点 head, 已知所有节点的值都不一样
	返回其中最大的且符合搜索二叉树条件的最大拓扑结构的大小
*/

func biggestTopoSize(head *ds.BTNode[int]) int {
	if head == nil {
		return 0
	}

	max := maxTopo(head, head)
	max = ds.Max[int](max, biggestTopoSize(head.Right))
	max = ds.Max[int](max, biggestTopoSize(head.Left))

	return max
}

// 以 node 为根的树中有多少节点是以 head 为根节点的 BST 中的节点
func maxTopo(head, node *ds.BTNode[int]) int {
	if head != nil && node != nil && isBSTNode(head, node) {
		return maxTopo(head, node.Left) + maxTopo(head, node.Right) + 1
	}
	return 0
}

// 判断 node 是否是以 head 为根节点的 BST 中的一个节点
func isBSTNode(head, node *ds.BTNode[int]) bool {
	if head == nil {
		return false
	}
	if head == node {
		return true
	}
	// 往 head 的子树中去寻找
	var nextNode *ds.BTNode[int]
	if head.Val > node.Val {
		nextNode = head.Left
	} else {
		nextNode = head.Right
	}
	return isBSTNode(nextNode, node)
}
