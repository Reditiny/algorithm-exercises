package main

import ds "algorithm-exercises/0_data_structure"

/**
给定一个矩阵, 从左上角开始每次只能向右或者向下走， 最后到达右下角的位置,
路径上所有的数字累加起来就是路径和， 返回所有的路径中最小的路径和
*/

func minPathSum(matrix [][]int) int {
	dp := make([]int, len(matrix[0]))
	dp[0] = matrix[0][0]
	for i := 1; i < len(matrix[0]); i++ {
		dp[i] = dp[i-1] + matrix[0][i]
	}
	// dp[i][j] 位置 (i,j) 的最小路经和，上一步（上和左两种情况）也应该是最小
	// dp[i][j] = min(dp[i-1][j],dp[i][j-1]) + matrix[i][j]
	// 可以通过滚动数组压缩第一维
	for i := 1; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if j == 0 {
				dp[j] += matrix[i][j]
			} else {
				dp[j] = ds.Min(dp[j-1], dp[j]) + matrix[i][j]
			}
		}
	}
	return dp[len(matrix[0])-1]
}
