package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClimbStairs(t *testing.T) {
	t.Run("recursive", func(t *testing.T) {
		assert.Equal(t, 1, climbStairsRe(1))
		assert.Equal(t, 2, climbStairsRe(2))
		assert.Equal(t, 3, climbStairsRe(3))
		assert.Equal(t, 5, climbStairsRe(4))
		assert.Equal(t, 8, climbStairsRe(5))
	})

	t.Run("dynamic programming", func(t *testing.T) {
		assert.Equal(t, 1, climbStairsDP(1))
		assert.Equal(t, 2, climbStairsDP(2))
		assert.Equal(t, 3, climbStairsDP(3))
		assert.Equal(t, 5, climbStairsDP(4))
		assert.Equal(t, 8, climbStairsDP(5))
	})
}

func climbStairsRe(n int) int {
	return dfs(n)
}

func dfs(i int) int {
	if i <= 1 {
		return 1
	}

	return dfs(i-1) + dfs(i-2)
}

func climbStairsDP(n int) int {
	if n <= 1 {
		return 1
	}

	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
