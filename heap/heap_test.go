package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHeap(t *testing.T) {
	h := NewHeap()
	h.Push(5)
	h.Push(3)
	h.Push(6)
	h.Push(1)
	h.Push(4)
	h.Push(2)
	assert.Equal(t, 1, h.Pop())
	assert.Equal(t, 2, h.Pop())
	assert.Equal(t, 3, h.Pop())
	assert.Equal(t, 4, h.Pop())
	assert.Equal(t, 5, h.Pop())
	assert.Equal(t, 6, h.Pop())
}
