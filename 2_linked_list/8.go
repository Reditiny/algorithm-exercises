package main

import ds "algorithm-exercises/0_data_structure"

/*
	将单向链表按某值划分成左边小、中间相等、右边大的形式
	给定一个单向链表的头节点 head, 节点的值类型是整型， 再给定一个整数 pivot。
	实现一个调整链表的函数， 将链表调整为左部分小于 pivot，右部分大于 pivot 的节点
	例如: 链表 9->0->4->5->1 , pivot=3
	调整后链表可以是 1->0->4->9->5, 也可以是 0->1->9->5->4
	进阶: 原问题的要求之上再增加如下两个要求
		• 要求每部分里的节点从左到右的顺序与原链表中节点的先后次序一致
		• 如果链表长度为 N, 时间复杂度请达到 O(N), 额外空间复杂度请达到 O(1)
*/

func listPartition(head *ds.LinkNode[int], pivot int) *ds.LinkNode[int] {
	smallHead := &ds.LinkNode[int]{}
	smallTail := smallHead
	equalHead := &ds.LinkNode[int]{}
	equalTail := equalHead
	bigHead := &ds.LinkNode[int]{}
	bigTail := equalHead
	work := head
	for work != nil {
		next := work.Next
		work.Next = nil
		if work.Val < pivot {
			smallTail.Next = work
			smallTail = work
		} else if work.Val > pivot {
			bigTail.Next = work
			bigTail = work
		} else {
			equalTail.Next = work
			equalTail = work
		}
		work = next
	}
	smallTail.Next = equalHead.Next
	equalTail.Next = bigHead.Next
	return smallHead.Next
}
