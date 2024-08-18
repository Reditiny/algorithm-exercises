package main

/**
给定一个字符串 str, 返回 str 的最长无重复字符子串的长度
*/

func maxUniqueStrLength(str string) string {
	n := len(str)
	if n == 0 {
		return ""
	}
	charCount := make(map[byte]int)
	j := 0
	max := 0
	var ans = ""
	for i := range str {
		charCount[str[i]]++
		for charCount[str[i]] > 1 {
			charCount[str[j]]--
			j++
		}
		if i-j+1 > max {
			max = i - j + 1
			ans = str[j : i+1]
		}
	}
	return ans
}
