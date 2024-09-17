package main

import ds "algorithm-exercises/0_data_structure"

/**
给定一个矩阵 matrix, 其中的值有正、有负、有 0, 返回子矩阵的最大累加和
*/

func maxSubMatrixSum(matrix [][]int) int {
	// 压缩矩阵的维度后就能将 maxSubMatrixSum 转化为 maxSubArrSum
	// 可考虑依次选择若干连续行压缩成一行，然后找这一行的 maxSubArrSum
	n := len(matrix)
	if n == 0 {
		return 0
	}
	m := len(matrix[0])
	if m == 0 {
		return 0
	}
	ans := 0
	for i1 := 0; i1 < n; i1++ {
		curArr := make([]int, m)
		for i2 := i1; i2 < n; i2++ {
			// 第 i1 行到第 i2 行合并成一行
			for j := 0; j < m; j++ {
				curArr[j] += matrix[i2][j]
			}
			// 求这一行的 maxSubArrSum
			ans = ds.Max(ans, maxSubArrSum(curArr))
		}
	}
	return ans
}
