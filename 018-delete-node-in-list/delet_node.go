/*
@Time : 2019-11-18 13:17
@Author : lfn
@File : delet_node
*/

package _18_delete_node_in_list

type ListNode struct {
	val int
	next *ListNode
}

func deleteNode(head, target *ListNode) {
	if head == nil || target == nil {
		return
	}
	if head == target && head.next == nil {
		head = nil
		return
	}
	if target.next == nil {
		searchAndDeleteFromHead(head, target)
		return
	}
	target.val = target.next.val
	target.next = target.next.next
	target.next.next = nil
}

func searchAndDeleteFromHead(head, target *ListNode) {
	cur := head
	for cur.next != target {
		cur = cur.next
	}
	cur.next = cur.next.next
	cur.next.next = nil
}







