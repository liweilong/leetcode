package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// 给定一组非负整数，重新排列它们的顺序使之组成一个最大的整数。
//
// 示例 1:
//
// 输入: [10,2]
// 输出: 210
// 示例 2:
//
// 输入: [3,30,34,5,9]
// 输出: 9534330
// 说明: 输出结果可能非常大，所以你需要返回一个字符串而不是整数。

func largestNumber(nums []int) string {
	strNums := make([]string, 0)
	for _, num := range nums {
		strNums = append(strNums, strconv.Itoa(num))
	}
	sort.Strings(strNums)
	l := len(nums)
	ss := make([]string, l, l)
	for i := 0; i < l; i++ {
		ss[l-i-1] = strNums[i]
	}

	return strings.Join(ss, "")
}

func main() {
	nums := []int{3, 30, 34, 5, 9}
	fmt.Println(largestNumber(nums))
}
