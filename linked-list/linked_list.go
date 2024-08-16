package linkedlist

import "errors"

type Node struct {
	previous *Node
	Value    interface{}
	next     *Node
}

type List struct {
	head   *Node
	tail   *Node
	length int
}

func NewList(elements ...interface{}) *List {
	l := &List{}
	for i := 0; i < len(elements); i++ {
		l.Push(elements[i])
	}
	return l
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.previous
}

func (l *List) Unshift(v interface{}) {
	if l.head == nil {
		l.head = &Node{Value: v}
		l.tail = l.head
	} else {
		l.head.previous = &Node{Value: v, next: l.head}
		l.head = l.head.previous
	}
	l.length++
}

func (l *List) Push(v interface{}) {
	if l.head == nil {
		l.head = &Node{Value: v}
		l.tail = l.head
	} else {
		l.tail.next = &Node{Value: v, previous: l.tail}
		l.tail = l.tail.next
	}
	l.length++
}

func (l *List) Shift() (interface{}, error) {
	if l.length == 0 {
		return nil, errors.New("list is empty")
	}
	first := l.head
	if l.length == 1 {
		l.head = nil
		l.tail = nil
	} else {
		l.head = l.head.next
		l.head.previous = nil
	}
	l.length--
	return first.Value, nil
}

func (l *List) Pop() (interface{}, error) {
	if l.length == 0 {
		return nil, errors.New("list is empty")
	}
	last := l.tail
	if l.length == 1 {
		l.head = nil
		l.tail = nil
	} else {
		l.tail = l.tail.previous
		l.tail.next = nil
	}
	l.length--
	return last.Value, nil
}

func (l *List) Reverse() {
	first := l.head
	current := l.head
	var prev *Node

	for current != nil {
		next := current.next
		current.next = prev
		current.previous = next
		prev = current
		current = next
	}

	l.head = l.tail
	l.tail = first
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}
