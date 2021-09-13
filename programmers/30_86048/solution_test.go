package programmers_30_86048

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_solution(t *testing.T) {
	assert.Equal(t, []int{0, 1, 1}, solution([]int{1, 3, 2}, []int{1, 2, 3}))
	assert.Equal(t, []int{2, 2, 1, 3}, solution([]int{1, 4, 2, 3}, []int{2, 1, 3, 4}))
	assert.Equal(t, []int{1, 1, 2}, solution([]int{3, 2, 1}, []int{2, 1, 3}))
	assert.Equal(t, []int{2, 2, 2}, solution([]int{3, 2, 1}, []int{1, 3, 2}))
	assert.Equal(t, []int{2, 2, 0, 2}, solution([]int{1, 4, 2, 3}, []int{2, 1, 4, 3}))
}
