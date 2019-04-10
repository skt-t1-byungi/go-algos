package stack

import "testing"

func TestStack(t *testing.T) {
	stack := New()
	if !stack.isEmpty() {
		t.Errorf("Expected true, but %t", stack.isEmpty())
	}

	stack.Push(10)
	if v, _ := stack.Peek(); v.(int) != 10 {
		t.Errorf("Expected 10, but %d", v)
	}
	if stack.isEmpty() {
		t.Errorf("Expected false, but %t", stack.isEmpty())
	}

	stack.Push(5)
	if v, _ := stack.Peek(); v.(int) != 5 {
		t.Errorf("Expected 5, but %d", v)
	}

	if v, _ := stack.Pop(); v.(int) != 5 {
		t.Errorf("Expected 5, but %d", v)
	}
	if v, _ := stack.Pop(); v.(int) != 10 {
		t.Errorf("Expected 10, but %d", v)
	}
	if _, err := stack.Pop(); err == nil {
		t.Error("Expected OutOfBoundsError, but nil")
	}
	if !stack.isEmpty() {
		t.Errorf("Expected true, but %t", stack.isEmpty())
	}
}
