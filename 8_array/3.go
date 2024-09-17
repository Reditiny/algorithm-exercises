package main

import "fmt"

/**
给定一个矩阵 matrix, 按照 "之" 字形的方式打印这个矩阵， 例如：
1  2  3  4
5  6  7  8
9 10 11 12
“之”字形打印的结果为：1, 2, 5, 9, 6, 3, 4, 7, 10, 11, 8, 12
*/

func printMatrixZigZag(matrix [][]int) {

	moves := [][]int{{-1, 1}, {1, -1}}

	n := len(matrix)
	if n == 0 {
		return
	}

	m := len(matrix[0])
	if m == 0 {
		return
	}

	i, j := 0, 0
	d := 0
	for {
		fmt.Print(matrix[i][j], " ")
		if i == n-1 && j == m-1 {
			break
		}
		nextI, nextJ := i+moves[d][0], j+moves[d][1]
		if d == 0 {
			if nextJ >= m {
				nextJ = m - 1
				nextI += 2
				d = 1
			} else if nextI < 0 {
				nextI = 0
				d = 1
			}
		}

		if d == 1 {
			if nextI >= n {
				nextI = n - 1
				nextJ += 2
				d = 0
			} else if nextJ < 0 {
				nextJ = 0
				d = 0
			}
		}
		i, j = nextI, nextJ
	}
}
