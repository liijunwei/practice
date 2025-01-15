package p0

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	for _, r := range "hello" {
		fmt.Println(r, string(r))
	}
}

func TestSlidingWindow(t *testing.T) {
	nums := []int{2, 3, 1, 2, 4, 3}
	s := 7
	result := minSubArrayLen(s, nums)
	fmt.Println(result) // 输出：2
}

func minSubArrayLen(s int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	left, right := 0, 0
	sum := 0
	minLen := n + 1

	for right < n {
		sum += nums[right]

		for sum >= s {
			minLen = min(minLen, right-left+1)
			sum -= nums[left]
			left++
		}

		right++
		println("left", left, "right", right)
	}

	if minLen == n+1 {
		return 0
	}

	return minLen
}

func Test2(t *testing.T) {
	left := 0
	maxLen := 0
	s := "abcabcbb"

	charIdx := make(map[byte]int)

	for right := 0; right < len(s); right++ {
		if i, ok := charIdx[s[right]]; ok {
			// old := left
			left = max(left, i+1)
			// println("move left", "old", old, "new", left)
		}
		charIdx[s[right]] = right
		maxLen = max(maxLen, right-left+1)
		println(s[right], string(s[right]), "left", left, "right", right, "maxLen", maxLen, "charIdx", toJSON(charIdx))
	}

	println(maxLen)
}

func toJSON(d any) string {
	data, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}
	return string(data)
}
