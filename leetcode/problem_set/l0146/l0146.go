package main

import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	Memory map[int]*list.Element
	Lst    *list.List
	Len    int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Memory: map[int]*list.Element{},
		Lst:    list.New(),
		Len:    capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.Memory[key]; ok {
		this.Lst.MoveToFront(v)
		return v.Value.(map[string]int)["value"]
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.Memory[key]; ok {
		v.Value.(map[string]int)["value"] = value
		this.Lst.MoveToFront(v)
	} else {
		if this.Lst.Len() >= this.Len {
			backEle := this.Lst.Back()
			delete(this.Memory, backEle.Value.(map[string]int)["key"])
			this.Lst.Remove(backEle)
		}
		this.Memory[key] = this.Lst.PushFront(map[string]int{"key": key, "value": value})
	}
}

func main() {
	obj := Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	fmt.Println(obj.Get(1)) // 1
	obj.Put(3, 3)
	fmt.Println(obj.Get(2)) // -1
	obj.Put(4, 4)
	fmt.Println(obj.Get(1)) // -1
	fmt.Println(obj.Get(3)) // 3
	fmt.Println(obj.Get(4)) // 4
}
