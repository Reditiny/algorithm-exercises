package main

import (
	"sort"
	"strings"
)

/**
给定一个字符串类型的数组 strs, 请找到一种拼接顺序， 使得将所有的字符串拼接起来
组成的大写字符串是所有可能性中字典顺序最小的， 并返回这个大写字符串
*/

func lowestString(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	sort.Slice(strs, func(i, j int) bool {
		return strs[i]+strs[j] < strs[j]+strs[i]
	})

	builder := strings.Builder{}
	for i := range strs {
		for j := range strs[i] {
			builder.WriteByte(UpLetter(strs[i][j]))
		}
	}
	return builder.String()
}

func UpLetter(b byte) byte {
	if b >= 'a' && b <= 'z' {
		return b - 32
	}
	return b
}
