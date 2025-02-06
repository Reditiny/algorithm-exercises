package main

import (
	"fmt"
	"math/rand"
)

/**
给定一个长度为 N 且没有重复元素的数组，实现函数等概率随机打印其中 M 个数
1. 相同的数不要重复打印
2. 时间复杂度为O(M)，额外空间复杂度为O(1)
3. 可以改变 arr 数组
*/

func randomMFromN(arr []int, m int) {
	cur := len(arr)
	for i := 0; i < m; i++ {
		// [0,1,2,...,cur-1] 随机打印一个
		idx := rand.Intn(cur)
		fmt.Println(arr[idx])
		// 打印完 arr[idx] 后将之放到末尾
		arr[idx], arr[cur-1] = arr[cur-1], arr[idx]
		cur--
	}
}
