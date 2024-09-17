package main

import ds "algorithm-exercises/0_data_structure"

/**
给定一个数组 arr, 返回子数组的最大累加和
要求时间复杂度为O(N), 额外空间复杂度为O(1)
*/

func maxSubArrSum(arr []int) int {
	n := len(arr)
	if n == 0 {
		return 0
	}
	// 最终答案的子数组的任意前缀数组之和均不小于 0
	ans := 0
	cur := 0
	for i := 0; i < n; i++ {
		// 依次累加得到前缀并更新最大值
		cur += arr[i]
		ans = ds.Max(ans, cur)
		// 如果当前前缀和小于 0，那么这一段前缀不会成为最终答案数组的前缀
		if ans < 0 {
			ans = 0
		}
	}
	return ans
}
