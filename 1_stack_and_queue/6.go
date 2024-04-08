package main

import "fmt"

/*
	用栈来求解汉诺塔问题
	汉诺塔问题比较经典， 这里修改一下游戏规则：现在限制不能从最左侧的塔直接移动到最右侧，
	也不能从最右侧直接移动到最左侧， 而是必须经过中间。求当塔有 N 层的时候，
	打印最优移动过程和最优移动总步数
*/

var move = [2][2]string{
	{"Move %d from left to mid", "Move %d from mid to right"},
	{"Move %d from right to mid", "Move %d from mid to left"},
}

func hanoi(n int, dir int) int {
	if n == 1 {
		fmt.Println(fmt.Sprintf(move[dir][0], n))
		fmt.Println(fmt.Sprintf(move[dir][1], n))
		return 2
	}
	// 递归过程 n 个塔
	// 1. n-1 个塔 左->中->右
	// 2. 1 个塔   左->中
	// 3. n-1 个塔 右->中->左
	// 4. 1 个塔   中->右
	// 5. n-1 个塔 左->中->右
	ans := 0
	ans += hanoi(n-1, dir)
	ans += 1
	fmt.Println(fmt.Sprintf(move[dir][0], n))
	// 注意方向的改变，通过 1-dir 实现 0 和 1 的变化
	ans += hanoi(n-1, 1-dir)
	ans += 1
	fmt.Println(fmt.Sprintf(move[dir][1], n))
	ans += hanoi(n-1, dir)
	return ans
}
