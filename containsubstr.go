package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	l1 := len(s1)
	l2 := len(s2)
	nums1 := make([]int, 128)
	nums2 := make([]int, 128)
	cs := make([]uint8, 0)
	for i := 0; i < l1; i++ {
		if nums1[s1[i]] == 0 {
			cs = append(cs, s1[i])
		}
		nums1[s1[i]]++
	}
	for i := 0; i < l2; i++ {
		nums2[s2[i]]++
		if i >= l1-1 {
			if compare(nums1, nums2, cs) {
				return true
			}
			nums2[s2[i-l1+1]]--
		}
	}
	return false
}

func compare(nums1 []int, nums2 []int, cs []uint8) bool {
	for _, c := range cs {
		if nums1[c] != nums2[c] {
			return false
		}
	}
	return true
}

func main() {
	s1 := "ab"
	s2 := ""
	fmt.Println(checkInclusion(s1, s2))
}
