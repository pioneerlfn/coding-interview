/*
@Time : 2019-11-12 11:56
@Author : lfn
@File : next_node
*/

package _07_next_node_in_binary_tree

type BinaryTree struct {
	Val byte
	Left *BinaryTree
	Right *BinaryTree
	Parent *BinaryTree
}

func findNextNodeValue(root, node *BinaryTree) byte {
	res := findNextNode(root, node)
	return res.Val
}

func findNextNode(root, node *BinaryTree) *BinaryTree {
	if root == nil || node == nil {
		return nil
	}
	cur := node
	// If a node has right child tree,
	// then the next node is the leftmost of the right child tree.
	if node.Right != nil {
		cur = node.Right
		for cur.Left != nil {
			cur = cur.Left
		}
		return cur
	} else if node == node.Parent.Left {
		// don't have a right tree
		// and the node is the left child of it's parent
		// then the next node is it's patent.
		return node.Parent
	} else {
		// don't have right child tree
		// and is neither the left child of it's parent.
		cur = node.Parent
		for cur.Parent.Left != cur && cur.Parent != nil {
			cur = cur.Parent
		}
		return cur.Parent
	}
}
