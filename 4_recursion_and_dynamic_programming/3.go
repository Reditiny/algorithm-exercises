package main

import (
	ds "algorithm-exercises/0_data_structure"
)

/**
给定数组 coins, 所有的值都为正数且不重复。每个值代表一种面值的货币，每种面值的货币可以使用任意张
再给定一个整数 aim 代表要找的钱数， 求组成 aim 的最少货币数
*/

const UnChangeable int = 10000

func coinChanges1(coins []int, aim int) int {
	dp := make([]int, aim+1)
	dp[0] = 0
	// dp[i] i 元钱用的最少数，用了一个 j 货币
	// dp[i] = min(dp[i-coins[j]] + 1)
	for i := 1; i <= aim; i++ {
		dp[i] = UnChangeable
		for j := 0; j < len(coins); j++ {
			if i >= coins[j] && dp[i-coins[j]] < UnChangeable {
				dp[i] = ds.Min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	if dp[aim] >= UnChangeable {
		return -1
	}
	return dp[aim]
}
