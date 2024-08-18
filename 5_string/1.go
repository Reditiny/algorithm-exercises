package main

/**
请实现函数判断两个字符串是否互为变形词
*/

func isTransWords(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	count := make([]int, 128)
	for i := 0; i < len(str2); i++ {
		count[str1[i]-'a']++
		count[str2[i]-'a']--
	}

	for i := 0; i < 128; i++ {
		if count[i] != 0 {
			return false
		}
	}

	return true
}
