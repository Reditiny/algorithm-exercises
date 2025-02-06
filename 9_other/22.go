package main

/**
一群孩子做游戏，现在请你根据游戏得分来发糖果，要求如下：
1. 每个孩子不管得分多少，起码分到 1 个糖果。
2. 任意两个相邻的孩子之间，得分较多的孩子必须拿多一些的糖果，得分相同则拿同样的糖果。
给定一个数组 arr 代表得分数组，请返回最少需要多少糖果
*/

func dispenseCandy(score []int) int {
	n := len(score)
	if n <= 1 {
		return n
	}
	candy := make([]int, len(score))
	for i := 1; i < len(score); i++ {
		if score[i-1] < score[i] && candy[i-1] >= candy[i] {
			candy[i] = candy[i-1] + 1
		} else if score[i-1] == score[i] && candy[i-1] != candy[i] {
			candy[i] = candy[i-1]
		}
	}

	for i := len(score) - 2; i >= 0; i-- {
		if score[i] > score[i+1] && candy[i] <= candy[i+1] {
			candy[i] = candy[i+1] + 1
		} else if score[i] == score[i+1] && candy[i] != candy[i+1] {
			if candy[i] > candy[i+1] {
				candy[i+1] = candy[i]
			} else {
				candy[i] = candy[i+1]
			}
		}
	}

	ans := len(score)
	for i := range candy {
		ans += candy[i]
	}

	return ans
}
