package main

/**
升序数组 arr 可能经过一次旋转处理，也可能没有，且 arr 可能存在重复的数。
例如，有序数组［1,2,3,4,5,6,7] 可以旋转处理成［4,5,6,7,1,2,3］
给定一个可能旋转过的有序数组 arr, 返回 arr 中的最小值
*/

func findMin(arr []int) int {
	left, right := 0, len(arr)-1
	// 虽然不完全有序，但是也可以使用二分
	for left < right {
		// 排除重复数的干扰
		for left < right && arr[left] == arr[right] {
			left++
		}
		// 说明从 left 到 right 是有序的，left 为最小索引
		if arr[left] < arr[right] {
			return arr[left]
		}

		mid := left + (right-left)/2
		if arr[left] > arr[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return arr[left]
}
