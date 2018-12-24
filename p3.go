package main

import (
	"fmt"
)

// https://leetcode.com/problems/longest-substring-without-repeating-characters/

// Given a string, find the length of the longest substring without repeating characters.
//
// Example 1:
//
// Input: "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.
// Example 2:
//
// Input: "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.
// Example 3:
//
// Input: "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
// Note that the answer must be a substring, "pwke" is a subsequence and not a substring.

// 利用滑动窗口计算，indexes用于记录每个字母开始计算长度的位置(a...a]
func lengthOfLongestSubstring(s string) int {
	max := 0
	n := len(s)
	indexes := make([]int, 128)
	for i, j := 0, 0; j < n; j++ {
		if indexes[s[j]] > i {
			i = indexes[s[j]]
		}
		if j-i+1 > max {
			max = j - i + 1
		}
		indexes[s[j]] = j + 1
	}
	return max
}

func main() {
	s1 := "abcabcbb"
	s2 := "bbbbb"
	s3 := "pwwkew"
	s4 := ""
	s5 := " "
	s6 := "abb"
	s7 := "abba"

	fmt.Println(lengthOfLongestSubstring(s1))
	fmt.Println(lengthOfLongestSubstring(s2))
	fmt.Println(lengthOfLongestSubstring(s3))
	fmt.Println(lengthOfLongestSubstring(s4))
	fmt.Println(lengthOfLongestSubstring(s5))
	fmt.Println(lengthOfLongestSubstring(s6))
	fmt.Println(lengthOfLongestSubstring(s7))
}
