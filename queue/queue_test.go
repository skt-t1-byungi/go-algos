package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	queue := New()
	assert.True(t, queue.IsEmpty())

	queue.Enqueue(4, 2, 5)

	assert.Equal(t, 3, queue.Len())
	assert.Equal(t, 4, queue.Peak())
	assert.Equal(t, 4, queue.Dequeue())
	assert.Equal(t, 2, queue.Len())
	assert.Equal(t, 2, queue.Dequeue())
	assert.Equal(t, 1, queue.Len())
	assert.Equal(t, 5, queue.Dequeue())
	assert.True(t, queue.IsEmpty())
	assert.Nil(t, queue.Dequeue())
}
