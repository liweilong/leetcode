package main

import "fmt"

// https://leetcode.com/problems/two-sum/

// Given an array of integers, return indices of the two numbers such that they add up to a specific target.
//
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
//
// Example:
//
// Given nums = [2, 7, 11, 15], target = 9,
//
// Because nums[0] + nums[1] = 2 + 7 = 9,
// return [0, 1].

// 时间复杂度O(n) 空间复杂度O(n)
// 注意[3,3] 6 这种情况
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := numMap[target-nums[i]]; ok && numMap[target-nums[i]] != i {
			return []int{numMap[target-nums[i]], i}
		}
		numMap[nums[i]] = i
	}
	return nil
}

func main() {
	fmt.Print(twoSum([]int{3, 2, 4}, 6))
}
