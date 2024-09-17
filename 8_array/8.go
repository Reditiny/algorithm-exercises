package main

import (
	ds "algorithm-exercises/0_data_structure"
	"math"
)

/**
如果一个数组在排序之后， 每相邻两个数差的绝对值都为1， 则该数组为可整合数组
例如， [5,3,4,6,2] 排序之后为 [2,3,4,5,6]，这个数组为可整合数组
给定一个整型数组 arr， 请返回其中最大可整合子数组的长度
*/

func getLIL(arr []int) int {
	n := len(arr)
	if n <= 1 {
		return n
	}

	m := make(map[int]struct{})

	ans := 0

	for i := 0; i < n; i++ {
		curMin := math.MaxInt
		curMax := math.MinInt
		for j := i; j < n; j++ {
			// arr[i:j+1] 有重复则不可整合
			if _, ok := m[arr[j]]; ok {
				break
			}
			m[arr[j]] = struct{}{}
			curMin = ds.Min(curMin, arr[j])
			curMax = ds.Max(curMax, arr[j])
			// 最大与最小之差等于数组长度-1，则可整合
			if curMax-curMin == j-i && j-i+1 > ans {
				ans = j - i + 1
			}
		}
		m = make(map[int]struct{})
	}

	return ans
}
