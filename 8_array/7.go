package main

/**
给定一个有 NxM 的整型矩阵 matrix 和一个整数 K
matrix 的每一行和每一列都是有序的。
实现一个函数， 判断 K 是否在 matrix 中
*/

func hasK(matrix [][]int, k int) bool {
	n := len(matrix)
	if n == 0 {
		return false
	}
	m := len(matrix[0])
	if m == 0 {
		return false
	}
	i, j := 0, m-1
	for i < n && j >= 0 {
		if matrix[i][j] > k {
			j--
		} else if matrix[i][j] < k {
			i--
		} else {
			return true
		}
	}

	return false
}
