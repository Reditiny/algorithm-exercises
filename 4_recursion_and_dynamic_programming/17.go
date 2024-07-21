package main

/**
在 NxN 的棋盘上要摆 N 个皇后， 要求任何两个皇后不同行、不同列，也不在同一条斜线上。
给定一个整数 N 返回皇后的摆法有多少种
*/

func nQueen(n int) int {
	ans := 0

	col := make([]bool, n)
	diagonal1 := make([]bool, 2*n-1)
	diagonal2 := make([]bool, 2*n-1)

	var queen func(level int)
	queen = func(level int) {
		if level == n {
			ans++
		}
		for j := 0; j < n; j++ {
			if col[j] || diagonal1[level+j] || diagonal2[n-1-level+j] {
				continue
			}
			col[j] = true
			diagonal1[level+j] = true
			diagonal2[n-1-level+j] = true
			queen(level + 1)
			col[j] = false
			diagonal1[level+j] = false
			diagonal2[n-1-level+j] = false
		}
	}
	queen(0)
	return ans
}
