/*
@Time : 2019-11-11 15:28
@Author : lfn
@File : build_binary_tree
*/

package _07_build_BST

type BinaryTree struct {
	Val int
	Left *BinaryTree
	Right *BinaryTree
}

func buildAndPostTraverse(pre, in []int) []int {
	tree := buildFromPreAndInorder(pre, in)
	return postTraverse(tree)
}

func buildFromPreAndInorder(pre, in []int) *BinaryTree {
	if len(pre) != len(in) || len(pre) == 0 {
		return nil
	}
	root := &BinaryTree{pre[0], nil, nil}
	if len(pre) == 1 {
		return root
	}
	var i int
	for ; i < len(in); i++  {
		if root.Val == in[i] {
			break
		}
	}
	preLeft, preRight := pre[1:i+1], pre[i+1:]
	inorderLeft, inorderRight := in[:i], in[i+1:]
	left := buildFromPreAndInorder(preLeft, inorderLeft)
	right := buildFromPreAndInorder(preRight, inorderRight)
	root.Left, root.Right = left, right
	return root
}

func postTraverse(root *BinaryTree) []int {
	if root == nil {
		return nil
	}
	var res []int
	res = append(res, postTraverse(root.Left)...)
	res = append(res, postTraverse(root.Right)...)
	res = append(res, root.Val)
	return res
}
