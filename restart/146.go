package main

import (
	"container/list"
	"sync"
)

type LRUCache struct {
	cap   int
	mutex sync.Mutex
	cache map[int]*list.Element
	ll    *list.List
}

type Meter struct {
	key, val int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap:   capacity,
		mutex: sync.Mutex{},
		cache: make(map[int]*list.Element),
		ll:    list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.cache[key]; ok {
		// 取值并提前
		this.ll.MoveToFront(v)
		return v.Value.(*Meter).val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.cache[key]; ok {
		// 更新值并且提前
		v.Value.(*Meter).val = value
		this.ll.MoveToFront(v)
		return
	}
	// 判断容量、创建新值
	if this.cap == this.ll.Len() {
		element := this.ll.Back()
		delete(this.cache, element.Value.(*Meter).key)
		this.ll.Remove(element)
	}
	front := this.ll.PushFront(&Meter{key: key, val: value})
	this.cache[key] = front
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
