package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	s = strings.Trim(s, " ")
	l := len(s)
	rs := make([]uint8, l)
	for i := 0; i < l; i++ {
		start := i
		for i < l && s[i] != 32 {
			i++
		}
		for j, k := i-1, start; j >= start; j-- {
			rs[k] = s[j]
			k++
		}
		if i != l {
			rs[i] = s[i]
		}
	}
	fmt.Println(string(rs))
	for i := 0; i < l/2; i++ {
		t := rs[i]
		rs[i] = rs[l-i-1]
		rs[l-i-1] = t
	}
	return string(rs)
}

func main() {
	s := "  the sky   is blue  "
	fmt.Println(reverseWords(s))
}
