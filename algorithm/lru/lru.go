package main

import "fmt"

// double node
type doubleNode struct {
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

func (l *doubleLinkList) DeleteTail() int {
	tail := l.Head.Prev
	prev := tail.Prev
	prev.Next = l.Head
	l.Head.Prev = prev
	tail.Prev = nil
	tail.Next = nil
	l.Len -= 1
	return tail.Val
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

type LRU struct {
	Cap    int
	Memory map[int]*doubleNode
	Lst    *doubleLinkList
}

func NewLRU(cap int) *LRU {
	var l LRU
	l.Cap = cap
	l.Memory = map[int]*doubleNode{}
	l.Lst = NewDoubleLinkList()
	return &l
}

func (l *LRU) Get(val int) int {
	// if val in cache
	if n, ok := l.Memory[val]; ok {
		l.Lst.DeleteNode(n)
		l.Lst.InsertHead(n)
	} else {
		// if cache is full
		if l.Lst.Len >= l.Cap {
			tailVal := l.Lst.DeleteTail()
			delete(l.Memory, tailVal)
		}
		node := doubleNode{Val: val}
		l.Lst.InsertHead(&node)
		l.Memory[val] = &node
	}
	//l.Print()
	return val
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
	lru.Get(1)
	lru.Get(2)
	lru.Get(3)
	lru.Get(4)
	lru.Get(2)
	lru.Get(3)
	lru.Get(5)
	//lru.Print()
}
