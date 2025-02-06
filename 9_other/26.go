package main

import ds "algorithm-exercises/0_data_structure"

/**
给定两个有序数组长度都为 N 求两个数组中所有数的中位数
时间复杂度 O(N) 空间复杂度 O(1)
*/

func GetMedian(arr1 []int, arr2 []int) int {

	l1, r1 := 0, len(arr1)
	l2, r2 := 0, len(arr2)
	var mid1, mid2 int
	offset := 0

	for l1 < r1 {
		mid1 = l1 + (r1-l1)/2
		mid2 = l2 + (r2-l2)/2

		// 元素个数为奇数， 则 offset 为 0
		// 元素个数为偶数， 则 offset 为 1
		if (r1-l1)%2 == 1 {
			offset = 0
		} else {
			offset = 1
		}

		if arr1[mid1] > arr2[mid2] {
			// 比 arr1[mid1] 小的数至少有一半，因此下次迭代在 [l1,mid1]中
			// 比 arr2[mid2] 大的数至少有一半，因此下次迭代在 [mid2,r2]中
			// 因为 mid 的计算是向下取整，因此在偶数个时 [mid,r] 的区间需要偏移成 [mid+1,r]
			r1 = mid1
			l2 = mid2 + offset
		} else if arr1[mid1] < arr2[mid2] {
			l1 = mid1 + offset
			r2 = mid2
		} else {
			return arr1[mid1]
		}
	}
	return ds.Min(arr1[l1], arr2[l2])
}
