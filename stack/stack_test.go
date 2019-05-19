package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack(t *testing.T) {
	stack := New()
	assert.True(t, stack.IsEmpty())

	stack.Push(10, 5)
	assert.Equal(t, 2, stack.Len())
	assert.Equal(t, 5, stack.Peek())
	assert.Equal(t, 5, stack.Pop())
	assert.Equal(t, 10, stack.Pop())
	assert.Equal(t, 0, stack.Len())
	assert.Nil(t, stack.Pop())
}
