package main

import "fmt"

/**
给定一个整型数组 arr, 打印其中出现次数大于一半的数

进阶，再给定一个整数 K，打印所有出现次数大于 N/K 的数
*/

func printHalfMajor(arr []int) {
	n := len(arr)
	if n == 0 {
		return
	}
	// 找出一个候选数
	var cur int
	count := 0
	for i := 0; i < n; i++ {
		if count == 0 {
			cur = arr[i]
		} else {
			if arr[i] == cur {
				count++
			} else {
				count--
			}
		}
	}

	count = 0
	for i := 0; i < n; i++ {
		if arr[i] == cur {
			count++
		}
	}

	if count > n/2 {
		fmt.Println(cur)
	}
}

func printKMajor(arr []int, k int) {
	n := len(arr)
	if n == 0 {
		return
	}
	// 找出 k 个候选数
	count := make([]int, k)
	cur := make([]int, k)
	for i := 0; i < n; i++ {
		found := false
		for j := 0; j < k; j++ {
			if count[j] != 0 && cur[j] == arr[i] {
				count[j]++
				found = true
			}
		}
		if found {
			continue
		}
		for j := 0; j < k; j++ {
			if count[j] == 0 {
				count[j] = 1
				cur[j] = arr[i]
				found = true
			}
		}
		if found {
			continue
		}
		for j := 0; j < k; j++ {
			count[j]--
		}
	}

	for i := 0; i < k; i++ {
		count[i] = 0
	}

	for i := 0; i < n; i++ {
		for j := 0; j < k; j++ {
			if arr[i] == cur[j] {
				count[j]++
				break
			}
		}
	}

	for i := 0; i < k; i++ {
		if count[i] < n/k {
			fmt.Println(cur[i])
		}
	}
}
