package main

import "fmt"

/**
给定一个整数 n，代表汉诺塔游戏中从小到大放置的 n 个圆盘， 假设开始时所有的圆盘都放在左边的柱子上
想按照汉诺塔游戏的要求把所有的圆盘都移到右边的柱子上。实现函数打印最优移动轨迹
*/

func hanoi(n int, from, mid, to string) {
	if n == 0 {
		return
	}
	if n == 1 {
		fmt.Printf("move from %s to %s\n", from, to)
		return
	}
	hanoi(n-1, from, to, mid)
	hanoi(1, from, mid, to)
	hanoi(n-1, mid, from, to)
}

/**
给定一个整型数组 air, 其中只含有 1、2 和 3, 代表所有圆盘目前的状态， 1 代表左柱，2 代表中柱， 3 代表右柱
如 arr=[3,3,2,1], 代表第 1 个圆盘在右柱上、第 2 个圆盘在右柱上、第 3 个圆盘在中柱上、第 4 个圆盘在左柱上
返回 arr 这种状态是最优移动轨迹中的第几个状态。如果 arr 代表的状态不是最优移动轨迹过程中出现的状态， 则返回 -1
*/

func step(arr []int) int {
	if len(arr) == 0 {
		return -1
	}
	return process(arr, len(arr)-1, 1, 2, 3)
}

// curLastIndex 当前递归轮次的最大圆盘
func process(arr []int, curLastIndex, from, mid, to int) int {
	if curLastIndex == -1 {
		return 0
	}
	if arr[curLastIndex] == mid {
		// 最大圆盘不可能出现在 mid 柱上
		return -1
	} else if arr[curLastIndex] == from {
		// 最大圆盘在 from 上，说明此时正处在将最大圆盘之上的圆盘从 from 移到 mid 的过程
		return process(arr, curLastIndex-1, from, to, mid)
	} else {
		// 最大圆盘在 to 上，说明至少已经移了 (2^n)-1 次，且最大圆盘之上的圆盘处在从 mid 移到 to 的过程
		res := process(arr, curLastIndex-1, mid, from, to)
		if res == -1 {
			return -1
		}
		return (1 << curLastIndex) + res
	}
}
