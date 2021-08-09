package programmers_30_83201

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_solution(t *testing.T) {
	assert.Equal(t, "FBABD", solution([][]int{
		{100, 90, 98, 88, 65},
		{50, 45, 99, 85, 77},
		{47, 88, 95, 80, 67},
		{61, 57, 100, 80, 65},
		{24, 90, 94, 75, 65},
	}))
	assert.Equal(t, "DA", solution([][]int{
		{50, 90},
		{50, 87},
	}))
	assert.Equal(t, "CFD", solution([][]int{
		{70, 49, 90},
		{68, 50, 38},
		{73, 31, 100},
	}))
}
