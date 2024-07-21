package main

import ds "algorithm-exercises/0_data_structure"

/**
给定数组 arr, arr[i]=k 代表可以从位置 i 向右跳 1〜k 个距离
比如，arr[2]=3, 代表从位置 2 可以跳到位置 3、位置 4 或位置 5。
如果从位置 0 出发， 返回最少跳几次能跳到 arr 最后的位置上
*/

func jumpCount(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	// 当前跳数，当前能到达的最大距离，再跳一次能到达的最大距离
	count, curMaxDis, nextMaxDis := 0, 0, 0
	for i := 0; i < len(arr); i++ {
		if curMaxDis < i {
			// 当前到不了 i，需要再跳一次，并更新当前能到达的最大距离
			count++
			curMaxDis = nextMaxDis
		}
		nextMaxDis = ds.Max(nextMaxDis, i+arr[i])
	}
	return count
}
