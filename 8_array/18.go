package main

/**
在数组中找到一个局部最小的位置
时间复杂度为 O(logN)、空间复杂度为 O(1)
*/

func findMinIndex(arr []int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}
	if n == 1 {
		return 0
	}
	if arr[0] < arr[1] {
		return 0
	}
	if arr[n-1] < arr[n-2] {
		return n - 1
	}
	l, r := 1, n-2
	for l < r {
		mid := l + (r-l)/2
		if arr[mid] > arr[mid-1] {
			r = mid - 1
		} else if arr[mid] > arr[mid+1] {
			l = mid + 1
		} else {
			return mid
		}
	}
	return l
}
