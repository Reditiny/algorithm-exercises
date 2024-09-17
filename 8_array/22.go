package main

/**
给定一个数组 arr, 返回不包含本位置值的累乘数组
不使用除法
时间复杂度为 O(N)
除需要返回的结果数组外， 额外空间复杂度为0(1)
*/

func multiplyArr(arr []int) []int {
	n := len(arr)
	if n == 0 {
		return nil
	}
	// cur 表示 arr[0...i-1] 的累乘
	ans := make([]int, n)
	cur := 1
	for i := 0; i < n; i++ {
		// ans[i] 表示 arr[0...i-1] 的累乘
		ans[i] = cur
		cur *= arr[i]
	}
	// cur 表示 arr[i+1...n-1] 的累乘
	cur = 1
	for i := n - 1; i >= 0; i-- {
		// ans[i] 表示 arr[0...i-1] 的累乘 * arr[i+1...n-1] 的累乘
		ans[i] *= cur
		cur *= arr[i]
	}
	return ans
}
