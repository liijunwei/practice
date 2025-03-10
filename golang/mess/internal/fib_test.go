package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFib(t *testing.T) {
	assert.Equal(t, fib_dp(2), fib_recur(2))
	assert.Equal(t, fib_dp(3), fib_recur(3))
	assert.Equal(t, fib_dp(40), fib_cache(40))
}

var N = 40

// cd mess/internal && go test -benchmem -run=^$ -bench ^Benchmark_fib . -v
func Benchmark_fib_dp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib_dp(N)
	}
}

func Benchmark_fib_recur(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib_recur(N)
	}
}

func Benchmark_fib_cache(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib_cache(N)
	}
}

func fib_dp(n int) int {
	if n <= 1 {
		return n
	}

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	// spew.Dump(dp)
	return dp[n]
}

func fib_recur(n int) int {
	if n <= 1 {
		return n
	}

	return fib_recur(n-1) + fib_recur(n-2)
}

func fib_cache(n int) int {
	cache := make(map[int]int)
	return fib(n, cache)
}

func fib(n int, cache map[int]int) int {
	if n <= 1 {
		return n
	}

	if _, ok := cache[n]; ok {
		return cache[n]
	}

	cache[n] = fib(n-1, cache) + fib(n-2, cache)

	return cache[n]
}
