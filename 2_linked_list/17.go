package main

import ds "algorithm-exercises/0_data_structure"

/*
	链表节点值类型为 int 型， 给定一个链表中的节点 node, 但不给定整个链表的头节点。
	如何在链表中删除 node? 请实现这个函数， 并分析这么会出现哪些问题。
	要求时间复杂度为O(1)
*/

func removeNodeWithoutHead(node *ds.LinkNode[int]) {
	if node == nil {
		return
	}
	if node.Next == nil {
		panic("cannot remove last node")
	}
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
