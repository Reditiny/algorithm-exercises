package main

/**
给定两个有序数组 arr1 和 arr2, 再给定一个整数 k 返回所有的数中第 k 小的数
*/

func getKth(arr1 []int, arr2 []int, k int) int {

	var longs, shorts []int
	if len(arr1) > len(arr2) {
		longs = arr1
		shorts = arr2
	} else {
		longs = arr2
		shorts = arr1
	}

	lenLong, lenShort := len(longs), len(shorts)
	if k < 1 || k > lenLong+lenShort {
		return -1
	}

	if k <= lenShort {
		// k 比两个数据的长度都小时，第 k 小的数是两数组前 k 个数的中位数
		return GetMedian(shorts[0:k], longs[0:k])
	}

	if k > lenLong {
		// k 比两个数据的长度都大时
		if shorts[k-lenLong-1] >= longs[lenLong-1] {
			// short 中有 k-lenLong-1 个数比 shorts[k-lenLong-1] 小，且 longs 中有 lenLong 个数比 shorts[k-lenLong-1] 小
			// 因此一共有 k-lenLong-1+lenLong=k-1 个数比 shorts[k-lenLong-1] 小
			return shorts[k-lenLong-1]
		}
		if longs[k-lenShort-1] >= shorts[lenShort-1] {
			// long 中有 k-lenShort-1 个数比 longs[k-lenShort-1] 小，且 short 中有 lenShort 个数比 longs[k-lenShort-1] 小
			// 因此一共有 k-lenShort-1+lenShort=k-1 个数比 longs[k-lenShort-1] 小
			return longs[k-lenShort-1]
		}
		// 上述两种情况都不满足时，说明 shorts[k-lenLong-1] 和 longs[k-lenShort-1] 两个数都比第 k 个数小
		// 第 k 个数是 shorts[k-lenLong:] 和 longs[k-lenShort:] 的中位数
		return GetMedian(shorts[k-lenLong:], longs[k-lenShort:])
	}

	// lenShort < k <= lenLong，第 k 小的数不可能出现在 longs[0:k-lenShort] 中
	return GetMedian(shorts, longs[k-lenShort:])
}
