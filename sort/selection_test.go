package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSelect(t *testing.T) {
	assert.Equal(t, []int{1, 2}, Selection([]int{2, 1}))
	assert.Equal(t, []int{1, 2, 3, 4, 5}, Selection([]int{2, 5, 3, 4, 1}))
}
