package main

/**
设计一种缓存结构， 该结构在构造时确定大小为 K, 并有 get 和 set两个功能，要求
1. set 和 get 方法的时间复杂度为 O(1)。
2. 某个 key 的 set 或 get 操作一旦发生， 认为这个 key 的记录成了最经常使用的。
3. 当缓存的大小超过 K 时， 移除最不经常使用的记录， 即 set 或 get 最久远的。
*/

type Node struct {
	Key  int
	Val  int
	Pre  *Node
	Next *Node
}

type LRUCache struct {
	size  int
	cache map[int]*Node
	head  *Node
	tail  *Node
}

func NewLRUCache(size int) *LRUCache {
	head := &Node{}
	tail := &Node{}
	head.Next = tail
	tail.Pre = head
	return &LRUCache{
		size:  size,
		cache: make(map[int]*Node),
		head:  head,
		tail:  tail,
	}
}

func (l *LRUCache) moveToHead(node *Node) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
	l.addHead(node)
}

func (l *LRUCache) addHead(node *Node) {
	node.Next = l.head.Next
	node.Pre = l.head
	l.head.Next.Pre = node
	l.head.Next = node
}

func (l *LRUCache) removeLast() {
	delete(l.cache, l.tail.Pre.Key)
	l.tail.Pre.Pre.Next = l.tail
	l.tail.Pre = l.tail.Pre.Pre
}

func (l *LRUCache) Get(key int) (int, bool) {
	if _, ok := l.cache[key]; !ok {
		return 0, false
	}
	l.moveToHead(l.cache[key])
	return l.cache[key].Val, true
}

func (l *LRUCache) Set(key, value int) {
	if node, ok := l.cache[key]; ok {
		node.Val = value
		l.moveToHead(node)
		return
	}

	if len(l.cache) >= l.size {
		l.removeLast()
	}
	node := &Node{
		Key: key,
		Val: value,
	}
	l.cache[key] = node
	l.addHead(node)
}
