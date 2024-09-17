package main

/**
给定一个无序整型数组 arr, 找到数组中未出现的最小正整数
时间复杂度 O(N), 空间复杂度为O(1)
*/

func minNumber(arr []int) int {
	n := len(arr)
	for i := 0; i < n; i++ {
		// 如果 arr[i] 在 [1...n] 区间内， 则需要将 arr[i] 放到 arr[arr[i]-1] 的位置上
		for arr[i] >= 1 && arr[i] <= n {
			// 已在正确位置上，或无需交换
			if arr[i] == i+1 || arr[arr[i]-1] == arr[i] {
				break
			}
			arr[i], arr[arr[i]-1] = arr[arr[i]-1], arr[i]
		}
		if arr[i] != i+1 {
			arr[i] = -1
		}
	}
	// 找到第一个不在正确位置上的元素
	for i := 0; i < n; i++ {
		if arr[i] != i+1 {
			return i + 1
		}
	}
	// 1...n 都在正确位置上，说明第一个缺失的是 n+1
	return n + 1
}

func main() {
	arr := []int{3, 1, 5, 5, 5}
	println(minNumber(arr))
}
