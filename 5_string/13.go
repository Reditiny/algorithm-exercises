package main

/**
给定一个字符串 str, 判断是不是整体有效的括号字符串

进阶：给定一个括号字符串 str, 返回最长的有效括号子串
*/

func isValid(str string) bool {
	if len(str) == 0 {
		return true
	}

	count := 0
	for i := range str {
		if str[i] == '(' {
			count++
		} else if str[i] == ')' {
			count--
			if count < 0 {
				return false
			}
		} else {
			return false
		}
	}
	return count == 0
}

func maxValidStr(str string) string {
	if len(str) == 0 {
		return ""
	}
	dp := make([]int, len(str))
	// dp[i] 表示以 str[i] 结尾的最长有效括号子串的长度
	max := 0
	pos := -1
	for i := 1; i < len(str); i++ {
		if str[i] == ')' {
			// 找到 dp[i-1] 最大有效括号子串的前一个字符
			matchPos := i - dp[i-1] - 1
			if matchPos >= 0 && str[matchPos] == '(' {
				// 如果匹配，则 dp[i] 的长度至少为 dp[i-1] + 2
				dp[i] = dp[i-1] + 2
				if matchPos > 0 {
					// matchPos 前面位置依然可能有有效括号子串
					dp[i] += dp[matchPos-1]
				}
			}
		}
		if dp[i] > max {
			max = dp[i]
			pos = i
		}
	}
	return str[pos+1-max : pos+1]
}
