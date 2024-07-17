package main

import ds "algorithm-exercises/0_data_structure"

/**
Tarjan 算法与并查集解决二叉树节点间最近公共祖先的批量查询问题
二叉树的节点数为 N, 查询语句的条数为 M，要求时间复杂度 O(N+M)
*/

func tarJanQuery(head ds.BTNode[int], queries [][]ds.BTNode[int]) []int {
	// 并查集
	sets := ds.NewDisjointSet()
	// 用于存储节点的祖先节点（已遍历到到）
	ancestorMap := make(map[int]int)
	// 用于存储查询语句，key 为节点值，value 为与该节点相关的查询语句的节点值列表
	queryMap := make(map[int][]int)
	// 用于存储查询语句，key 为节点值，value 为与该节点相关的查询语句的查询语句索引列表
	indexMap := make(map[int][]int)
	ans := make([]int, len(queries))

	for i := 0; i < len(queries); i++ {
		queryMap[queries[i][0].Val] = append(queryMap[queries[i][0].Val], queries[i][1].Val)
		queryMap[queries[i][1].Val] = append(queryMap[queries[i][1].Val], queries[i][0].Val)
		indexMap[queries[i][0].Val] = append(indexMap[queries[i][0].Val], i)
		indexMap[queries[i][1].Val] = append(indexMap[queries[i][1].Val], i)
	}

	var setAns func(head *ds.BTNode[int])
	setAns = func(head *ds.BTNode[int]) {
		if head == nil {
			return
		}
		setAns(head.Left)
		// 子树遍历完后合并集合，确定祖先
		if head.Left != nil {
			sets.Union(head.Left.Val, head.Val)
			ancestorMap[sets.Find(head.Val)] = head.Val
		}

		setAns(head.Right)

		if head.Right != nil {
			sets.Union(head.Right.Val, head.Val)
			ancestorMap[sets.Find(head.Val)] = head.Val
		}

		// 处理和当前节点相关的查询语句
		curQuery := queryMap[head.Val]
		curIndex := indexMap[head.Val]
		for i := 0; i < len(curQuery); i++ {
			// 该节点当前已遍历到的祖先节点即为两者最近公共祖先
			ancestor, ok := ancestorMap[sets.Find(curQuery[i])]
			if ok {
				ans[curIndex[i]] = ancestor
			}
		}

	}

	setAns(&head)
	return ans
}
