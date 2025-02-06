package main

/**
设计一种结构， 在该结构中有如下三个功能，时间复杂度都是 O(1)：
• insert(key)：将某个 key 加入到该结构， 做到不重复加入
• delete(key)：将原本在结构中的某个 key 移除
• getRandom()：等概率随机返回结构中的任何一个 key
*/

type RandomPool struct {
	m map[string]struct{}
}

func (r RandomPool) insert(key string) {
	r.m[key] = struct{}{}
}

func (r RandomPool) delete(key string) {
	delete(r.m, key)
}

func (r RandomPool) getRandom() string {
	for k, _ := range r.m {
		return k
	}
	return ""
}
