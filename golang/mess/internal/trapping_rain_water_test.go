package internal

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO fix this
// https://leetcode.cn/problems/trapping-rain-water/?envType=study-plan-v2&envId=top-100-liked
func TestTrap(t *testing.T) {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	assert.Equal(t, 6, solution(height))
	assert.Equal(t, 9, solution([]int{4, 2, 0, 3, 2, 5}))
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

	fmt.Println("left", left)
	fmt.Println("right", right)

	result := 0

	for i := 1; i < len(height)-1; i++ {
		result += min(left[i], right[i]) - height[i]
	}

	fmt.Println(result)

	return result
}
