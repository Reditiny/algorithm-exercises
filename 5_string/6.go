package main

import "strings"

/**
给定三个字符串 str、 from 和 to, 把 str 中所有 from 的子串全部替换成 to 字符串
对连续出现 from 的部分要求只替换成一个 to 字符串， 返回最终的结果字符
*/

func replaceStr(str, from, to string) string {

	temp := strings.Builder{}

	n := len(str)
	m := len(from)

	for i := 0; i < n; i++ {
		if i <= n-m && str[i:i+m] == from {
			temp.WriteByte('\000')
			i += m - 1
		} else {
			temp.WriteByte(str[i])
		}
	}

	ans := strings.Builder{}
	tempStr := temp.String()

	for i := 0; i < len(tempStr); i++ {
		if tempStr[i] == '\000' {
			if i == 0 || tempStr[i-1] != '\000' {
				ans.WriteString(to)
			}
		} else {
			ans.WriteByte(tempStr[i])
		}
	}

	return ans.String()
}
