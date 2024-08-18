package main

/**
新类型字符的定义如下：
1. 新类型字符是长度为 1 或者 2 的字符串。
2. 表现形式可以仅是小写字母，如"e"；\;大写字母+小写字母 例如"Ab";大写字母+大写字母，如"DC"
现在给定一个字符串 str, str 一定是若干新类型字符正确组合的结果。比如 "kaCCBi"，
由新类型字符 "k"、"a"、"CC"和"Bi"拼成。再给定一个整数 k 代表 str 中的位置。请返回 k 位置的新类型字符
*/

func newCharAt(str string, k int) string {
	if len(str) == 0 || k < 0 || k >= len(str) {
		return ""
	}
	// 从 k - 1 位置开始向左计算连续的大写字母数量
	count := 0
	for i := k - 1; i >= 0 && (str[i] >= 'A' && str[i] <= 'Z'); i-- {
		count++
	}
	if count%2 == 1 {
		return str[k-1 : k+1]
	} else {
		if str[k] >= 'A' && str[k] <= 'Z' {
			return str[k : k+2]
		} else {
			return str[k : k+1]
		}
	}
}
