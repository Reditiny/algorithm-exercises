package main

import ds "algorithm-exercises/0_data_structure"

/**
给定一个二维数组 map， 含义是一张地图，如下矩阵
  -2  -3  3
  -5  -10 1
   0  30 -5
骑士从左上角出发，每次只能向右或向下走， 最后到达右下角见到公主
负数说明此处有怪兽，要让骑士损失血量，否则代表此处有血瓶，能让骑士回血
为了保证骑士能见到公主， 初始血量至少是多少
*/

func getInitHeath(graph [][]int) int {
	n := len(graph)
	m := len(graph[0])
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}
	dp[n][m] = 1
	if graph[n-1][m-1] < 0 {
		dp[n][m] -= graph[n-1][m-1]
	}

	// dp[i][j] 到达(i,j)处的最小血量
	// 如果 (i,j) 为怪兽，则下一步至少要满足 dp[i+1][j]、dp[i][j+1] 其中之一
	// 若为血包，还要额外保证当前血量至少为 1
	// dp[i][j] = ds.Min(dp[i+1][j], dp[i][j+1]) - graph[i-1][j-1]

	for i := n; i >= 0; i-- {
		for j := m; j > 0; j-- {
			if i == n && j == m {
				continue
			} else if i != n && j == m {
				dp[i][j] = dp[i+1][j] - graph[i-1][j-1]
			} else if i == n && j != m {
				dp[i][j] = dp[i][j+1] - graph[i-1][j-1]
			} else {
				dp[i][j] = ds.Min(dp[i+1][j], dp[i][j+1]) - graph[i-1][j-1]
			}
			if dp[i][j] < 0 {
				dp[i][j] = 1
			}
		}
	}

	return dp[0][0]
}
