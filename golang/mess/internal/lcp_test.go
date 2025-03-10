package internal

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLCP(t *testing.T) {
	// longestCommonPrefix([]string{"flower", "flow", "flight"})
	// longestCommonPrefix([]string{"dog", "racecar", "car"})
	// assert.Equal(t, "fl", longestCommonPrefix([]string{"flower", "flow", "flight", "fl"}))
	assert.Equal(t, "a", longestCommonPrefix([]string{"ab", "a"}))
	// assert.Equal(t, "", longestCommonPrefix([]string{"dog", "racecar", "car"}))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	s0 := strs[0]
	// 一列一列遍历
	for j := range s0 {
		for _, s := range strs[1:] {
			fmt.Println("loop", "s0", s0, "j", j, "s", s)
			if j == len(s) || s[j] != s0[j] {
				fmt.Println("return", "s0", s0, "j", j, "s", s, "s0[:j]", s0[:j])
				return s0[:j]
			}
		}
	}
	return s0

	// 作者：灵茶山艾府
	// 链接：https://leetcode.cn/problems/longest-common-prefix/solutions/2801713/jian-dan-ti-jian-dan-zuo-pythonjavaccgoj-478q/
}
