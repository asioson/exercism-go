// package linkedlist implements routines on a doubly linked list structure
package linkedlist

import "errors"

type Node struct {
	Val        interface{}
	next, prev *Node
}

type List struct {
	first, last *Node
}

var ErrEmptyList = errors.New("empty list")

// NewList generates a new list given a slice of values
func NewList(args ...interface{}) (l *List) {
	l = new(List)
	for _, v := range args {
		l.PushBack(v)
	}
	return l
}

// Next returns the address of the next Node object
func (n *Node) Next() *Node {
	if n == nil {
		return nil
	}
	return n.next
}

// Prev returns the address of the previous Node object
func (n *Node) Prev() *Node {
	if n == nil {
		return nil
	}
	return n.prev
}

// PushFront inserts a new item v at the front of the list
func (l *List) PushFront(v interface{}) {
	item := new(Node)
	item.Val = v
	if l.First() == nil {
		l.first = item
		l.last = item
	} else {
		item.next = l.First()
		l.First().prev = item
		l.first = item
	}
}

// PushBack inserts a new item v at the end of the list
func (l *List) PushBack(v interface{}) {
	item := new(Node)
	item.Val = v
	if l.Last() == nil {
		l.last = item
		l.first = item
	} else {
		item.prev = l.Last()
		l.Last().next = item
		l.last = item
	}
}

// PopFront removes the item at the front of the list and returns its value
func (l *List) PopFront() (interface{}, error) {
	if l == nil {
		return nil, ErrEmptyList
	} else {
		if l.First() == nil {
			return nil, ErrEmptyList
		} else {
			item := l.First()
			if l.First() == l.Last() {
				l.first = nil
				l.last = nil
			} else {
				l.first = l.First().Next()
				l.First().prev = nil
			}
			return item.Val, nil
		}
	}
}

// PopBack removes the item at the end of the list and returns its value
func (l *List) PopBack() (interface{}, error) {
	if l == nil {
		return nil, ErrEmptyList
	} else {
		if l.Last() == nil {
			return nil, ErrEmptyList
		} else {
			item := l.Last()
			if l.First() == l.Last() {
				l.first = nil
				l.last = nil
			} else {
				l.last = l.Last().Prev()
				l.Last().next = nil
			}
			return item.Val, nil
		}
	}
}

// Reverse reverses the arrangement of the values in the list
func (l *List) Reverse() {
	if l == nil {
		return
	}
	item := l.First()
	if item == nil {
		return
	}
	for item != l.Last() {
		item.next, item.prev = item.Prev(), item.Next()
		item = item.Prev()
	}
	item.next, item.prev = item.Prev(), item.Next()
	l.first, l.last = l.Last(), l.First()
}

// First returns the address of the first item in the list
func (l *List) First() *Node {
	return l.first
}

// Last returns the address of the last item in the list
func (l *List) Last() *Node {
	return l.last
}
