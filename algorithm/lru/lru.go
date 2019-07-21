package main

import "fmt"

// double node
type doubleNode struct {
	Key  string
	Val  int
	Prev *doubleNode
	Next *doubleNode
}

// double link list
type doubleLinkList struct {
	Len  int
	Head *doubleNode
}

func NewDoubleLinkList() *doubleLinkList {
	var l doubleLinkList
	var dummyNode doubleNode
	dummyNode.Prev = &dummyNode
	dummyNode.Next = &dummyNode
	l.Head = &dummyNode
	return &l
}

func (l *doubleLinkList) InsertHead(node *doubleNode) {
	next := l.Head.Next
	l.Head.Next = node
	node.Prev = l.Head
	node.Next = next
	next.Prev = node
	l.Len += 1
}

func (l *doubleLinkList) DeleteTail() string {
	tail := l.Head.Prev
	prev := tail.Prev
	prev.Next = l.Head
	l.Head.Prev = prev
	tail.Prev = nil
	tail.Next = nil
	l.Len -= 1
	return tail.Key
}

func (l *doubleLinkList) DeleteNode(node *doubleNode) {
	prev := node.Prev
	next := node.Next
	prev.Next = next
	next.Prev = prev
	node.Prev = nil
	node.Next = nil
	l.Len -= 1
}

type memory map[string]*doubleNode

type LRU struct {
	Cap    int
	Memory memory
	Lst    *doubleLinkList
}

func NewLRU(cap int) *LRU {
	var l LRU
	l.Cap = cap
	l.Memory = memory{}
	l.Lst = NewDoubleLinkList()
	return &l
}

func (l *LRU) Push(key string, val int) {
	// if key in cache
	if n, ok := l.Memory[key]; ok {
		l.Lst.DeleteNode(n)
		l.Lst.InsertHead(n)
	} else {
		// if cache is full
		if l.Lst.Len >= l.Cap {
			tailKey := l.Lst.DeleteTail()
			delete(l.Memory, tailKey)
		}
		node := doubleNode{Key: key, Val: val}
		l.Lst.InsertHead(&node)
		l.Memory[key] = &node
	}
	//l.Print()
}

func (l *LRU) Get(key string) (b bool, val int) {
	if n, ok := l.Memory[key]; ok {
		l.Lst.DeleteNode(n)
		l.Lst.InsertHead(n)
		b = true
		val = n.Val
	}
	//l.Print()
	return
}

func (l *LRU) Print() {
	cur := l.Lst.Head.Next
	for cur != l.Lst.Head {
		fmt.Print(cur.Val, " ")
		cur = cur.Next
	}
	fmt.Println()
}

func main() {
	lru := NewLRU(3)
	lru.Push("1", 1)
	lru.Push("2", 2)
	lru.Push("3", 3)
	lru.Push("4", 4)
	lru.Push("2", 2)
	fmt.Println(lru.Get("4"))
	fmt.Println(lru.Get("5"))
	lru.Print()
}
