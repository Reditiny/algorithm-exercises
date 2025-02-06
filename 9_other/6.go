package main

import "fmt"

/**
一张纸条放在桌子上，下边向上方对折 1 次，压出折痕后展开，此时折痕是凹下去的， 即折痕突起的方向指向纸条的背面。
连续对折 2 次， 压出折痕后展开， 此时有三条折痕， 从上到下依次是下折痕、下折痕和上折痕
给定一个输入参数 N， 代表对折 N 次， 请从上到下打印所有折痕的方向
*/

const (
	Up   = "up"
	Down = "down"
)

// 1.    	 				down
// 2.			down 		down 			up
// 3.	down 	down 	up 	down 	down 	up 	up
// 第一次折痕是 down，后续再折会在每条折痕的两边产生 down 和 up
func printAllFolds(n int) {
	process(1, n, Down)
}

func process(cur, n int, act string) {
	if cur > n {
		return
	}
	process(cur+1, n, Down)
	fmt.Println(act)
	process(cur+1, n, Up)
}
