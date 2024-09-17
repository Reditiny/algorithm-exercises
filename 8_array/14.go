package main

/**
给定一个长度为 N 的整型数组 arr, 其中有 N 个互不相等的自然数 1〜N
请实现 arr 的排序，但是不要把下标 0〜N-1 位置上的数通过直接赋值的方式替换成 1〜N
时间复杂度为O(N), 额外空间复杂度为O(1)
*/

func sortEArr(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		// arr[i] 最终值应该是 i+1
		// 若当前 arr[i] != i+1，则将 arr[i] 放到自己应该在的位置上，即 arr[i]-1
		for arr[i] != i+1 {
			arr[i], arr[arr[i]-1] = arr[arr[i]-1], arr[i]
		}
	}
}
