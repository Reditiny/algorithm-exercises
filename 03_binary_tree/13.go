package main

import ds "algorithm-exercises/0_data_structure"

/**
根据后序数组重建搜索二叉树
*/

func buildTreeByPostOrder(postOrder []int) *ds.BTNode[int] {
	return build(postOrder, 0, len(postOrder)-1)
}

func build(postOrder []int, postStart, postEnd int) *ds.BTNode[int] {
	if postStart > postEnd {
		return nil
	}
	// 后序遍历的最后一个节点为根节点，左子树的节点都小于根节点，右子树的节点都大于根节点
	curRoot := postOrder[postEnd]
	nextStar := postStart
	for postOrder[nextStar] < curRoot {
		nextStar++
	}

	root := &ds.BTNode[int]{Val: curRoot}
	root.Left = build(postOrder, postStart, nextStar-1)
	root.Right = build(postOrder, nextStar, postEnd-1)
	return root
}
