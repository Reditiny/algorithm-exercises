package main

import ds "algorithm-exercises/0_data_structure"

/*
	假设链表中每一个节点的值都在 0〜9 之间， 那么链表整体就可以代表一个整数。
	例如： 9->3->7, 可以代表整数 937。
	给定两个这种链表的头节点 head1 和 head2, 请生成代表两个整数相加值的结果链表。
	例如： 链表 1 为 9->3->7, 链表 2 为 6->3, 结果链表为 1->0->0->0
*/

func sumLink(head1, head2 *ds.LinkNode[int]) *ds.LinkNode[int] {
	// 翻转两个链表，便于从个位开始相加
	dummyHead1 := &ds.LinkNode[int]{}
	work := head1
	for work != nil {
		next := work.Next
		work.Next = dummyHead1.Next
		dummyHead1.Next = work
		work = next
	}
	dummyHead2 := &ds.LinkNode[int]{}
	work = head2
	for work != nil {
		next := work.Next
		work.Next = dummyHead2.Next
		dummyHead2.Next = work
		work = next
	}
	// 一次相加 记录进位
	dummyHead3 := &ds.LinkNode[int]{}
	work1 := dummyHead1.Next
	work2 := dummyHead2.Next
	carry := 0
	for work1 != nil || work2 != nil {
		curSum := carry
		if work1 != nil {
			curSum += work1.Val
			work1 = work1.Next
		}
		if work2 != nil {
			curSum += work2.Val
			work2 = work2.Next
		}
		carry = curSum / 10
		dummyHead3.Next = &ds.LinkNode[int]{Next: dummyHead3.Next, Val: curSum % 10}
	}
	if carry != 0 {
		dummyHead3.Next = &ds.LinkNode[int]{Next: dummyHead3.Next, Val: carry}
	}
	return dummyHead3.Next
}
