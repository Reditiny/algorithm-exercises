package main

import (
	ds "algorithm-exercises/0_data_structure"
)

/**
给定两个字符串 str1 和 str2, 返回两个字符串的最长公共子序列
*/

func commonSubSeq(str1, str2 string) string {
	n := len(str1)
	m := len(str2)
	dp := getCommonSubSeqDp(str1, str2)
	cur := dp[n][m]
	ans := make([]byte, cur)

	for cur > 0 {
		if n > 0 && dp[n][m] == dp[n-1][m] {
			n--
		} else if m > 0 && dp[n][m] == dp[n][m-1] {
			m--
		} else {
			// 此时说明 dp[n][m] == dp[n-1][m-1]
			// 有 str1[n-1] == str2[m-1]，且 str1[n-1] 在最长公共子序列上
			cur--
			ans[cur] = str1[n-1]
			n--
			m--
		}
	}
	return string(ans)
}

// 获得最长公共子序列的 dp 数组
func getCommonSubSeqDp(str1, str2 string) [][]int {
	n := len(str1)
	m := len(str2)

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}

	// dp[i][j] 前 i 和前 j 个元素的最长公共子序列
	// dp[i][j] = max(dp[i-1][j],dp[i][j-1],dp[i-1][j-1]+1/0)
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			dp[i][j] = ds.Max(dp[i-1][j], dp[i][j-1])
			if str1[i-1] == str2[j-1] {
				dp[i][j] = ds.Max(dp[i-1][j-1]+1, dp[i][j])
			}
		}
	}

	return dp
}
