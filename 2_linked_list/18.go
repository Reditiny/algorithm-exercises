package main

import (
	ds "algorithm-exercises/0_data_structure"
	"math"
)

/*
	一个环形单链表从头节点 head 开始不降序， 同时由最后的节点指回头节点。
	给定这样头节点 head 和整数 num, 请生成节点值为 num 的新节点
	插入到这个环形链表中， 保证调整后的链表依然有序
*/

// head 为值最小的点
func insertToLinkList(head *ds.LinkNode[int], num int) *ds.LinkNode[int] {
	if head == nil {
		newHead := &ds.LinkNode[int]{Val: num}
		newHead.Next = newHead
		return newHead
	}
	// 断开 tail -> head 的链
	work := head
	for work.Next != head {
		work = work.Next
	}
	work.Next = nil
	tail := work
	// 找到插入位置的前驱
	dummyHead := &ds.LinkNode[int]{Val: math.MinInt, Next: head}
	work = dummyHead
	for work != nil {
		if work.Next != nil && work.Val <= num && num <= work.Next.Val {
			break
		}
		work = work.Next
	}
	// 插入后还原链
	newNode := &ds.LinkNode[int]{Val: num}
	if work == nil {
		tail.Next = newNode
		newNode.Next = dummyHead.Next
	} else {
		newNode.Next = work.Next
		work.Next = newNode
		tail.Next = dummyHead.Next
	}
	return dummyHead.Next
}

// head 为随机点
func insertToLinkList2(head *ds.LinkNode[int], num int) *ds.LinkNode[int] {
	if head == nil {
		newHead := &ds.LinkNode[int]{Val: num}
		newHead.Next = newHead
		return newHead
	}
	ans := head
	// 断开链时找到最小值的点
	tail := head
	work := head
	for work.Next != head {
		if work.Val > work.Next.Val {
			tail = work
			head = work.Next
			break
		}
		work = work.Next
	}
	tail = work
	tail.Next = nil
	// 找到插入位置的前驱
	dummyHead := &ds.LinkNode[int]{Val: math.MinInt, Next: head}
	work = dummyHead
	for work != nil {
		if work.Next != nil && work.Val < num && num <= work.Next.Val {
			break
		}
		work = work.Next
	}
	// 插入后还原链
	newNode := &ds.LinkNode[int]{Val: num}
	if work == nil {
		tail.Next = newNode
		newNode.Next = dummyHead.Next
	} else {
		newNode.Next = work.Next
		work.Next = newNode
		tail.Next = dummyHead.Next
	}

	return ans
}
