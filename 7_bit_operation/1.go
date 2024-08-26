package main

/**
如何不用任何额外变量交换两个整数的值？
*/

func swapBit(a, b int) (int, int) {
	// ^ 为异或运算符，相同为 0，不同为 1
	// a ^ 0 = a
	// a ^ a = 0
	a = a ^ b
	b = a ^ b
	a = a ^ b
	return a, b
}

func swapAdd(a, b int) (int, int) {
	a = a + b
	b = a - b
	a = a - b
	return a, b
}
