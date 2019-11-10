/*
@Time : 2019-11-10 21:03
@Author : lfn
@File : reverse_print_list
*/

package _06_print_list

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func reversePrintWithStack(head *ListNode) string {
	if head == nil {
		return ""
	}
	if head.Next == nil {
		return strconv.Itoa(head.Val)
	}

	cur := head
	var stack []int
	for cur != nil {
		stack = append(stack, cur.Val)
		cur = cur.Next
	}
	var sb strings.Builder
	for i := len(stack)-1; i >= 0; i-- {
		sb.WriteString(strconv.Itoa(stack[i]))
		sb.WriteString("->")
	}
	res := sb.String()
	return res[:len(res)-2]
}

func reversePrintRecursive(head *ListNode) string {
	var sb strings.Builder
	var printRecursive func(*ListNode, *strings.Builder)
	printRecursive = func(head *ListNode, sb *strings.Builder)  {
		if head == nil {
			return
		}
		if head.Next == nil {
			sb.WriteString(strconv.Itoa(head.Val))
			return
		}
		printRecursive(head.Next, sb)
		sb.WriteString("->")
		sb.WriteString(strconv.Itoa(head.Val))
	}
	printRecursive(head, &sb)
	tmp := sb.String()
	return tmp
}