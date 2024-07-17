package main

import (
	ds "algorithm-exercises/0_data_structure"
	"bytes"
	"strconv"
)

/*
	二叉树被记录成文件叫作二叉树的序列化,通过文件内容重建原来二叉树叫作二叉树的反序列化。
	给定一棵二叉树的头节点 head, 值的类型为 int。设计一种二叉树序列化和反序列化的方案
*/

func serializeBTree(root *ds.BTNode[int]) []byte {
	data := make([]byte, 0)
	queue := ds.NewQueue[*ds.BTNode[int]](10)
	queue.EnQueue(root)
	for !queue.Empty() {
		curSize := queue.Size()
		for i := 0; i < curSize; i++ {
			curNode := queue.Front()
			queue.DeQueue()
			if curNode == nil {
				// # 代表节点为空
				data = append(data, '#')
			} else {
				data = append(data, []byte(strconv.Itoa(curNode.Val))...)
				queue.EnQueue(curNode.Left)
				queue.EnQueue(curNode.Right)
			}
			// ! 为节点分隔符
			data = append(data, '!')
		}
	}
	return data
}

func deserializeBTree(data BTreeData) (root *ds.BTNode[int]) {
	root = data.nextNode()
	if root == nil {
		return
	}
	queue := ds.NewQueue[*ds.BTNode[int]](10)
	queue.EnQueue(root)
	for !queue.Empty() {
		curSize := queue.Size()
		for i := 0; i < curSize; i++ {
			curNode := queue.Front()
			queue.DeQueue()
			curNode.Left = data.nextNode()
			if curNode.Left != nil {
				queue.EnQueue(curNode.Left)
			}
			curNode.Right = data.nextNode()
			if curNode.Right != nil {
				queue.EnQueue(curNode.Right)
			}
		}
	}
	return
}

type BTreeData struct {
	data []byte
}

func (b *BTreeData) nextNode() *ds.BTNode[int] {
	if len(b.data) == 0 {
		return nil
	}

	index := bytes.Index(b.data, []byte{'!'})
	if index == -1 {
		return nil
	}
	defer func() {
		b.data = b.data[index+1:]
	}()
	if b.data[index-1] == '#' {
		return nil
	}
	val, err := strconv.Atoi(string(b.data[:index]))
	if err != nil {
		return nil
	}
	return &ds.BTNode[int]{Val: val}
}
