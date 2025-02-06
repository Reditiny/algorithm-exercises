package main

import "math/rand"

/**
1. 不能使用任何额外的随机机制，请用 rand1To5 实现等概率随机产生 1〜7 的随机函数 rand1To7

2. 不能使用任何额外的随机机制，请用 randO1p 实现等概率随机产生 1 〜6 的随机函数 rand1To6

3. 不能使用任何额外的随机机制，请用rand1ToM(m) 实现等概率随机产生 1〜n 的随机函数 rand1ToN，入参为 n，m
*/

// 等概率随机产生 1~5 的随机函数
func rand1To5() int {
	return rand.Int()%5 + 1
}

// 以 p 概率产生 0, 以 1-p 概率产生 1 的随机函数
func rand10p() int {
	p := 0.83
	if rand.Float64() < p {
		return 0
	}
	return 1
}

// 等概率随机产生 1〜M 的随机函数
func rand1ToM(m int) int {
	return rand.Int()%m + 1
}

func rand1To7() int {
	for {
		// rand1To5()-1 等概率产生 0, 1, 2, 3, 4
		// (rand1To5()-1)*5 等概率产生 0, 5, 10, 15, 20
		// 两者的和产生 0, 1, 2, 3, 4, 5, 6, 10, 11, 12, 13, 14, 15, 16, 17, 20, 21, 22, 23, 24
		num := (rand1To5()-1)*5 + rand1To5() - 1
		if num < 21 {
			return num%7 + 1
		}
	}
}

func rand1To6() int {
	// rand10p() 以 p 概率产生 0, 以 1-p 概率产生 1
	// 产生 10 和 01 的概率都是 p*(1-p)，由此可得等概率产生 0 和 1 的随机函数
	rand0To1 := func() int {
		for {
			num1 := rand10p()
			num2 := rand10p()
			if num1 == 1 && num2 == 0 {
				return 0
			}
			if num1 == 0 && num2 == 1 {
				return 1
			}
		}
	}
	// rand0To1 等概率产生 0 1
	// 2*rand0To1() 等概率产生 0 2
	// 2*rand0To1() + rand0To1() 等概率产生 0 1 2 3
	// 4*rand0To1() 等概率产生 0 4
	// 4*rand0To1() + 2*rand0To1() + rand0To1() 等概率产生 0 1 2 3 4 5 6 7
	for {
		num := 4*rand0To1() + 2*rand0To1() + rand0To1()
		if num < 7 {
			return num + 1
		}
	}
}

// 对于 rand1ToN 来说，rand1ToN - 1 等概率产生 0, 1, 2, ..., N-1
// rand1ToN-1 + (rand1ToN-1)*N 等概率产生 0, 1, 2, ..., N^2-1
// rand1ToN-1 + (rand1ToN-1)*N + (rand1ToN-N)*N^2 等概率产生 0, 1, 2, ..., N^3-1
func rand1ToN(n, m int) int {
	if n <= m {
		// 若 n <= m，找到小于等于 m 的 n 的最大倍数
		threshold := m - m%n
		for {
			// rand1ToM(m) - 1 等概率产生 0, 1, 2, ..., threshold, threshold+1, ..., m-1
			num := rand1ToM(m) - 1
			if num < threshold {
				return num%n + 1
			}
		}
	} else {
		// 当 n > m 时，找到大于等于 n 的 m 的幂次 maxValue
		powerCount := 1
		work := n
		for work > m {
			work /= m
			powerCount++
		}
		maxValue := 1
		for i := 0; i < powerCount; i++ {
			maxValue *= m
		}
		// 找到小于等于 maxValue 的 n 的最大倍数
		threshold := maxValue - maxValue%n
		for {
			// 等概率产生 0, 1, 2, ..., maxValue-1
			base := 1
			num := 0
			for i := 0; i < powerCount; i++ {
				num += (rand1ToM(m) - 1) * base
				base *= m
			}
			if num < threshold {
				return num%n + 1
			}
		}
	}
}
