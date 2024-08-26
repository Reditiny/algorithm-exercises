package main

/**
给定一个整型数组 arr 和一个大于 1 的整数 k
己知 arr 中只有 1 个数出现了 1 次， 其他的数都出现了 k 次， 请返回只出现了 1 次的数
*/

func printOddTimesNumK(arr []int, k int) int {
	// 用一个 32 位的数组记录每个数的每一位的和
	bitCount := make([]int, 32)
	for _, num := range arr {
		for i := 0; i < 32; i++ {
			bitCount[i] += num >> i & 1
			bitCount[i] %= k
		}
	}

	res := 0
	for i := 0; i < 32; i++ {
		res |= bitCount[i] << i
	}

	return res
}
