package main

import (
	ds "algorithm-exercises/0_data_structure"
)

/*
	给定两个单链表的头节点 head1 和 head2，可能有环，也可能无环
	两个链表可能相交，也可能不相交。
	请实现一个函数，如果两个链表相交，请返回相交的第一个节点；如果不相交，返回 null 即可
	要求：如果链表长度分别为 N 和 M，时间复杂度请 O(N+M), 空间复杂度 O(1)
*/

func getIntersectNode(head1, head2 *ds.LinkNode[int]) *ds.LinkNode[int] {
	// 判断是否有环
	node1 := isLooped(head1)
	node2 := isLooped(head2)
	// 1. 都有环
	// 		1.1 同一个环同一个入口
	//		1.2 同一个环不同入口
	// 		1.3 不同的环
	if node1 != nil && node2 != nil {
		return getIntersectNodeWithLoop(head1, head2, node1, node2)
	}
	// 2. 都无环
	if node1 == nil && node2 == nil {
		return getIntersectNodeWithoutLoop(head1, head2, nil)
	}
	// 3. 一个有环一个无环不可能相交
	return nil

}

// 链表判环 返回入环点
// 证明：设入环前链长为 x 环大小为 r 快慢指针第一次相遇点距入环点 y
// 慢指针 x + m * r + y 快指针 x + n * r + y
// => (x + m * r + y) * 2 = x + n * r + y
// =>  x + y = t * r
// 当前慢指针距离入环点 y 再走 x 步到达入环点
func isLooped(head *ds.LinkNode[int]) *ds.LinkNode[int] {
	// 1. 快指针一次走两步 慢指针一次走一步 有环则必在环内相遇
	fast := head
	slow := head
	for fast != nil {
		fast = fast.Next
		if fast == slow {
			break
		}
		slow = slow.Next
		if fast == nil {
			return nil
		}
		fast = fast.Next

	}
	// 2. 快指针回到头节点 快慢每次都走一步 两者在环入口相遇
	fast = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}

func getIntersectNodeWithoutLoop(head1, head2, end *ds.LinkNode[int]) *ds.LinkNode[int] {
	work1 := head1
	count1 := 0
	for work1 != end {
		count1++
		work1 = work1.Next
	}
	work2 := head2
	count2 := 0
	for work2 != end {
		count2++
		work2 = work2.Next
	}
	work1 = head1
	work2 = head2
	for count1 > count2 {
		count1--
		work1 = work1.Next
	}
	for count1 < count2 {
		count2--
		work2 = work2.Next
	}
	for work1 != work2 {
		work1 = work1.Next
		work2 = work2.Next
	}
	return work1
}

func getIntersectNodeWithLoop(head1, head2, node1, node2 *ds.LinkNode[int]) *ds.LinkNode[int] {
	if node1 == node2 {
		return getIntersectNodeWithoutLoop(head1, head2, node1)
	}
	if inSameLoop(node1, node2) {
		return node1
	}
	return nil
}

func inSameLoop(node1, node2 *ds.LinkNode[int]) bool {
	work := node1.Next
	for work != node1 {
		if work == node2 {
			return true
		}
		work = work.Next
	}
	return false
}
