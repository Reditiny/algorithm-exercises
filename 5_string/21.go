package main

/**
给定字符串 str, 其中绝对不含有字符'.'和'*' 再给定字符串 exp, 其中可以含有'.'或'*',
'*'字符不能是 exp 的首字符，并且任意两个'*'字符不相邻，exp 中的'.'代表任何一个字符，
exp 中的'*'字符表示前一个字符可以有 0 个或者多个。请写一个函数， 判断 str 是否能被 exp 匹配
*/

func match(str, exp string) bool {
	if !valid(str, exp) {
		return false
	}
	return process(str, exp, 0, 0)
}

func valid(str, exp string) bool {
	for i := range str {
		if str[i] == '.' || str[i] == '*' {
			return false
		}
	}
	for i := range exp {
		if exp[i] == '*' && (i == 0 || exp[i-1] == '*') {
			return false
		}
	}
	return true
}

// str[i..] 能否被 exp[j..] 匹配
func process(str, exp string, i, j int) bool {
	if j == len(exp) {
		return i == len(str)
	}
	if j == len(exp)-1 || exp[j+1] != '*' {
		// exp 的最后一个字符，或是 exp 的下一个字符不是 '*'，本次匹配只匹配掉 exp 的一个字符
		return i < len(str) && (str[i] == exp[j] || exp[j] == '.') && process(str, exp, i+1, j+1)
	}
	// exp 的下一个字符是 '*'，本次匹配匹配掉 exp[j] 和 exp[j+1]
	for i < len(str) && (str[i] == exp[j] || exp[j] == '.') {
		if process(str, exp, i, j+2) {
			return true
		}
		i++
	}
	return process(str, exp, i, j+2)
}
