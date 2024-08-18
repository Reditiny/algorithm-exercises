package main

import "sort"

/**
判断字符数组中是否所有的字符都只出现过一次
1. 实现时间复杂度为 O(N)的方法。
2. 在保证额外空间复杂度为O(1)的前提下， 请实现时间复杂度尽量低的方法。
*/

func isUnique1(arr []byte) bool {
	count := make([]int, 256)
	for i := range arr {
		count[arr[i]]++
		if count[arr[i]] > 1 {
			return false
		}
	}

	return true
}

func isUnique2(arr []byte) bool {
	n := len(arr)
	if n <= 1 {
		return true
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	last := arr[0]
	for i := 1; i < n; i++ {
		if arr[i] == last {
			return false
		}
		last = arr[i]
	}
	return true
}
