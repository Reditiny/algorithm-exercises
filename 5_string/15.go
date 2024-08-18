package main

/**
给定一个整数 N 求由"0"字符与T"字符组成的长度为N 的所有字符串中，
满足"0"字符的左边必有"1"字符的字符串数量。
*/

// N = 	 1，2，3，4，5，6， 7， 8
// ans = 1，2，3，5，8，13，21，34
func getNum(N int) int {
	if N < 1 {
		return 0
	}
	if N == 1 {
		return 1
	}
	if N == 2 {
		return 2
	}
	pre := 1
	cur := 2
	for i := 3; i <= N; i++ {
		cur, pre = cur+pre, cur
	}
	return cur
}
