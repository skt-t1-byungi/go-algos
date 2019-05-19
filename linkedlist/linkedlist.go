package linkedlist

type LinkedList struct {
	head *Node
	tail *Node
	len  int
}

type Node struct {
	prev  *Node
	next  *Node
	value interface{}
}

func New() *LinkedList {
	return &LinkedList{nil, nil, 0}
}

func (ll *LinkedList) each(iteratee func(node *Node)) {
	for node := ll.head; node != nil; node = node.next {
		iteratee(node)
	}
}

func (ll LinkedList) Slice() []interface{} {
	slice := make([]interface{}, ll.len)
	i := -1
	ll.each(func(node *Node) {
		i++
		slice[i] = node.value
	})
	return slice
}

func (ll *LinkedList) Len() int {
	return ll.len
}

func (ll *LinkedList) Append(values ...interface{}) {
	head, last := createNodes(values)
	if ll.len == 0 {
		ll.head = head
	} else {
		ll.tail.next = head
		head.prev = ll.tail
	}
	ll.tail = last
	ll.len += len(values)
}

func (ll *LinkedList) Prepend(values ...interface{}) {
	if ll.len == 0 {
		ll.Append(values...)
	} else {
		head, last := createNodes(values)
		last.next = ll.head
		ll.head.prev = last
		ll.head = head
		ll.len += len(values)
	}
}

func (ll *LinkedList) At(idx int) interface{} {
	node := ll.nodeAt(idx)
	if node != nil {
		return node.value
	}
	return nil
}

func (ll *LinkedList) First() interface{} {
	return ll.At(0)
}

func (ll *LinkedList) Last() interface{} {
	return ll.At(ll.len - 1)
}

func (ll *LinkedList) nodeAt(idx int) *Node {
	if idx < 0 || idx >= ll.len {
		return nil
	}

	var node *Node
	if idx < ll.len>>1 {
		node = ll.head
		for i := 0; i < idx; i++ {
			node = node.next
		}
	} else {
		node = ll.tail
		for i := ll.len - 1; i > idx; i-- {
			node = node.prev
		}
	}
	return node
}

func (ll *LinkedList) InsertAt(idx int, values ...interface{}) {
	idx = withNegativeIndex(ll.len, idx)

	if idx == 0 {
		ll.Prepend(values...)
		return
	}
	if idx == ll.len {
		ll.Append(values...)
		return
	}
	if idx > ll.len {
		ll.Append(make([]interface{}, idx-ll.len)...)
		ll.Append(values...)
		return
	}

	head, last := createNodes(values)
	oldNode := ll.nodeAt(idx)
	prevNode := oldNode.prev
	head.prev = prevNode
	last.next = oldNode
	oldNode.prev = last
	prevNode.next = head
	ll.len += len(values)
}

func (ll *LinkedList) RemoveFirst() bool {
	return ll.removeEdge(func() {
		newHead := ll.head.next
		ll.head.next = nil
		newHead.prev = nil
		ll.head = newHead
	})
}

func (ll *LinkedList) RemoveLast() bool {
	return ll.removeEdge(func() {
		newTail := ll.tail.prev
		ll.tail.prev = nil
		newTail.next = nil
		ll.tail = newTail
	})
}

func (ll *LinkedList) removeEdge(remover func()) bool {
	if ll.len == 0 {
		return false
	}
	if ll.len == 1 {
		ll.head = nil
		ll.tail = nil
	} else {
		remover()
	}
	ll.len--
	return true
}

func (ll *LinkedList) RemoveAt(idx int) bool {
	idx = withNegativeIndex(ll.len, idx)
	if idx > ll.len-1 {
		return false
	}

	if idx == 0 {
		return ll.RemoveFirst()
	}
	if idx == ll.len-1 {
		return ll.RemoveLast()
	}

	node := ll.nodeAt(idx)
	prevNode, nextNode := node.prev, node.next
	prevNode.next, nextNode.prev = nextNode, prevNode
	node.prev, node.next = nil, nil
	ll.len--
	return true
}

func (ll *LinkedList) SetAt(idx int, val interface{}) {
	idx = withNegativeIndex(ll.len, idx)
	if idx > ll.len {
		ll.InsertAt(idx, val)
		return
	}

	node := ll.nodeAt(idx)
	node.value = val
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

func (ll *LinkedList) IsEmpty() bool {
	return ll.len == 0
}

func (ll *LinkedList) Empty() {
	ll.len = 0
	ll.head = nil
	ll.tail = nil
}

func createNodes(values []interface{}) (*Node, *Node) {
	var head *Node
	var last *Node
	for _, value := range values {
		curr := &Node{nil, nil, value}
		if last == nil {
			head = curr
		} else {
			curr.prev = last
			last.next = curr
		}
		last = curr
	}
	return head, last
}

func withNegativeIndex(len, idx int) int {
	if idx < 0 {
		if len == 0 {
			idx = 0
		} else {
			idx = len + idx%len
		}
	}
	return idx
}
