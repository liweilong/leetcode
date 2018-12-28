package main

import (
	"fmt"
	"strconv"
)

func addStrings(num1 string, num2 string) string {
	l1, l2 := len(num1), len(num2)
	rs := make([]int, 0)
	extra := 0
	for i := 0; i < l1 || i < l2; i++ {
		n1, n2 := 0, 0
		if i < l1 {
			n1 = int(num1[l1-i-1] - 48)
		}
		if i < l2 {
			n2 = int(num2[l2-i-1] - 48)
		}
		sum := n1 + n2 + extra
		rs = append(rs, sum%10)
		extra = sum / 10
	}
	if extra > 0 {
		rs = append(rs, extra)
	}
	str := ""
	lr := len(rs)
	for i := 0; i < lr; i++ {
		n := rs[lr-i-1]
		if str == "" && n == 0 {
			continue
		}
		str += strconv.Itoa(n)
	}
	if str == "" {
		return "0"
	}
	return str
}

func main() {
	num1, num2 := "155", "841"
	fmt.Println(addStrings(num1, num2))
}
