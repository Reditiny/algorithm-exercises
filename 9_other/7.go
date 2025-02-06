package main

import "math/rand"

/**
有一个机器按自然数序列的方式吐出球 (1 号球 2 号球 3 号球)
一个袋子最多只能装下 K 个球，并且除袋子以外，没有更多的空间。
设计一种选择方式，使得当机器吐出第 N 号球的时候 (N>K)
袋子中的球数是 K 个， 同时可以保证每一个球被选进袋子的概率都是 K/N
*/

func getKNumsRand(k, n int) {
	bag := make([]int, k)
	for i := 0; i < n; i++ {
		if i < k {
			bag[i] = i + 1
		} else {
			// 判断这个球是否被选中
			if rand.Intn(i+1) < k {
				// 选中后判断要替换哪个球
				bag[rand.Intn(k)] = i + 1
			}
		}
	}
}
