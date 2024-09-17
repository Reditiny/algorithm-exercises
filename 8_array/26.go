package main

import ds "algorithm-exercises/0_data_structure"

/**
给定一个整型数组 arr, 返回排序后的相邻两数的最大差值
*/

func maxGap(arr []int) int {
	n := len(arr)
	if n <= 1 {
		return 0
	}
	// 找到数组中的最大值和最小值
	maxValue := arr[0]
	minValue := arr[0]
	for i := 1; i < n; i++ {
		maxValue = ds.Max(maxValue, arr[i])
		minValue = ds.Min(minValue, arr[i])
	}
	if maxValue == minValue {
		return 0
	}
	// 有 n 个数， 准备 n+1 个桶，则必有桶为空
	// 桶内数的最大差一定小于相隔空桶的最小差，因此只需记录桶内最大值和最小值即可
	hasNum := make([]bool, n+1)
	maxs := make([]int, n+1)
	mins := make([]int, n+1)
	for i := 0; i < n; i++ {
		// 确定 arr[i] 放入几号桶
		bid := (arr[i] - minValue) * n / (maxValue - minValue)
		hasNum[bid] = true
		maxs[bid] = ds.Max(maxs[bid], arr[i])
		mins[bid] = ds.Min(mins[bid], arr[i])
	}
	// 0 号桶一定有，依次往前找两个“相邻”非空空桶，计算两个桶之间的差值
	ans := 0
	lastMax := maxs[0]
	for i := 1; i <= n; i++ {
		if hasNum[i] {
			ans = ds.Max(ans, mins[i]-lastMax)
			lastMax = maxs[i]
		}
	}
	return ans
}
