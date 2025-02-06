package main

import (
	ds "algorithm-exercises/0_data_structure"
	"strings"
)

/**
给定一个 32 位整数 num, 写两个函数分别返回 num 的英文与中文表达字符串
*/

func getChineseNumStr(num int32) string {
	if num == 0 {
		return "零"
	}

	primes := []string{"十", "百", "千", "万", "十", "百", "千", "亿", "十"}
	numsMap := map[int32]string{
		0: "零",
		1: "一",
		2: "二",
		3: "三",
		4: "四",
		5: "五",
		6: "六",
		7: "七",
		8: "八",
		9: "九",
	}
	neg := num < 0
	num = ds.Abs(num)

	stack := ds.NewStack[string](50)
	p := 0
	for num > 0 {
		cur := num % 10

		switch numsMap[cur] {
		case "零":
			if stack.Empty() || stack.Top() == "万" || stack.Top() == "亿" {
				break
			}
			stack.Pop()
			if !stack.Empty() && stack.Top() == "零" {
				stack.Pop()
			}
			if stack.Empty() || stack.Top() == "万" || stack.Top() == "亿" {
				break
			}
			stack.Push(numsMap[cur])
		default:
			stack.Push(numsMap[cur])
		}

		num /= 10
		if num > 0 {
			if primes[p] == "亿" && !stack.Empty() && stack.Top() == "万" {
				stack.Pop()
			}
			stack.Push(primes[p])
			p++
		}
	}

	ans := strings.Builder{}
	if neg {
		ans.WriteString("负")
	}

	for !stack.Empty() {
		ans.WriteString(stack.Top())
		stack.Pop()
	}

	return ans.String()
}
