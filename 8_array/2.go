package main

/**
给定一个 NXN 的矩阵 matrix, 把这个矩阵调整成顺时针转动 90 度
*/

func rotateMatrix(matrix [][]int) {
	n := len(matrix)
	if n == 0 {
		return
	}

	for i := 0; i < n/2; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j], matrix[n-i-1][j] = matrix[n-i-1][j], matrix[i][j]
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
