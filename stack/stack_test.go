package stack

import "testing"

func TestStack(t *testing.T) {
	stack := New()
	if !stack.IsEmpty() {
		t.Errorf("Expected true, but %t", stack.IsEmpty())
	}

	stack.Push(10)
	stack.Push(5)
	if stack.Size() != 2 {
		t.Errorf("Expected 2, but %d", stack.Size())
	}

	if v, _ := stack.Peek(); v.(int) != 5 {
		t.Errorf("Expected 5, but %d", v)
	}
	if v, _ := stack.Pop(); v.(int) != 5 {
		t.Errorf("Expected 5, but %d", v)
	}
	if v, _ := stack.Pop(); v.(int) != 10 {
		t.Errorf("Expected 10, but %d", v)
	}
	if _, ok := stack.Pop(); !ok {
		t.Errorf("Expected false, but %t", ok)
	}
	if !stack.IsEmpty() {
		t.Errorf("Expected true, but %t", stack.IsEmpty())
	}
}
