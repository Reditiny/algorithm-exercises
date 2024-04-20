package main

import ds "algorithm-exercises/0_data_structure"

/*
	在单链表和双链表中删除倒数第 K 个节点
	分别实现两个函数， 一个可以删除单链表中倒数第 K 个节点，
	另一个可以删除双链表中倒数第 K 个节点
*/

func removeReciprocal(head *ds.LinkNode[int], k int) {
	if k <= 0 {
		return
	}
	// 单链表需要找到待删除节点的前驱节点
	dummyHead := &ds.LinkNode[int]{Next: head}
	fast := dummyHead
	for k >= 0 {
		if fast == nil {
			return
		}
		fast = fast.Next
		k--
	}
	slow := dummyHead
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
}

func removeReciprocal2(head *ds.BiLinkNode[int], k int) {
	if k <= 0 {
		return
	}
	// 双链表需要找到待删除节点（或找其前驱或后继都可以）
	fast := head
	for k > 0 {
		if fast == nil {
			return
		}
		fast = fast.Next
		k--
	}
	slow := head
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	if slow.Pre != nil {
		slow.Pre.Next = slow.Next
	}
	if slow.Next != nil {
		slow.Next.Pre = slow.Pre
	}
}
