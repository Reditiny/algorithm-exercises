package main

import (
	ds "algorithm-exercises/0_data_structure"
	"fmt"
)

/*
	打印两个有序链表的公共部分
	给定两个有序链表的头指针 head1 和 head2, 打印两个链表的公共部分。
*/

func printCommonPart(head1, head2 *ds.LinkNode[int]) {
	work1 := head1
	work2 := head2
	for work1 != nil && work2 != nil {
		if work1.Val < work2.Val {
			work1 = work1.Next
		} else if work1.Val > work2.Val {
			work2 = work2.Next
		} else {
			fmt.Println(work1.Val)
			work1 = work1.Next
			work2 = work2.Next
		}
	}
}
