package main

import ds "algorithm-exercises/0_data_structure"

/*
	给定两个有序单链表的头节点 head1 和 head2, 请合并两个有序链表，
	合并后的链表依然有序， 并返回合并后链表的头节点
*/

func mergeLinkList(head1, head2 *ds.LinkNode[int]) *ds.LinkNode[int] {
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}
	dummyHead := &ds.LinkNode[int]{}
	tail := dummyHead

	work1 := head1
	work2 := head2
	for work1 != nil && work2 != nil {
		var cur *ds.LinkNode[int]
		if work1.Val < work2.Val {
			cur = work1
			work1 = work1.Next
		} else {
			cur = work2
			work2 = work2.Next
		}
		cur.Next = nil
		tail.Next = cur
		tail = cur
	}
	if work1 != nil {
		tail.Next = work1
	}
	if work2 != nil {
		tail.Next = work2
	}
	return dummyHead.Next
}
