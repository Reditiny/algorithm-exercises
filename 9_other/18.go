package main

import (
	ds "algorithm-exercises/0_data_structure"
)

/**
判断一个数是否是回文数，负数取绝对值
*/

func isPalindromicNumber(num int) bool {
	absNum := ds.Abs(num)
	if absNum < 10 {
		return true
	}

	help := 10
	for absNum/help > 10 {
		help *= 10
	}

	// 依次比较最高位和最地位，然后移除
	for absNum >= 10 {
		high := absNum / help
		low := absNum % 10
		if high != low {
			return false
		}

		absNum %= help
		absNum -= low
		absNum /= 10
		help /= 100
	}
	return true
}
