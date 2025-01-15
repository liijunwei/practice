package twosum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwosum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9

	result := twoSum(nums, target)
	assert.Equal(t, []int{0, 1}, result)

	nums = []int{2, 7, 11, 15}
	target = 10
	assert.Panics(t, func() { twoSum(nums, target) })
}

func twoSum(nums []int, target int) []int {
	cache := make(map[int]int)

	for i, m := range nums {
		complement := target - m
		fmt.Println("target", target, "m", m, "complement", complement, "cache", cache)
		if n, ok := cache[complement]; ok {
			return []int{n, i}
		}
		cache[m] = i
	}

	// depends on the edge case handling
	panic("not found")
}

func twoSumV2(nums []int, target int) []int {
	cache := make(map[int]int)

	for i, m := range nums {
		cache[m] = i
	}

	for i, m := range nums {
		if n, ok := cache[target-m]; ok && n != i {
			return []int{i, n}
		}
	}

	return nil
}
