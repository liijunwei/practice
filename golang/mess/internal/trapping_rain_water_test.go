package internal

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrap(t *testing.T) {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	assert.Equal(t, 6, solution(height))
}

func solution(height []int) int {
	left := make([]int, len(height))
	right := make([]int, len(height))

	for i := 1; i < len(height); i++ {
		left[i] = max(left[i-1], height[i])
	}

	for i := len(height) - 2; i >= 0; i-- {
		right[i] = max(right[i+1], height[i])
	}

	fmt.Println(left)
	fmt.Println(right)

	result := 0

	for i := 1; i < len(height)-1; i++ {
		result += min(left[i], right[i]) - height[i]
	}

	fmt.Println(result)

	return result
}
