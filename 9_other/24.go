package main

import ds "algorithm-exercises/0_data_structure"

/**
设计一个堆结构
1. 没有扩容的负担。
2. 可以生成小根堆， 也可以生成大根堆。
3. 包含 getHead 方法，返回当前堆顶的值，复杂度 O(1)
4. 包含 getSize 方法，返回当前堆的大小，复杂度 O(1)
5. 包含 add(x)方法，即向堆中新加元素 x， 复杂度 O(logN)
6. 包含 popHead 方法， 即删除并返回堆顶的值， 复杂度 O(logN)
*/

type LinkHeap struct {
	root     *ds.HeapNode[int]
	lastNode *ds.HeapNode[int]
	size     int
	compare  func(node1, node2 *ds.HeapNode[int]) bool
}

func NewLinkHeap(compare func(cur, target *ds.HeapNode[int]) bool) *LinkHeap {
	return &LinkHeap{
		compare: compare,
	}
}

func (lh *LinkHeap) GetHead() (int, bool) {
	if lh.size == 0 || lh.root == nil {
		return 0, false
	}

	return lh.root.Val, true
}

func (lh *LinkHeap) GetSize() int {
	if lh.size == 0 || lh.root == nil {
		return 0
	}
	return lh.size
}

func (lh *LinkHeap) Add(x int) {
	// todo find last
}

func (lh *LinkHeap) PopHead() (int, bool) {
	if lh.GetSize() == 0 {
		return 0, false
	}

	res := lh.root.Val
	lh.root.Val, lh.lastNode.Val = lh.lastNode.Val, lh.root.Val
	lh.down(lh.root)
	// todo find last
	return res, true
}

func (lh *LinkHeap) up(node *ds.HeapNode[int]) {
	if node.Par != nil && lh.compare(node, node.Par) {
		node.Par.Val, node.Val = node.Val, node.Par.Val
		lh.up(node.Par)
	}
}

func (lh *LinkHeap) down(node *ds.HeapNode[int]) {
	if (node.Left != nil && !lh.compare(node, node.Left)) || (node.Right != nil && !lh.compare(node, node.Right)) {
		var work *ds.HeapNode[int]
		if node.Left == nil {
			work = node.Right
		}
		if node.Right == nil {
			work = node.Left
		}
		if work == nil {
			if lh.compare(node.Right, node.Left) {
				work = node.Right
			} else {
				work = node.Left
			}
		}
		node.Val, work.Val = work.Val, node.Val
		lh.down(work)
	}
}
