package linkedlist

import (
	"errors"
)

var OutOfRangeError = errors.New("OutOfRangeError")

type linkedList struct {
	head *node
	len  int
}

type node struct {
	next  *node
	value interface{}
}

func New() *linkedList {
	return &linkedList{nil, 0}
}

func (l *linkedList) Append(val interface{}) {
	var idx int

	if l.len == 0 {
		idx = 0
	} else {
		idx = l.len
	}

	_ = l.InsertAt(idx, val)
}

func (l *linkedList) InsertAt(idx int, val interface{}) error {
	if idx < 0 || idx > l.len {
		return OutOfRangeError
	}

	newNode := &node{nil, val}
	l.len++

	if l.len == 1 {
		l.head = newNode
		return nil
	}

	curr := l.nodeByIndex(idx - 1)
	if curr.next != nil {
		newNode.next = curr.next
	}
	curr.next = newNode

	return nil
}

func (l *linkedList) nodeByIndex(idx int) *node {
	curr := l.head
	for ; idx > 0; idx-- {
		curr = curr.next
	}
	return curr
}

func (l *linkedList) At(idx int) (interface{}, error) {
	if l.isOutOfRange(idx) {
		return nil, OutOfRangeError
	}
	return l.nodeByIndex(idx).value, nil
}

func (l *linkedList) isOutOfRange(idx int) bool {
	return idx < 0 || idx >= l.len
}

func (l *linkedList) RemoveAt(idx int) error {
	if l.isOutOfRange(idx) {
		return OutOfRangeError
	}

	curr := l.nodeByIndex(idx - 1)

	var tail *node
	if curr.next.next != nil {
		tail = curr.next.next
	}
	curr.next = tail
	l.len--

	return nil
}

func (l *linkedList) Size() int {
	return l.len
}

func (l *linkedList) Each(iteratee func(interface{}, int)) {
	curr := l.head
	for i := 0; i < l.len; i++ {
		iteratee(curr.value, i)
		curr = curr.next
	}
}
