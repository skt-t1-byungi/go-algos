package linkedlist

import "errors"

var OutOfBoundsError = errors.New("OutOfBoundsError")

type LinkedList struct {
	head *Node
	len  int
}

type Node struct {
	next  *Node
	value interface{}
}

func New() *LinkedList {
	return &LinkedList{nil, 0}
}

func (ll *LinkedList) Append(val interface{}) {
	_ = ll.InsertAt(ll.len, val)
}

func (ll *LinkedList) InsertAt(at int, val interface{}) error {
	if at < 0 || at > ll.len {
		return OutOfBoundsError
	}

	curr := &Node{nil, val}
	if at == 0 {
		ll.head = curr
	} else {
		prev := ll.nodeAt(at - 1)
		if prev.next != nil {
			curr.next = prev.next
		}
		prev.next = curr
	}
	ll.len++

	return nil
}

func (ll *LinkedList) nodeAt(at int) *Node {
	curr := ll.head
	for ; at > 0; at-- {
		curr = curr.next
	}
	return curr
}

func (ll *LinkedList) At(at int) (interface{}, error) {
	if ll.isOutOfBounds(at) {
		return nil, OutOfBoundsError
	}
	return ll.nodeAt(at).value, nil
}

func (ll *LinkedList) isOutOfBounds(i int) bool {
	return i < 0 || i >= ll.len
}

func (ll *LinkedList) RemoveAt(at int) error {
	if ll.isOutOfBounds(at) {
		return OutOfBoundsError
	}

	var currRef **Node
	if at == 0 {
		currRef = &ll.head
	} else {
		currRef = &ll.nodeAt(at - 1).next
	}

	var tail *Node
	if curr := *currRef; curr.next != nil {
		tail = curr.next
		curr.next = nil
	}
	*currRef = tail
	ll.len--

	return nil
}

func (ll *LinkedList) Size() int {
	return ll.len
}

func (ll *LinkedList) Each(iteratee func(interface{}, int)) {
	curr := ll.head
	for i := 0; i < ll.len; i++ {
		iteratee(curr.value, i)
		curr = curr.next
	}
}

func (ll *LinkedList) Slice() []interface{} {
	r := make([]interface{}, ll.Size())
	ll.Each(func(val interface{}, i int) {
		r[i] = val
	})
	return r
}
