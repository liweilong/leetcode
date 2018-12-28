package main

import "fmt"

// 编写一个函数来查找字符串数组中的最长公共前缀。
//
// 如果不存在公共前缀，返回空字符串 ""。
//
// 示例 1:
//
// 输入: ["flower","flow","flight"]
// 输出: "fl"
// 示例 2:
//
// 输入: ["dog","racecar","car"]
// 输出: ""
// 解释: 输入不存在公共前缀。
// 说明:
//
// 所有输入只包含小写字母 a-z 。

func longestCommonPrefix(strs []string) string {
	if strs == nil || len(strs) <= 0 {
		return ""
	}
	prefix := make([]byte, 0)
	end := false
	for j := 0; !end; j++ {
		if len(strs[0]) <= j {
			end = true
			break
		}
		c := strs[0][j]
		for _, s := range strs {
			if len(s) <= j || s[j] != c {
				end = true
				break
			}
		}
		if !end {
			prefix = append(prefix, c)
		}
	}
	return string(prefix)
}

func main() {
	strs := []string{"f", "flaaa", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}
