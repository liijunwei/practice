package main

import (
	"fmt"
	"testing"

	assert1 "github.com/stretchr/testify/assert"
)

// https://leetcode.cn/problems/maximum-average-subarray-i/
//
// 给你一个由 n 个元素组成的整数数组 nums 和一个整数 k 。
// 请你找出平均数最大且 长度为 k 的连续子数组，并输出该最大平均数。
func TestFindMaxAverage(t *testing.T) {
	// assert1.Equal(t, 12.75, findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4))
	// assert1.Equal(t, 5.0, findMaxAverage([]int{5}, 1))
	assert1.Equal(t, 2.8, findMaxAverage([]int{4, 0, 4, 3, 3}, 5)) // TODO figure out why this case wont pass?
}

func findMaxAverage(nums []int, k int) float64 {
	assert(k > 0)

	presumresult := presum(nums)
	n := len(presumresult)

	// range sum: sum(L, R) = prefix[R] - prefix[L-1]
	//
	// raw:    [1 2 3 4 9]
	// presum: [1 3 6 10 19]
	//          0   k
	//              k-2   k
	// sum(0,0+k)
	// sum(1,1+k)
	// sum(2,2+k)
	// ...
	// sum(n-k,n-k+k)

	max := nums[0]

	for i := 0; i < n-k; i++ {
		tmp := presumresult[i+k] - presumresult[i]

		if tmp > max {
			max = tmp
		}

		fmt.Println("i", i, "tmp", tmp, "max", max, "presumresult", presumresult)
	}

	return float64(max) / float64(k)
}
