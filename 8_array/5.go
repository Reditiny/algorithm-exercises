package main

/**
给定一个无序数组 arr, 求出需要排序的最短子数组长度。
例如 arr=[1, 5, 3, 4, 2, 6, 7] 返回 4, 因为只有[5, 3, 4, 2]需要排序
*/

func minSortLength(arr []int) int {
	n := len(arr)
	if n <= 1 {
		return 0
	}
	// 如果一个数比它右边的最小值大，那么这个数一定要排序
	// 如果一个数比它左边的最大值小，那么这个数一定要排序

	// 找到最左边的需要排序的数
	noMinIndex := -1
	curMin := arr[n-1]
	for i := n - 2; i >= 0; i-- {
		if arr[i] > curMin {
			noMinIndex = i
		} else {
			curMin = arr[i]
		}
	}

	if noMinIndex == -1 {
		return 0
	}
	// 找到最右边的需要排序的数
	noMaxIndex := -1
	curMax := arr[0]
	for i := 1; i < n; i++ {
		if arr[i] < curMax {
			noMaxIndex = i
		} else {
			curMax = arr[i]
		}
	}

	return noMaxIndex - noMinIndex + 1
}
