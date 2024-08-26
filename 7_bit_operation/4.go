package main

/**
给定一个整型数组 arr 其中只有一个数出现了奇数次， 其他的数都出现了偶数次，打印这个数

有两个数出现了奇数次， 其他的数都出现了偶数次，打印这两个数
*/

func printOddTimesNum1(arr []int) int {
	res := 0
	for _, num := range arr {
		res ^= num
	}
	return res
}

func printOddTimesNum2(arr []int) (int, int) {
	// 先异或所有数，得到两个奇数次数的异或结果
	res := 0
	for _, num := range arr {
		res ^= num
	}
	// 找到 res 中最右侧的 1
	rightOne := res & (^res + 1)
	// 将数组分为两组，一组 rightOne 位为 1，一组不为 1
	// 两组数分别异或，得到两个奇数次数的数
	ans1, ans2 := 0, 0
	for _, num := range arr {
		if num&rightOne == 0 {
			ans1 ^= num
		} else {
			ans2 ^= num
		}
	}
	return ans1, ans2
}
