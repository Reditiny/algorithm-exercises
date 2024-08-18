package main

import (
	"math"
)

/**
给定一个字符串 str, 如果 str 符合日常书写的整数形式， 并且属于 32 位整数的范围，
返回 str 所代表的整数值， 否则返回 0
*/

func strToInt32(str string) int32 {
	n := len(str)

	var ans, mul, sig int64

	if n == 0 || n > 11 || str[0] == '0' {
		return 0
	}

	sig = 1
	if str[0] == '-' {
		sig = -1
		str = str[1:]
	}

	ans = 0
	mul = 1
	for i := len(str) - 1; i >= 0; i-- {
		ans += int64(str[i]-'0') * mul
		mul *= 10
	}

	ans *= sig

	if ans > math.MaxInt32 || ans < math.MinInt32 {
		return 0
	}

	return int32(ans)
}
