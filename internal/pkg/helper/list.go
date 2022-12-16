package helper

import (
	"fmt"
)

type Node struct {
	prev  *Node
	next  *Node
	value interface{}
}

type List struct {
	head *Node
	tail *Node
}

func (l *List) InsertReversed(value interface{}) {
	// New element to insert
	newList := &Node{
		prev:  l.tail,
		value: value,
	}

	// In case of a non-empty list, we must point at the current node as the "next" of the former tail, then we make
	// this new node the list's tail.
	if l.tail != nil {
		l.tail.next = newList
	}
	l.tail = newList

	// We must refresh the list's head.
	tmpList := l.tail
	for tmpList.prev != nil {
		tmpList = tmpList.prev
	}
	l.head = tmpList
}

func (l *List) Insert(value interface{}) {
	// New element to insert
	newList := &Node{
		next:  l.head,
		value: value,
	}

	// In case of a non-empty list, we must point at the current node as the "previous" of the former head, then we make
	// this new node the list's head.
	if l.head != nil {
		l.head.prev = newList
	}
	l.head = newList

	// We must refresh the list's tail.
	tmpList := l.head
	for tmpList.next != nil {
		tmpList = tmpList.next
	}
	l.tail = tmpList
}

func (l *List) Display() {
	list := l.head
	for list != nil {
		fmt.Printf("%+v ->", list.value)
		list = list.next
	}
	fmt.Println()
}

func (l *List) Pop() (result interface{}) {
	if l.head == nil {
		return nil
	} else {
		result = l.head.value

		l.head = l.head.next
	}

	return result
}
