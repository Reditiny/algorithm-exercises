package main

/**
给定两个字符串 str 和 match 如果字符串 str 中含有子串 match, 则返回 match 在 str 中的开始位置, 不含有则返回 -1
*/

func kmp(str, match string) int {
	next := getNext(match)
	j := 0
	for i := 0; i < len(str); i++ {
		for j > 0 && str[i] != match[j] {
			j = next[j-1]
		}
		if str[i] == match[j] {
			j++
		}

		if j == len(match) {
			return i - len(match)
		}
	}
	return -1
}

func getNext(match string) []int {
	n := len(match)
	next := make([]int, n)
	j := 0
	for i := 1; i < n; i++ {
		for j > 0 && match[i] != match[j] {
			j = next[j-1]
		}

		if match[i] == match[j] {
			j++
		}

		next[i] = j
	}

	return next
}
