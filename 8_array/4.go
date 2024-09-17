package main

import (
	"sort"
)

/**
给定一个无序的整型数组 arr, 找到其中最小的 k 个数
要求时间复杂度为 O(Nlogk)
*/

type maxHeap struct {
	data    []int
	maxSize int
	size    int
}

func (h *maxHeap) Len() int {
	return h.size
}

func (h *maxHeap) Pop() {
	if h.size == 0 {
		return
	}
	h.size--
	h.data[0], h.data[h.size] = h.data[h.size], h.data[0]
	h.Down(0)
}

func (h *maxHeap) Down(i int) {
	left, right := 2*i+1, 2*i+2
	if left >= h.size {
		return
	}
	if right >= h.size {
		if h.data[i] < h.data[left] {
			h.data[i], h.data[left] = h.data[left], h.data[i]
			h.Down(left)
		}
		return
	}
	if h.data[left] > h.data[right] {
		if h.data[i] < h.data[left] {
			h.data[i], h.data[left] = h.data[left], h.data[i]
			h.Down(left)
		}
	} else {
		if h.data[i] < h.data[right] {
			h.data[i], h.data[right] = h.data[right], h.data[i]
			h.Down(right)
		}
	}
}

func (h *maxHeap) Push(j int) {
	if h.size == h.maxSize {
		if h.data[0] > j {
			h.data[0] = j
			h.Down(0)
		}
		return
	}
	h.data[h.size] = j
	h.size++
	h.Up(h.size - 1)
}

func (h *maxHeap) Up(i int) {
	if i == 0 {
		return
	}
	parent := (i - 1) / 2
	if h.data[parent] < h.data[i] {
		h.data[parent], h.data[i] = h.data[i], h.data[parent]
		h.Up(parent)
	}
}

// heap 时间复杂度 O(Nlogk)
func getLeastKNumbers1(arr []int, k int) []int {
	n := len(arr)
	if n == 0 || k <= 0 || k > n {
		return nil
	}

	heap := &maxHeap{
		data:    make([]int, k),
		size:    0,
		maxSize: k,
	}

	for i := 0; i < n; i++ {
		heap.Push(arr[i])
	}

	ans := make([]int, k)
	for i := 0; i < k; i++ {
		ans[i] = heap.data[i]
	}
	return ans
}

// BFPRT 时间复杂度 O(N)
func getLeastKNumbers2(arr []int, k int) []int {
	n := len(arr)
	if n == 0 || k <= 0 || k > n {
		return nil
	}

	// 找到第 k 小的数
	kth := BFPRTFindKth(arr, k)
	ans := make([]int, 0, k)
	for i := 0; i < n; i++ {
		if arr[i] < kth {
			ans = append(ans, arr[i])
		}
	}

	for len(ans) < k {
		ans = append(ans, kth)
	}

	return ans
}

// BFPRTFindKth 找到第 k 小的数
func BFPRTFindKth(arr []int, k int) int {
	if len(arr) == 1 || k == 0 {
		return arr[0]
	}
	// 1. 将 arr 分成 n/5 组，每组 5 个数，最后一组不足 5 个数
	// 2. 对每组进行排序，找到每组的中位数，组成一个新的数组 middleArr
	middleArr := make([]int, 0)
	for i := 0; i < len(arr); i += 5 {
		end := i + 5
		if end > len(arr) {
			end = len(arr)
		}
		sort.Slice(arr[i:end], func(i, j int) bool {
			return arr[i] < arr[j]
		})
		middleArr = append(middleArr, arr[(i+end-1)/2])
	}
	// 3. 递归调用 BFPRTFindKth 找到 middleArr 的中位数 middleOfMiddle
	middleOfMiddle := BFPRTFindKth(middleArr, (len(middleArr))/2)
	// 4. 以 middleOfMiddle 为基准将 arr 分成三部分，left, middle, right
	left, right := make([]int, 0), make([]int, 0)
	middleCount := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] < middleOfMiddle {
			left = append(left, arr[i])
		} else if arr[i] > middleOfMiddle {
			right = append(right, arr[i])
		} else {
			middleCount++
		}
	}
	// 5. 判断 k 在哪一部分，递归调用
	// 5.1 如果 left 的长度大于等于 k，说明第 k 小的数在 left 中
	if len(left) >= k {
		return BFPRTFindKth(left, k)
	}
	// 5.2 如果 left 的长度加上 middleCount 大于等于 k，说明第 k 小的数就是 middleOfMiddle
	if len(left)+middleCount >= k {
		return middleOfMiddle
	}
	// 5.3 如果 left 的长度加上 middleCount 小于 k，说明第 k 小的数在 right 中
	return BFPRTFindKth(right, k-len(left)-middleCount)
}
