package main

import (
	"cmp"
	"fmt"
)

type OrderableFunc[T any] func(t1, t2 T) int

type Node[T any] struct {
	val         T
	left, right *Node[T]
}
type Tree[T any] struct {
	f    OrderableFunc[T]
	root *Node[T]
}

func NewTree[T any](f OrderableFunc[T]) *Tree[T] {
	return &Tree[T]{
		f: f,
	}
}

func (t *Tree[T]) Add(val T) {
	t.root = t.root.Add(t.f, val)
}

func (t *Tree[T]) Contains(val T) bool {
	return t.root.Contains(t.f, val)
}
func (n *Node[T]) Add(f OrderableFunc[T], val T) *Node[T] {
	if n == nil {
		return &Node[T]{
			val: val,
		}
	}
	switch f(val, n.val) {
	case -1:
		n.left = n.left.Add(f, val)
	case 1:
		n.right = n.right.Add(f, val)
	}
	return n
}

func (n *Node[T]) Contains(f OrderableFunc[T], val T) bool {
	if n == nil {
		return false
	}
	switch f(val, n.val) {
	case -1:
		return n.left.Contains(f, val)
	case 1:
		return n.right.Contains(f, val)
	}
	return true
}

type Person struct {
	Name string
	Age  int
}

func OrderPeople(p1, p2 Person) int {
	r := cmp.Compare(p1.Name, p2.Name)
	if r == 0 {
		r = cmp.Compare(p1.Age, p2.Age)
	}
	return r
}

func main() {
	t2 := NewTree(OrderPeople)
	t2.Add(Person{"Bob", 30})
	t2.Add(Person{"Maria", 35})
	t2.Add(Person{"Bob", 50})
	fmt.Println(t2.Contains(Person{"Bob", 30}))
	fmt.Println(t2.Contains(Person{"Fred", 25}))

}
