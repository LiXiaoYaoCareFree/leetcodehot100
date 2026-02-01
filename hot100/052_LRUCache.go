package main

import (
	"container/list"
	"fmt"
)

/*
题目：LRU 缓存 (LRU Cache)
描述：请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
实现 LRUCache 类：
LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
输入示例：
LRUCache 2
put 1 1
put 2 2
get 1
put 3 3
get 2
put 4 4
get 1
get 3
get 4
输出示例：
null null null 1 null -1 null -1 3 4
*/

type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	list     *list.List
}

type entry struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		list:     list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	if elem, ok := this.cache[key]; ok {
		this.list.MoveToFront(elem)
		return elem.Value.(*entry).value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if elem, ok := this.cache[key]; ok {
		this.list.MoveToFront(elem)
		elem.Value.(*entry).value = value
	} else {
		if this.list.Len() == this.capacity {
			delete(this.cache, this.list.Back().Value.(*entry).key)
			this.list.Remove(this.list.Back())
		}
		elem := this.list.PushFront(&entry{key, value})
		this.cache[key] = elem
	}
}

func main() {
	// Simple driver code for the example
	// Input format parsing for this complex interaction is custom
	// Assuming a simplified command sequence for demonstration
	// Capacity
	var cap int
	fmt.Scan(&cap)
	obj := Constructor(cap)

	var op string
	for {
		if _, err := fmt.Scan(&op); err != nil {
			break
		}
		if op == "put" {
			var k, v int
			fmt.Scan(&k, &v)
			obj.Put(k, v)
			fmt.Print("null ")
		} else if op == "get" {
			var k int
			fmt.Scan(&k)
			fmt.Printf("%d ", obj.Get(k))
		}
	}
	fmt.Println()
}
