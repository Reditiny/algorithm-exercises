package main

/**
给定两个 32 位整数 a 和 b, 可正、可负、可 0。
不能使用算术运算符， 实现加减乘除运算
*/

func bitAdd(a, b int) int {
	for b != 0 {
		//	(a & b) << 1 为进位
		carry := (a & b) << 1
		// a ^ b 为不进位的和
		a = a ^ b
		// 继续加上进位
		b = carry
	}
	return a
}

func bitSubtract(a, b int) int {
	// a - b = a + (-b)
	// -b = ^b + 1 补码，取反加一
	return bitAdd(a, bitAdd(^b, 1))
}

func bitMultiply(a, b int) int {
	// a * b = a + a + ... + a 共 b 个 a
	// 乘法转化为多次加法
	// a << 1 为 a * 2， a << 2 为 a * 4
	// b & 1 为 b 的最低位，b >> 1 为 b 的所有位右移一位
	res := 0
	for i := 0; b != 0; i++ {
		if b&1 == 1 {
			// 当前最低位为 1，加上 a * 2^i
			res = bitAdd(res, a<<i)
		}
		b = b >> 1
	}
	return res
}

func bitDivide(a, b int) int {
	if b == 0 {
		panic("divided by zero")
	}
	// a = b * x + y
	// x = a / b, y = a % b
	// 通过移位来逼近 x
	// a = (b << 1) * x + y
	// x = (a - y) / (b << 1)
	res := 0
	for a >= b {
		// 找到最大的 x，使得 b << x <= a
		// 1 << i 为 2^i
		for i := 0; a >= b<<i; i++ {
			res = res | 1<<i
			a = bitSubtract(a, b<<i)
		}
	}
	return res
}
