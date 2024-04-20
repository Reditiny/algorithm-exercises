package main

import ds "algorithm-exercises/0_data_structure"

/*
	分别实现反转单向链表和反转双向链表的函数
	如果链表长度为 N, 时间复杂度要求为O(N)，额外空间复杂度要求为O(1)
*/

func reverse1(head *ds.LinkNode[int]) *ds.LinkNode[int] {
	dummyHead := &ds.LinkNode[int]{}
	work := head
	for work != nil {
		next := work.Next
		work.Next = dummyHead.Next
		dummyHead.Next = work
		work = next
	}
	return dummyHead.Next
}

func reverse2(head *ds.BiLinkNode[int]) *ds.BiLinkNode[int] {
	dummyHead := &ds.BiLinkNode[int]{}
	work := head
	for work != nil {
		next := work.Next
		work.Next = dummyHead.Next
		if dummyHead.Next != nil {
			dummyHead.Next.Pre = work
		}
		dummyHead.Next = work
		work.Pre = dummyHead
		work = next
	}
	return dummyHead.Next
}
