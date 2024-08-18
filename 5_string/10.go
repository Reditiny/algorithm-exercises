package main

/**
给定一个字符类型的数组 chars[], chars 右半区全是空字符，左半区不含有空字符
现在想将左半区中所有的空格字符替换成"%20"假设 chars 右半区足够大，可以满足替换所需要的空间，请完成替换函数。
替换函数的时间复杂度为 O(N)，额外空间复杂度为 O(1)
*/

func replaceSpace(chars []byte) {
	if len(chars) == 0 {
		return
	}

	j := len(chars) - 1
	for i := len(chars) - 1; i >= 0; i-- {
		if chars[i] == '0' {
			continue
		} else if chars[i] == ' ' {
			chars[j] = '0'
			chars[j-1] = '2'
			chars[j-2] = '%'
			j -= 3
		} else {
			chars[j] = chars[i]
			j--
		}
	}
}
