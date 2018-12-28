package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	l := len(nums)
	sort.Ints(nums)
	for i := 0; i < l-2; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		sum := -nums[i]

		for j, k := i+1, l-1; j < k; {
			t := nums[j] + nums[k]
			if t == sum {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				j++
				k--
			} else if t > sum {
				k--
			} else {
				j++
			}
			for j < k && k+1 < l && nums[k] == nums[k+1] {
				k--
			}
			for j < k && j-1 > i && nums[j] == nums[j-1] {
				j++
			}
		}
	}
	return result
}

func main() {
	//nums := []int{-2,0,0,2,2}
	//nums := []int{-1, 0, 1, 2, -1, -4}
	nums := []int{0, 0, 0, 0}
	result := threeSum(nums)
	for _, r := range result {
		fmt.Println(r)
	}
}
