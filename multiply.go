package main

import (
	"fmt"
	"strconv"
)

// 给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。
//
// 示例 1:
//
// 输入: num1 = "2", num2 = "3"
// 输出: "6"
// 示例 2:
//
// 输入: num1 = "123", num2 = "456"
// 输出: "56088"
// 说明：
//
// num1 和 num2 的长度小于110。
// num1 和 num2 只包含数字 0-9。
// num1 和 num2 均不以零开头，除非是数字 0 本身。
// 不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。

func multiply(num1 string, num2 string) string {
	l1, l2 := len(num1), len(num2)
	if l1 < l2 {
		temp := num1
		num1 = num2
		num2 = temp
		t := l1
		l1 = l2
		l2 = t
	}
	nums := make([]string, 0)
	for i := 0; i < l2; i++ {
		s := m0(num1, int(num2[l2-i-1])-48)
		nums = append(nums, s)
	}
	return add(nums)
}

func m0(num1 string, num int) string {
	l := len(num1)
	extra := 0
	rs := make([]int, 0)
	for i := 0; i < l; i++ {
		n := num1[l-i-1] - 48
		t := int(n)*num + extra
		rs = append(rs, t%10)
		extra = t / 10
	}
	if extra != 0 {
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

func add(nums []string) string {
	result := make([]int, 0)
	extra := 0
	v0 := uint8(48)
	for i := 0; ; i++ {
		hasNum := false
		sum := 0
		for j, s := range nums {
			if i-j < 0 {
				break
			}
			if i-j < len(s) {
				sum += int(s[len(s)-(i-j)-1] - v0)
				hasNum = true
			}
		}
		if !hasNum && extra == 0 {
			break
		}
		sum += extra
		extra = sum / 10
		result = append(result, sum%10)
	}
	str := ""
	lr := len(result)
	for i := 0; i < lr; i++ {
		n := result[lr-i-1]
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
	num1, num2 := "140", "721"

	fmt.Println(multiply(num1, num2))
}
