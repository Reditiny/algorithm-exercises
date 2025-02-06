package main

import (
	ds "algorithm-exercises/0_data_structure"
	"fmt"
)

/**
给定 String 类型的数组, 再给定整数 k 请按照排名顺序打印出现次数前 k 名的字符串
时间复杂度 O(Nlogk)
*/

type strCount struct {
	str   string
	count int
}

func printTopK(arr []string, k int) {
	count := make(map[string]int)
	for i := range arr {
		count[arr[i]]++
	}

	// 维护一个小根堆
	minPQ := ds.NewPriorityQueue[strCount](func(a, b strCount) int {
		return b.count - a.count
	})

	for str, cnt := range count {
		// 堆中最多只有 k 个元素
		if minPQ.Size() < k {
			minPQ.Push(strCount{str, cnt})
		} else {
			// 仅当当前元素的计数多于堆顶元素的计数时才入队，调整堆的复杂度为 O(logk)
			top := minPQ.Top()
			if top.count < cnt {
				minPQ.Pop()
				minPQ.Push(strCount{str, cnt})
			}
		}
	}

	ans := make([]string, 0, minPQ.Size())

	for !minPQ.Empty() {
		top := minPQ.Top()
		ans = append(ans, top.str)
		minPQ.Pop()
	}

	for i := len(ans) - 1; i >= 0; i-- {
		println(ans[i])
	}
}

/**
【进阶】
设计并实现 TopKRecord 结构，可以不断地向其中加入字符串，并且可以随时打印加入次数最多前 k 个字符串：
1. k 在 TopKRecord 实例生成时指定， 并且不再变化 (构造函数的参数)。
2. 含有 add(String str) 方法, 复杂度 O(logk)。
3. 含有 printTopK() 方法, 不要求严格按排名顺序打印, 复杂度 O(k)
*/

type TopKRecord struct {
	pq   *ds.PriorityQueue[strCount]
	cnt  map[string]int
	inPq map[string]bool
	k    int
}

func newTopKRecord(k int) TopKRecord {
	return TopKRecord{
		pq: ds.NewPriorityQueue[strCount](func(a, b strCount) int {
			return b.count - a.count
		}),
		cnt:  make(map[string]int),
		inPq: make(map[string]bool),
		k:    k,
	}
}

func (tk *TopKRecord) add(str string) {
	tk.cnt[str]++
	if tk.inPq[str] {
		// 新元素在堆中时需要调整堆，复杂度为 O(logk)
		temp := make([]strCount, 0)
		for !tk.pq.Empty() {
			top := tk.pq.Top()
			tk.pq.Pop()
			if top.str == str {
				top.count++
			}
			temp = append(temp, top)
			if top.str == str {
				break
			}
		}
		for i := 0; i < len(temp); i++ {
			tk.pq.Push(temp[i])
		}
	} else {
		// 新元素不在堆中时，如果堆未满则直接入堆，否则比较堆顶元素和当前元素的计数
		if tk.pq.Size() < tk.k {
			tk.pq.Push(strCount{str, tk.cnt[str]})
			tk.inPq[str] = true
		} else {
			top := tk.pq.Top()
			if tk.cnt[str] > top.count {
				tk.pq.Pop()
				delete(tk.inPq, top.str)
				tk.pq.Push(strCount{str, tk.cnt[str]})
				tk.inPq[str] = true
			}
		}
	}
}

func (tk *TopKRecord) printTopK() {
	tk.pq.Range(func(s strCount) {
		fmt.Println(s.str, " ", s.count)
	})
}
