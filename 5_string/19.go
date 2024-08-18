package main

import ds "algorithm-exercises/0_data_structure"

/**
给定字符串 str1 和 str2, 求 str1 的子串中含有 str2 所有字符的最小子串长度
*/

func minContainStrLength(str1, str2 string) int {
	if len(str1) == 0 || len(str2) == 0 || len(str1) < len(str2) {
		return 0
	}
	// 记录 str2 中每个字符出现的次数
	charCount := make(map[byte]int)
	for i := range str2 {
		charCount[str2[i]]--
	}

	j := 0
	ans := len(str1) + 1
	for i := range str1 {
		charCount[str1[i]]++
		for allGEZero(charCount) {
			ans = ds.Min(ans, i-j+1)
			charCount[str1[j]]--
			j++
		}
	}
	if ans == len(str1)+1 {
		return 0
	}
	return ans
}

func allGEZero(charCount map[byte]int) bool {
	for _, v := range charCount {
		if v < 0 {
			return false
		}
	}
	return true
}
