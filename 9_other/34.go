package main

import "math"

/**
一条直线上有居民点， 邮局只能建在居民点上
给定一个有序整型数组 arr, 每个值表示居民点的一维坐标, 再给定一个正数 num, 表示邮局数量。
选择 num 个居民点建立 num 个邮局， 使所有的居民点到邮局的总距离最短， 返回最短的总距离
*/

func minDistance(arr []int, num int) int {
	n := len(arr)
	if n == 0 || num <= 0 {
		return 0
	}

	// w[i][j] 表示在 arr[i] 到 arr[j] 之间建立一个邮局的最短距离
	// [i,j] 奇数个居民点时，邮局建在中间的居民点，偶数个居民点时，邮局建在中间两个居民点的任意一个
	w := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		w[i] = make([]int, n+1)
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			// 假设 [i,j] 之间有奇数个居民点，则邮局在 (i+j)/2 处，[i,j-1] 之间有偶数个居民点，邮局在 (i+j-1)/2 或 (i+j)/2 处
			// [i,j] 之间有偶数个居民点时同理
			// 因此 w[i][j] 与 w[i][j-1] 的邮局距离只多了最后一个居民点，其他居民点到邮局的距离没有变化
			w[i][j] = w[i][j-1] + arr[j] - arr[(i+j)/2]
		}
	}

	// dp[i][j] 表示在 arr[0] 到 arr[j] 之间建立 i+1 个邮局的最短距离
	dp := make([][]int, num)
	for i := 0; i < num; i++ {
		dp[i] = make([]int, n)
	}

	// 初始化只建一个邮局
	for j := 0; j < n; j++ {
		dp[0][j] = w[0][j]
	}

	// dp[i][j] = min{dp[i-1][k] + w[k+1][j]} (0 <= k < j)
	for i := 1; i < num; i++ {
		for j := i; j < n; j++ {
			res := math.MaxInt
			for k := 0; k < j; k++ {
				res = min(res, dp[i-1][k]+w[k+1][j])
			}
			dp[i][j] = res
		}
	}

	return dp[num-1][n-1]
}
