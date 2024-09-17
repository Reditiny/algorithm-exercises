package main

import "fmt"

/**
给定一个整型矩阵 matrix, 请按照转圈的方式打印它。
例如：
	 1  2  3  4
	 5  6  7  8
	 9 10 11 12
	13 14 15 16

打印结果为： 1, 2, 3, 4, 8, 12, 16, 15, 14, 13, 9, 5, 6, 7, 11, 10
*/

func printMatrix(matrix [][]int) {

	moves := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	n := len(matrix)
	if n == 0 {
		return
	}
	m := len(matrix[0])
	if m == 0 {
		return
	}

	var printM func(sI, sJ, eI, eJ int)
	printM = func(sI, sJ, eI, eJ int) {
		if sI > eI || sJ > eJ {
			return
		}

		if sI == eI {
			for j := sJ; j <= eJ; j++ {
				fmt.Print(matrix[sI][j], " ")
			}
		} else if sJ == eJ {
			for i := sI; i <= eI; i++ {
				fmt.Print(matrix[i][sJ], " ")
			}
		} else {
			curI, curJ := sI, sJ
			fmt.Print(matrix[curI][curJ], " ")
			for _, move := range moves {
				for {
					nextI, nextJ := curI+move[0], curJ+move[1]
					if nextI < sI || nextI > eI || nextJ < sJ || nextJ > eJ {
						break
					}
					if nextI == sI && nextJ == sJ {
						break
					}
					curI, curJ = nextI, nextJ
					fmt.Print(matrix[curI][curJ], " ")
				}
			}
			printM(sI+1, sJ+1, eI-1, eJ-1)
		}
	}
	printM(0, 0, n-1, m-1)
}
