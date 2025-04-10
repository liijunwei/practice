package algo

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.cn/problems/trapping-rain-water/?envType=study-plan-v2&envId=top-100-liked
func TestTrap(t *testing.T) {
	assert.Equal(t, 6, solution([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	assert.Equal(t, 9, solution([]int{4, 2, 0, 3, 2, 5}))
	assert.Equal(t, 0, solution([]int{4}))
	assert.Equal(t, 0, solution([]int{}))
	assert.Equal(t, 0, solution([]int{4}))
	assert.Equal(t, 0, solution([]int{4, 3}))
	assert.Equal(t, 1, solution([]int{4, 2, 3}))
	assert.Equal(t, 0, solution([]int{4, 3, 3}))
}

func solution(height []int) int {
	if len(height) == 0 {
		return 0
	}

	left := make([]int, len(height))
	right := make([]int, len(height))

	left[0] = height[0]
	right[len(height)-1] = height[len(height)-1]

	for i := 1; i < len(height); i++ {
		left[i] = max(left[i-1], height[i])
	}

	for i := len(height) - 2; i >= 0; i-- {
		right[i] = max(right[i+1], height[i])
	}

	debug("left", left)
	debug("right", right)

	resultArr := make([]int, len(height))
	resultArr[0] = 0
	resultArr[len(height)-1] = 0

	// head and tail are not included, because they are not trapped
	for i := 1; i < len(height)-1; i++ {
		resultArr[i] = min(left[i], right[i]) - height[i]
	}

	debug("resultArr", resultArr)

	result := 0
	for i := 0; i < len(resultArr); i++ {
		result += resultArr[i]
	}

	debug("result", result)
	debug()

	return result
}

func debug(s ...any) {
	if os.Getenv("DEBUG") == "" {
		return
	}

	fmt.Println(s...)
}
