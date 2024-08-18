package main

import (
	"strings"
)

/**
给定一个字符串 str 和一个整数 k 如果 str 中正好有连续的 k 个 'O' 字符出现时
把 k 个连续的 'O' 字符去除， 返回处理后的字符串
*/

func removeKZero(str string, k int) string {
	n := len(str)
	if n <= 0 || k <= 0 || k > n {
		return str
	}
	ans := strings.Builder{}
	newStr := "a" + str + "a"
	n += 2
	i, j := 0, k+1
	for j < n {
		if newStr[i] != '0' && newStr[j] != '0' {
			allZero := true
			for v := i + 1; v < j; v++ {
				if newStr[v] != '0' {
					allZero = false
				}
			}
			if !allZero {
				ans.WriteString(newStr[i:j])
			} else {
				ans.WriteByte(newStr[i])
			}
			i = j
			j = i + k + 1
		} else {
			ans.WriteByte(newStr[i])
			i += 1
			j += 1
		}
	}

	if i < n {
		ans.WriteString(newStr[i:])
	}

	return ans.String()[1 : ans.Len()-1]
}
