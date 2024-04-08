package main

import (
	ds "algorithm-exercises/0_data_structure"
	"time"
)

/*
	实现一种狗猫队列的结构， 要求如下：
	• 用户可以调用 add 方法将 cat 类或 dog 类的实例放入队列中；
	• 用户可以调用 pollAll 方法， 将队列中所有的实例按照进队列的先后顺序依次弹出；
	• 用户可以调用 pollDog 方法， 将队列中 dog 类的实例按照进队列的先后顺序依次弹出；
	• 用户可以调用 poIlCat 方法，将队列中 cat 类的实例按照进队列的先后顺序依次弹出；
	• 用户可以调用 isEmpty 方法， 检查队列中是否还有 dog 或 cat 的实例；
	• 用户可以调用 isDogEmpty 方法， 检查队列中是否有 dog 类的实例；
	• 用户可以调用 isCatEmpty 方法， 检查队列中是否有 cat 类的实例
*/

type Pet interface {
	Name() string
}

type Dog struct {
	EnqueueTime time.Time // 时间戳额外记录入队时间，也可以封装在外层从而不改变 dog 的结构
}

func (d *Dog) Name() string {
	return "dog"
}

type Cat struct {
	EnqueueTime time.Time
}

func (c *Cat) Name() string {
	return "cat"
}

type CatDogQueue struct {
	catQueue ds.Queue[*Cat]
	dogQueue ds.Queue[*Dog]
}

func (cdq *CatDogQueue) add(p Pet) {
	if _, ok := p.(*Dog); ok {
		dog := p.(*Dog)
		dog.EnqueueTime = time.Now()
		cdq.dogQueue.EnQueue(dog)
	} else {
		cat := p.(*Cat)
		cat.EnqueueTime = time.Now()
		cdq.catQueue.EnQueue(cat)
	}
}

func (cdq *CatDogQueue) pollAll() Pet {
	var cat *Cat
	var dog *Dog
	if !cdq.dogQueue.Empty() {
		dog = cdq.dogQueue.Front()
	}
	if !cdq.catQueue.Empty() {
		cat = cdq.catQueue.Front()
	}
	if cat == nil {
		cdq.dogQueue.DeQueue()
		return dog
	}
	if dog == nil {
		cdq.catQueue.DeQueue()
		return cat
	}
	if cat.EnqueueTime.Before(dog.EnqueueTime) {
		cdq.catQueue.DeQueue()
		return cat
	} else {
		cdq.dogQueue.DeQueue()
		return dog
	}
}

func (cdq *CatDogQueue) pollDog() *Dog {
	dog := cdq.dogQueue.Front()
	cdq.dogQueue.DeQueue()
	return dog
}

func (cdq *CatDogQueue) pollCar() *Cat {
	cat := cdq.catQueue.Front()
	cdq.catQueue.DeQueue()
	return cat
}

func (cdq *CatDogQueue) isEmpty() bool {
	return cdq.dogQueue.Empty() && cdq.catQueue.Empty()
}

func (cdq *CatDogQueue) isDogEmpty() bool {
	return cdq.dogQueue.Empty()
}

func (cdq *CatDogQueue) isCatEmpty() bool {
	return cdq.catQueue.Empty()
}
