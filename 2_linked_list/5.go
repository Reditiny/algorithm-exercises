package main

import ds "algorithm-exercises/0_data_structure"

/*
	给定一个单向链表的头节点 head, 以及两个整数 from 和 to, 在单向链表上把第 from
	个节点到第 to 个节点这一部分进行反转。
	例如：
	1->2->3->4->5->null , from=2 , to=4
	调整结果为： 1->4->3->2->5->null
	再如：
	1->2->3->null, from=1, to=3
	调整结果为： 3->2->1->null
	1. 如果链表长度为N 时间复杂度要求为 0(N)，额外空间复杂度要求为0(1)。
	2. 如果不满足 1<=from<=to<=N, 则不用调整。
*/

func reversePart(head *ds.LinkNode[int], from, to int) *ds.LinkNode[int] {
	dummyHead := &ds.LinkNode[int]{Next: head}
	work := dummyHead
	var start, end *ds.LinkNode[int]
	count := 1
	for work != nil {
		if count == from {
			start = work
		}
		if count == to+1 {
			end = work
		}
		count++
		work = work.Next
	}
	if start == nil || end == nil {
		return dummyHead.Next
	}
	work = start.Next
	start.Next = end.Next
	end.Next = nil
	for work != nil {
		next := work.Next
		work.Next = start.Next
		start.Next = work
		work = next
	}
	return dummyHead.Next
}
