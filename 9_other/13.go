package main

/**
给定一个路径数组 paths 表示一张图, 只有一个根节点 paths[i] == i 且无环
实现一个函数返回，离根节点距离的统计数组
*/

func statistic(paths []int) {

	n := len(paths)

	var distance func(index int) int
	distance = func(index int) int {
		if paths[index] >= n {
			return paths[index] / n
		}
		if paths[index] == index {
			return 0
		}
		parentDis := distance(paths[index])
		paths[index] += (parentDis + 1) * n
		return 1 + parentDis
	}

	// 1. 原地将数组变为距离数组，即 paths[i] 为 i 到根节点距离
	for i := 0; i < n; i++ {
		// paths[i] == i 表示根节点，paths[i] < 0 表示已知距离
		if paths[i] != i && paths[i] >= 0 {
			pre := i
			cur := paths[i]
			// 记录本次起点
			paths[i] = -1
			// 迭代寻找根节点或已知距离的点，同时记录下回头的路
			for paths[cur] != cur && paths[cur] >= 0 {
				next := paths[cur]
				paths[cur] = pre
				pre = cur
				cur = next
			}
			var val int
			if paths[cur] == cur {
				val = -1
			} else {
				val = paths[cur] - 1
			}
			// 回头依次标记距离
			cur = pre
			for paths[cur] != -1 {
				next := paths[cur]
				paths[cur] = val
				val--
				cur = next
			}
			paths[cur] = val
		}
	}

	for i := 0; i < n; i++ {
		if paths[i] == i {
			paths[i] = 0
		} else {
			paths[i] = -paths[i]
		}
	}

	// 2. 原地将距离数组转为距离统计数组
	for i := 0; i < n; i++ {
		paths[paths[i]%n] += n
	}

	for i := 0; i < n; i++ {
		paths[i] /= n
	}
}
