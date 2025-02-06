package main

import (
	ds "algorithm-exercises/0_data_structure"
	"math/rand"
)

/**
rand.Float32() 等概率随机返回一个在 [0,1) 范围上的数
实现一个函数返回［0,x) 区间上的数出现的概率为 x^k (0<x≤1)
*/

func randomOnK(k int) float64 {
	val := 1.0
	for i := 0; i < k; i++ {
		val = ds.Max(val, rand.Float64())
	}
	return val
}
