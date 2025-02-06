package main

import ds "algorithm-exercises/0_data_structure"

/**
给定一个正数数组 arr, 其中所有的值都为整数, 以下是最小不可组成和的概念:
1. 把 arr 每个子集求和, 其中最小的记为 min, 最大的记为 max
2. 区间[min,max］上, 若有数不可由 arr 某一子集求和得到, 其中的最小值为 arr 的最小不可组成和
3. 如果所有的数都可以由 arr 的某一子集求和得到, 则 max+1 是 arr 的最小不可组成和
*/

func minUnformedSum(arr []int) int {
	n := len(arr)
	if n == 0 {
		return 0
	}

	minSum := arr[0]
	maxSum := 0
	for i := range arr {
		minSum = ds.Min(minSum, arr[i])
		maxSum += arr[i]
	}

	// dp[j] 表示 j 是否为组成和
	// 对于 arr[i]，dp[j] = dp[j-arr[i]]
	dp := make([]bool, maxSum+1)
	dp[0] = true
	for i, _ := range arr {
		for j := 0; j+arr[i] <= maxSum; j++ {
			if dp[j] {
				dp[j+arr[i]] = true
			}
		}
	}

	for i := range dp {
		if !dp[i] {
			return i
		}
	}

	return maxSum + 1
}
