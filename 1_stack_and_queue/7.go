package main

import ds "algorithm-exercises/0_data_structure"

/*
	生成窗口最大值数组
	有一个整型数组 arr 长度为 n 和一个大小为 w 的窗口从数组的最左边滑到最右边,
	窗口每次向右边滑一个位置一共产生 n-w+1 个窗口的最大值
	请实现一个函数。
	• 输入： 整型数组 arr, 窗口大小为 w。
	• 输出： 一个长度为 n-w+1 的数组 res, res[i]表示每一种窗口状态下的最大值
*/

// [4,3,5,4,3,3,6,7]
func maxValueInWindow(arr []int, w int) []int {
	n := len(arr)
	ans := make([]int, 0)
	deque := ds.NewDeque[int](3)
	for i := 0; i < w; i++ {
		for !deque.Empty() && arr[i] > arr[deque.Back()] {
			deque.PopBack()
		}
		deque.PushBack(i)
	}

	for i := w; i < n; i++ {
		ans = append(ans, deque.Front())
		// 保持单调减
		for !deque.Empty() && arr[i] > arr[deque.Back()] {
			deque.PopBack()
		}
		deque.PushBack(i)
		// 窗口长度淘汰
		if deque.Front() == i-w {
			deque.PopFront()
		}
	}
	ans = append(ans, deque.Front())

	return ans
}
