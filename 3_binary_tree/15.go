package main

import ds "algorithm-exercises/0_data_structure"

/**
通过有序数组生成平衡搜索二叉树
*/

func buildBSTByArr(arr []int) *ds.BTNode[int] {
	n := len(arr)
	if n == 0 {
		return nil
	}
	node := &ds.BTNode[int]{Val: arr[n/2]}
	node.Left = buildBSTByArr(arr[:n/2])
	node.Right = buildBSTByArr(arr[n/2+1:])
	return node
}
