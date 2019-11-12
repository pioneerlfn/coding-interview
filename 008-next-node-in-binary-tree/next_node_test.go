/*
@Time : 2019-11-12 13:52
@Author : lfn
@File : next_node_test
*/

package _07_next_node_in_binary_tree

import (
	"reflect"
	"testing"
)

func TestFindNextNode(t *testing.T) {
	a := &BinaryTree{'a', nil, nil, nil}
	b := &BinaryTree{'b', nil, nil, a}
	c := &BinaryTree{'c', nil, nil, a}
	d := &BinaryTree{'d', nil, nil, b}
	e := &BinaryTree{'e', nil, nil, b}
	f := &BinaryTree{'f', nil, nil, c}
	g := &BinaryTree{'g', nil, nil, c}
	h := &BinaryTree{'h', nil, nil, e}
	i := &BinaryTree{'i', nil, nil, e}
	a.Left, a.Right = b, c
	b.Left, b.Right = d, e
	c.Left, c.Right = f, g
	e.Left, e.Right = h, i

	tests := map[string]struct{
		root *BinaryTree
		target *BinaryTree
		want byte
	}{
		"without right and is the left":{a, h, 'e'},
		"without right and is the right":{a, i, 'a'},
		"has right":{a, e, 'i'},
		"root":{a, a, 'f'},
		"null":{a, nil, '\x00'},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := findNextNodeValue(tt.root, tt.target)
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("want=%v, got=%v\n", tt.want, got)
			}
		})
	}
}









