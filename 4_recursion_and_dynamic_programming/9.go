package main

import ds "algorithm-exercises/0_data_structure"

/**
给定两个字符串 str1 和 str2, ic、de 和 re, 分别代表插入、删除和替换一个字符的代价，
返回将 str1 编辑成 str2 的最小代价
*/

func editDistance(str1, str2 string, ic, de, re int) int {
	n := len(str1)
	m := len(str2)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}
	// dp[i][j] str1[:i] 变成 str2[:j] 的代价
	// dp[i][j] = dp[i-1][j-1] if str1[i] == str2[j]
	// dp[i][j] = min(d[i-1][j] + de, dp[i][j-1] + ic, dp[i-1][j-1] + re)

	// 从 i 到 0，需要删除 i 个
	for i := 1; i <= n; i++ {
		dp[i][0] = i * de
	}
	// 从 0 到 j，需要插入 j 个
	for j := 1; j < m; j++ {
		dp[0][j] = j * ic
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = ds.Min(ds.Min(dp[i][j-1]+ic, dp[i-1][j]+de), dp[i-1][j-1]+re)
			}
		}
	}
	return dp[n][m]
}
