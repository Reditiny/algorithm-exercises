package main

import (
	ds "algorithm-exercises/0_data_structure"
	"math"
)

/*
	给定一个无序单链表的头节点 head, 实现单链表的选择排序。
	额外空间复杂度为O(1)
*/

func selectSortForLinkList(head *ds.LinkNode[int]) *ds.LinkNode[int] {
	dummyHead := &ds.LinkNode[int]{Next: head, Val: math.MinInt}
	newHead := &ds.LinkNode[int]{}
	// 选择排序
	// 与数组不同的是要选到指定节点的前驱节点
	for true {
		curVal := math.MinInt
		cur := dummyHead
		pre := dummyHead
		work := dummyHead.Next
		for work != nil {
			if work.Val > curVal {
				cur = pre
				curVal = work.Val
			}
			pre = work
			work = work.Next
		}
		if cur.Next == nil {
			break
		}
		next := cur.Next
		cur.Next = cur.Next.Next
		next.Next = newHead.Next
		newHead.Next = next
	}
	return newHead.Next
}
