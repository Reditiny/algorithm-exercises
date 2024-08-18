package main

import (
	ds "algorithm-exercises/0_data_structure"
	"math"
)

/**
给定一个字符串数组 strs, 再给定两个字符串 str1 和 str2, 返回在 strs 中 str1 与 str2 的
最小距离， 如果 str1 或 str2 为null, 或不在 strs 中， 返回-1
*/

func minDistance(strs []string, str1, str2 string) int {
	if len(strs) == 0 || str1 == "" || str2 == "" {
		return -1
	}

	if str1 == str2 {
		return 0
	}

	lastStr1Index := -1
	lastStr2Index := -1
	min := math.MaxInt
	for i := range strs {
		if strs[i] == str1 {
			if lastStr2Index != -1 {
				min = ds.Min(min, i-lastStr2Index)
			}
			lastStr1Index = i
		} else if strs[i] == str2 {
			if lastStr1Index != -1 {
				min = ds.Min(min, i-lastStr1Index)
			}
			lastStr2Index = i
		}
	}
	if min == math.MaxInt {
		return -1
	}
	return min
}
