package main

/**
如果一个字符串 str, 把字符串 str 前面任意的部分挪到后面形成的字符串叫作 str 的旋转词
给定两个字符串 str1 和 str2, 请判断 str1 和 str2 是否互为旋转词
*/

func isRotateWords(str1, str2 string) bool {
	return false
}

func KMP(str1, str2 string) bool {
	return false
}

func getNextArray(str string) []int {
	next := make([]int, len(str))
	next[0] = -1
	next[1] = 0
	i := 2
	cn := 0
	for i < len(str) {
		if str[i-1] == str[cn] {
			cn++
			next[i] = cn
			i++
		} else if cn > 0 {
			cn = next[cn]
		} else {
			next[i] = 0
			i++
		}
	}
	return next
}
