package main

/**
在二叉树中找到一个节点的后继节点
type node struct {
	Val   int
	Left  *node
	Right *node
	Parent *node
}
只给一个在二叉树中的某个节点 node, 实现返回中序遍历的序列中 node 的后继节点
*/

type node struct {
	Val    int
	Left   *node
	Right  *node
	Parent *node
}

/**
             6
  		   /   \
		  3     9
		 / \   / \
        1   4 8  10
		 \       /
		  2     7
*/

func getSuccessorNode(node *node) *node {
	if node == nil {
		return nil
	}
	// 如果有右孩子，后继节点是右子树的最左节点
	if node.Right != nil {
		work := node.Right
		for work.Left != nil {
			work = work.Left
		}
		return work
	}
	// 没有右孩子，如果是父节点的左孩子，后继节点是父节点
	if node.Parent != nil && node.Parent.Left == node {
		return node.Parent
	}
	// 没有右孩子，如果是父节点的右孩子，向上遍历，直到找到一个节点是其父节点的左孩子
	work := node
	for work.Parent != nil && work.Parent.Right == work {
		work = work.Parent
	}
	return work.Parent
}
