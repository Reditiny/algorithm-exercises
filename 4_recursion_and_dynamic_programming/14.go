package main

import ds "algorithm-exercises/0_data_structure"

/**
给定一个整型数组 arr, 代表数值不同的纸牌排成一条线。玩家 A 和玩家 B 依次拿走每张纸牌
玩家 A 先拿，玩家 B 后拿，每次只能拿走最左或最右的纸牌，请返回最后获胜者的分数
*/

func winnerScore(arr []int) int {
	n := len(arr)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return arr[1]
	}

	dpFirst := make([][]int, n)
	for i := 0; i < n; i++ {
		dpFirst[i] = make([]int, n)
	}
	dpSecond := make([][]int, n)
	for i := 0; i < n; i++ {
		dpSecond[i] = make([]int, n)
	}

	for j := 0; j < n; j++ {
		dpFirst[j][j] = arr[j]
		for i := j - 1; i >= 0; i-- {
			dpFirst[i][j] = ds.Max(arr[i]+dpSecond[i+1][j], arr[j]+dpSecond[i][j-1])
			dpSecond[i][j] = ds.Min(dpFirst[i+1][j], dpFirst[i][j-1])
		}
	}

	return ds.Max(dpFirst[0][n-1], dpSecond[0][n-1])
}
