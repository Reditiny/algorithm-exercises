package main

import ds "algorithm-exercises/0_data_structure"

/*
	求最大子矩阵的大小
	给定一个整型矩阵 map, 其中的值只有 0 和 1 两种
	求其中全是 1 的所有矩形区域中，最大的矩形的面积
*/

func maxMatrix(matrix [][]int) int {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	n := len(matrix)
	m := len(matrix[0])

	ans := 0
	height := make([]int, m)
	for i := 0; i < n; i++ {
		// 依次以每一行为底，去当前可得到的最大面积
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				height[j] = 0
			} else {
				height[j]++
			}
		}
		ans = ds.Max(ans, getArea(height))
	}

	return ans
}

func getArea(height []int) int {
	n := len(height)
	// 单调栈求每个高度左右第一个比自身低的索引
	leftFirstGreater := make([]int, n)
	stack := ds.NewStack[int](n)
	for i := 0; i < n; i++ {
		for !stack.Empty() && height[i] < height[stack.Top()] {
			stack.Pop()
		}
		if stack.Empty() {
			leftFirstGreater[i] = -1
		} else {
			leftFirstGreater[i] = stack.Top()
		}
		stack.Push(i)
	}
	stack.Clear()
	rightFirstGreater := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		for !stack.Empty() && height[i] < height[stack.Top()] {
			stack.Pop()
		}
		if stack.Empty() {
			rightFirstGreater[i] = -1
		} else {
			rightFirstGreater[i] = stack.Top()
		}
		stack.Push(i)
	}
	// 依次以自身为高，求最大面积
	ans := 0
	for i := 0; i < n; i++ {
		ans = ds.Max(ans, height[i]*(rightFirstGreater[i]-leftFirstGreater[i]-1))
	}
	return ans
}
