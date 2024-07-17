package main

/**
二叉树共 n 个节点，返回可能的二叉树结构有多少个
*/

func countTree(n int) int {

	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		for c := 0; c < i; c++ {
			dp[i] += dp[c] * dp[i-1-c]
		}
	}

	return dp[n]
}
