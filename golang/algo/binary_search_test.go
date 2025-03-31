package algo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{1, 2, 3, 4, 5}, 3, 2},
		{[]int{1, 2, 3, 4, 5}, 6, -1},
		{[]int{}, 1, -1},
	}

	for _, tt := range tests {
		got1 := binarySearch1(tt.nums, tt.target)
		assert.Equal(t, tt.want, got1)

		got2 := binarySearch2(tt.nums, tt.target)
		assert.Equal(t, tt.want, got2)
	}
}

func binarySearch1(nums []int, target int) int {
	var rank func(nums []int, lo, hi int) int

	rank = func(nums []int, lo, hi int) int {
		if hi < lo {
			return -1
		}

		mid := lo + (hi-lo)/2
		if nums[mid] < target {
			lo = mid + 1
		} else if nums[mid] > target {
			hi = mid - 1
		} else {
			return mid
		}

		return rank(nums, lo, hi)
	}

	return rank(nums, 0, len(nums)-1)
}

// assmume nums are sorted in ascending order
func binarySearch2(nums []int, target int) int {
	lo, hi := 0, len(nums)-1

	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] > target {
			hi = mid - 1
		} else if nums[mid] < target {
			lo = mid + 1
		} else {
			return mid
		}
	}

	return -1
}
