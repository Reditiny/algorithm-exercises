package main

import (
	ds "algorithm-exercises/0_data_structure"
)

/*
	构造数组的 MaxTree
	一个数组的 MaxTree 定义如下。
	• 数组必须没有重复元素。
	• MaxTree 是一棵二叉树， 数组的每一个值对应一个二叉树节点。
	• 包括 MaxTree 树在内且在其中的每一棵子树上， 值最大的节点都是树的头
	给定一个没有重复元素的数组 arr, 写出生成这个数组的 MaxTree 的函数
	要求如果数组长度为N, 则时间复杂度为 O(N)、额外空间复杂度为 O(N)。
*/

func MakeMaxTree(arr []int) *ds.BTNode {
	n := len(arr)
	nodes := make([]*ds.BTNode, n)
	for i := 0; i < n; i++ {
		nodes[i] = &ds.BTNode{Val: arr[i]}
	}

	// 单调栈确定每个值左右两边第一个比自己大的值
	stack := ds.NewStack[*ds.BTNode](n)
	leftFirstGreater := make(map[*ds.BTNode]*ds.BTNode)
	for i := 0; i < n; i++ {
		for !stack.Empty() && stack.Top().Val < nodes[i].Val {
			stack.Pop()
		}
		if stack.Empty() {
			leftFirstGreater[nodes[i]] = nil
		} else {
			leftFirstGreater[nodes[i]] = stack.Top()
		}
		stack.Push(nodes[i])
	}
	stack.Clear()
	rightFirstGreater := make(map[*ds.BTNode]*ds.BTNode)
	for i := n - 1; i >= 0; i-- {
		for !stack.Empty() && stack.Top().Val < nodes[i].Val {
			stack.Pop()
		}
		if stack.Empty() {
			rightFirstGreater[nodes[i]] = nil
		} else {
			rightFirstGreater[nodes[i]] = stack.Top()
		}
		stack.Push(nodes[i])
	}
	// 构造树
	// 每一个数的父节点是它左边第一个比它大的数和它右边第一个比它大的数中较小的那个
	// 如果某个数左右都没有比它大的数， 说明它是根节点
	var root *ds.BTNode
	for i := 0; i < n; i++ {
		cur := nodes[i]
		lfg := leftFirstGreater[cur]
		rfg := rightFirstGreater[cur]
		var parent *ds.BTNode
		if lfg == nil && rfg == nil {
			root = cur
			continue
		} else if lfg == nil {
			parent = rfg
		} else if rfg == nil {
			parent = lfg
		} else {
			if rfg.Val < lfg.Val {
				parent = rfg
			} else {
				parent = lfg
			}
		}
		if parent.Right == nil {
			parent.Right = cur
		} else {
			parent.Left = cur
		}
	}
	return root
}
