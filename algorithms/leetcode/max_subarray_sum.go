package main

import "errors"

func maxSubarraySum(arr []int, k int) (int, error) {
	if k > len(arr) || k <= 0 {
		return 0, errors.New("invalid window size")
	}

	// first window windowSum
	windowSum := 0
	for i := range k {
		windowSum += arr[i]
	}

	maxSum := windowSum
	for i := k; i < len(arr); i++ {
		windowSum = windowSum - arr[i-k] + arr[i]
		if windowSum > maxSum {
			maxSum = windowSum
		}
	}
	return maxSum, nil
}
