package main

import ds "algorithm-exercises/0_data_structure"

/*
	调整搜索二叉树中两个错误的节点
	一棵二叉树原本是搜索二叉树，但是其中有两个节点调换了位置，使得这棵二叉树不再是搜索二叉树，请找到这两个错误节点并返回
	进阶：在结构上完全交换两个节点的位置，而不是交换两个节点的值
*/

func getWrongNodes(head *ds.BTNode[int]) []*ds.BTNode[int] {
	ans := make([]*ds.BTNode[int], 2)
	if head == nil {
		return ans
	}
	var last *ds.BTNode[int]
	count := 0
	// 1 5 3 4 2	有两次降序
	// 1 2 4 3 5	只有一次
	inOrderIter(head, func(cur *ds.BTNode[int]) {
		if last != nil && last.Val > cur.Val {
			if count == 0 {
				// 第一次出现降序时同时记录两个节点
				ans[count] = last
				count++
				ans[count] = cur
			} else {
				// 如果有第二次降序，覆盖之前的第二个节点即可
				ans[count] = cur
			}
		}
		last = cur
	})
	return ans
}

/*
				r
				4
			2	    6
	      1   3   5   7
*/
func fixBST(head *ds.BTNode[int]) *ds.BTNode[int] {
	wrongNodes := getWrongNodes(head)
	// 根节点也可能被换，添加一个根节点的父节点，便于处理交换根节点
	rootParent := &ds.BTNode[int]{Left: head}
	parentNodes := make([]*ds.BTNode[int], 2)
	inOrderIter(rootParent, func(cur *ds.BTNode[int]) {
		for i := 0; i < 2; i++ {
			if cur.Left == wrongNodes[i] || cur.Right == wrongNodes[i] {
				parentNodes[i] = cur
			}
		}
	})
	// 调整子树
	wrongNodes[0].Left, wrongNodes[0].Right, wrongNodes[1].Left, wrongNodes[1].Right = wrongNodes[1].Left, wrongNodes[1].Right, wrongNodes[0].Left, wrongNodes[0].Right
	// 调整父节点
	if parentNodes[0] == wrongNodes[1] {
		// case1. 1,2 互换
		if parentNodes[1].Left == wrongNodes[1] {
			parentNodes[1].Left = wrongNodes[0]
		} else {
			parentNodes[1].Right = wrongNodes[0]
		}
		if parentNodes[0].Left == wrongNodes[0] {
			wrongNodes[0].Left = wrongNodes[1]
		} else {
			wrongNodes[0].Right = wrongNodes[1]
		}
	} else if parentNodes[1] == wrongNodes[0] {
		// case2. 2,3 互换
		if parentNodes[0].Left == wrongNodes[0] {
			parentNodes[0].Left = wrongNodes[1]
		} else {
			parentNodes[0].Right = wrongNodes[1]
		}
		if parentNodes[1].Left == wrongNodes[0] {
			wrongNodes[1].Left = wrongNodes[0]
		} else {
			wrongNodes[1].Right = wrongNodes[0]
		}
	} else {
		// case. 不相邻节点互换
		if parentNodes[0].Left == wrongNodes[0] {
			parentNodes[0].Left = wrongNodes[1]
		} else {
			parentNodes[0].Right = wrongNodes[1]
		}
		if parentNodes[1].Left == wrongNodes[1] {
			parentNodes[1].Left = wrongNodes[0]
		} else {
			parentNodes[1].Right = wrongNodes[0]
		}
	}
	return rootParent.Left
}
