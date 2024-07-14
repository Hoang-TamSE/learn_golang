package main

import "fmt"

type Node[T comparable] struct {
	val      T
	nextNode *Node[T]
}

type List[T comparable] struct {
	Head *Node[T]
	Tail *Node[T]
}

func (l *List[T]) Add(val T) {
	n := &Node[T]{
		val: val,
	}
	if l.Head == nil {
		l.Head = n
		l.Tail = n
		return
	}
	l.Tail.nextNode = n
	l.Tail = l.Tail.nextNode
}

func (l *List[T]) Insert(val T, p int) {
	node := &Node[T]{
		val: val,
	}
	if l.Head == nil {
		l.Head = node
		l.Tail = node
		return
	}
	if p <= 0 {
		node.nextNode = l.Head
		l.Head = node
		return
	}
	curNode := l.Head
	for i := 1; i < p; i++ {
		if curNode.nextNode == nil {
			curNode.nextNode = node
			l.Tail = curNode.nextNode
			return
		}
		curNode = curNode.nextNode
	}
	node.nextNode = curNode.nextNode
	curNode.nextNode = node
	if l.Tail == curNode {
		l.Tail = node
	}
	return
}

func (l *List[T]) Index(val T) int {
	i := 0
	for curNode := l.Head; curNode != nil; curNode = curNode.nextNode {
		if curNode.val == val {
			return i
		}
		i++
	}
	return -1
}

func main() {
	l := &List[int]{}
	l.Add(5)
	l.Add(10)
	fmt.Println(l.Index(5))
	fmt.Println(l.Index(10))
	fmt.Println(l.Index(20))

	l.Insert(100, 0)
	fmt.Println(l.Index(5))
	fmt.Println(l.Index(10))
	fmt.Println(l.Index(20))
	fmt.Println(l.Index(100))

	l.Insert(200, 1)
	fmt.Println(l.Index(5))
	fmt.Println(l.Index(10))
	fmt.Println(l.Index(200))
	fmt.Println(l.Index(20))
	fmt.Println(l.Index(100))

	for curNode := l.Head; curNode != nil; curNode = curNode.nextNode {
		fmt.Println(curNode.val)
	}

	l.Insert(300, 10)
	for curNode := l.Head; curNode != nil; curNode = curNode.nextNode {
		fmt.Println(curNode.val)
	}

	l.Add(400)
	for curNode := l.Head; curNode != nil; curNode = curNode.nextNode {
		fmt.Println(curNode.val)
	}

	l.Insert(500, 6)
	for curNode := l.Head; curNode != nil; curNode = curNode.nextNode {
		fmt.Println(curNode.val)
	}
}
