package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertion(t *testing.T) {
	assert.Equal(t, []int{1, 2}, Insertion([]int{2, 1}))
	assert.Equal(t, []int{1, 2, 3, 4, 5}, Insertion([]int{2, 5, 3, 4, 1}))
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, Insertion([]int{6, 2, 8, 7, 5, 3, 4, 9, 1}))
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, Insertion([]int{5, 1, 8, 4, 9, 3, 2, 7, 6}))
}
