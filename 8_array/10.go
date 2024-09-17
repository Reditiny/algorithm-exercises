package main

/**
给定一个数组 arr，无序，均为正数，求 arr 的所有子数组中元素和为 k 的最长子数组长度
*/

func maxLengthEqualKAllPositive(arr []int, k int) int {
	sum := 0
	j := -1
	ans := -1
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		for sum > k {
			sum -= arr[j+1]
			j++
		}
		if sum == k && i-j > ans {
			ans = i - j
		}
	}

	return ans
}
