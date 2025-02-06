package main

import (
	ds "algorithm-exercises/0_data_structure"
)

/**
有一个源源不断地吐出整数的数据流，假设你有足够的空间来保存吐出的
请设计一个数据结构可以随时取得之前吐出所有数的中位数
*/

type MedianHolder struct {
	minHalf *ds.PriorityQueue[int] // 大根堆维护小的一半，奇数个时中位数存在此处
	maxHalf *ds.PriorityQueue[int] // 小根堆维护大的一半
	size    int
}

func NewMedianHolder() *MedianHolder {
	return &MedianHolder{
		minHalf: ds.NewPriorityQueue[int](func(a, b int) int {
			return a - b
		}),
		maxHalf: ds.NewPriorityQueue[int](func(a, b int) int {
			return b - a
		}),
	}
}

func (mh *MedianHolder) getMedian() (int, bool) {
	if mh.size == 0 {
		return 0, false
	}

	if mh.size%2 == 0 {
		return (mh.minHalf.Top() + mh.maxHalf.Top()) / 2, true
	} else {
		return mh.minHalf.Top(), true
	}
}

func (mh *MedianHolder) push(x int) {

	if mh.minHalf.Size() == 0 {
		mh.minHalf.Push(x)
		mh.size++
		return
	}

	maxOfMinHalf := mh.minHalf.Top()

	if x <= maxOfMinHalf {
		mh.minHalf.Push(x)
	} else {
		mh.maxHalf.Push(x)
	}

	for mh.minHalf.Size() > mh.maxHalf.Size()+1 {
		mh.maxHalf.Push(mh.minHalf.Top())
		mh.minHalf.Pop()
	}

	for mh.maxHalf.Size() > mh.minHalf.Size() {
		mh.minHalf.Push(mh.maxHalf.Top())
		mh.maxHalf.Pop()
	}
	mh.size++
}
