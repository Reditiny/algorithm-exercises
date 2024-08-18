package main

/**
假设组成所有单词的字符仅是 "a" 〜 "z"，请实现字典树结构，包含以下四个主要功能
1. 添加 word, 可重复添加
2. 删除 word, 如果 word 添加过多次， 仅删除一个
3. 查询 word 是否在字典树中
4. 返回以字符串 pre 为前缀的单词数量
*/

type TrieNode struct {
	// 指向下一个字符的指针
	m map[byte]*TrieNode
	// 以当前节点结尾的单词数量
	end int
	// 以当前节点为前缀的单词数量
	path int
}

type Trie struct {
	root *TrieNode
}

func (t *Trie) insert(word string) {
	curNode := t.root
	curNode.path++
	for i := range word {
		if curNode.m[word[i]] == nil {
			curNode.m[word[i]] = &TrieNode{m: make(map[byte]*TrieNode)}
		}
		curNode = curNode.m[word[i]]
		curNode.path++
	}
	curNode.end++
}

func (t *Trie) search(word string) bool {
	curNode := t.root
	for i := range word {
		if curNode.m[word[i]] == nil {
			return false
		}
		curNode = curNode.m[word[i]]
	}
	return curNode.end > 0
}

func (t *Trie) delete(word string) {
	if t.search(word) {
		curNode := t.root
		curNode.path--
		for i := range word {
			curNode = curNode.m[word[i]]
			curNode.path--
		}
		curNode.end--
	}
}

func (t *Trie) prefixNumber(pre string) int {
	curNode := t.root
	for i := range pre {
		if curNode.m[pre[i]] == nil {
			return 0
		}
		curNode = curNode.m[pre[i]]
	}
	return curNode.path
}
