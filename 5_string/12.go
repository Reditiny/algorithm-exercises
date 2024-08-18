package main

import ds "algorithm-exercises/0_data_structure"

/**
给定一个字符串 str, 如果可以在 str 的任意位置添加字符
请返回最少添加几个字符，让 str 整体都是回文字符串的一种结果
*/

func getPalindrome(str string) string {

	// dp[i][j] 表示 str[i..j] 最少添加几个字符可以让 str[i..j] 整体都是回文字符串
	dp := getDp(str)
	ans := make([]byte, len(str)+dp[0][len(str)-1])
	i := 0
	j := len(str) - 1
	ansI := 0
	ansJ := len(ans) - 1
	for i <= j {
		if str[i] == str[j] {
			ans[ansI] = str[i]
			ans[ansJ] = str[j]
			i++
			j--
		} else if dp[i][j-1] < dp[i+1][j] {
			ans[ansI] = str[j]
			ans[ansJ] = str[j]
			j--
		} else {
			ans[ansI] = str[i]
			ans[ansJ] = str[i]
			i++
		}
		ansI++
		ansJ--
	}
	return string(ans)
}

func getDp(str string) [][]int {
	dp := make([][]int, len(str))
	for i := range dp {
		dp[i] = make([]int, len(str))
	}

	n := len(str)
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if i == j {
				dp[i][j] = 0
			} else if i+1 == j {
				if str[i] == str[j] {
					dp[i][j] = 0
				} else {
					dp[i][j] = 1
				}
			} else {
				if str[i] == str[j] {
					dp[i][j] = dp[i+1][j-1]
				} else {
					dp[i][j] = ds.Min(dp[i+1][j], dp[i][j-1]) + 1
				}
			}
		}
	}
	return dp
}
