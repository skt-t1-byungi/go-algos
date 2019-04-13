package linkedlist

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
	ll.InsertAt(ll.len, val)
}

func (ll *LinkedList) InsertAt(at int, val interface{}) bool {
	if at < 0 || at > ll.len {
		return false
	}

	var ptr **Node
	if at == 0 {
		ptr = &ll.head
	} else {
		ptr = &ll.nodeAt(at - 1).next
	}

	curr := &Node{nil, val}
	if next := *ptr; next != nil {
		curr.next = next
	}
	*ptr = curr
	ll.len++

	return true
}

func (ll *LinkedList) nodeAt(at int) *Node {
	curr := ll.head
	for ; at > 0; at-- {
		curr = curr.next
	}
	return curr
}

func (ll *LinkedList) At(at int) (interface{}, bool) {
	if ll.isOutOfBounds(at) {
		return nil, false
	}
	return ll.nodeAt(at).value, true
}

func (ll *LinkedList) isOutOfBounds(i int) bool {
	return i < 0 || i >= ll.len
}

func (ll *LinkedList) RemoveAt(at int) bool {
	if ll.isOutOfBounds(at) {
		return false
	}

	var ptr **Node
	if at == 0 {
		ptr = &ll.head
	} else {
		ptr = &ll.nodeAt(at - 1).next
	}

	var tail *Node
	if curr := *ptr; curr.next != nil {
		tail = curr.next
		curr.next = nil
	}
	*ptr = tail
	ll.len--

	return true
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

func (ll *LinkedList) Find(finder func(val interface{}, i int) bool) (interface{}, int) {
	curr := ll.head
	for i := 0; i < ll.len; i++ {
		if finder(curr.value, i) {
			return curr.value, i
		}
		curr = curr.next
	}
	return nil, -1
}
