package main

import ds "algorithm-exercises/0_data_structure"

/*
判断 t1 树中是否有与 t2 树拓扑结构完全相同的子树
t1 有 N 个节点，t2 有 M 个节点，复杂度 O(N+M)
*/
func isSubTree(t1 *ds.BTNode[int], t2 *ds.BTNode[int]) bool {
	tree1 := serializeBTree(t1)
	tree2 := serializeBTree(t2)
	// 若满足条件，则序列化后 t2 一定是 t1 的子串
	return kmp(tree1, tree2)
}

// kmp 算法
func kmp(str1, str2 []byte) bool {
	next := getNext(str2)
	j := 0
	for i := 0; i < len(str1); i++ {
		for j != -1 && str1[i] != str2[j] {
			j = next[j]
		}
		if j == len(str2)-1 {
			return true
		}
	}
	return false
}

func getNext(p []byte) []int {
	if len(p) == 0 {
		return nil
	}
	next := make([]int, len(p))
	next[0] = -1
	k := -1
	for i := 1; i < len(p); i++ {
		for k != -1 && p[k+1] != p[i] {
			k = next[k]
		}
		if p[k+1] == p[i] {
			k++
		}
		next[i] = k
	}
	return next
}
