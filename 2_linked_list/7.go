package main

import (
	ds "algorithm-exercises/0_data_structure"
)

/*
	给定一个链表的头节点 head, 请判断该链表是否为回文结构。
	例如：
	1->2->1, 返回 true
	1->2->2->1, 返回 true
	15->6->15, 返回 true
	1->2->3, 返回 false
	进阶：如果链表长度为N, 时间复杂度达到 O(N)，额外空间复杂度达到 0(1)。
*/

func isPalindrome(head *ds.LinkNode[int]) bool {
	if head == nil || head.Next == nil {
		return true
	}
	dummyHead := &ds.LinkNode[int]{Next: head}
	// 找中点
	fast := dummyHead
	slow := dummyHead
	for fast != nil {
		fast = fast.Next
		if fast == nil {
			break
		}
		slow = slow.Next
		fast = fast.Next
	}
	// 翻转后一半
	dummyHead.Next = nil
	work := slow.Next
	slow.Next = nil
	for work != nil {
		next := work.Next
		work.Next = dummyHead.Next
		dummyHead.Next = work
		work = next
	}
	// 翻回来
	defer func() {
		work = dummyHead.Next
		for work != nil {
			next := work.Next
			work.Next = slow.Next
			slow.Next = work
			work = next
		}
	}()
	// 判断是否回文
	work1 := head
	work2 := dummyHead.Next
	for work1 != nil && work2 != nil {
		if work1.Val != work2.Val {
			return false
		}
		work1 = work1.Next
		work2 = work2.Next
	}
	return true
}
