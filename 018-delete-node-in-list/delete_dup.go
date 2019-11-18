/*
@Time : 2019-11-18 14:49
@Author : lfn
@File : delete_dup
*/

package _18_delete_node_in_list

func deleteDup(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{0, head}
	cur := dummy.next
	for cur.next != nil {
		if cur.next.val == cur.val {
			cur.next = cur.next.next
		} else {
			cur = cur.next
		}
	}
	return dummy.next
}
