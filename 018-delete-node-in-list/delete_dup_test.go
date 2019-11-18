/*
@Time : 2019-11-18 14:55
@Author : lfn
@File : delete_dup_test
*/

package _18_delete_node_in_list

import (
	"reflect"
	"testing"
)

func TestDeleteDup(t *testing.T) {
	n1 := &ListNode{1, nil}
	n2 := &ListNode{1, nil}
	n3 := &ListNode{1, nil}
	n4 := &ListNode{1, nil}
	n5 := &ListNode{2, nil}
	n6 := &ListNode{3, nil}
	n7 := &ListNode{3, nil}
	n8 := &ListNode{4, nil}
	n1.next = n2
	n2.next = n3
	n3.next = n4
	n4.next = n5
	n5.next = n6
	n6.next = n7
	n7.next = n8

	tests := map[string]struct{
		head *ListNode
		want []int
	}{
		"0":{n1, []int{1,2,3,4}},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			deDup := deleteDup(tt.head)
			var got []int
			for cur := deDup; cur != nil; cur = cur.next {
				got = append(got, cur.val)
			}
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("want=%v, got=%v\n", tt.want, got)
			}
		})
	}
}
