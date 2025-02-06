package main

import ds "algorithm-exercises/0_data_structure"

/**
给定一个长度为 N (N>1) 的整型数组 arr, 可以划分成左右两个部分
左部分为arr[0..K], 右部分为arr[K+1..N], K 取值的范围是[0,N-1]。
求左部分中的最大值减去右部分最大值的绝对值的最大值
*/

// O(N) 时间复杂度 O(N) 空间复杂度
func maxABS(arr []int) int {
	leftMax := make([]int, len(arr))
	leftMax[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		leftMax[i] = ds.Max(leftMax[i-1], arr[i])
	}
	rightMax := make([]int, len(arr))
	rightMax[len(arr)-1] = arr[len(arr)-1]
	for i := len(arr) - 2; i >= 0; i-- {
		rightMax[i] = ds.Max(rightMax[i+1], arr[i])
	}
	ans := 0
	for i := 0; i < len(arr)-1; i++ {
		if ds.Abs(leftMax[i]-rightMax[i+1]) > ans {
			ans = ds.Abs(leftMax[i] - rightMax[i+1])
		}
	}
	return ans
}

// O(N) 时间复杂度 O(1) 空间复杂度
func maxABS2(arr []int) int {
	// 全集最大值
	maxValue := arr[0]
	for i := 1; i < len(arr); i++ {
		maxValue = ds.Max(maxValue, arr[i])
	}
	// 左右两部分都是越少最大值越小，因此只需要分成 [0] + [1, N-1] 和 [0, N-2] + [N-1] 两种情况
	return maxValue - ds.Min(arr[0], arr[len(arr)-1])
}
