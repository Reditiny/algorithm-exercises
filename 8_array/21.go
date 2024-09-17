package main

/**
给定一个 NxN 的矩阵 matrix, 在这个矩阵中， 只有 0 和 1 两种值，
返回边框全是 1 的最大正方形的边长长度
*/

func maxSquareLength(matrix [][]int) int {
	n := len(matrix)
	// right[i][j] 表示从 matrix[i][j] 到 matrix[i][n-1] 有多少个连续的 1
	right := make([][]int, n)
	for i := 0; i < n; i++ {
		right[i] = make([]int, n)
	}
	// down[i][j] 表示从 matrix[i][j] 到 matrix[n-1][j] 有多少个连续的 1
	down := make([][]int, n)
	for i := 0; i < n; i++ {
		down[i] = make([]int, n)
	}
	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if matrix[i][j] == 0 {
				continue
			}
			if j == n-1 {
				right[i][j] = 1
			} else {
				right[i][j] = 1 + right[i][j+1]
			}
			if i == n-1 {
				down[i][j] = 1
			} else {
				down[i][j] = 1 + down[i+1][j]
			}
		}
	}
	ans := 0
	// 以 matrix[i][j] 为左上角的正方形的最大边长
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			curMaxLen := min(right[i][j], down[i][j])
			for k := curMaxLen; k > ans; k-- {
				if right[i+k-1][j] >= k && down[i][j+k-1] >= k {
					ans = k
					break
				}
			}
		}
	}
	return ans
}
