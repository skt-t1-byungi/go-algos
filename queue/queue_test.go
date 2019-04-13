package queue

import "testing"

func TestQueue(t *testing.T) {
	queue := New()
	if !queue.IsEmpty() {
		t.Errorf("Expected true, but %t", queue.IsEmpty())
	}

	queue.Enqueue(4)
	queue.Enqueue(2)
	queue.Enqueue(5)
	if queue.Size() != 3 {
		t.Errorf("Expected 3, but %d", queue.Size())
	}
	if v, _ := queue.Peak(); v != 4 {
		t.Errorf("Expected 4, but %d", v)
	}

	if v, _ := queue.Dequeue(); v != 4 {
		t.Errorf("Expected 4, but %d", v)
	}
	if v, _ := queue.Dequeue(); v != 2 {
		t.Errorf("Expected 4, but %d", v)
	}
	if v, _ := queue.Dequeue(); v != 5 {
		t.Errorf("Expected 4, but %d", v)
	}
	if _, ok := queue.Dequeue(); ok {
		t.Errorf("Expected false, but %t", ok)
	}
}
