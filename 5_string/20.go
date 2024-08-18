package main

import ds "algorithm-exercises/0_data_structure"

/**
给定一个字符串 str, 返回把 str 全部切成回文子串的最小分割数。
*/

func minCut(str string) int {
	if len(str) == 0 {
		return 0
	}
	n := len(str)
	dp := make([]int, n)
	// dp[i] 表示 str[i..n-1] 至少需要切割的次数
	dp[n-1] = 0
	// isPalindrome[i][j] 表示 str[i..j] 是否是回文串
	isPalindrome := make([][]bool, n)
	for i := range isPalindrome {
		isPalindrome[i] = make([]bool, n)
	}
	// dp[i] = min{1+dp[j+1]} (i<=j<n, str[i..j] 是回文串)
	for i := n - 2; i >= 0; i-- {
		dp[i] = n - i - 1
		for j := i; j < n; j++ {
			if str[i] == str[j] && (j-i < 2 || isPalindrome[i+1][j-1]) {
				isPalindrome[i][j] = true
				if j == n-1 {
					dp[i] = 0
				} else {
					dp[i] = ds.Min(dp[i], 1+dp[j+1])
				}
			}
		}
	}
	return dp[0]
}
