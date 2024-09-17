package main

import "fmt"

/**
给定排序数组 arr 和整数 k, 不重复打印 arr 中所有相加和为 k 的不降序二元组

给定排序数组 arr 和整数 k, 不重复打印 arr 中所有相加和为 k 的不降序三元组
*/

func printAllUniqueBiPair(arr []int, k int, hook func()) {
	n := len(arr)
	if n == 0 {
		return
	}
	i, j := 0, n-1

	for i < j {
		if arr[i]+arr[j] < k {
			i++
		} else if arr[i]+arr[j] > k {
			j--
		} else {
			// 避免重复
			if !((i != 0 && arr[i] == arr[i-1]) || (j != n-1 && arr[j] == arr[j+1])) {
				fmt.Print("[")
				hook()
				fmt.Print(arr[i], ",", arr[j], "]")
				fmt.Print("\n")
			}
			i++
			j--
		}
	}
}

func printAllUniqueTrPair(arr []int, k int) {
	n := len(arr)
	if n == 0 {
		return
	}
	for i := 0; i < n; i++ {
		if i != 0 && arr[i] == arr[i-1] {
			continue
		}
		printAllUniqueBiPair(arr[i+1:], k-arr[i], func() {
			fmt.Print(arr[i], ",")
		})
	}
}
