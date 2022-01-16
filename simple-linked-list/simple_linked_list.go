package linkedlist

import "fmt"

type Element struct {
	value int
	next  *Element
}

type List struct {
	head *Element
}

// New converts a slice of ints to singly linked list
func New(input []int) *List {
	var list *List
	list = new(List)
	for _, v := range input {
		list.Push(v)
	}
	return list
}

// Size returns the number of elements in the singly linked list
func (l *List) Size() int {
	count := 0
	if l != nil {
		head := l.head
		for head != nil {
			count++
			head = head.next
		}
	}
	return count
}

// Push inserts a new element in front of the singly linked list
func (l *List) Push(element int) {
	var e *Element
	e = new(Element)
	e.value = element
	if l.head != nil {
		e.next = l.head
	}
	l.head = e
}

// Pop removes (and then returns) the front element of the singly linked list
func (l *List) Pop() (int, error) {
	if l.head == nil {
		return 0, fmt.Errorf("empty list")
	} else {
		v := l.head.value
		l.head = l.head.next
		return v, nil
	}
}

// Array converts a singly linked list to a slice of ints
func (l *List) Array() []int {
	var a []int
	head := l.head
	for head != nil {
		a = append([]int{head.value}, a...)
		head = head.next
	}
	return a
}

// Reverse reverses the order of elements in the singly linked list
func (l *List) Reverse() *List {
	var newList *List
	newList = new(List)
	for l.head != nil {
		v, _ := l.Pop()
		newList.Push(v)
	}
	return newList
}
