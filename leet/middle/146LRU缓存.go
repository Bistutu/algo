package main

import (
	"container/list"
	"fmt"
)

func main() {
	cache := Constructor(2)

	cache.Put(1, 0) // 缓存是 {1=0}
	cache.Put(2, 2) // 缓存是 {1=0, 2=2}

	fmt.Print(cache.Get(1), "\t") // 返回 0

	cache.Put(3, 3) // 逐出键 2，缓存是 {1=0, 3=3}

	fmt.Print(cache.Get(2), "\t") // 返回 -1 (未找到)

	cache.Put(4, 4) // 逐出键 1，缓存是 {4=4, 3=3}

	fmt.Print(cache.Get(1), "\t") // 返回 -1 (未找到)
	fmt.Print(cache.Get(3), "\t") // 返回 3
	fmt.Print(cache.Get(4), "\t") // 返回 4
}

type LRUCache struct {
	capacity int
	cache    map[int]*list.Element // 缓存
	list     *list.List            // 顺序
}

type pair struct {
	key, value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element, capacity),
		list:     list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	if ele, ok := this.cache[key]; ok {
		// 移动至队首
		this.list.MoveToFront(ele)
		return ele.Value.(pair).value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	// 如果存在，则只是更新
	if element, ok := this.cache[key]; ok {
		this.list.MoveToFront(element)
		element.Value = pair{key, value}
	} else {
		if this.capacity == this.list.Len() { // 容量达到则删除一个元素
			back := this.list.Back()
			if back != nil {
				this.list.Remove(back)
				delete(this.cache, back.Value.(pair).key)
			}
		}
		this.cache[key] = this.list.PushFront(pair{key, value})
	}
}
