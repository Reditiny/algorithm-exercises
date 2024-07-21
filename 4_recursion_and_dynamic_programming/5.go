package main

import ds "algorithm-exercises/0_data_structure"

/**
给定数组 arr, 返回 arr 的最长递增子序列
*/

func getList(arr []int) []int {
	dpList, maxVal := getDPListAndMaxLen(arr)
	// 得到 dp 数组以及最长个数时，从后往前将 dp 数组中相应值取出即是最长递增子序列
	ans := make([]int, len(dpList))
	idx := len(ans) - 1
	for i := len(dpList) - 1; i >= 0; i-- {
		if dpList[i] == maxVal {
			ans[idx] = arr[i]
			maxVal--
			idx--
		}
	}
	return ans
}

func getDPListAndMaxLen(arr []int) ([]int, int) {
	n := len(arr)
	dp := make([]int, n)
	if n == 0 {
		return dp, n
	}
	// dp[i] 前 i 个元素的最长递增子序列
	// dp[i] = max(dp[j] + 1) for j < i and arr[j] < arr[i]
	dp[0] = 1
	ans := 1
	for i := 1; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] {
				dp[i] = ds.Max(dp[i], dp[j]+1)
			}
		}
		ans = ds.Max(ans, dp[i])
	}
	return dp, ans
}
