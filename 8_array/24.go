package main

import ds "algorithm-exercises/0_data_structure"

/**
用一个整型矩阵 matrix 表示一个网络，1 代表有路，0 代表无路，每一个位置只要不越界，都有上下左右 4 个方向，
求从最左上角到最右下角的最短通路值
*/

var move = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

type point struct {
	i, j int
}

func minPathValue(matrix [][]int) int {
	n := len(matrix)
	if n == 0 {
		return -1
	}
	m := len(matrix[0])
	if m == 0 {
		return -1
	}

	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}
	if matrix[0][0] == 0 {
		return -1
	}
	queue := ds.NewQueue[point](10)
	queue.EnQueue(point{0, 0})
	visited[0][0] = true
	ans := 1
	found := false
	// BFS
	for !queue.Empty() {
		// 当前层的所有节点
		curSize := queue.Size()
		for k := 0; k < curSize; k++ {
			cur := queue.Front()
			queue.DeQueue()
			if cur.i == n-1 && cur.j == m-1 {
				found = true
				break
			}
			// 上下左右四个方向
			for s := 0; s < 4; s++ {
				nextI := cur.i + move[s][0]
				nextJ := cur.j + move[s][1]
				if nextI >= 0 && nextI < n && nextJ >= 0 && nextJ < m && matrix[nextI][nextJ] == 1 && !visited[nextI][nextJ] {
					visited[nextI][nextJ] = true
					queue.EnQueue(point{nextI, nextJ})
				}
			}
		}
		if found {
			break
		}
		ans++
	}
	if found {
		return ans
	}
	return -1

}
