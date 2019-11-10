/*
@Time : 2019-11-10 21:16
@Author : lfn
@File : print_test
*/

package _06_print_list

import (
	"testing"
)

func TestReversePrintListWithStack(t *testing.T) {
	n1 := &ListNode{1, nil}
	n2 := &ListNode{2, nil}
	n3 := &ListNode{3, nil}
	n4 := &ListNode{4, nil}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4

	tests := map[string]struct{
		head *ListNode
		want string
	}{
		"simple": {n1,"4->3->2->1"},
		"null": {nil, ""},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := reversePrintWithStack(tt.head)
			if got != tt.want {
				t.Errorf("want=%s, got=%s\n", tt.want, got)
			}
		})
	}
}

func TestReversePrintListRecursive(t *testing.T) {
	n1 := &ListNode{1, nil}
	n2 := &ListNode{2, nil}
	n3 := &ListNode{3, nil}
	n4 := &ListNode{4, nil}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4

	tests := map[string]struct{
		head *ListNode
		want string
	}{
		"simple": {n1,"4->3->2->1"},
		"null": {nil, ""},
		"one node": {n4, "4"},
		"two node": {n3, "4->3"},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := reversePrintRecursive(tt.head)
			if got != tt.want {
				t.Errorf("want=%s, got=%s\n", tt.want, got)
			}
		})
	}
}
