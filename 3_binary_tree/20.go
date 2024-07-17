package main

import (
	ds "algorithm-exercises/0_data_structure"
)

/**
已知一棵二叉树的所有节点值都不同 给定这棵二叉树正确的先序、中序和后序数组。
请分别用三个函数实现任意两种数组结合重构原来的二叉树，并返回重构二叉树的头节点
*/

func buildTreeByPreAndIn(pre, in []int) *ds.BTNode[int] {
	if len(pre) == 0 {
		return nil
	}
	head := &ds.BTNode[int]{Val: pre[0]}
	mid := 0
	for i := 0; i < len(in); i++ {
		if in[i] == pre[0] {
			mid = i
			break
		}
	}
	head.Left = buildTreeByPreAndIn(pre[1:mid+1], in[:mid])
	head.Right = buildTreeByPreAndIn(pre[mid+1:], in[mid+1:])
	return head
}

func buildTreeByInAndPost(in, post []int) *ds.BTNode[int] {
	n := len(post)
	if n == 0 {
		return nil
	}
	head := &ds.BTNode[int]{Val: post[n-1]}
	mid := 0
	for i := 0; i < len(in); i++ {
		if in[i] == post[n-1] {
			mid = i
			break
		}
	}
	head.Left = buildTreeByInAndPost(in[:mid], post[:mid])
	head.Right = buildTreeByInAndPost(in[mid+1:], post[mid:n-1])
	return head
}

// 除叶节点外其他节点均有两个孩子的树才能通过先序和后序重构
func buildTreeByPreAndPost(pre, post []int) *ds.BTNode[int] {
	if len(pre) == 0 || len(post) == 0 {
		return nil
	}

	head := &ds.BTNode[int]{Val: pre[0]}
	if len(pre) == 1 {
		return head
	}

	mid := 1
	for mid < len(pre) && pre[mid] != post[len(post)-2] {
		mid++
	}

	head.Left = buildTreeByPreAndPost(pre[1:mid], post[:mid-1])
	head.Right = buildTreeByPreAndPost(pre[mid:], post[mid-1:len(post)-1])

	return head
}
