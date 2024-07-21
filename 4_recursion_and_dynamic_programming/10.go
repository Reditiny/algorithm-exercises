package main

/**
给定三个字符串 str1 str2 和 aim, 如果 aim 包含且仅包含来自 str1 和 str2 的所有字符，
而且在 aim 中属于 str1 的字符之间保持原来在 str1 中的顺序，属于 str2 的字符之间保持原来在 str2 中的顺序
那么称 aim 是 str1 和 str2 的交错组成。
实现一个函数， 判断 aim 是否是 str1 和 str2 交错组成
*/

func isCross(str1, str2, aim string) bool {
	n := len(str1)
	m := len(str2)
	v := len(aim)
	if n+m != v {
		return false
	}

	dp := make([][]bool, n+1)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, m+1)
	}

	// dp[i][j] str1[:i] 与 str2[:j] 是否可以组成 aim[:i+j]
	// dp[i][j] = (str1[i-1] == aim[i+j-1] && dp[i-1][j]) || (str2[j-1] == aim[i+j-1] && dp[i][j-1])
	dp[0][0] = true
	for i := 1; i < n; i++ {
		dp[i][0] = str1[0:i] == aim[0:i]
	}
	for j := 1; j < m; j++ {
		dp[0][j] = str2[0:j] == aim[0:j]
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if str1[i-1] == aim[i+j-1] {
				dp[i][j] = dp[i-1][j] || dp[i][j]
			}
			if str2[j-1] == aim[i+j-1] {
				dp[i][j] = dp[i][j-1] || dp[i][j]
			}
		}
	}

	return dp[n][m]
}
