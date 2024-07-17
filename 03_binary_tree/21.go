package main

/**
己知一棵二叉树所有的节点值都不同， 给定这棵树正确的先序和中序数组，
不要重建整棵树， 而是通过这两个数组直接生成正确的后序数组
*/

func buildPostArrByPreAndIn(pre, in []int) []int {
	if len(pre) == 0 {
		return nil
	}
	if len(pre) == 1 {
		return pre
	}
	mid := 0
	for i := 0; i < len(in); i++ {
		if in[i] == pre[0] {
			mid = i
			break
		}
	}
	left := buildPostArrByPreAndIn(pre[1:mid+1], in[:mid])
	right := buildPostArrByPreAndIn(pre[mid+1:], in[mid+1:])
	return append(append(left, right...), pre[0])
}
