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
	s := make([]interface{}, ll.len)
	i := -1
	ll.each(func(node *Node) {
		i++
		s[i] = node.value
	})
	return s
}

func (ll *LinkedList) Len() int {
	return ll.len
}

func (ll *LinkedList) Append(val interface{}) {
	newNode := &Node{nil, nil, val}
	var ptr **Node
	if ll.len == 0 {
		ptr = &ll.head
	} else {
		ptr = &ll.tail.next
		newNode.prev = ll.tail
	}
	*ptr = newNode
	ll.tail = newNode
	ll.len++
}

func (ll *LinkedList) Prepend(val interface{}) {
	if ll.len == 0 {
		ll.Append(val)
	} else {
		prevHead := ll.head
		ll.head = &Node{nil, prevHead, val}
		prevHead.prev = ll.head
		ll.len++
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

func (ll *LinkedList) InsertAt(idx int, val interface{}) {
	idx = withNegativeIndex(ll.len, idx)

	if idx == 0 {
		ll.Prepend(val)
		return
	} else if idx == ll.len {
		ll.Append(val)
		return
	}

	if idx > ll.len {
		for i := ll.len; i < idx; i++ {
			ll.Append(nil)
		}
		ll.Append(val)
		return
	}

	oldNode := ll.nodeAt(idx)
	prevNode := oldNode.prev
	newNode := &Node{prevNode, oldNode, val}
	oldNode.prev = newNode
	prevNode.next = newNode
	ll.len++
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

func (ll *LinkedList) isEmpty() bool {
	return ll.len == 0
}

func (ll *LinkedList) Empty() {
	ll.len = 0
	ll.head = nil
	ll.tail = nil
}
