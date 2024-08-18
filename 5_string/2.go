package main

/**
给定一个字符串 str, 求其中全部数字串所代表的数字之和
忽略小数点字符；如果紧贴数字子串的左侧出现字符"-"，当连续出现的数量为奇数时，则数字视为负，偶数时为正
*/

type Status int

const (
	NotStart    Status = 1
	StartNumber Status = 2
	StartMinus  Status = 3
)

func strNumSum(str string) int {
	n := len(str)
	if n == 0 {
		return 0
	}
	ans := 0

	cur := 0
	mul := 1
	status := NotStart

	for j := n - 1; j >= 0; j-- {
		if status == NotStart {
			if str[j] >= '0' && str[j] <= '9' {
				status = StartNumber
				cur += int(str[j]-'0') * mul
				mul *= 10
			}
		} else if status == StartNumber {
			if str[j] >= '0' && str[j] <= '9' {
				cur += int(str[j]-'0') * mul
				mul *= 10
			} else if str[j] == '-' {
				status = StartMinus
				cur = -cur
			} else {
				status = NotStart
				ans += cur
				cur = 0
				mul = 1
			}
		} else if status == StartMinus {
			if str[j] == '-' {
				cur = -cur
			} else if str[j] >= '0' && str[j] <= '9' {
				status = NotStart
				ans += cur
				cur = 0
				mul = 1
				j++
			} else {
				status = NotStart
				ans += cur
				cur = 0
				mul = 1
			}
		}
	}

	return ans

}
