package main

/**
给定两个 32 位整数 a 和 b, 返回 a 和 b 中较大的
不用任何比较判断
*/

func getMax1(a, b int) int {
	// 32 位整数，最高位为符号位
	// k = 1 时 a - b 为负数，k = 0 时 a - b 为正数
	k := (a - b) >> 31 & 1
	return a - k*(a-b)
}
