package main

import "fmt"

// https://leetcode.com/problems/median-of-two-sorted-arrays/

// There are two sorted arrays nums1 and nums2 of size m and n respectively.
//
// Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).
//
// You may assume nums1 and nums2 cannot be both empty.
//
// Example 1:
//
// nums1 = [1, 3]
// nums2 = [2]
//
// The median is 2.0
// Example 2:
//
// nums1 = [1, 2]
// nums2 = [3, 4]
//
// The median is (2 + 3)/2 = 2.5

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1, n2 := len(nums1), len(nums2)
	tn := n1 + n2

	return float64(tn)
}

func main() {
	nums1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	nums2 := []int{3, 4, 5, 6, 7, 8, 9}

	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
