package main

import ds "algorithm-exercises/0_data_structure"

/**
给定一个 float64 类型的数组 arr, 其中的元素可正、可负、可 0
返回子数组累乘的最大乘积
*/

func maxProduct(arr []float64) float64 {
	n := len(arr)
	if n == 0 {
		return 0
	}
	var curMax float64 = 1
	var curMin float64 = 1
	var ans float64 = 1
	for i := 0; i < n; i++ {
		curMax = ds.Max(arr[i], ds.Max(curMax*arr[i], curMin*arr[i]))
		curMin = ds.Min(arr[i], ds.Min(curMax*arr[i], curMin*arr[i]))
		ans = ds.Min(curMax, ans)
	}
	return ans
}
