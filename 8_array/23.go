package main

/**
给定一个有序数组 arr, 调整 arr 使得这个数组的左半部分没有重复元素且升序， 而不用保证右部分是否有序。

给定一个数组 arr, 其中只可能含有 0、1、2 三个值， 请实现 arr 的排序。
*/

func leftUnique(arr []int) {
	n := len(arr)
	if n <= 2 {
		return
	}
	// uIndex 表示 [0...uIndex] 区间内的元素都是唯一且有序的
	uIndex := 0
	for i := 1; i < n; i++ {
		// 如果当前元素和有序唯一区间的最后一个元素不相等，那么将当前元素加入有序唯一区间
		if arr[i] != arr[uIndex] {
			uIndex++
			arr[uIndex], arr[i] = arr[i], arr[uIndex]
		}
	}
}

func sort012(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}
	zeroToOneCount := 0
	twoToOneCount := 0
	i, j := 0, n-1
	// 快排思想，将 0 放在左边，2 放在右边
	for i < j {
		for i < j && arr[i] != 2 {
			// 将左边的 1 直接变成 0
			if arr[i] == 1 {
				zeroToOneCount++
				arr[i] = 0
			}
			i++
		}
		for i < j && arr[j] != 0 {
			// 将右边的 1 直接变成 2
			if arr[j] == 1 {
				twoToOneCount++
				arr[j] = 2
			}
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
	}
	// 还原上述直接被覆盖的 1
	for zeroToOneCount > 0 {
		if arr[i] == 0 {
			arr[i] = 1
			zeroToOneCount--
		}
		i--
	}

	for twoToOneCount > 0 {
		if arr[j] == 2 {
			arr[j] = 1
			twoToOneCount--
		}
		j++
	}
}
