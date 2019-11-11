/*
@Time : 2019-11-11 15:39
@Author : lfn
@File : build_binary_tree_test
*/

package _07_build_BST

import (
	"reflect"
	"testing"
)

func TestBuildBinaryTree(t *testing.T) {
	tests := map[string]struct{
		pre []int
		inorder []int
		want []int
	}{
		"1": {[]int{1, 2, 4, 7, 3, 5, 6, 8},
			[]int{4, 7, 2, 1, 5, 3, 8, 6},
			[]int{7, 4, 2, 5, 8, 6, 3, 1}},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			tree := buildFromPreAndInorder(tt.pre, tt.inorder)
			got := PostTranverse(tree)
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("want=%v, got=%v\n", tt.want, got)
			}
		})
	}
}

func PostTranverse(root *BinaryTree) []int {
	if root == nil {
		return nil
	}
	var res []int
	res = append(res, PostTranverse(root.Left)...)
	res = append(res, PostTranverse(root.Right)...)
	res = append(res, root.Val)
	return res
}
