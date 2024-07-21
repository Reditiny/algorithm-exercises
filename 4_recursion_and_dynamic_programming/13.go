package main

/**
给定一个只由 0、1、&、｜和 ^ 五种字符组成的字符串 express, 再给定一个布尔值 desired。
返回 express 能有多少种组合方式，可以达到 desired 的结果。
*/

func desiredCount(express string, desired bool) int {
	if !valid(express) {
		return 0
	}

	n := len(express)
	count := (n + 1) / 2

	dpTrue := make([][]int, count)
	for i := 0; i < count; i++ {
		dpTrue[i] = make([]int, count)
	}
	dpFalse := make([][]int, count)
	for i := 0; i < count; i++ {
		dpFalse[i] = make([]int, count)
	}

	// dpTrue[i][j] 从第 i 到第 j 个字符表达式为 true 的个数
	// dpFalse[i][j] 从第 i 到第 j 个字符表达式为 false 的个数
	for i := 0; i < n; i += 2 {
		if express[i] == '1' {
			dpTrue[i/2][i/2] = 1
		} else {
			dpFalse[i/2][i/2] = 1
		}
	}

	for i := 0; i < count; i++ {
		for j := i + 1; j < count; j++ {
			for k := 2*i + 1; k < 2*j; k++ {
				switch express[k] {
				case '&':
					dpTrue[i][j] = dpTrue[i][k/2] * dpTrue[k/2+1][j]
					dpFalse[i][j] = dpFalse[i][k/2]*dpFalse[k/2+1][j] + dpTrue[i][k/2]*dpFalse[k/2+1][j] + dpFalse[i][k/2]*dpTrue[k/2+1][j]
				case '|':
					dpTrue[i][j] = dpTrue[i][k/2]*dpTrue[k/2+1][j] + dpTrue[i][k/2]*dpFalse[k/2+1][j] + dpFalse[i][k/2]*dpTrue[k/2+1][j]
					dpFalse[i][j] = dpFalse[i][k/2] * dpFalse[k/2+1][j]
				case '^':
					dpTrue[i][j] = dpTrue[i][k/2]*dpFalse[k/2+1][j] + dpFalse[i][k/2]*dpTrue[k/2+1][j]
					dpFalse[i][j] = dpTrue[i][k/2]*dpTrue[k/2+1][j] + dpFalse[i][k/2]*dpFalse[k/2+1][j]
				}
			}
		}
	}

	if desired {
		return dpTrue[0][count-1]
	}

	return dpFalse[0][count-1]
}

func valid(express string) bool {
	if len(express)%2 == 0 {
		return false
	}
	for i := 0; i < len(express); i += 2 {
		if express[i] != '1' && express[i] != '0' {
			return false
		}
	}
	for i := 1; i < len(express); i += 2 {
		if express[i] != '&' && express[i] != '|' && express[i] != '^' {
			return false
		}
	}
	return true
}
