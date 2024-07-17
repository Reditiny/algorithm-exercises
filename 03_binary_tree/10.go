package main

import (
	ds "algorithm-exercises/0_data_structure"
)

/*
	判断 t1 树是否包含 t2 树全部的拓扑结构
*/

func containsTopo(head1, head2 *ds.BTNode[int]) bool {

	contained := false

	preOrderIter(head1, func(cur *ds.BTNode[int]) {
		if contained {
			return
		}
		if checkContainsTopo(cur, head2) {
			contained = true
		}
	})

	return contained
}

func checkContainsTopo(head1, head2 *ds.BTNode[int]) bool {
	if head2 == nil {
		return true
	}
	if head1 == nil || head1.Val != head2.Val {
		return false
	}
	return checkContainsTopo(head1.Right, head2.Right) && checkContainsTopo(head1.Left, head2.Left)
}
