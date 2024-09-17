package main

/**
给定一个长度不小于 2 的数组 arr, 实现一个函数调整 arr,
要么让所有的偶数下标都是偶数， 要么让所有的奇数下标都是奇数
*/

func modifyArr(arr []int) {
	n := len(arr)
	i, j := 0, 1
	for i < n && j < n {
		// 找到偶数下标不是偶数的位置
		for i < n && arr[i]%2 == 0 {
			i += 2
		}
		// 找到奇数下标不是奇数的位置
		for j < n && arr[j]%2 == 1 {
			j += 2
		}
		// 尝试交换
		if i < n && j < n {
			arr[i], arr[j] = arr[j], arr[i]
			i += 2
			j += 2
		}
	}
}
