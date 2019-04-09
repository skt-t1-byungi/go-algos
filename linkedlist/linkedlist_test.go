package linkedlist

import (
	"reflect"
	"testing"
)

func TestLinkedList(t *testing.T) {
	ll := New()
	if ll.Size() != 0 {
		t.Errorf("Expected 0, but %d", ll.Size())
	}

	ll.Append(3)
	ll.Append(1)
	ll.Append(2)

	if v, _ := ll.At(0); v.(int) != 3 {
		t.Errorf("Expected 3, but %d", v)
	}

	if v, _ := ll.At(1); v.(int) != 1 {
		t.Errorf("Expected 1, but %d", v)
	}

	if ll.Size() != 3 {
		t.Errorf("Expected 3, but %d", ll.Size())
	}

	if slice := ll.Slice(); !reflect.DeepEqual(slice, []interface{}{3, 1, 2}) {
		t.Errorf("Expected [3 1 2], but %v", slice)
	}

	_ = ll.RemoveAt(1)
	if ll.Size() != 2 {
		t.Errorf("Expected 2, but %d", ll.Size())
	}

	if slice := ll.Slice(); !reflect.DeepEqual(slice, []interface{}{3, 2}) {
		t.Errorf("Expected [3 2], but %v", slice)
	}

	_ = ll.InsertAt(1, 5)
	if slice := ll.Slice(); !reflect.DeepEqual(slice, []interface{}{3, 5, 2}) {
		t.Errorf("Expected [3 5 2], but %v", slice)
	}
}
