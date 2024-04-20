package main

import ds "algorithm-exercises/0_data_structure"

/*
	深拷贝含有随机指针节点的链表
	type LinkNodeWithRand[T any] struct {
		Val T
		Next *LinkNodeWithRand[T]
		Rand *LinkNodeWithRand[T]
	}
	相比正常单链表多了一个 Rand 指针，其可以指向链表中任意节点或nil
	进阶: 时间复杂度为O(N), 空间复杂度 O(N)
*/

func copyLink(head *ds.LinkNodeWithRand[int]) *ds.LinkNodeWithRand[int] {
	if head == nil {
		return nil
	}
	// 首先从左到右遍历链表，拷贝每个节点存入当前节点和下个节点之间
	// 1->2->3->nil  =>  1->1'->2->2'->3->3'->nil
	work := head
	for work != nil {
		newNode := &ds.LinkNodeWithRand[int]{
			Val:  work.Val,
			Next: work.Next,
		}
		work.Next = newNode
		work = newNode.Next
	}
	// 调整新节点的 Rand 指针 new.Rand = origin.Rand.Next
	work = head
	for work != nil {
		next := work.Next.Next
		if work.Rand != nil {
			work.Next.Rand = work.Rand.Next
		}
		work = next
	}
	// 分离新节点
	dummyHead := &ds.LinkNodeWithRand[int]{}
	tail := dummyHead
	work = head
	for work != nil {
		next := work.Next.Next
		tail.Next = work.Next
		tail = tail.Next
		tail.Next = nil
		work.Next = next
		work = next
	}
	return dummyHead.Next
}
