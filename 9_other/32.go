package main

import (
	ds "algorithm-exercises/0_data_structure"
	"math"
)

/**
一座大楼有 0〜N 层， 地面算作第 0 层， 最高的一层为第 N 层。
己知棋子从第 0 层掉落肯定不会摔碎， 从第 i 层（0<i<=N）掉落可能会摔碎
给定整数 N 作为楼层数， 再给定整数 K 作为棋子数，
返回如果想找到棋子不会摔碎的最高层数， 即最差情况下的最少次数。一次只能扔一个棋子
*/

// 递归 复杂度 O(N!)
func solutionRecursion(n, k int) int {
	// 设 s(n,k) 为 n 层楼 k 个棋子最差情况下的最少次数
	var s func(n, k int) int
	s = func(n, k int) int {
		if n == 0 || k == 0 {
			return 0
		}
		if k == 1 {
			return n
		}
		res := math.MaxInt
		// 棋子扔在 i 层楼上
		for i := 1; i <= n; i++ {
			// 棋子碎了，剩下 k-1 个棋子，继续在 i-1 层楼上测试
			r1 := s(i-1, k-1)
			// 棋子没碎，剩下 k 个棋子，继续在 n-i 层楼上测试
			r2 := s(n-i, k)
			// 最差情况下的最少次数
			res = min(res, max(r1, r2)+1)
		}
		return res
	}
	return s(n, k)
}

// 动态规划 复杂度 O(N^2K)
func solutionDP(n, k int) int {
	// 设 dp(n,k) 为 n 层楼 k 个棋子最差情况下的最少次数
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
	}
	for i := 1; i <= n; i++ {
		dp[i][1] = i
	}
	for i := 1; i <= k; i++ {
		dp[0][i] = 0
	}
	// 状态转移方程 dp(i,j) = min{max{dp(t-1,j-1),dp(i-t,j)}+1} 1<=t<=i
	for i := 1; i <= n; i++ {
		for j := 2; j <= k; j++ {
			res := math.MaxInt
			for t := 1; t <= i; t++ {
				r1 := dp[t-1][j-1]
				r2 := dp[i-t][j]
				res = ds.Min(res, ds.Max(r1, r2)+1)
			}
			dp[i][j] = res
		}
	}
	return dp[n][k]
}

// 动态规划 复杂度 O(N^2K) 优化空间复杂度
func solutionDP2(n, k int) int {
	// 设 dp(n,k) 为 n 层楼 k 个棋子最差情况下的最少次数
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	for i := 1; i <= n; i++ {
		dp[i][1] = i
	}
	for i := 1; i <= k; i++ {
		dp[0][i%2] = 0
	}
	// 状态转移方程 dp(i,j) = min{max{dp(t-1,j-1),dp(i-t,j)}+1} 1<=t<=i
	for i := 1; i <= n; i++ {
		for j := 2; j <= k; j++ {
			res := math.MaxInt
			for t := 1; t <= i; t++ {
				r1 := dp[t-1][(j-1)%2]
				r2 := dp[i-t][j%2]
				res = ds.Min(res, ds.Max(r1, r2)+1)
			}
			dp[i][j%2] = res
		}
	}
	return dp[n][k]
}
