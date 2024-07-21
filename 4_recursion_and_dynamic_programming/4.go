package main

/**
给定数组 coins, 所有的值都为正数且不重复。每个值代表一种面值的货币， 每种面值的货币可以使用任意张
再给定一个整数 aim 代表要找的钱数，求换钱有多少种方法（组合数）
*/

func coinChanges2(coins []int, aim int) int {
	dp := make([]int, aim+1)
	dp[0] = 1
	// dp[i][j] 前 i 种货币组成 j 元的方法数
	// dp[i][j] = dp[i-1][j-coins[i]] + dp[i-1][j-2*coins[i]] + ... + dp[i-1][j-n*coins[i]]
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= aim; j++ {
			dp[j] += dp[j-coins[i]]
		}
	}
	return dp[aim]
}
