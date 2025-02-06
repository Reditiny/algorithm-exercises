package main

import "math"

/**
1 到 n 中 1 出现的次数
n=5: 1, 2, 3, 4, 5 中 1 出现了 1 次
n=11: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11 中 1 出现了 4 次
*/

func oneCount(num int) int {
	if num < 1 {
		return 0
	}

	n := lenOfNum(num)
	powerOfBase10 := int(math.Pow10(n))
	firstNum := num / powerOfBase10

	// 最高位为 1 的个数，若最高位大于 1 则有 powerOfBase10 否则为 num % powerOfBase10 + 1
	// 例如 123 则有 100 101 ... 123 (123%100 + 1 个)
	// 234 则有 100 101 ... 199 (100 个)
	firstOneCount := powerOfBase10
	if firstNum == 1 {
		firstOneCount = num%powerOfBase10 + 1
	}

	// 最高位数字 * 除去最高位后剩下的位数 * 某一位固定是 1 其他位均可是 0～9
	otherOneCount := firstNum * (n - 1) * (powerOfBase10 / 10)

	// 零头递归计算
	return firstOneCount + otherOneCount + oneCount(num%powerOfBase10)
}

func lenOfNum(num int) int {
	ans := 0
	for num != 0 {
		ans++
		num /= 10
	}
	return ans
}
