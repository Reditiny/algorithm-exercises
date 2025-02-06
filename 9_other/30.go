package main

import ds "algorithm-exercises/0_data_structure"

/**
给定一个字符串 str, 返回 str 中最长回文子串的长度
*/

// dp 时间复杂度 O(n^2) 空间复杂度 O(n^2)
func longestPalindrome(str string) int {
	n := len(str)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		dp[i][i] = 1
		if i < n-1 && str[i] == str[i+1] {
			dp[i][i+1] = 2
		}
	}

	ans := 1
	for i := n - 1; i >= 0; i-- {
		for j := i + 2; j < n; j++ {
			if str[i] == str[j] {
				if dp[i+1][j-1] == 0 {
					dp[i][j] = 0
				} else {
					dp[i][j] = dp[i+1][j-1] + 2
				}
			}
			ans = ds.Max(ans, dp[i][j])
		}
	}

	return ans
}

// Manacher 算法 todo
