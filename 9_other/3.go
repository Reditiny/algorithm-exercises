package main

/**
1. 给定一个非负整数 N, 返回 M 结果的末尾为 0 的数量。

2. 给定一个非负整数 N , 用二进制数表达 N! 的结果， 返回最低位的 1 在哪个位置上
*/

// 5 * 2 = 10
// 因子 2 的数目比因子 5 的数目多，所以只需要统计因子 5 的数目
// Z = N/5 + N/5^2 + N/5^3 + ... + N/5^k
func zeroCount(n int) int {
	count := 0
	for n > 0 {
		n /= 5
		count += n
	}
	return count
}

// 每个因子 2 都可以让最低位的 1 向左移动一位
// Z = N/2 + N/2^2 + N/2^3 + ... + N/2^k
func lowBit(n int) int {
	count := 0
	for n > 0 {
		n /= 2
		count += n
	}
	return count
}
