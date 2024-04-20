package main

import ds "algorithm-exercises/0_data_structure"

/*
	据说著名犹太历史学家 Josephus 有过以下故事：在罗马人占领乔塔帕特后，39个犹太
	人与 Josephus 及他的朋友躲到一个洞中， 39 个犹太人决定宁愿死也不要被敌人抓到，
	于是决定了一个自杀方式，41 个人排成一个圆圈，由第 1 个人开始报数，报数到 3 的人就自杀,
	然后再由下一个人重新报 1, 报数到 3 的人再自杀， 这样依次下去， 直到剩下最后一个人时，
	那个人可以自由选择自己的命运。这就是著名的约瑟夫问题。
	现在请用单向环形链表描述该结构并呈现整个过程

	输入： 一个环形单向链表的头节点 head 和报数的值 m。
	返回： 最后生存下来的节点， 且这个节点自己组成环形单向链表， 其他节点都删掉。
	进阶：如果链表节点数为M 想在时间复杂度为 O(N)时完成原问题的要求， 该怎么实现？
*/

func josephusLast(head *ds.LinkNode[int], m int) *ds.LinkNode[int] {
	n := 0
	work := head
	for work != nil {
		n++
		work = work.Next
	}
	last := getLive(n, m)
	work = head
	for last > 1 {
		work = work.Next
		last--
	}
	work.Next = work
	return work
}

// n 个节点返回幸存节点编号
// 找到 getLive(n) 与 getLive(n-1) 的关系
// 设从第一个节点开始报数，第 i 个节点报数为 (i-1)%m+1
// 删除第 m 个节点后，原第 m + 1 个节点变为第一个节点
// 1 2 3 .... m-1 m m+1 m+2 .. n   old
// ...... n-2 n-1 _  1   2 .....   new
// old = (new+m-1)%n + 1
func getLive(n, m int) int {
	if n == 1 {
		return 1
	}
	return (getLive(n-1, m)+m-1)%n + 1
}
