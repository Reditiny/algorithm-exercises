package main

import (
	"strings"
)

/**
字符串和数字的对应关系
chs=[A,B,C,...,Z]
A -> 1
B -> 2
...
Z -> 26
AA -> 27
AB -> 28
...

chs=[A,B,C]
A -> 1
B -> 2
C -> 3
AA -> 4
...

实现字符串与数字的转换
*/

func StrToNum(chars []byte, str string) int {
	charToNum := make(map[byte]int)
	for i := range chars {
		charToNum[chars[i]] = i + 1
	}
	prime := len(chars)
	base := 1
	ans := 0
	for i := len(str) - 1; i >= 0; i-- {
		ans += charToNum[str[i]] * base
		base *= prime
	}
	return ans
}

func NumToStr(chars []byte, num int) string {
	numToChar := make(map[int]byte)
	for i := range chars {
		numToChar[i+1] = chars[i]
	}
	prime := len(chars)
	ansChars := make([]byte, 0)
	for num != 0 {
		ansChars = append(ansChars, numToChar[num%prime])
		num /= prime
	}

	ans := strings.Builder{}
	for i := len(ansChars) - 1; i >= 0; i-- {
		ans.WriteByte(ansChars[i])
	}

	return ans.String()
}
