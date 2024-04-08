package main

import (
	ds "algorithm-exercises/0_data_structure"
)

/*
	最大值减去最小值小于或等于 num 的子数组数量
	给定数组 arr 和整数 num, 共返回有多少个子数组满足如下情况：
	max(arr[i..j]) - min(arr[i..j]) <= num
	如果数组长度为N 请实现时间复杂度为O(N)的解法
*/

func subArrCount(arr []int, num int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}

	n := len(arr)
	// 单调双端队列维护窗口中的最大值与最小值的索引
	maxDeque := ds.NewDeque[int](n)
	minDeque := ds.NewDeque[int](n)
	// 双指针滑动窗口，以 i 为窗口左端，探查最远的右端
	// 虽然两层 for 但是 j 不会回溯一直再增加，复杂度为 O(N)
	j, ans := 0, 0
	for i := 0; i < n; i++ {
		for j < n {
			for !maxDeque.Empty() && arr[maxDeque.Back()] <= arr[j] {
				maxDeque.PopBack()
			}
			maxDeque.PushBack(i)
			for !minDeque.Empty() && arr[minDeque.Back()] >= arr[j] {
				minDeque.PopBack()
			}
			minDeque.PushBack(i)
			if arr[maxDeque.Front()]-arr[minDeque.Front()] > num {
				// 当 arr[i..j] 不满足条件时，后续的 arr[i..j+t] 都不会满足
				break
			}
			j++
		}
		ans += j - i
		// 本轮右端确定后，窗口左端移动
		if maxDeque.Front() == i {
			maxDeque.PopFront()
		}
		if minDeque.Front() == i {
			minDeque.PopFront()
		}
	}

	return 0
}
