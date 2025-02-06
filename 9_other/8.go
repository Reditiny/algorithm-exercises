package main

/**
哈希表常见的三个操作是 put、get 和 containsKey, 而且这三个操作的时间复杂度为 O(1)。
现在想加一个 setAll 功能，就是把所有记录的 value 都设成统一的值
请设计并实现这种有 setAll 功能的哈希表，并且 put、get、containsKey 和 setAll 四个操作的时间复杂度O(1)
*/

type MyHashMap struct {
	mainMap  map[int]int
	auxMap   map[int]struct{}
	allValue int
}

func NewMyHashMap() *MyHashMap {
	return &MyHashMap{
		mainMap: make(map[int]int),
		auxMap:  make(map[int]struct{}),
	}
}

func (m *MyHashMap) put(key, value int) {
	m.mainMap[key] = value
	m.auxMap[key] = struct{}{}
}

func (m *MyHashMap) get(key int) (int, bool) {
	if _, ok := m.mainMap[key]; ok {
		return m.mainMap[key], true
	}
	if _, ok := m.auxMap[key]; ok {
		return m.allValue, true
	}
	return -1, false
}

func (m *MyHashMap) remove(key int) {
	delete(m.mainMap, key)
	delete(m.auxMap, key)
}

func (m *MyHashMap) containsKey(key int) bool {
	if _, ok := m.mainMap[key]; ok {
		return true
	}
	if _, ok := m.auxMap[key]; ok {
		return true
	}
	return false
}

func (m *MyHashMap) setAll(value int) {
	m.allValue = value
	m.mainMap = make(map[int]int)
}
