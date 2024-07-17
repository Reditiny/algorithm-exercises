package main

import (
	ds "algorithm-exercises/0_data_structure"
	"fmt"
)

/*
	二叉树的按层打印与 ZigZag 打印
*/

func printByLevel(head *ds.BTNode[int]) {
	if head == nil {
		return
	}
	queue := ds.NewQueue[*ds.BTNode[int]](10)
	queue.EnQueue(head)
	for !queue.Empty() {
		count := queue.Size()
		for i := 0; i < count; i++ {
			cur := queue.Front()
			queue.DeQueue()
			fmt.Print(cur.Val)
			if cur.Left != nil {
				queue.EnQueue(cur.Left)
			}
			if cur.Right != nil {
				queue.EnQueue(cur.Right)
			}
		}
	}
}

func printByZigZag(head *ds.BTNode[int]) {
	if head == nil {
		return
	}
	deque := ds.NewDeque[*ds.BTNode[int]](100)
	deque.PushBack(head)
	leftToRight := true

	for !deque.Empty() {
		count := deque.Size()
		for i := 0; i < count; i++ {
			if leftToRight {
				cur := deque.Front()
				deque.PopFront()
				fmt.Print(cur)
				if cur.Left != nil {
					deque.PushBack(cur.Left)
				}
				if cur.Right != nil {
					deque.PushBack(cur.Right)
				}
			} else {
				cur := deque.Back()
				deque.PopBack()
				fmt.Print(cur)
				if cur.Right != nil {
					deque.PushFront(cur.Right)
				}
				if cur.Left != nil {
					deque.PushFront(cur.Left)
				}
			}
		}
		leftToRight = !leftToRight
	}
}
