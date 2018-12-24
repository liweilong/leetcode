package main

import "fmt"

// https://leetcode.com/problems/add-two-numbers/

// You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.
//
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.
//
// Example:
//
// Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
// Output: 7 -> 0 -> 8
// Explanation: 342 + 465 = 807.

type ListNode struct {
	Val  int
	Next *ListNode
}

// 注意大数运算

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	n1 := l1
	n2 := l2
	sum := 0
	result := &ListNode{}
	r := result
	for n1 != nil || n2 != nil || sum != 0 {
		v1 := 0
		v2 := 0
		if n1 != nil {
			v1 = n1.Val
			n1 = n1.Next
		}
		if n2 != nil {
			v2 = n2.Val
			n2 = n2.Next
		}
		tmp := v1 + v2 + sum
		extra := tmp / 10
		num := tmp % 10
		r.Val = num
		sum = extra
		if extra > 0 || n1 != nil || n2 != nil {
			r.Next = &ListNode{}
			r = r.Next
		}
	}
	return result
}

func main() {
	r1 := []int{9, 8, 7}
	r2 := []int{4, 5, 6}
	l1 := &ListNode{}
	n1 := l1
	for i := 0; i < len(r1); i++ {
		n1.Val = r1[i]
		if i != len(r1)-1 {
			n1.Next = &ListNode{}
			n1 = n1.Next
		}
	}
	n1 = l1
	for n1 != nil {
		fmt.Print(n1.Val)
		n1 = n1.Next
	}
	fmt.Println()

	l2 := &ListNode{}
	n2 := l2
	for i := 0; i < len(r2); i++ {
		n2.Val = r2[i]
		if i != len(r2)-1 {
			n2.Next = &ListNode{}
			n2 = n2.Next
		}
	}
	n2 = l2
	for n2 != nil {
		fmt.Print(n2.Val)
		n2 = n2.Next
	}
	fmt.Println()

	l := addTwoNumbers(l1, l2)
	n := l
	fmt.Printf("result:")

	for n != nil {
		fmt.Print(n.Val)
		n = n.Next
	}
}
