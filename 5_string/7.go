package main

import (
	"strconv"
	"strings"
)

/**
给定一个字符串 str, 返回 str 的统计字符串
例如"aaabbadd"的统计字符串为"a3b2a1d2"
*/

func getCharCount(str string) string {
	n := len(str)
	if n == 0 {
		return ""
	}

	ans := strings.Builder{}

	lastChar := str[0]
	count := 1
	for i := 1; i < n; i++ {
		if str[i] == lastChar {
			count++
		} else {
			ans.WriteByte(lastChar)
			ans.WriteString(strconv.Itoa(count))
			lastChar = str[i]
			count = 1
		}
	}

	ans.WriteByte(lastChar)
	ans.WriteString(strconv.Itoa(count))
	return ans.String()
}
