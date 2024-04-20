package main

import (
	ds "algorithm-exercises/0_data_structure"
)

/*
	给定一个单链表的头部节点 head, 链表长度为N,
	如果N为偶数， 那么前N/2个节点算作左半区后 N/2 个节点算作右半区；
	如果N为奇数，那么前 N/2 个节点算作左半区，后N/2+1 个节点算作右半区
	左半区从左到右依次记为L1->L2->...， 右半区从左到右依次记为R1->R2->...，
	请将单链表调整成 L1->R1->L2->R2->...的形式
*/

func mergeLR(head *ds.LinkNode[int]) *ds.LinkNode[int] {
	count := 0
	work := head
	for work != nil {
		count++
		work = work.Next
	}
	if count <= 3 {
		return head
	}

	work = head
	for i := 0; i < (count/2)-1; i++ {
		work = work.Next
	}
	head2 := work.Next
	work.Next = nil

	dummyHead := &ds.LinkNode[int]{}
	tail := dummyHead
	work1 := head
	work2 := head2
	for work1 != nil && work2 != nil {
		next1 := work1.Next
		next2 := work2.Next
		tail.Next = work1
		work1.Next = work2
		tail = work2
		work1 = next1
		work2 = next2
	}
	return dummyHead.Next
}
