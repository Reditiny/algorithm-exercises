package main

import (
	ds "algorithm-exercises/0_data_structure"
)

/*
	给定一个链表的头节点 head 和一个整数 num, 请实现函数将值为 num 的节点全部删除
*/

func removeNum(head *ds.LinkNode[int], num int) *ds.LinkNode[int] {
	dummyHead := &ds.LinkNode[int]{Next: head}
	pre := dummyHead
	cur := head
	for cur != nil {
		next := cur.Next
		if cur.Val == num {
			pre.Next = next
		} else {
			pre = cur
		}
		cur = next
	}
	return dummyHead.Next
}
