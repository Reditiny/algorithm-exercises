package main

import ds "algorithm-exercises/0_data_structure"

/**
给定一个无序数组 arr, 其中元素可正、可负、可 0,
整数 k 求 arr 子数组中累加和 <= k 的最长长度
*/

func maxLengthLessEqualK(arr []int, k int) int {
	// 生成前缀和数组
	sumPrefix := make([]int, len(arr))
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		sumPrefix[i] = sum
	}
	// 调整前缀和，使其单调递增
	adjustSumPrefix := make([]int, len(sumPrefix))
	copy(adjustSumPrefix, sumPrefix)
	for i := 1; i < len(adjustSumPrefix); i++ {
		if adjustSumPrefix[i] < adjustSumPrefix[i-1] {
			adjustSumPrefix[i] = adjustSumPrefix[i-1]
		}
	}
	// 对于前缀和 sum，找到最小的 i 使得 adjustSumPrefix[i] >= sum-k
	ans := 0
	for i := 0; i < len(sumPrefix); i++ {
		sum = sumPrefix[i]
		need := sum - k
		l, r := 0, i
		j := i
		for l < r {
			mid := l + (r-l)/2
			if adjustSumPrefix[mid] >= need {
				j = mid
				r = mid
			} else {
				l = mid + 1
			}
		}
		if adjustSumPrefix[j] >= need {
			ans = ds.Max(ans, i-j+1)
		}
	}

	return ans
}
