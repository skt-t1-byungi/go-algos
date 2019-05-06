package linkedlist

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func ExampleLinkedList_Append() {
	ll := New()
	ll.Append(3)
	ll.Append(2)
	ll.Append(5)
	printList(ll)
	// output:
	// 3,nil,2
	// 2,3,5
	// 5,2,nil
}

func ExampleLinkedList_Prepend() {
	ll := New()
	ll.Prepend(3)
	ll.Prepend(2)
	ll.Prepend(5)
	printList(ll)
	// output:
	// 5,nil,2
	// 2,5,3
	// 3,2,nil
}

func TestLinkedList_At(t *testing.T) {
	ll := New()
	ll.Append(3)
	ll.Append(2)
	ll.Append(5)
	assert.Equal(t, 3, ll.At(0).(int))
	assert.Equal(t, 2, ll.At(1).(int))
	assert.Equal(t, 5, ll.At(2).(int))
	assert.Nil(t, ll.At(-1))
	assert.Nil(t, ll.At(3))
	ll.Append(4)
	assert.Equal(t, 4, ll.At(3).(int))
}

func BenchmarkLinkedList_At(b *testing.B) {
	prepare := func() *LinkedList {
		ll := New()
		for i := 0; i < 20000; i++ {
			ll.Append(0)
		}
		return ll
	}
	b.Run("At(9999) - size:20000", func(b *testing.B) {
		ll := prepare()
		b.ResetTimer()
		for i := 0; i < 10000; i++ {
			ll.At(9999)
		}
	})
	b.Run("At(19999) - size:20000", func(b *testing.B) {
		ll := prepare()
		b.ResetTimer()
		for i := 0; i < 10000; i++ {
			ll.At(19999)
		}
	})
}

func TestLinkedList_Slice(t *testing.T) {
	ll := New()
	for _, n := range []int{3, 1, 2} {
		ll.Append(n)
	}
	assert.Equal(t, []interface{}{3, 1, 2}, ll.Slice())
}

func TestLinkedList_InsertAt(t *testing.T) {
	ll := New()
	ll.InsertAt(0, 3)
	ll.InsertAt(0, 1)
	ll.InsertAt(-1, 2)
	ll.InsertAt(1, 4)
	assert.Equal(t, []interface{}{1, 4, 2, 3}, ll.Slice())

	ll.InsertAt(7, 5)
	assert.Equal(t, []interface{}{1, 4, 2, 3, nil, nil, nil, 5}, ll.Slice())
}

func TestLinkedList_RemoveFirst(t *testing.T) {
	ll := New()
	assert.False(t, ll.RemoveFirst())
	ll.Append(3)
	ll.Append(2)
	assert.Equal(t, 2, ll.Len())
	assert.Equal(t, 3, ll.First().(int))
	assert.True(t, ll.RemoveFirst())
	assert.Equal(t, 1, ll.Len())
	assert.Equal(t, 2, ll.First().(int))
}

func TestLinkedList_RemoveLast(t *testing.T) {
	ll := New()
	assert.False(t, ll.RemoveLast())
	ll.Append(3)
	ll.Append(2)
	assert.Equal(t, 2, ll.Len())
	assert.Equal(t, 2, ll.Last().(int))
	assert.True(t, ll.RemoveLast())
	assert.Equal(t, 1, ll.Len())
	assert.Equal(t, 3, ll.Last().(int))
}

func TestLinkedList_RemoveAt(t *testing.T) {
	ll := New()
	for _, n := range []int{3, 7, 1, 2, 6, 5} {
		ll.Append(n)
	}
	assert.Equal(t, []interface{}{3, 7, 1, 2, 6, 5}, ll.Slice())
	assert.True(t, ll.RemoveAt(-2))
	assert.Equal(t, []interface{}{3, 7, 1, 2, 5}, ll.Slice())
	assert.True(t, ll.RemoveAt(2))
	assert.Equal(t, []interface{}{3, 7, 2, 5}, ll.Slice())
	assert.True(t, ll.RemoveAt(0))
	assert.Equal(t, []interface{}{7, 2, 5}, ll.Slice())
	assert.True(t, ll.RemoveAt(2))
	assert.Equal(t, []interface{}{7, 2}, ll.Slice())
	assert.False(t, ll.RemoveAt(2))
}

func TestLinkedList_Find(t *testing.T) {
	ll := New()
	for _, n := range []int{3, 7, 1, 2, 6, 5} {
		ll.Append(n)
	}
	found, index := ll.Find(func(val interface{}, i int) bool {
		n := val.(int)
		return n > 5 && n < 7
	})
	assert.Equal(t, 6, found)
	assert.Equal(t, 4, index)
}

func TestLinkedList_SetAt(t *testing.T) {
	ll := New()
	ll.Append(3)
	ll.Append(2)
	ll.SetAt(1, 5)
	assert.Equal(t, []interface{}{3, 5}, ll.Slice())
	ll.SetAt(4, 1)
	assert.Equal(t, []interface{}{3, 5, nil, nil, 1}, ll.Slice())
}

func TestLinkedList_Len(t *testing.T) {
	ll := New()
	assert.Equal(t, 0, ll.Len())
	ll.Append(1)
	assert.Equal(t, 1, ll.Len())
}

func printList(ll *LinkedList) {
	get := func(n *Node) string {
		if n != nil {
			return strconv.Itoa(n.value.(int))
		}
		return "nil"
	}
	ll.each(func(node *Node) {
		fmt.Printf("%d,%s,%s\n", node.value, get(node.prev), get(node.next))
	})
}
