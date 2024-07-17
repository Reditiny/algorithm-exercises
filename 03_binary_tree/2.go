package main

import (
	ds "algorithm-exercises/0_data_structure"
	"fmt"
)

/*
	给定一棵二叉树的头节点 head, 按照如下标准实现二叉树边界节点的逆时针打印。
	1. 头节点为边界节点
	2. 叶节点为边界节点
	3. 如果节点在其所在的层中是最左或最右的，那么也是边界节点
*/

func printBorder(root *ds.BTNode[int], f func(*ds.BTNode[int])) {
	if root == nil {
		return
	}
	nodesArr := make([][]int, 0)
	queue := ds.NewQueue[*ds.BTNode[int]](10)
	queue.EnQueue(root)
	level := 0
	for !queue.Empty() {
		curSize := queue.Size()
		nodesArr[level] = make([]int, curSize)
		for i := 0; i < curSize; i++ {
			curNode := queue.Front()
			queue.DeQueue()
			nodesArr[level][i] = curNode.Val
			if curNode.Left != nil {
				queue.EnQueue(curNode.Left)
			}
			if curNode.Right != nil {
				queue.EnQueue(curNode.Right)
			}
		}
		level++
	}
	// 左边界
	for i := 0; i < len(nodesArr)-1; i++ {
		fmt.Println(nodesArr[i][0])
	}
	// 叶节点
	for j := 0; j < len(nodesArr[len(nodesArr)-1]); j++ {
		fmt.Println(nodesArr[len(nodesArr)-1][j])
	}
	// 右边界
	for i := len(nodesArr) - 2; i > 0; i-- {
		fmt.Println(nodesArr[i][len(nodesArr[i])-1])
	}
}

func levelOrder(root *ds.BTNode[int], f func(*ds.BTNode[int])) {
	if root == nil {
		return
	}
	queue := ds.NewQueue[*ds.BTNode[int]](10)
	queue.EnQueue(root)
	for !queue.Empty() {
		curSize := queue.Size()
		for i := 0; i < curSize; i++ {
			curNode := queue.Front()
			queue.DeQueue()
			f(curNode)
			if curNode.Left != nil {
				queue.EnQueue(curNode.Left)
			}
			if curNode.Right != nil {
				queue.EnQueue(curNode.Right)
			}
		}
	}
}
