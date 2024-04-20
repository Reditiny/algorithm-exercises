package main

import ds "algorithm-exercises/0_data_structure"

/*
	给定一个无序单链表的头节点 head, 删除其中值重复出现的节点。
	例如： 1->2->3->3->4->2->1->1 删除值重复的节点之后为 1->2->3->4
	请按以下要求实现两种方法。
	方法1 时间复杂度 O(N)
	方法2 空间复杂度 O(1)
*/

func removeRep1(head *ds.LinkNode[int]) *ds.LinkNode[int] {
	for i := head; i != nil; i = i.Next {
		for j := i; j != nil; {
			if j.Next != nil && j.Next.Val == i.Val {
				j.Next = j.Next.Next
			} else {
				j = j.Next
			}
		}
	}
	return head
}

func removeRep2(head *ds.LinkNode[int]) *ds.LinkNode[int] {
	set := make(map[int]struct{})
	dummyHead := &ds.LinkNode[int]{Next: head}
	for i := dummyHead; i != nil; {
		if i.Next == nil {
			break
		}
		_, ok := set[i.Next.Val]
		if ok {
			i.Next = i.Next.Next
		} else {
			i = i.Next
			set[i.Val] = struct{}{}
		}
	}
	return dummyHead.Next
}
