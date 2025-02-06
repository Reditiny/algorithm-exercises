package main

import (
	ds "algorithm-exercises/0_data_structure"
	"math"
)

/**
给定一个整型数组 arr, 数组中的每个值都为正数， 表示完成一幅画作需要的时间，
再给定一个整数 num 表示画匠的数量， 每个画匠只能画连在一起的画作。
所有的画家并行工作， 请返回完成所有的画作需要的最少时间
*/

func drawDP(arr []int, m int) int {
	n := len(arr)
	if n == 0 || m == 0 {
		return 0
	}

	prefixSum := make([]int, n+1)
	sum := 0
	l := 0
	for i := 1; i <= n; i++ {
		l = ds.Max(l, arr[i])
		sum += arr[i]
		prefixSum[i+1] = sum
	}

	// 如何画家人数多，则每人只画一幅即可
	if m >= n {
		return l
	}

	// dp(i,j) i 个画家画 j 幅画的最少时间
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for j := 1; j <= n; j++ {
		dp[1][j] = prefixSum[j]
	}

	// 状态转移 dp(i,j) = min{max{dp(i-1,t), prefixSum[j]-prefixSum[t])}} 0<=t<=j
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			res := math.MaxInt
			for t := 0; t <= j; t++ {
				r1 := dp[i-1][t] // 第一维可以压缩
				r2 := prefixSum[j] - prefixSum[t]
				res = ds.Min(res, ds.Max(r1, r2))
			}
			dp[i][j] = res
		}
	}

	return dp[n][m]
}

func drawBisection(arr []int, m int) int {
	n := len(arr)
	if n == 0 || m == 0 {
		return 0
	}

	// 前缀和
	prefixSum := make([]int, n+1)
	l, r := 0, 0
	sum := 0
	for i := 1; i <= n; i++ {
		l = ds.Max(l, arr[i])
		sum += arr[i]
		prefixSum[i+1] = sum
	}
	r = sum

	// 如何画家人数多，则每人只画一幅即可
	if m >= n {
		return l
	}

	for l < r {
		mid := l + (r-l)/2
		if getNum(arr, mid) > m {
			r = mid - 1
		} else {
			l = mid
		}
	}
	return l
}

func getNum(arr []int, limit int) int {
	ans := 0
	sum := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > limit {
			return math.MaxInt
		}
		sum += arr[i]
		if sum > limit {
			ans++
			sum = arr[i]
		}
	}
	return ans
}
