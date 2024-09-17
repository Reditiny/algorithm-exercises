package main

import ds "algorithm-exercises/0_data_structure"

/**
1. 给定一个无序数组 arr, 其中元素可正、可负、可 0
给定一个整数 k，求 arr 所有的子数组中累加和为 k 的最长子数组长度

2. 给定一个无序数组 arr, 其中元素可正、可负、可 0
求 arr 所有的子数组中正数与负数个数相等的最长子数组长度

3. 给定一个无序数组 arr, 其中元素只是 1 或 0
求 arr 所有的子数组中 0 和 1 个数相等的最长子数组长度
*/

func maxLengthEqualK(arr []int, k int) int {
	// 记录前缀和第一次出现的位置
	sumFirstPos := make(map[int]int)
	sum := 0
	ans := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		// 记录 sum 第一次出现的位置
		if _, ok := sumFirstPos[sum]; !ok {
			sumFirstPos[sum] = i
		}
		// 如果 sum-k 出现过，则说明 sum-k 的位置到 sum 的位置之间的子数组和为 k
		if pos, ok := sumFirstPos[sum-k]; ok {
			ans = ds.Max(ans, i-pos)
		}
	}
	return ans
}

func maxLengthEqualPositiveNegative(arr []int) int {
	// 正数计为 1，负数计为 -1，记录前缀和第一次出现的位置
	sumFirstPos := make(map[int]int)
	sum := 0
	ans := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > 0 {
			sum += 1
		} else if arr[i] < 0 {
			sum -= 1
		}
		// 记录 sum 第一次出现的位置
		if _, ok := sumFirstPos[sum]; !ok {
			sumFirstPos[sum] = i
		}
		// 如果 -sum 出现过，则说明 sum 的位置到 -sum 的位置之间的子数组正负个数相同
		if pos, ok := sumFirstPos[-sum]; ok {
			ans = ds.Max(ans, i-pos)
		}
	}
	return ans
}

func maxLengthEqual01(arr []int) int {
	// 1 计为 1，0 计为 -1，记录前缀和第一次出现的位置
	sumFirstPos := make(map[int]int)
	sum := 0
	ans := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == 1 {
			sum += 1
		} else {
			sum -= 1
		}
		// 记录 sum 第一次出现的位置
		if _, ok := sumFirstPos[sum]; !ok {
			sumFirstPos[sum] = i
		}
		// 如果 -sum 出现过，则说明 sum 的位置到 -sum 的位置之间的子数组 0 1 个数相同
		if pos, ok := sumFirstPos[-sum]; ok {
			ans = ds.Max(ans, i-pos)
		}
	}
	return ans
}
