package linkedlist

import (
	"reflect"
	"testing"
)

func TestLinkedList(t *testing.T) {
	l := New()
	l.Append(3)
	l.Append(1)
	l.Append(2)

	if l.Size() != 3 {
		t.Errorf("Expected 3, but %d", l.Size())
	}

	r := make([]int, 0)
	l.Each(func(v interface{}, _ int) {
		r = append(r, v.(int))
	})

	if !reflect.DeepEqual(r, []int{3, 1, 2}) {
		t.Errorf("Expected [3 1 2], but %v", r)
	}
}
