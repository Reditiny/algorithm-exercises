package main

import ds "algorithm-exercises/0_data_structure"

/**
给定无序数组 arr, 返回其中最长的连续序列的长度
*/

func longestConsecutive(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	ans := 1
	consecutiveCount := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		if _, ok := consecutiveCount[arr[i]]; !ok {
			consecutiveCount[arr[i]] = 1
			// 向下合并
			if _, ok := consecutiveCount[arr[i]-1]; ok {
				cur := merge(consecutiveCount, arr[i]-1, arr[i])
				ans = ds.Max(ans, cur)
			}
			// 向上合并
			if _, ok := consecutiveCount[arr[i]+1]; ok {
				cur := merge(consecutiveCount, arr[i], arr[i]+1)
				ans = ds.Max(ans, cur)
			}
		}
	}
	return ans
}

func merge(m map[int]int, less, more int) int {
	// 找到 less 的下边界和 more 的上边界
	left := less - m[less] + 1
	right := more + m[more] - 1
	// 得到合并后的长度，并更新两个边界的值，以便后序更新
	length := right - left + 1
	m[left] = length
	m[right] = length
	return length
}
