package main

/**
给定一个字符串 str，全部由数字字符组成，如果 str 中某一个或某相邻两个字符组成的子串值在 1〜26 之间， 则这个子串可以转换为一个字母
规定"1"转换为"A"，"2"转换为"B"，求 str 有多少种不同的转换结果
*/

func ConvertCount(str string) int {
	n := len(str)
	// 0 开头不合法
	if n <= 1 || str[0] == '0' {
		return n
	}
	// 0 前面的数必须为 1 或 2
	for i := 1; i < n; i++ {
		if str[i] == '0' && (str[i-1] > '2' || str[i-1] == '0') {
			return 0
		}
	}

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	// dp[i] 前 i 个字符有多少种转化结果
	// 如果 str[i] == 0, 则必须与 str[i-1] 组合
	// 如果 str[i-1] == 0, 则 str[i] 不能与之组合
	for i := 2; i <= n; i++ {
		if str[i-1] == '0' {
			dp[i] = dp[i-2]
		} else if str[i-2] == '0' {
			dp[i] = dp[i-1]
		} else {
			dp[i] = dp[i-1] + dp[i-2]
		}
	}

	return dp[n]
}
