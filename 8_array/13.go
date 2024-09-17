package main

/**
计算数组的小和
对于数组 arr，arr[i] 左边小于等于 arr[i] 元素累加起来记为 s[i]，所有 s[i] 累加起来就是数组的小和
*/

func smallSum(arr []int) int {

	ans := 0
	// 利用 mergeSort 计算小和
	// 左右两个有序数组，计算左边数组中比右边数组中的元素小的元素的和
	var mergeSort func(arr []int, start, end int)
	mergeSort = func(arr []int, start, end int) {
		if start >= end {
			return
		}
		mid := start + (end-start)/2
		mergeSort(arr, start, mid)
		mergeSort(arr, mid+1, end)

		arr1 := make([]int, mid-start+1)
		copy(arr1, arr[start:mid+1])
		arr2 := make([]int, end-mid)
		copy(arr2, arr[mid+1:end+1])
		i, j := 0, 0
		idx := start
		for i < len(arr1) && j < len(arr2) {
			if arr1[i] <= arr2[j] {
				arr[idx] = arr1[i]
				i++
				// 如果 arr1[i] <= arr2[j]，则 arr1[i] 之后的元素都比 arr2[j] 大
				ans += arr1[i] * (len(arr2) - j + 1)
			} else {
				arr[idx] = arr2[j]
				j++
			}
			idx++
		}

		for i < len(arr1) {
			arr[idx] = arr1[i]
			i++
			idx++
		}

		for j < len(arr2) {
			arr[idx] = arr2[j]
			j++
			idx++
		}
	}

	return ans
}
