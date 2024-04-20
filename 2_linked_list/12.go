package main

import ds "algorithm-exercises/0_data_structure"

/*
	给定一个单链表的头节点 head 实现一个调整单链表的函数，
	使得每 K 个节点之间逆序， 如果最后不够 K 个节点一组， 则不调整最后几个节点。
	例如：链表 1->2->3->4->5->6->7->8->null , K=3
	调整为： 3->2->1->6->5->4->7->8->null 其中 7、8 不调整因为不够一组
	时间复杂度 O(N) 空间复杂度 O(1)
*/

func reverseEveryK(head *ds.LinkNode[int], k int) *ds.LinkNode[int] {
	if k < 2 {
		return head
	}
	work := head
	count := 0
	for work != nil {
		count++
		work = work.Next
	}
	dummyHead := &ds.LinkNode[int]{}
	curHead := dummyHead
	tail := dummyHead
	count = count / k
	work = curHead
	// 每次翻 k 个
	// 通过 curHead 与 tail 控制迭代
	for i := 0; i < count; i++ {
		var next *ds.LinkNode[int]
		for j := 0; j < k; j++ {
			next = work.Next
			work.Next = curHead.Next
			curHead.Next = work
			work = next
			if tail == curHead {
				tail = curHead.Next
			}
		}
		curHead = tail
	}
	curHead.Next = work
	return dummyHead.Next
}
