package main

import (
	ds "algorithm-exercises/0_data_structure"
	"fmt"
)

/**
有 N 个长度不一的数组， 所有的数组都是有序的
请从大到小打印这 N 个数组整体最大的前 K 个数
*/

func printTopK(arrs [][]int, k int) {

	pq := ds.NewPriorityQueue[int](func(a, b int) int {
		return a - b
	})

	n := len(arrs)
	indexes := make([]int, n)
	// 将每个数组的最大值放入优先级队列
	for i := 0; i < n; i++ {
		indexes[i] = len(arrs[i]) - 1
		if indexes[i] >= 0 {
			pq.Push(arrs[i][indexes[i]])
		}
	}
	// 依次弹出优先级队列的最大值，并将该值所在数组的下一个值放入优先级队列
	for i := 0; i < k; i++ {
		if pq.Empty() {
			break
		}
		cur := pq.Top()
		pq.Pop()
		for j := 0; j < n; j++ {
			if indexes[j] >= 0 && arrs[j][indexes[j]] == cur {
				indexes[j]--
				if indexes[j] >= 0 {
					pq.Push(arrs[j][indexes[j]])
				}
				break
			}
		}
		fmt.Println(cur)
	}
}
