package main

import (
	ds "algorithm-exercises/0_data_structure"
	"strconv"
)

/**
给定一个字符串 str, str 表示一个算式，可能有整数、加减乘除符号和左右括号，返回公式的计算结果
*/

var calculateMap = map[byte]func(a, b int) int{
	'+': func(a, b int) int {
		return a + b
	},
	'-': func(a, b int) int {
		return a - b
	},
	'*': func(a, b int) int {
		return a * b
	},
	'/': func(a, b int) int {
		return a / b
	},
}

var priorityMap = map[byte]int{
	'+': 1,
	'-': 1,
	'*': 2,
	'/': 2,
	'(': 3,
}

func calculate(str string) int {
	symbol := ds.NewStack[byte](len(str))
	number := ds.NewStack[int](len(str))
	lastPos := -1
	for i := range str {
		// 处理数字，遇到符号时，将符号前的数字入栈
		if str[i] == '+' || str[i] == '-' || str[i] == '*' || str[i] == '/' || str[i] == '(' || str[i] == ')' {
			if str[i] == '(' || (lastPos >= 0 && str[lastPos] == ')') {
				// 符号前无数字的情况
			} else {
				if lastPos+1 == i {
					// 特殊情况，如 (-1) 计算时变为 (0-1)
					number.Push(0)
				} else {
					n, _ := strconv.Atoi(str[lastPos+1 : i])
					number.Push(n)
				}
			}
			lastPos = i
		}
		// 处理符号
		if str[i] == '(' {
			symbol.Push(str[i])
		} else if str[i] == ')' {
			top := symbol.Top()
			for top != '(' {
				second := number.Top()
				number.Pop()
				first := number.Top()
				number.Pop()
				number.Push(calculateMap[top](first, second))
				symbol.Pop()
				top = symbol.Top()
			}
			symbol.Pop()
		} else if str[i] == '+' || str[i] == '-' || str[i] == '*' || str[i] == '/' {
			if symbol.Empty() {
				symbol.Push(str[i])
			} else {
				top := symbol.Top()
				if top == '(' {
					symbol.Push(str[i])
				} else {
					if priorityMap[str[i]] <= priorityMap[top] {
						second := number.Top()
						number.Pop()
						first := number.Top()
						number.Pop()
						number.Push(calculateMap[top](first, second))
						symbol.Pop()
						symbol.Push(str[i])
					} else {
						symbol.Push(str[i])
					}
				}
			}
		}
	}

	last, _ := strconv.Atoi(str[lastPos+1 : len(str)])
	number.Push(last)

	for !symbol.Empty() {
		second := number.Top()
		number.Pop()
		first := number.Top()
		number.Pop()
		number.Push(calculateMap[symbol.Top()](first, second))
		symbol.Pop()
	}

	return number.Top()
}
