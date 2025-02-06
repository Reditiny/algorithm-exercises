package main

import (
	"fmt"
	"sync"
)

/**
己知一个消息流会不断地吐出整数 1~N 但不一定按照顺序吐出。
如果上次打印的数为 i, 当 i+1 出现时,打印 i+1 及其之后接收过连续的,直到 1〜N全部接收并打印完
请设计这种接收并打印的结构
*/

type MessageBox struct {
	receivedMsg sync.Map
	lastMsg     int
}

func (mb *MessageBox) receive(i int) {
	if i <= mb.lastMsg {
		return
	}
	mb.receivedMsg.Store(i, struct{}{})
}

func (mb *MessageBox) printing() {
	for {
		if _, ok := mb.receivedMsg.Load(mb.lastMsg + 1); ok {
			mb.lastMsg++
			mb.receivedMsg.Delete(mb.lastMsg)
			fmt.Println(mb.lastMsg)
		}
	}
}
