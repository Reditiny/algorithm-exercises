package main

import ds "algorithm-exercises/0_data_structure"

/**
给定两个有序数组 arr1 和 arr2, 再给定一个整数 k
返回来自 arr1 和 arr2 的两个数相加和最大的前 k 个
两个数必须分别来自两个数组
*/

func getKMaxSum(arr1 []int, arr2 []int, k int) []int {
	n, m := len(arr1), len(arr2)
	if n == 0 || m == 0 || k <= 0 {
		return nil
	}
	// 优先级队列存放两个下标, 优先级为两个数的和
	pq := ds.NewPriorityQueue[[2]int](func(a, b [2]int) int {
		return (arr1[a[0]] + arr2[a[1]]) - (arr1[b[0]] + arr2[b[1]])
	})

	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}
	// arr1[i] + arr2[j] 一定比 arr1[i] + arr2[j-1] 和 arr1[i-1] + arr2[j] 大
	// 因此从右下角开始向左上角遍历 arr1[n-1] 和 arr2[m-1]
	pq.Push([2]int([]int{n - 1, m - 1}))
	visited[n-1][m-1] = true
	var ans []int
	for i := 0; i < k; i++ {
		// 取出最大和的下标后, 将其左边和上边的下标放入优先级队列
		cur := pq.Top()
		pq.Pop()
		ans = append(ans, arr1[cur[0]]+arr2[cur[1]])
		if cur[0] > 0 && !visited[cur[0]-1][cur[1]] {
			pq.Push([2]int([]int{cur[0] - 1, cur[1]}))
			visited[cur[0]-1][cur[1]] = true
		}
		if cur[1] > 0 && !visited[cur[0]][cur[1]-1] {
			pq.Push([2]int([]int{cur[0], cur[1] - 1}))
			visited[cur[0]][cur[1]-1] = true
		}
	}
	return ans
}
