package main

import (
	"errors"
	"math"
	"testing"
)

var errInvalidWindow = errors.New("invalid window size")

// brute-force reference: try every subarray of length k, return max sum
// oracle implementation
func maxSubarraySumBrute(arr []int, k int) (int, error) {
	if k > len(arr) || k <= 0 {
		return 0, errInvalidWindow
	}
	maxSum := math.MinInt
	for i := 0; i <= len(arr)-k; i++ {
		sum := 0
		for j := i; j < i+k; j++ {
			sum += arr[j]
		}
		if sum > maxSum {
			maxSum = sum
		}
	}
	return maxSum, nil
}

func TestInvalidWindow(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		k    int
	}{
		{"k zero", []int{1, 2, 3}, 0},
		{"k negative", []int{1, 2, 3}, -1},
		{"k larger than array", []int{1}, 2},
		{"empty array", []int{}, 1},
		{"empty array k zero", []int{}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := maxSubarraySum(tt.arr, tt.k)
			if err == nil {
				t.Fatal("expected error, got nil")
			}
		})
	}
}

func TestExhaustiveSmallArrays(t *testing.T) {
	// Generate all arrays of length 0..7 with values in [-3, 3]
	for arrLen := 0; arrLen <= 7; arrLen++ {
		arrays := generateAllArrays(arrLen, -3, 3)
		for _, arr := range arrays {
			for k := 1; k <= arrLen+1; k++ {
				got, gotErr := maxSubarraySum(arr, k)
				want, wantErr := maxSubarraySumBrute(arr, k)

				if (gotErr == nil) != (wantErr == nil) {
					t.Errorf("arr=%v k=%d: got err=%v, want err=%v", arr, k, gotErr, wantErr)
					continue
				}
				if gotErr == nil && got != want {
					t.Errorf("arr=%v k=%d: got=%d, want=%d", arr, k, got, want)
				}
			}
		}
	}
}

func TestValidEdgeCases(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		k    int
		want int
	}{
		{"single element k=1", []int{5}, 1, 5},
		{"k equals array length", []int{3, -2, 4}, 3, 5},
		{"all negative", []int{-5, -1, -3}, 2, -4},
		{"all positive", []int{1, 2, 3, 4}, 2, 7},
		{"max at beginning", []int{10, 1, 1, 1}, 2, 11},
		{"max at end", []int{1, 1, 1, 10}, 2, 11},
		{"with zeros", []int{0, 0, 5, 0}, 2, 5},
		{"overflow safety", []int{math.MaxInt32, math.MinInt32, 0}, 2, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := maxSubarraySum(tt.arr, tt.k)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("got=%d, want=%d", got, tt.want)
			}
		})
	}
}

// generateAllArrays returns all arrays of given length with each element in [lo, hi].
func generateAllArrays(length int, lo, hi int) [][]int {
	if length == 0 {
		return [][]int{{}}
	}
	rangeLen := hi - lo + 1
	total := 1
	for i := 0; i < length; i++ {
		total *= rangeLen
	}
	result := make([][]int, 0, total)
	for i := 0; i < total; i++ {
		arr := make([]int, length)
		n := i
		for j := 0; j < length; j++ {
			arr[j] = lo + n%rangeLen
			n /= rangeLen
		}
		result = append(result, arr)
	}
	return result
}
